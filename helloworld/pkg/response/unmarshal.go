package response

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//Unmarshal 对请求的resp进行解析操作
//resp: http请求返回的response对象
//value: 根据参数个数判断，第一个用于接收请求成功的数据，第二个用于接收失败的数据
func Unmarshal(resp *http.Response, value ...interface{}) *InnerError {
	if resp == nil {
		return &InnerError{Status: http.StatusBadRequest, Msg: "response参数信息为空"}
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 { //http status不为200，表示请求失败
		ierr := &InnerError{Status: resp.StatusCode}
		bts, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			ierr.Msg = err.Error()
		} else {
			ierr.Msg = string(bts)
		}
		return ierr
	}
	response := &Response{}
	length := len(value)
	if length > 0 {
		response.Data = value[0]
	}
	if length > 1 {
		response.Err = value[1]
	}
	err := json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return &InnerError{Status: http.StatusInternalServerError, Msg: err.Error()}
	}
	if !response.Result {
		return &InnerError{Status: response.Code, Msg: response.Message}
	}
	return nil
}
