package util

import "encoding/json"

// APIResult marlon's application api result
type APIResult struct {
	Code    int         `json:"code"`
	Message interface{} `json:"msg"`
	Data    interface{} `json:"data"`
}

// NewAPIResult 返回一个response对象
func NewAPIResult() *APIResult {
	return &APIResult{
		Code:    0,
		Message: ""}
}

// Success 成功对象
func (a *APIResult) Success(data interface{}) *APIResult {
	a.Data = data
	return a
}

// Failed 失败对象
func (a *APIResult) Failed(code int, message interface{}, data interface{}) *APIResult {
	a.Code = code
	a.Message = message
	a.Data = data
	return a
}

// ToJSON return a json byte
func (a *APIResult) ToJSON() (result []byte, err error) {
	return json.Marshal(a)
}
