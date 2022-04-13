package jwt

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type ZJWT struct {
	PrivateKey []byte
}

func NewZJWT(key string) *ZJWT {
	return &ZJWT{PrivateKey: []byte(key)}
}

// CreateJwt 创建jwt加密数据
// exp:超市时间  m：加密的数据
func (z *ZJWT) CreateJwt(exp int, m map[string]interface{}) (string, error) {
	jwtMap := jwt.MapClaims{}
	for key, val := range m {
		jwtMap[key] = val
	}
	jwtMap["exp"] = time.Now().Add(time.Minute * time.Duration(exp)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtMap) //组装数据
	tokenString, err := token.SignedString(z.PrivateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseJWT 解析jwt加密
//token:加密后的信息
func (z *ZJWT) ParseJWT(token string) (jwt.MapClaims, error) {
	parse, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return z.PrivateKey, nil
	})
	if err != nil {
		return nil, err
	}
	return parse.Claims.(jwt.MapClaims), nil
}
