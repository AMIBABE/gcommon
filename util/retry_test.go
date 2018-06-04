package util

import (
	"errors"
	"testing"
	"time"
)

func TestRetry(t *testing.T) {
	{
		var i = 3
		Retry(3, time.Second, func() error {
			i--
			return errors.New("")
		})

		if i != 0 {
			t.Error("retry方法没有进行每次重试,在失败情况下")
		}
	}

	{
		var i = 3
		Retry(3, time.Second, func() error {
			i--
			return nil
		})

		if i != 2 {
			t.Error("retry方法没有停止,在执行成功的情况下")
		}
	}
}
