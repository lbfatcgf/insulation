package jwt_util

import (
	"crypto"
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 自定义声明结构
type PayloadClaims struct {
	Payload string `json:"payload"`
	jwt.RegisteredClaims
}
type AdminJwt struct {
	priKey crypto.PrivateKey
	pubKey crypto.PublicKey
	secret []byte
	option *JwtOptions
}

func NewAdminJwt() *AdminJwt {
	return &AdminJwt{}
}

func (j *AdminJwt) SetSecret(option JwtOptions) error {
	switch option.SecretType {
	case "rsa":
		pri, err := loadRsaPrivateKey(option.PriKeyPath)
		if err != nil {
			panic(err)
		}
		pub, err := loadRsaPublicKey(option.PubKeyPath)
		if err != nil {
			panic(err)
		}
		j.priKey = pri
		j.pubKey = pub

	case "ecdsa":
		pri, err := loadEcdsaPrivateKey(option.PriKeyPath)
		if err != nil {
			panic(err)
		}
		pub, err := loadEcdsaPublicKey(option.PubKeyPath)
		if err != nil {
			panic(err)
		}
		j.priKey = pri
		j.pubKey = pub

	default:
		j.secret = []byte(option.Secret)
	}
	j.option = &option
	return nil
}

func (j *AdminJwt) GenerateToken(payload string) (string, error) {
	claims := PayloadClaims{
		Payload: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.option.ExpireDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "insulation admin",
		},
	}

	// 根据SecretType选择算法
	var token *jwt.Token
	switch j.option.SecretType {
	case "rsa":
		token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
		return token.SignedString(j.priKey)
	case "ecdsa":
		token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
		return token.SignedString(j.priKey)
	default: // 默认使用HMAC对称加密
		token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		return token.SignedString(j.secret)
	}
}

// 验证token
// bearerToken: 带Bearer的token 'Bearer xxxx'
// 返回值：payload, error: payload为json字符串
func (j *AdminJwt) ValidateToken(bearToken string) (string, error) {
	tokenString := strings.TrimPrefix(bearToken, "Bearer ")
	token, err := jwt.ParseWithClaims(tokenString, &PayloadClaims{}, func(token *jwt.Token) (interface{}, error) {
		switch j.option.SecretType {
		case "rsa":
			return j.pubKey, nil
		case "ecdsa":
			return j.pubKey, nil
		default: // 默认使用HMAC对称加密
			return j.secret, nil
		}
	})
	if err != nil {
		return ``, err
	}
	if claims, ok := token.Claims.(*PayloadClaims); ok && token.Valid {
		return claims.Payload, nil
	}
	return ``, errors.New(`jwt un Valid`)
}
