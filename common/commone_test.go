package common

import (
	"github.com/zx/box/log"
	"testing"
)

func TestCommon(t *testing.T) {
	resp := NewResp()
	resp.SetCode(1120).SetMsg(`success`).SetData(map[string]interface{}{
		"code": 200,
		"msg":  "success",
	})

	log.Logger.Debug(resp)
}
