package DataTypes

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	Request1 struct {
		Cmd     string `json:"cmd" valid:"required"`
		Request string `json:"request" valid:"required"`
	}
	Point1 struct {
		ID           bson.ObjectId `json:"_id"`
		Code         string        `json:"code" bson:"code"`
		Item1Value   string        `json:"item_1_value" bson:"item_1_value"`
		Item2Value   string        `json:"item_2_value" bson:"item_2_value"`
		Item3Value   string        `json:"item_3_value" bson:"item_3_value"`
		Item3Current string        `json:"item_3_current" bson:"item_3_current"`
		Reason       string        `json:"reason" bson:"reason"`
		Step         string        `json:"step" bson:"step"`
		Foundation   string        `json:"foundation" bson:"foundation"`
		Village      string        `json:"village" bson:"village"`
		State        string        `json:"state" bson:"state"`
		City         string        `json:"city" bson:"city"`
		Households   string        `json:"households" bson:"households"`
		Lat          string        `json:"lat" bson:"lat"`
		Long         string        `json:"long" bson:"long"`
		Images       []string      `json:"images" bson:"images"`
		Zone         string        `json:"zone"`
		Active       bool          `json:"active"`
	}

	States1 []State

	State1 struct {
		ID          bson.ObjectId `json:"_id" bson:"_id"`
		Name        string        `json:"name" bson:"name"`
		DisplayName string        `json:"display_name" bson:"display_name"`
		Lat         string        `json:"lat" bson:"lat"`
		Long        string        `json:"long" bson:"long"`
		Cites       []City        `json:"cites" bson:"cites"`
		Trash       bool          `json:"trash" bson:"trash"`
	}
	NewState1 struct {
		ID          string `json:"id" valid:"optional,alphanum,length(24|24)"`
		Name        string `json:"name" bson:"name" valid:"required"`
		DisplayName string `json:"display_name" bson:"display_name" valid:"required"`
		Lat         string `json:"lat" bson:"lat" valid:"required"`
		Long        string `json:"long" bson:"long" valid:"required"`
	}

	City1 struct {
		CityID      bson.ObjectId `json:"city_id" bson:"city_id"`
		Name        string        `json:"name" bson:"name"`
		DisplayName string        `json:"display_name" bson:"display_name"`
		Lat         string        `json:"lat" bson:"lat"`
		Long        string        `json:"long" bson:"long"`
		Villages    []Village     `json:"villages" bson:"villages"`
		Trash       bool          `json:"trash" bson:"trash"`
	}
	NewCity1 struct {
		ID          string `json:"id" valid:"optional,alphanum,length(24|24)"`
		Name        string `json:"name" valid:"required"`
		DisplayName string `json:"display_name" valid:"required"`
		Lat         string `json:"lat" valid:"required"`
		Long        string `json:"long" valid:"required"`
		State       string `json:"state" valid:"required,alphanum,length(24|24)"`
	}

	OID struct {
		OID string `json:"uoid"`
	}

	Village1 struct {
		VillageID   bson.ObjectId `json:"village_id" bson:"village_id"`
		Name        string        `json:"name" bson:"name"`
		DisplayName string        `json:"display_name" bson:"display_name"`
		Lat         string        `json:"lat" bson:"lat"`
		Long        string        `json:"long" bson:"long"`
		Trash       bool          `json:"trash" bson:"trash"`
	}
	NewVillage1 struct {
		ID          string `json:"id" valid:"optional,alphanum,length(24|24)"`
		Name        string `json:"name" valid:"required"`
		DisplayName string `json:"display_name" valid:"required"`
		Lat         string `json:"lat" valid:"required"`
		Long        string `json:"long" valid:"required"`
		City        string `json:"city" valid:"required,alphanum,length(24|24)"`
		State       string `json:"state" valid:"required,alphanum,length(24|24)"`
	}
	PointsRequest1 struct {
		Zone string `json:"zone" valid:"required,alphanum,length(24|24)"`
	}
	ImageRequest struct {
		Code string `json:"code" valid:"required"`
	}

	CreatePaymentRequest struct {
		Code   string `json:"code" valid:"required,numeric"`
		Amount string `json:"amount" valid:"required,numeric"`
		Date   string `json:"date" valid:"required,numeric"`
		Time   string `json:"time" valid:"required,numeric"`
		Secret string `json:"secret"  valid:"required"`
	}
	CreatePaymentResponse1 struct {
		Amount  string `json:"amount" bson:"amount"`
		Code    string `json:"code" bson:"code"`
		Date    string `json:"date" bson:"date"`
		Time    string `json:"time" bson:"time"`
		Status  string `json:"status" bson:"status"`
		Info    string `json:"info" bson:"info"`
		OrderID string `json:"order_id" bson:"order_id"`
		PayerID string `json:"payer_id" bson:"payer_id"`
	}
	PaymentStatus struct {
		OrderID string `json:"order_id" valid:"required"`
		Status  string `json:"status" valid:"required"`
		Info    string `json:"info"`
		Secret  string `json:"secret"`
	}
	PaymentStatusResponse1 struct {
		Result string `json:"result"`
	}
	PaymentSoapResponse struct {
		Response string `json:"response"`
	}
	LoginRequest1 struct {
		Username string `json:"username" valid:"required,alphanum"`
		Password string `json:"password" valid:"required"`
	}
	JwtToken struct {
		Token string `json:"token"`
	}
	Exception struct {
		Message string `json:"message"`
	}
	LoginInfo struct {
		Username string `json:"username"`
		Expire   string `json:"expire"`
	}
	RequestResult struct {
		Token   string `json:"token"`
		Message string `json:"message"`
	}
	RequestResultFile struct {
		Token   string      `json:"token"`
		Message string      `json:"message"`
		Payload interface{} `json:"payload"`
	}
	RequestResultState struct {
		Token   string `json:"token"`
		Message State  `json:"message"`
	}
	PointsRequestResult struct {
		Message string  `json:"message"`
		Points  []Point `json:"points"`
	}
	ResCode struct {
		ResCode string
	}
	Delete1 struct {
		Id     string `json:"id"`
		Object string `json:"object"`
		User   string `json:"user"`
		Time   string `json:"time"`
	}
	//---------------------------------------------------------------------afshin
	NewPoint1 struct {
		Userid       string   `json:"userid" bson:"userid" valid:"required"`
		Code         string   `json:"code" bson:"code" valid:"required,numeric"`
		Item1Value   string   `json:"item_1_value" bson:"item_1_value" valid:"required,numeric"`
		Item2Value   string   `json:"item_2_value" bson:"item_2_value" valid:"required,numeric"`
		Item3Value   string   `json:"item_3_value" bson:"item_3_value" valid:"required,numeric"`
		Item3Current string   `json:"item_3_current" bson:"item_3_current" valid:"required,numeric"`
		Reason       string   `json:"reason" bson:"reason" valid:"required"`
		Step         string   `json:"step" bson:"step" valid:"required"`
		Foundation   string   `json:"foundation" bson:"foundation" valid:"required"`
		Village      string   `json:"village" bson:"village"`
		State        string   `json:"state" bson:"state"`
		City         string   `json:"city" bson:"city"`
		Households   string   `json:"households" bson:"households" valid:"required"`
		Lat          string   `json:"lat" bson:"lat" valid:"required"`
		Long         string   `json:"long" bson:"long" valid:"required"`
		Images       []string `json:"images" bson:"images"`
		Zone         string   `json:"zone" bson:"zone" valid:"required,alphanum,length(24|24)"`
		Active       bool     `json:"active"`
	}
	//--------------------------------------------------------------
	Users1 []User
	User1  struct {
		ID         bson.ObjectId `bson:"_id`
		Name       string        `json:"name"`
		Family     string        `json:"family"`
		Username   string        `json:"username"`
		Password   string        `json:"password"`
		Email      string        `json:"email"`
		Phone      string        `json:"phone"`
		Mobile     string        `json:"mobile"`
		Roles      []Roles       `json:"roles"`
		Active     bool          `json:"active"`
		SuperAdmin bool          `json:"superadmin"`
		Admin      bool          `json:"admin"`
	}
	NewUser1 struct {
		Name       string `json:"name" bson:"name" valid:"required"`
		Family     string `json:"family" bson:"family" valid:"required"`
		Username   string `json:"username" bson:"username" valid:"required"`
		Password   string `json:"password" bson:"password" valid:"required"`
		Email      string `json:"email" bson:"email" valid:"email,required"`
		Phone      string `json:"phone" bson:"phone" valid:"-"`
		Mobile     string `json:"mobile" bson:"mobile" valid:"required"`
		Active     bool   `json:"active" bson:"active"`
		SuperAdmin bool   `json:"superadmin" bson:"superadmin"`
		Admin      bool   `json:"admin" bson:"admin"`
	}
	Roles1 struct {
		Editor bool   `json:"editor" bson:"editor"`
		Insert bool   `json:"insert" bson:"insert"`
		Delete bool   `json:"delete" bson:"delete"`
		Code   string `json:"code" bson:"code"`
	}
	NewRoles1 struct {
		UserID string  `json:"userid"`
		Roles  []Roles `json:"roles" bson:"roles"`
	}
	////////////////

	States []State

	State struct {
		ID          bson.ObjectId `json:"_id" bson:"_id"`
		Name        string        `json:"name" bson:"name"`
		DisplayName string        `json:"display_name" bson:"display_name"`
		Lat         string        `json:"lat" bson:"lat"`
		Long        string        `json:"long" bson:"long"`
		Cites       []City        `json:"cites" bson:"cites"`
		Trash       bool          `json:"trash" bson:"trash"`
	}
	NewState struct {
		ID          string `json:"id" `
		Name        string `json:"name" bson:"name" valid:"required"`
		DisplayName string `json:"display_name" bson:"display_name" valid:"required"`
		Lat         string `json:"lat" bson:"lat" valid:"required"`
		Long        string `json:"long" bson:"long" valid:"required"`
	}
	City struct {
		CityID      bson.ObjectId `json:"city_id" bson:"city_id"`
		Name        string        `json:"name" bson:"name"`
		DisplayName string        `json:"display_name" bson:"display_name"`
		Lat         string        `json:"lat" bson:"lat"`
		Long        string        `json:"long" bson:"long"`
		Villages    []Village     `json:"villages" bson:"villages"`
		Trash       bool          `json:"trash" bson:"trash"`
	}
	NewCity struct {
		UID         string        `json:"id"`
		ID          bson.ObjectId `json:"_id" bson:"_id"`
		Name        string        `json:"name" bson:"name" valid:"required"`
		DisplayName string        `json:"display_name" bson:"display_name" valid:"required"`
		Lat         string        `json:"lat" bson:"lat" valid:"required"`
		Long        string        `json:"long" bson:"long" valid:"required"`
		State       string        `json:"state"`
	}

	Village struct {
		VillageID   bson.ObjectId `json:"village_id" bson:"village_id"`
		Name        string        `json:"name" bson:"name"`
		DisplayName string        `json:"display_name" bson:"display_name"`
		Lat         string        `json:"lat" bson:"lat"`
		Long        string        `json:"long" bson:"long"`
		Trash       bool          `json:"trash" bson:"trash"`
	}
	NewVillage struct {
		ID          string `json:"id"`
		Name        string `json:"name" bson:"name"`
		DisplayName string `json:"display_name" bson:"display_name"`
		Lat         string `json:"lat" bson:"lat"`
		Long        string `json:"long" bson:"long"`
		City        string `json:"city"`
		State       string `json:"state"`
	}
	PointsRequest struct {
		Zone string `json:"zone" bson:"zone"`
	}

	Point struct {
		ID           bson.ObjectId `json:"_id" bson:"_id"`
		Code         string        `json:"code" bson:"code"`
		Item1Value   string        `json:"item_1_value" bson:"item_1_value"`
		Item2Value   string        `json:"item_2_value" bson:"item_2_value"`
		Item3Value   string        `json:"item_3_value" bson:"item_3_value"`
		Item3Current string        `json:"item_3_current" bson:"item_3_current"`
		Reason       string        `json:"reason" bson:"reason"`
		Step         string        `json:"step" bson:"step"`
		Foundation   string        `json:"foundation" bson:"foundation"`
		Village      string        `json:"village" bson:"village"`
		State        string        `json:"state" bson:"state"`
		City         string        `json:"city" bson:"city"`
		Households   string        `json:"households" bson:"households"`
		Lat          string        `json:"lat" bson:"lat"`
		Long         string        `json:"long" bson:"long"`
		Images       []string      `json:"images" bson:"images"`
		Zone         bson.ObjectId `json:"zone" bson:"zone"`
		Active       bool          `json:"active" bson:"active"`
	}

	UpdatePointStep struct {
		Code  string `json:"code" bson:"code" valid:"required,numeric"`
		Step  string `json:"step" bson:"step" valid:"required"`
		Year  string `json:"year" bson:"year"`
		Month string `json:"month" bson:"month"`
		Day   string `json:"day" bson:"day"`
		Time  string `json:"time" bson:"time"`
	}
	CreatePaymentRequest1 struct {
		Code   string `json:"code"`
		Amount string `json:"amount"`
		Date   string `json:"date"`
		Time   string `json:"time"`
	}
	CreatePaymentResponse struct {
		Code   string `json:"code"`
		Amount string `json:"amount"`
		Date   string `json:"date"`
		Time   string `json:"time"`
		Info   string `json:"info"`
	}
	Payment struct {
		ID      bson.ObjectId
		Amount  string        `json:"amount" bson:"amount"`
		Code    string        `json:"code" bson:"code"`
		Date    string        `json:"date" bson:"date"`
		Time    string        `json:"time" bson:"time"`
		Status  string        `json:"status" bson:"status"`
		Info    string        `json:"info" bson:"info"`
		PointID bson.ObjectId `json:"point_id" bson:"point_id"`
		OrderID string        `json:"order_id" bson:"order_id"`
		PayerID string        `json:"payer_id" bson:"payer_id"`
	}
	PaymentStatus1 struct {
		OrderID string `json:"order_id"`
		Status  string `json:"status"`
		Info    string `json:"info"`
	}
	PaymentStatusResponse struct {
		Result string `json:"result"`
	}
	LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	Delete struct {
		Id     string `json:"id"`
		Object string `json:"object"`
		User   string `json:"user"`
		Time   string `json:"time"`
	}
	UploadRequest struct {
		Code    string      `json:"code" valid:"required,alphanum,length(3|32)" form:"code"`
		Payload interface{} `json:"payload"  bson:"payload"`
	}
	//-------------------------------------------------------------------------afshin
	NewPoint struct {
		Userid       string   `json:"userid" bson:"userid"`
		Code         string   `json:"code" bson:"code" valid:"required,numeric"`
		Item1Value   string   `json:"item_1_value" bson:"item_1_value" valid:"required,numeric"`
		Item2Value   string   `json:"item_2_value" bson:"item_2_value" valid:"required,numeric"`
		Item3Value   string   `json:"item_3_value" bson:"item_3_value" valid:"required,numeric"`
		Item3Current string   `json:"item_3_current" bson:"item_3_current" valid:"required,numeric"`
		Reason       string   `json:"reason" bson:"reason" valid:"required"`
		Step         string   `json:"step" bson:"step" valid:"required"`
		Foundation   string   `json:"foundation" bson:"foundation" valid:"required"`
		Village      string   `json:"village" bson:"village"`
		State        string   `json:"state" bson:"state"`
		City         string   `json:"city" bson:"city"`
		Households   string   `json:"households" bson:"households" valid:"required"`
		Lat          string   `json:"lat" bson:"lat" valid:"required"`
		Long         string   `json:"long" bson:"long" valid:"required"`
		Images       []string `json:"images" bson:"images"`
		Zone         string   `json:"zone" bson:"zone" valid:"required,alphanum,length(24|24)"`
		Active       bool     `json:"active" bson:"active"`
	}
	//------------------------------------------------------------------

	Users2      []UserRespons
	UserRespons struct {
		Name             string        `json:"name" bson:"name"`
		Family           string        `json:"family" bson:"family"`
		Username         string        `json:"username" bson:"username"`
		Password         string        `json:"password" bson:"password"`
		Email            string        `json:"email" bson:"email"`
		Phone            string        `json:"phone" bson:"phone"`
		Mobile           string        `json:"mobile" bson:"mobile"`
		Roles            []Roles       `json:"roles" bson:"roles"`
		Active           bool          `json:"active" bson:"active"`
		Superadmin       bool          `json:"superadmin" bson:"superadmin"`
		Admin            bool          `json:"admin" bson:"admin"`
		StateID          bson.ObjectId `json:"state_id" bson:"state_id"`
		CityID           bson.ObjectId `json:"city_id" bson:"city_id"`
		Trash            bool          `json:"trash" bson:"trash"`
		StateName        string        `json:"state_name" bson:"state_name" `
		StateDisplayName string        `json:"state_display_name" bson:"state_display_name"`
		CityName         string        `json:"city_name" bson:"city_name"`
		CityDisplayName  string        `json:"city_display_name" bson:"city_display_name"`
	}
	Users []User

	User struct {
		Name       string        `json:"name" bson:"name"`
		Family     string        `json:"family" bson:"family"`
		Username   string        `json:"username" bson:"username"`
		Password   string        `json:"password" bson:"password"`
		Email      string        `json:"email" bson:"email"`
		Phone      string        `json:"phone" bson:"phone"`
		Mobile     string        `json:"mobile" bson:"mobile"`
		Roles      []Roles       `json:"roles" bson:"roles"`
		Active     bool          `json:"active" bson:"active"`
		Superadmin bool          `json:"superadmin" bson:"superadmin"`
		Admin      bool          `json:"admin" bson:"admin"`
		StateID    bson.ObjectId `json:"state_id" bson:"state_id"`
		CityID     bson.ObjectId `json:"city_id" bson:"city_id"`
		Trash      bool          `json:"trash" bson:"trash"`
	}
	NewUser struct {
		Name       string        `json:"name" bson:"name" valid:"required"`
		Family     string        `json:"family" bson:"family" valid:"required"`
		Username   string        `json:"username" bson:"username" valid:"required"`
		Password   string        `json:"password" bson:"password" valid:"required"`
		Email      string        `json:"email" bson:"email" valid:"email,required"`
		Phone      string        `json:"phone" bson:"phone" valid:"-"`
		Mobile     string        `json:"mobile" bson:"mobile" valid:"-"`
		Active     bool          `json:"active" bson:"active"`
		Superadmin bool          `json:"superadmin" bson:"superadmin"`
		Admin      bool          `json:"admin" bson:"admin"`
		StateID    bson.ObjectId `json:"state_id" bson:"state_id"`
		CityID     bson.ObjectId `json:"city_id" bson:"city_id"`
		Trash      bool          `json:"trash" bson:"trash"`
	}
	NewUserUpdate struct {
		Baseusername string        `json:"baseusername"bson:"baseusername" valid:"required"`
		Name         string        `json:"name" bson:"name" valid:"required"`
		Username     string        `json:"username" bson:"username" valid:"required"`
		Family       string        `json:"family" bson:"family" valid:"required"`
		Email        string        `json:"email" bson:"email" valid:"email,required"`
		StateID      bson.ObjectId `json:"state_id" bson:"state_id" valid:"required"`
		CityID       bson.ObjectId `json:"city_id" bson:"city_id" valid:"required"`
	}
	NewPasswordUpdate struct {
		Username    string `json:"username" bson:"username" valid:"required"`
		NewPassword string `json:"newpassword" bson:"newpassword" valid:"required"`
	}
	DeleteRequest struct {
		Username string `json:"username" bson:"username" valid:"required"`
	}
	GetFileRequest struct {
		Code string `json:"code" bson:"code" valid:"required"`
	}
	DeleteFileRequest struct {
		Code string `json:"code" bson:"code" valid:"required"`
		Path string `json:"path" bson:"path" valid:"required"`
	}
	GetFileResponse struct {
		Path []string `json:"path"`
	}
	Roles struct {
		Editor bool   `json:"editor" bson:"editor"`
		Insert bool   `json:"insert" bson:"insert"`
		Delete bool   `json:"delete" bson:"delete"`
		Code   string `json:"code" bson:"code"`
	}
	NewRoles struct {
		UserID string  `json:"userid"`
		Roles  []Roles `json:"roles" bson:"roles"`
	}
	RegisterRequest struct {
		Name     string `json:"name" bson:"name"`
		Family   string `json:"family" bson:"family"`
		Username string `json:"username" bson:"username" valid:"required"`
		Password string `json:"password" bson:"password" valid:"required"`
		Email    string `json:"email" bson:"email" valid:"required"`
		Phone    string `json:"phone" bson:"phone"`
		Mobile   string `json:"mobile" bson:"mobile"`
		Active   bool   `json:"active" bson:"active"`
		Admin    bool   `json:"admin" bson:"admin"`
		Trash    bool   `json:"trash" bson:"trash"`
	}
	UserStruct struct {
		Name     string `json:"name" bson:"name"`
		Family   string `json:"family" bson:"family"`
		Username string `json:"username" bson:"username"`
		Password string `json:"password" bson:"password"`
		Email    string `json:"email" bson:"email"`
		Phone    string `json:"phone" bson:"phone"`
		Mobile   string `json:"mobile" bson:"mobile"`
		Active   bool   `json:"active" bson:"active"`
		Admin    bool   `json:"admin" bson:"admin"`
		Trash    bool   `json:"trash" bson:"trash"`
	}
	NurseStruct struct {
		Name     string `json:"name" bson:"name"`
		Family   string `json:"family" bson:"family"`
		Username string `json:"username" bson:"username"`
		Password string `json:"password" bson:"password"`
		Email    string `json:"email" bson:"email"`
		Phone    string `json:"phone" bson:"phone"`
		Mobile   string `json:"mobile" bson:"mobile"`
		Active   bool   `json:"active" bson:"active"`
		Admin    bool   `json:"admin" bson:"admin"`
		Trash    bool   `json:"trash" bson:"trash"`
	}
	VerifyRequest struct {
		VerifyID   string `json:"verify_id" bson:"verify_id"`
		VerifyCode string `json:"verify_code" bson:"verify_code"`
		Email      string `json:"email" bson:"email"`
		Token      string `json:"token" bson:"token"`
	}
	RegisterResponse struct {
		VerifyID string `json:"verify_id" bson:"verify_id"`
	}
	VerifyResponse struct {
		Token string `json:"token" bson:"token"`
	}

	Response struct {
		Result  string      `json:"result"`
		Message string      `json:"message"`
		Payload interface{} `json:"payload"`
	}
	LoginReq struct {
		Username string `json:"username" bson:"username"`
		Password string `json:"password" bson:"password"`
	}
	LoginResponse struct {
		Token string `json:"token" bson:"token"`
	}
	ReVerifyRequest struct {
		VerifyID string `json:"verify_id" bson:"verify_id"`
	}
	SendRequest struct {
		Price      int64  `json:"price" bson:"price"`
		Obligation int64  `json:"obligation" bson:"obligation"`
		Size       string `json:"cylinder" bson:"size"`
		Percent    string `json:"discount" bson:"percent"`
		HasDamage  string `json:"hasDamage" bson:"has_damage"`
		DamageType string `json:"damageType" bson:"damage_type"`
	}
	SendRequestBa struct {
		Obligation int64  `json:"obligation" bson:"obligation"`
		Model      string `json:"model" bson:"model"`
		Percent    string `json:"percent" bson:"percent"`
		A          int64  `json:"a" bson:"a"`
		B          int64  `json:"b" bson:"b"`
		C          int64  `json:"c" bson:"c"`
		D          int64  `json:"d" bson:"d"`
		F          int64  `json:"f" bson:"f"`
	}
	ActiveResponse struct {
		Active bool `json:"active" bson:"active"`
	}
	ResponsePrice struct {
		Data1 Data  `json:"data_1"`
		Data2 Data  `json:"data_2"`
		Data3 Data  `json:"data_3"`
		Data4 Data  `json:"data_4"`
		Data5 Data  `json:"data_5"`
		Data6 Data  `json:"data_6"`
		Data7 Data  `json:"data_7"`
		Sum   int64 `json:"sum"`
	}
	Data struct {
		Company string `json:"company"`
		Price   int64  `json:"price"`
		Base    int64  `json:"base"`
	}
	SendFile struct {
		FileID string `json:"file_id" bson:"file_id"`
		Token  string `json:"token" bson:"token"`
	}

	Request struct {
		Src string `json:"src"`
		Group string `json:"group"`
		Key string `json:"key"`
		Payload interface{} `json:"payload"`
	}
	CityResponse struct{
		City string `json:"city"`
	}
	CityResponseA struct{
		Payload interface{} `json:"payload"`
	}
)
