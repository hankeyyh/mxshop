package zerrors

type Code int32

func (c Code) Value() int32 {
	return int32(c)
}

func (c Code) String() string {
	msg, ok := codeMsg[c]
	if !ok {
		return "未知错误"
	}
	return msg
}

const (
	// 通用[0, 1000)
	UnknownError       Code = -1 // 兼容旧客户端错误码
	OK                 Code = 0
	InvalidParam       Code = 1
	IncompleteParam    Code = 2
	DBConnectErr       Code = 3
	DBExecuteErr       Code = 4
	DownStreamSvrError Code = 5
	RecordNotFound     Code = 6
	StringConvError    Code = 7
	StringSplitError   Code = 8
	JsonMarshalError   Code = 9
	JsonUnmarshalError Code = 10
	RedisConnectErr    Code = 11
	RedisExecuteErr    Code = 12
)

var codeMsg = map[Code]string{
	// 通用[0, 1000)
	OK:                 "",
	InvalidParam:       "参数不合法",
	IncompleteParam:    "参数不完整",
	DBConnectErr:       "数据库连接异常",
	DBExecuteErr:       "数据库执行异常",
	DownStreamSvrError: "下游服务异常",
	RecordNotFound:     "记录未找到",
	StringConvError:    "字符串类型转换失败",
	StringSplitError:   "字符串分割失败",
	JsonMarshalError:   "json序列化失败",
	JsonUnmarshalError: "json反序列化失败",
	RedisConnectErr:    "redis连接异常",
	RedisExecuteErr:    "redis执行异常",
}
