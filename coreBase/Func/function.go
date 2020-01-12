package functions

import (
	Configurations "coreBase/Configurations"
	DataTypes "coreBase/Type"
	Utils "coreBase/Utils"
	verify "coreBase/verify"
	"encoding/json"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
)

func ProcessRegister(request []byte) (DataTypes.VerifyContactResponse, error) {
	
	var vr DataTypes.VerifyContactRequest
	var user DataTypes.User
	var res DataTypes.VerifyContactResponse
	//unmarshal request
	err := json.Unmarshal([]byte(request), &user)
	if err != nil {
		return DataTypes.VerifyContactResponse{}, errors.New("err: error in unmarshal user")
	}
	//db
	session := Configurations.Mongodb.Session
	if session == nil {
		return DataTypes.VerifyContactResponse{}, errors.New("err: error in DB session")
	}
	//defer session.Close()
	// check mobile use

	coll := session.DB(Configurations.Configs.DatabaseName).C("users")
	
	//add user to db
	
	user.Verify = true
	err = coll.Insert(&user)
	if err != nil {
		fmt.Print(err)
		return DataTypes.VerifyContactResponse{}, errors.New("err: error in DB")
	}

	//token
	t, err := newToken(&user.Mobile)
	if err != nil {
		return DataTypes.VerifyContactResponse{}, errors.New("err: error in creat token")
	}
	//verify
	coll = session.DB(Configurations.Configs.DatabaseName).C("verify")
	vr.Code = Utils.CreateVerifyCodeString(4)
	vr.Verify = bson.NewObjectId().Hex()
	vr.Contact = user.Mobile
	vr.Token = t
	err = coll.Insert(&vr)
	if err != nil {
		return DataTypes.VerifyContactResponse{}, errors.New("err: error in DB")
	}

	//send sms
	go func() {
		err = verify.SendSms(user.Mobile, vr.Code)
		if err != nil {
			return
		}

	}()
	// response verify id
	res.Verify = vr.Verify

	return res, nil
}
func ProcessVerify(request []byte) (DataTypes.VerificationResponse, error) {
	var req DataTypes.VerificationRequest
	var vr DataTypes.VerifyContactRequest
	//unmarshal request
	err := json.Unmarshal([]byte(request), &req)
	if err != nil {
		return DataTypes.VerificationResponse{}, errors.New("err: error in unmarshal user")
	}
	//db
	session := Configurations.Mongodb.Session
	if session == nil {
		return DataTypes.VerificationResponse{}, errors.New("err: error in DB session")
	}
	//defer session.Close()
	// check code
	fmt.Println("verify id:", req.Verify)
	fmt.Println("code:", req.Code)
	coll := session.DB(Configurations.Configs.DatabaseName).C("verify")
	bs := bson.M{"$and": []bson.M{{"verify": req.Verify}, {"code": req.Code}}}
	err = coll.Find(bs).One(&vr)
	if err != nil {
		return DataTypes.VerificationResponse{}, errors.New("err: code not currect")
	}
	//update user verify
	coll = session.DB(Configurations.Configs.DatabaseName).C("users")

	bs = bson.M{"mobile": vr.Contact}

	//err = coll.Find(bs).One(&u)
	err = coll.Update(bs, bson.M{"$set": bson.M{"verify": true}})
	if err != nil {
		return DataTypes.VerificationResponse{}, errors.New("err: error in  user DB ")
	}
	fmt.Println("token", vr.Token)
	return DataTypes.VerificationResponse{
		Token: vr.Token,
	}, nil
}
func ProcessLogin(request []byte) (DataTypes.LoginResponse, error) {
	var req DataTypes.LoginRequest

	//unmarshal request
	err := json.Unmarshal([]byte(request), &req)
	if err != nil {
		return DataTypes.LoginResponse{}, errors.New("err: error in unmarshal user")
	}
	//db
	err = CheckUserLogin(req.Username, req.Password)
	if err != nil {
		return DataTypes.LoginResponse{}, err
	}
	//token
	t, err := newToken(&req.Username)
	if err != nil {
		return DataTypes.LoginResponse{}, errors.New("err: error in creat token")
	}
	//verify
	//Verify, err := AddVerifyCode(req.Mobile, t)
	//if err != nil {
	//	return DataTypes.VerifyContactResponse{}, err
	//}

	return DataTypes.LoginResponse{
	Token: t,
	}, nil
}


func CheckUserVerify(mobile string) error {
	var u DataTypes.User

	session := Configurations.Mongodb.Session
	if session == nil {
		return errors.New("err: error in db session")
	}
	//	defer session.Close()
	// check code
	coll := session.DB(Configurations.Configs.DatabaseName).C("users")
	bs := bson.M{"mobile": mobile}
	err := coll.Find(bs).One(&u)
	if err != nil {
		return errors.New("err:user not found")
	}
	if !u.Verify {
		return errors.New("err:user not verifed")
	}

	return nil
}

