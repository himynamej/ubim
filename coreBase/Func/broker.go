package functions

import DataTypes "coreBase/Type"

func RequestHandler(req DataTypes.Request) DataTypes.Response {
	var res DataTypes.Response
	if fn, ok := Operations[req.Group][req.Key]; ok {
		res = fn.Run(&req)
	} else {
		res.Result = "failed"
		res.Message =  "operation request is not valid"
	}

	return res
}
