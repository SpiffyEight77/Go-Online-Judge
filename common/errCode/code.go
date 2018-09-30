package errCode

const (
	SUCCESS      = 200
	BADREQUEST   = 400
	UNAUTHORIZED = 401
	NOT_FOUND    = 404
	ERROR        = 500

	TOKEN_MISSING = 1001
	INVALID_TOKEN = 1002
	TOKEN_TIMEOUT = 1003

	REQUEST_AUTH_CENTER_FAILED      = 2001
	REQUEST_AUTH_CENTER_PHONE_ERROR = 2001001

	REQUEST_SMS_VCODE_FAILED         = 3001
	REQUEST_SMS_VCODE_TYPE_ERROR     = 3001001
	REQUEST_SMS_VCODE_PHONE_NOTFOUND = 3001002

	REQUEST_PASSWD_RESET_FAILED = 4001

	REQUEST_USER_REGISTER_FAILED            = 5001
	REQUEST_USER_REGISTER_ALREADY_EXITS     = 5001001
	REQUEST_USER_REGISTER_INVIATE_NOT_EXITS = 5001002
	REQUEST_USERNAME_INVALID                = 5001003

	REQUEST_USER_LOGIN_FAILED = 6001

	REQUEST_PARAMETER_MISSING    = 0010001
	REQUEST_PARAMETER_TYPE_ERROR = 0010002
	REQUEST_FILE_SIZE_ERROR      = 0010003
	REQUEST_FILE_TYPE_ERROR      = 0010004

	REQUEST_EXP_FAILED          = 7001
	REQUEST_EXP_STATUS_FAILED   = 7001001
	REQUEST_EXP_GROUP_EXIST     = 7001002
	REQUEST_EXP_GROUP_NOT_EXIST = 7001003

	UNPUBLISH_FORBIDDEN = 8001001
)

// response code format: object + action + state if error happened or success if successful
// object: user / project / dataset / repo / experiment etc.
// action: login / registry / create / retrieve / update / delete / publish / un publish etc.
// state: success / fail etc.

const (
	objectUser     = "100"
	objectData     = "200"
	objectDataList = "201"
	objectProject  = "300"
	objectRepo     = "400"
	objectExp      = "500"
	objectResource = "600"
	objectParam    = "700"

	actionLogin     = "100"
	actionRegistry  = "200"
	actionCreate    = "300"
	actionRetrieve  = "400"
	actionUpdate    = "500"
	actionDelete    = "600"
	actionPublish   = "700"
	actionUnPublish = "800"

	actionCheck               = "900"
	actionExistCheck          = "901"
	actionNameCheck           = "902"
	actionVcodeCheck          = "903"
	actionCaptchaCheck        = "904"
	actionTokenCheck          = "905"
	actionPhoneCheck          = "906"
	actionInvitationCodeCheck = "907"
	actionStatusCheck         = "908"

	stateFail    = "100"
	success      = "200"
	stateError   = "300"
	stateMissing = "400"
	stateTimeout = "500"
)

const (
	Success = success

	//user
	UserLoginSuccess       = success
	UserNameCheckFailed    = objectUser + actionNameCheck + stateFail
	UserVcodeCheckFailed   = objectUser + actionVcodeCheck + stateFail
	UserCaptchaCheckFailed = objectUser + actionCaptchaCheck + stateFail

	UserTokenCheckError      = objectUser + actionTokenCheck + stateError
	UserTokenCheckMissing    = objectUser + actionTokenCheck + stateMissing
	UserTokenCheckTimeout    = objectUser + actionTokenCheck + stateTimeout
	UserPhoneCheckMissing    = objectUser + actionPhoneCheck + stateMissing
	UserPhoneCheckError      = objectUser + actionPhoneCheck + stateError
	UserInvitationCodeFailed = objectUser + actionNameCheck + stateFail

	//project
	ProjectCheckExistFailed = objectProject + actionExistCheck + stateMissing
	ProjectUnPublishFailed  = objectProject + actionUnPublish + stateFail

	//exp
	ExpCheckExistFailed  = objectExp + actionExistCheck + stateMissing
	ExpCheckStatusFailed = objectExp + actionStatusCheck + stateFail

	//data
	DataCreateSuccess       = success
	DataListRetrieveSuccess = success
	DataListRetrieveFailed  = objectDataList + actionRetrieve + stateFail
	DataCreateFail          = objectData + actionCreate + stateFail
	DataDeleteSuccess       = success

	DataExpCheckExistFailed = objectData + actionExistCheck + stateMissing
	// etc.

	//resource
	ResourceExpCheckExistFailed = objectResource + actionExistCheck + stateMissing

	//param
	RequestParameterError   = objectParam + actionCheck + stateError
	RequestParameterMissing = objectParam + actionCheck + stateMissing
)
