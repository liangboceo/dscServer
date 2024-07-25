package dto

type Result struct {
	Status  bool        `json:"status" doc:"状态"`
	Code    int         `json:"code" doc:"错误码"`
	Message string      `json:"message" doc:"错误信息"`
	Data    interface{} `json:"data" doc:"返回体"`
}

const NO_LOGIN = -1
const SERVER_ERROR = 500
const SUCCESS = 200

func Success(data interface{}) Result {
	result := Result{}
	result.Status = true
	result.Code = SUCCESS
	result.Message = ""
	result.Data = data
	return result
}
func SuccessMessage(data interface{}, message string) Result {
	result := Result{}
	result.Status = true
	result.Code = SUCCESS
	result.Message = message
	result.Data = data
	return result
}

func Failure(data interface{}) Result {
	result := Result{}
	result.Status = false
	result.Code = SERVER_ERROR
	result.Message = ""
	result.Data = data
	return result
}

func FailureMessage(data interface{}, message string) Result {
	result := Result{}
	result.Status = false
	result.Code = SERVER_ERROR
	result.Message = message
	result.Data = data
	return result
}
