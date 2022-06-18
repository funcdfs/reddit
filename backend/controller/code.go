package controller

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParameter
	CodeUserExists
	CodeUserNotFound
	CodeInvalidPassword
	CodeServerBusy
)

var codeMsgMap = map[ResCode]any{
	CodeSuccess:          "success",
	CodeInvalidParameter: "invalid_parameter",
	CodeUserExists:       "user_exists",
	CodeUserNotFound:     "user_not_found",
	CodeInvalidPassword:  "invalid_password",
	CodeServerBusy:       "server_busy",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg.(string)
}
