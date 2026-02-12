package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub         int `json:"sub"` // user id
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}

func CreateJwt(secret string, data Payload) (string, error) {
	// 1. Create and encode Header
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}
	headerBytes, _ := json.Marshal(header)
	fmt.Println("headerBytes:", headerBytes)
	headerB64 := base64UrlEncode(headerBytes)
	fmt.Println("headerB64:", headerB64 ,"---")

	// 2. Encode Payload
	byteArrData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	payloadB64 := base64UrlEncode(byteArrData)

	// 3. Create Message (Header.Payload)
	message := headerB64 + "." + payloadB64

	// 4. Generate Signature using HMAC-SHA256
	byteArrSecret := []byte(secret)
	byteArrMessage := []byte(message)

	h := hmac.New(sha256.New, byteArrSecret)
	h.Write(byteArrMessage)

	signature := h.Sum(nil)
	signatureB64 := base64UrlEncode(signature)

	// 5. Final JWT string
	jwt := headerB64 + "." + payloadB64 + "." + signatureB64

	return jwt, nil
}

