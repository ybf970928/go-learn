package constants

const (
	SUCCESS = 1
	ERROR   = 0

	ErrorAuthCheckTokenFail = 20001
)

var MsgFlags = map[int]string{
	SUCCESS: "ok",
	ERROR:   "fail",

	ErrorAuthCheckTokenFail: "token null",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]

	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
