package DataTypes

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	Post struct {
		ID           bson.ObjectId `bson:"_id"`
		UID          string        `json:"uid" bson:"uid" mapstructure:"uid"`
		UserID       string        `json:"user_id" bson:"user_id" ms:"user_id"`
		Files        []string      `json:"files" bson:"files" mapstructure:"files"`
		Content      string        `json:"content" bson:"content" mapstructure:"content"`
		AllowComment bool          `json:"allow_comment" bson:"allow_comment" mapstructure:"allow_comment"`
		AllowShare   bool          `json:"allow_share" bson:"allow_share" mapstructure:"allow_share"`
		AllowLike    bool          `json:"allow_like" bson:"allow_like" ms:"allow_like"`
		View         string        `json:"view" bson:"view" mapstructure:"view"`
		PostType     string        `json:"post_type" bson:"post_type" mapstructure:"post_type"`
		Client       string        `json:"client" bson:"client" mapstructure:"client"`
		Time         time.Time     `json:"time" bson:"time" mapstructure:"time"`
		Tags         []string      `json:"tags" bson:"tags" mapstructure:"tags"`
		PostID       string        `json:"post_id" bson:"post_id" mapstructure:"post_id"`
		Like         int64         `json:"like" bson:"like" mapstructure:"like"`
		Likes        []Liked       `json:"likes" bson:"likes" mapstructure:"likes"`
		Share        int64         `json:"share" bson:"share" mapstructure:"share"`
		Shares       []Shared      `json:"shares" bson:"shares" mapstructure:"shares"`
		Comment      int64         `json:"comment" bson:"comment" mapstructure:"comment"`
		Comments     []PostComment `json:"comments" bson:"comments" mapstructure:"comments"`
		ShopID       string        `json:"shop_id" bson:"shop_id" mapstructure:"shop_id"`
		Cost         int           `json:"cost" bson:"cost"`
	}
	PostComment struct {
		UserID  string     `json:"user_id" bson:"user_id" mapstructure:"user_id"`
		UID     string     `json:"uid" bson:"uid" mapstructure:"uid"`
		Content string     `json:"content" bson:"content" mapstructure:"content"`
		Tags    []string   `json:"tags" bson:"tags" mapstructure:"tags"`
		Like    int64      `json:"like" bson:"like" mapstructure:"like"`
		Likes   []Liked    `json:"likes" bson:"likes" mapstructure:"likes"`
		For     PostReplay `json:"for" bson:"for" mapstructure:"for"`
	}
	PostReplay struct {
		UserID string `json:"user_id" bson:"user_id" mapstructure:"user_id"`
		UID    string `json:"uid" bson:"uid" mapstructure:"uid"`
	}
	Liked struct {
		UserID string `json:"user_id" bson:"user_id" mapstructure:"user_id"`
		UID    string `json:"uid" bson:"uid" mapstructure:"uid"`
	}
	Shared struct {
		UserID string     `json:"user_id" bson:"user_id" mapstructure:"user_id"`
		UID    string     `json:"uid" bson:"uid" mapstructure:"uid"`
		For    PostReplay `json:"for" bson:"for" mapstructure:"for"`
	}
	Shop struct {
		ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Name        string        `json:"name" bson:"name"`
		Category    string        `json:"category" bson:"category" mapstructure:"category"`
		Admins      []Admin       `json:"admins" bson:"admins"`
		SuperUser   string        `json:"super_user" bson:"super_user"`
		Following   []string      `json:"following" bson:"following" mapstructure:"following"`
		Description string        `json:"description" bson:"description" mapstructure:"description"`
		ShoutOuts   []Post        `json:"shout_outs" bson:"shout_outs" mapstructure:"shout_outs"`
		OnSales     []Post        `json:"on_sale" bson:"on_sale" mapstructure:"on_sale"`
		ShopID      string        `json:"shop_id" bson:"shop_id" mapstructure:"shop_id"`
		PageID      string        `json:"page_id" bson:"page_id" mapstructure:"page_id"`
		Enable      bool          `json:"enable" bson:"enable"`
	}
	Category struct {
		Name string `json:"name" bson:"name"`
	}
	Admin struct {
		UID string `json:"uid" bson:"uid"`
	}
	// project implemented nats message type
	// base structure for holding requests
	Request struct {
		Group string
		Key   string
		ID    string
		// request payload for holding request json data
		Payload []byte `json:"payload"`
		UID     string
		Token   string
	}
	// login request payload
	UserLoginRequest struct {
		Password string `json:"password"`
		UserID   string `json:"user_id" mapstructure:"user_id"`
		Client   string `json:"client"`
		Extra    string `json:"extra"`
	}
	// base structure for holding register request as request payload
	RegisterRequest struct {
		Password  string `json:"password"`
		Country   string `json:"country"`
		Phone     string `json:"phone"`
		Mobile    string `json:"mobile"`
		Email     string `json:"email"`
		FirstName string `json:"first_name" mapstructure:"first_name"`
		SureName  string `json:"sure_name" mapstructure:"sure_name"`
		Username  string `json:"username"`
		Client    string `json:"client"`
		Extra     string `json:"extra"`
	}
	User struct {
		Password  string `json:"password"`
		Country   string `json:"country"`
		Phone     string `json:"phone"`
		Mobile    string `json:"mobile"`
		Email     string `json:"email"`
		FirstName string `json:"first_name" mapstructure:"first_name"`
		SureName  string `json:"sure_name" mapstructure:"sure_name"`
		Username  string `json:"username"`
		Client    string `json:"client"`
		Extra     string `json:"extra"`
		Verify    bool   `json:"verify"`
	}

	ClientInfo struct {
		CID    string        `json:"cid" bson:"cid,omitempty"`
		ID     bson.ObjectId `json:"id" bson:"_id,omitempty"`
		OS     string        `json:"os" bson:"os"`
		App    string        `json:"app" bson:"app"`
		Date   string        `json:"date"`
		Token  string        `json:"token" bson:"token,omitempty"`
		UserID string        `json:"user_id" bson:"user_id"  mapstructure:"user_id"`
		Extra  string        `json:"extra" bson:"extra"`
	}

	SendPasswordRequest struct {
		UserID string `json:"user_id" mapstructure:"user_id"`
		Client string `json:"client"`
		Extra  string `json:"extra"`
	}
	CheckExistingRequest struct {
		Check  string `json:"check"`
		For    string `json:"for"`
		Value  string `json:"value"`
		Client string `json:"client"`
		Extra  string `json:"extra"`
	}
	VerifyContactRequest struct {
		Verify  string `json:"verify" bson:"verify"`
		Code    string `json:"code" bson:"code"`
		Contact string `json:"contact" bson:"contact"`
		Client  string `json:"client" bson:"client"`
		Extra   string `json:"extra" bson:"extra"`
		Token   string `json:"token" bson:"token"`
	}
	VerificationRequest struct {
		Verify string `json:"verify" bson:"verify"`
		Code   string `json:"code" bson:"code"`
	}
	VerifyContactResponse struct {
		Verify string `json:"verify" bson:"verify_id"`
	}
	LoginResponse struct {
		Token string `json:"token" bson:"token"`
	}
	ResendVerifyCodesRequest struct {
		UserID string `json:"user_id" mapstructure:"user_id"`
		Client string `json:"client"`
		Extra  string `json:"extra"`
	}
	LoginResult struct {
		TwoFactor      string `json:"two_factor" mapstructure:"two_factor"`
		Result         string
		Token          string `json:"token"`
		VerificationID string `json:"verification_id"`
		UID            string `json:"uid"`
		Shop           []Shop `json:"shop"`
	}
	Response struct {
		Result  string      `json:"result"`
		Message string      `json:"message"`
		Payload interface{} `json:"payload"`
	}
	SendVerify struct {
		Type   string `json:"type"`
		Value  string `json:"value"`
		UserID string `json:"user_id" mapstructure:"user_id"`
		Code   string `json:"code"`
	}
	GetFilesRequest struct {
		Tag    string `json:"tag" mapstructure:"tag"`
		UserID string `json:"user_id" mapstructure:"user_id"`
		UID    string `json:"uid" mapstructure:"uid"`
	}
	GetFilePathRequest struct {
		FileID string `json:"file_id" mapstructure:"file_id"`
		Tag    string `json:"tag" mapstructure:"tag"`
	}
	VerificationResponse struct {
		Token string `json:"token"`
	}
	LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	CheckUserRequest struct {
		Token string `json:"token"`
	}
	RequestPayload struct {
		FromUser  string    `json:"from_user" bson:"from_user"`
		ToUser    string    `json:"to_user" bson:"to_user"`
		Time      time.Time `json:"time" bson:"time"`
		MessageID string    `json:"message_id" bson:"message_id"`
		Server    bool      `json:"server" bson:"server"`
		Deliver   bool      `json:"deliver" bson:"deliver"`
		Seen      bool      `json:"seen" bson:"seen"`
		//Payload MessagePayload `json:"payload" bson:"payload"`
		Edit    bool      `json:"edit" bson:"edit"`
		Forward string    `json:"forward" bson:"forward"`
		Expire  time.Time `json:"expire" bson:"expire"`
	}
	RequestFile struct {
		Name        string `json:"name"`
		FileID      string `json:"file_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Suf         string `json:"suf"`
	}
	CarN struct {
		Car string `json:"car"`
	}
	CarRequest struct {
		CarName []CarN `json:"carname"`
	}
)
