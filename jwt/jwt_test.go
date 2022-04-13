package jwt

import (
	"fmt"
	"testing"
)

func TestZJWT_CreateJwt(t *testing.T) {
	zjwt := NewZJWT("1123456789")
	m := make(map[string]interface{})
	m["id"] = 1
	m["name"] = "zx"
	jwt, err := zjwt.CreateJwt(20, m)
	if err != nil {
		panic(err)
	}
	fmt.Println(jwt)
	parseJWT, err := zjwt.ParseJWT(jwt)
	if err != nil {
		panic(err)
	}
	fmt.Println(parseJWT["name"])
}
