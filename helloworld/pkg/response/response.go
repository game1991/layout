package response

// Response ajax request response json struct
type Response struct {
	Result  bool        `json:"result"`         //标记请求是否成功
	Code    int         `json:"code"`           //错误code编码，code从1000开始，1000以内http status使用
	Message string      `json:"msg,omitempty"`  //状态信息提示
	Data    interface{} `json:"data,omitempty"` //请求成功后返回的数据信息
	Err     interface{} `json:"err,omitempty"`  //请求失败后返回的数据信息
}
