package errCode

var CodeFlags = map[int]string{
	SUCCESS:                                 "success",
	ERROR:                                   "fail",
	UNAUTHORIZED:                            "unauthorized",
	BADREQUEST:                              "request parameter error",
	REQUEST_PARAMETER_MISSING:               "request parameter missing",
	REQUEST_SMS_VCODE_FAILED:                "send vcode failed",
	REQUEST_USER_REGISTER_FAILED:            "user register failed",
	REQUEST_USER_LOGIN_FAILED:               "user login failed",
	REQUEST_PASSWD_RESET_FAILED:             "user password reset failed",
	TOKEN_MISSING:                           "token missing",
	INVALID_TOKEN:                           "token error",
	TOKEN_TIMEOUT:                           "token timeout",
	REQUEST_AUTH_CENTER_FAILED:              "user operation failed",
	REQUEST_PARAMETER_TYPE_ERROR:            "request type error, must be 1/2/3",
	REQUEST_FILE_SIZE_ERROR:                 "upload file size bigger than 500kb",
	REQUEST_FILE_TYPE_ERROR:                 "error photo format",
	REQUEST_USER_REGISTER_INVIATE_NOT_EXITS: "invitation code not exist",
	REQUEST_USERNAME_INVALID:                "用户名包含特殊字符",
	REQUEST_AUTH_CENTER_PHONE_ERROR:         "user phone error",
	REQUEST_SMS_VCODE_TYPE_ERROR:            "request vcode type error",
	REQUEST_SMS_VCODE_PHONE_NOTFOUND:        "phone not found",
	REQUEST_EXP_STATUS_FAILED:               "exp status error",
	REQUEST_EXP_GROUP_EXIST:                 "exp group already exist",
	REQUEST_EXP_GROUP_NOT_EXIST:             "exp group NOT exist",
	//AUTH_ERROR:                              "权限错误，没有相应的权限",
}

func ErrorString(code int) string {
	msg, ok := CodeFlags[code]
	if ok {
		return msg
	}
	return CodeFlags[ERROR]
}
