package response

//Error default err struct
type Error interface {
	Error() string
	Code() int
}

//InnerError 内部定义错误类型
type InnerError struct {
	Status int
	Msg    string
}

//Code 返回错误code码
func (i *InnerError) Code() int {
	return i.Status
}

//Error 返回错误信息
func (i *InnerError) Error() string {
	return i.Msg
}
