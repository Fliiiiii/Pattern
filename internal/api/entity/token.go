package entity

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"reforce.pattern/config"
	"reforce.pattern/pkg/controllers"
)

type token struct {
	aes      string
	Services []string `json:"service_ids"`
	UserID   string   `json:"user_id"`
	Expired  int64    `json:"expired"`
}

func Token(aes string) token {
	return token{aes: aes}
}

func (t *token) Decode() error {
	hexedToken, err := hex.DecodeString(t.aes)
	if err != nil {
		return err
	}
	block, err := aes.NewCipher(config.CFG.ReforceID.DecodeKey)
	if err != nil {
		return err
	}
	blockMode := cipher.NewCBCDecrypter(block, config.CFG.ReforceID.DecodeKey)
	blockMode.CryptBlocks(hexedToken, hexedToken)
	hexedToken = unPadding(hexedToken)

	err = json.Unmarshal(hexedToken, t)
	if err != nil {
		return err
	}
	return nil
}

func (t *token) Valid() bool {
	return controllers.StringsContain(t.Services, config.CFG.ReforceID.ServiceID)
}

func (t *token) ReforceID() string {
	return t.UserID
}

func unPadding(str []byte) []byte {
	n := len(str)
	count := int(str[n-1])
	newPaddingText := str[:n-count]
	return newPaddingText
}
