package util

import (
	"reflect"
	"testing"
)

func TestNewApiResult(t *testing.T) {
	response := NewAPIResult()
	if reflect.TypeOf(response).String() == "*APIResult" {
		t.Error(reflect.TypeOf(response))
	}

	if response.Code != 0 {
		t.Error("调用NewAPIResult方法得到的 APIResult.code 默认值不等于0, 与预期结果不符")
	}
}

func TestAPIResultSuccess(t *testing.T) {
	data := 1
	response := NewAPIResult().Success(data)
	if response.Data != data {
		t.Error("成功情况下,data与期望结果不符")
	}
}

func TestAPIResultFailed(t *testing.T) {
	code, msg, data := 1, "error", "data"

	response := NewAPIResult().Failed(code, msg, data)
	if response.Data != data {
		t.Error("失败情况下,data与期望结果不符")
	}

	if response.Code != code {
		t.Error("失败情况下,code与期望结果不符")
	}

	if response.Message != msg {
		t.Error("失败情况下,msg与期望结果不符")
	}
}

func TestAPIResultToJson(t *testing.T) {
	responseByte, err := NewAPIResult().Success(1).ToJSON()
	if err == nil && reflect.TypeOf(responseByte).String() != "[]uint8" {
		t.Error("APIResult序列化失败")
	}
}