func CheckUserLogin(username string, password string) error {
	var u DataTypes.User

	session := Configurations.Mongodb.Session
	if session == nil {
		return errors.New("err: error in db session")
	}
	//	defer session.Close()
	// check code
	coll := session.DB(Configurations.Configs.DatabaseName).C("users")
	bs := bson.M{"$and": []bson.M{{"username": username}, {"password": password}}}
	err := coll.Find(bs).One(&u)
	if err != nil {
		return errors.New("incorrect username or password")
	}
	if !u.Verify {
		return errors.New("err:user not verifed")
	}

	return nil
}
func AddVerifyCode(mobile string, t string) (string, error) {
	session := Configurations.Mongodb.Session
	if session == nil {
		return "", errors.New("err: error in db session")
	}
	var vr DataTypes.VerifyContactRequest
	coll := session.DB(Configurations.Configs.DatabaseName).C("verify")

	vr.Code = Utils.CreateVerifyCodeString(4)
	vr.Verify = bson.NewObjectId().Hex()
	vr.Contact = mobile
	vr.Token = t

	err := coll.Insert(&vr)
	if err != nil {
		return "", errors.New("err: error in DB insert")
	}

	return vr.Verify, nil
}
func ProcessCheckUser(request []byte) error {
	var req DataTypes.CheckUserRequest
	var u DataTypes.User
	//unmarshal request
	err := json.Unmarshal([]byte(request), &req)
	if err != nil {
		return errors.New("err: error in unmarshal user")
	}
	//db
	session := Configurations.Mongodb.Session
	if session == nil {
		return errors.New("err: error in DB session")
	}
	//defer session.Close()
	// check user token
	username,err:=CheckToken(&req.Token)
	
    
	coll := session.DB(Configurations.Configs.DatabaseName).C("users")
	bs := bson.M{"username": username}
	err = coll.Find(bs).One(&u)
	if err != nil {
		return errors.New("err: user not found")
	}
	if !u.Verify {
		return errors.New("err: user not active")
	}
	
	return nil
}
func ProcessAddRequest(request []byte) error {
	var req DataTypes.RequestPayload
	var m DataTypes.RequestPayload
	//unmarshal request
	err := json.Unmarshal([]byte(request), &req)
	if err != nil {
		return errors.New("err: error in unmarshal user")
	}
	//check data
	if req.MessageID == "" {
		return errors.New("err: error in message id")
	}
	//db
	session := Configurations.Mongodb.Session
	if session == nil {
		return errors.New("err: error in DB session")
	}
	//defer session.Close()
	// check request
	coll := session.DB(Configurations.Configs.DatabaseName).C("request")
	bs := bson.M{"message_id": req.MessageID}
	err = coll.Find(bs).One(&m)
	if err == nil {
		return errors.New("err:request  found")
	}

	coll = session.DB(Configurations.Configs.DatabaseName).C("request")

	err = coll.Insert(req)
	if err != nil {
		return errors.New("err: user not found")
	}
	return nil
}
func ProcessAddFile(request []byte) error {
	var req DataTypes.RequestFile
	var r DataTypes.RequestFile
	//unmarshal request
	err := json.Unmarshal([]byte(request), &req)
	if err != nil {
		return errors.New("err: error in unmarshal user")
	}
	//check data
	
	//db
	session := Configurations.Mongodb.Session
	if session == nil {
		return errors.New("err: error in DB session")
	}
	//defer session.Close()
	// check request
	coll := session.DB(Configurations.Configs.DatabaseName).C("files")
	

	bs := bson.M{"file_id": req.FileID}
	err = coll.Find(bs).One(&r)
	if err == nil {
		return errors.New("err: file  found")
	}
	
	err = coll.Insert(&req)
	if err != nil {
		fmt.Print(err)
		return  errors.New("err: error in DB")
	}

	
	return nil
}
func newToken(u *string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": u,
		"expire": time.Now().Add(15 * time.Minute),
		"time":   time.Now(),
	})
	tokenString, err := token.SignedString([]byte("advjuH@(F&H#(#FOUEGV&@(#Fv2982b9o&g397"))
	if err != nil {
		return "", errors.New("internal server error [jwt-p-sec]")
	}
	return tokenString, nil
}
func CheckToken(u *string) (string, error) {
	tokenString := *u  
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
    return []byte("advjuH@(F&H#(#FOUEGV&@(#Fv2982b9o&g397"), nil
	})
	
// ... error handling
	if err!=nil{
	return "",err
	}
	fmt.Println(token)
	s:=claims["username"].(string)
	
	return s,nil
}