package helper

import "strings"

// *response is used for static shape json return
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

// *Empry obejct is user when data doesnt want to be null on json
type EmptyObj struct{}

func BuildResponse(status bool, message string, data interface{}) Response {

	res := Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}

	return res

}

// *BuildErrorResponse method is to inejct data value to dynamic failed response
func BuildErrorResponse(Message string, err string, data interface{}) Response {
	splitError := strings.Split(err, "\n")

	res := Response{
		Status:  false,
		Message: Message,
		Errors:  splitError,
		Data:    data,
	}

	return res

}
