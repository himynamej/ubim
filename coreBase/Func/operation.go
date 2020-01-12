package functions

import (
	DataTypes "coreBase/Type"
)

// Operation implements an image transformation runnable interface
type OperationFunc func(req *DataTypes.Request) DataTypes.Response

type OperationsMap map[string]OperationFunc

var Operations = map[string]OperationsMap{
	"register": register,
	"user":     user,
	"file":     file,
}

var register = map[string]OperationFunc{
	"registerUser": registerUser,
	"verifiction":  verification,
	"login":        login,
	"test":         test,
}
var user = map[string]OperationFunc{
	"checkuser":  CheckUser,
	"addrequest": AddRequest,
}
var file = map[string]OperationFunc{
	"addfile": AddFile,
}

func (o OperationFunc) Run(req *DataTypes.Request) DataTypes.Response {
	return o(req)
}
func registerUser(req *DataTypes.Request) DataTypes.Response {

	/*
		r:=DataTypes.RegisterRequest{}
		fmt.Println("payload:",req.Payload)
		err := json.Unmarshal(req.Payload,&r)
	*/
	//err = mapstructure.Decode(req.Payload, &r)
	res, err := ProcessRegister(req.Payload)
	if err != nil {
		return DataTypes.Response{
			Message: err.Error(),
			Result:  "error",
		}
	}

	return DataTypes.Response{
		Message: "success",
		Result:  "ok",
		Payload: res,
	}

}
func verification(req *DataTypes.Request) DataTypes.Response {

	/*
		r:=DataTypes.RegisterRequest{}
		fmt.Println("payload:",req.Payload)
		err := json.Unmarshal(req.Payload,&r)
	*/
	//err = mapstructure.Decode(req.Payload, &r)
	res, err := ProcessVerify(req.Payload)
	if err != nil {
		return DataTypes.Response{
			Message: err.Error(),
			Result:  "error",
		}
	}

	return DataTypes.Response{
		Message: "success",
		Result:  "ok",
		Payload: res,
	}

}
func login(req *DataTypes.Request) DataTypes.Response {

	/*
		r:=DataTypes.RegisterRequest{}
		fmt.Println("payload:",req.Payload)
		err := json.Unmarshal(req.Payload,&r)
	*/
	//err = mapstructure.Decode(req.Payload, &r)
	res, err := ProcessLogin(req.Payload)
	if err != nil {
		return DataTypes.Response{
			Message: err.Error(),
			Result:  "error",
		}
	}

	return DataTypes.Response{
		Message: "success",
		Result:  "ok",
		Payload: res,
	}

}
func CheckUser(req *DataTypes.Request) DataTypes.Response {

	/*
		r:=DataTypes.RegisterRequest{}
		fmt.Println("payload:",req.Payload)
		err := json.Unmarshal(req.Payload,&r)
	*/
	//err = mapstructure.Decode(req.Payload, &r)
	err := ProcessCheckUser(req.Payload)
	if err != nil {
		return DataTypes.Response{
			Message: err.Error(),
			Result:  "error",
		}
	}
	return DataTypes.Response{
		Message: "success",
		Result:  "ok",
		Payload: nil,
	}

}
func AddRequest(req *DataTypes.Request) DataTypes.Response {

	/*
		r:=DataTypes.RegisterRequest{}
		fmt.Println("payload:",req.Payload)
		err := json.Unmarshal(req.Payload,&r)
	*/
	//err = mapstructure.Decode(req.Payload, &r)
	err := ProcessAddRequest(req.Payload)
	if err != nil {
		return DataTypes.Response{
			Message: err.Error(),
			Result:  "error",
		}
	}

	return DataTypes.Response{
		Message: "success",
		Result:  "ok",
		Payload: nil,
	}

}
func AddFile(req *DataTypes.Request) DataTypes.Response {

	/*
		r:=DataTypes.RegisterRequest{}
		fmt.Println("payload:",req.Payload)
		err := json.Unmarshal(req.Payload,&r)
	*/
	//err = mapstructure.Decode(req.Payload, &r)
	err := ProcessAddFile(req.Payload)
	if err != nil {
		return DataTypes.Response{
			Message: err.Error(),
			Result:  "error",
		}
	}

	return DataTypes.Response{
		Message: "success",
		Result:  "ok",
		Payload: nil,
	}

}
func test(req *DataTypes.Request) DataTypes.Response {

	/*
		r:=DataTypes.RegisterRequest{}
		fmt.Println("payload:",req.Payload)
		err := json.Unmarshal(req.Payload,&r)
	*/
	//err = mapstructure.Decode(req.Payload, &r)
	//err := ProcessAddFile(req.Payload)
	//if err != nil {
	//	return DataTypes.Response{
	//		Message: err.Error(),
	//		Result:  "error",
	//	}
	//}
	var car DataTypes.CarRequest

	car.CarName = append(car.CarName, DataTypes.CarN{
		Car:"SAMAND",
	},DataTypes.CarN{

		Car:"PEYCAN",
	},DataTypes.CarN{

		Car:"RENO",
	},DataTypes.CarN{

		Car:"PQ",
	},DataTypes.CarN{

		Car:"MASOOD",
	})

	return DataTypes.Response{
		Message: "success",
		Result:  "ok",
		Payload: car.CarName,
	}

}
