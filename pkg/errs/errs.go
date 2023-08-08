package errs

import "fmt"

type Error struct {
	StructName       string `json:"struct,omitempty"`
	FuncName         string `json:"function,"`
	Message          string `json:"message"`
	DeveloperMessage error  `json:"developer_message,omitempty"`
}

func (e Error) Error() string {
	var str string

	if e.StructName != "" {
		str = fmt.Sprintf("%s -", e.StructName)
	}

	str = fmt.Sprintf("%s %s : %s", str, e.FuncName, e.Message)

	if e.DeveloperMessage != nil {
		str = fmt.Sprintf("%s | DEV - %s", str, e.DeveloperMessage)
	}

	return str
}

func NewError(sn, fn, m string, de error) Error {
	return Error{
		StructName:       sn,
		FuncName:         fn,
		Message:          m,
		DeveloperMessage: de,
	}
}
