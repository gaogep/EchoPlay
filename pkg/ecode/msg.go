package ecode

var MsgFlags = map[int]string{
	Success:                    "Ok",
	Error:                      "Fail",
	InvalidParams:              "Invalid params",
	ErrorExistTag:              "Existed tag",
	ErrorNotExistTag:           "Tag does not exist",
	ErrorNotExistPost:          "Post does not exist",
	ErrorAuthCheckTokenFail:    "Auth check failed",
	ErrorAuthCheckTokenTimeOut: "Token time out",
	ErrorAuthToken:             "Wrong token",
	ErrorAuth:                  "Auth error",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[Error]
}
