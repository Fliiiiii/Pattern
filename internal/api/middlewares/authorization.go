package middlewares

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"reforce.pattern/config"
	"reforce.pattern/internal/api/entity"
	"reforce.pattern/pkg/logger"
	"reforce.pattern/pkg/mongodb"
	"time"
)

var cfg = config.CFG.ReforceID

// Authorization проверка авторизации пользователя через reforce_id
func Authorization(ctx *gin.Context, mdb *mongodb.Collections) error {
	aes, _ := ctx.Cookie("authorization")
	if aes == "" {
		logger.Warn("Пустое значение для cookie авторизации")
		return errors.New("empty authorization")
	}

	usr := new(entity.User)

	if !usr.Get(aes) {
		tkn := entity.Token(aes)
		if err := tkn.Decode(); err != nil {
			logger.Error("Не удалось получить данные из токена: %s", err.Error())
			return err
		}

		if !tkn.Valid() {
			logger.Warn("Нельзя использовать токен для проекта objectCard")
			return errors.New("wrong service")
		}

		usr.ReforceID = tkn.UserID
		usr.TokenExpire = tkn.Expired

		if !usr.Get(tkn.UserID) {
			if err := usr.Partner(mdb); err != nil {
				logger.Error("Ошибка получения партнёра: %s", err.Error())
				return err
			}
			usr.Set(usr.ReforceID)
		}
		usr.ReforceID = tkn.UserID
		usr.TokenExpire = tkn.Expired

		usr.Set(aes)
	}

	if usr.TokenExpire < time.Now().Unix() {
		code, str := refreshToken(ctx)
		if code != 200 {
			logger.Error("Ошибка обновления токена статус %d; ошибка %s", code, str)
			return errors.New(str)
		}
	}

	partner, _ := ctx.Cookie("partner")
	if partner == "" {
		logger.Warn("Пустое значение в cookie partner")
		partner = usr.PartnerID
		http.SetCookie(ctx.Writer, &http.Cookie{
			Name:     "partner",
			Value:    partner,
			Path:     "/",
			Domain:   config.CFG.Cookie.Domain,
			Expires:  time.Now().Add(time.Duration(config.CFG.Cookie.MaxAge) * time.Second),
			Secure:   false,
			HttpOnly: false,
		})
	} else {
		usr.PartnerID = partner
		usr.Set(aes)
		usr.Set(usr.ReforceID)
	}
	return nil
}

func refreshToken(ctx *gin.Context) (int, string) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/token/refresh", config.CFG.ReforceID.URL), nil)
	for _, cookie := range ctx.Request.Cookies() {
		req.AddCookie(cookie)
	}

	resp, _ := http.DefaultClient.Do(req)
	var answer struct {
		Error string `json:"error"`
		Token string `json:"token"`
	}
	body, _ := io.ReadAll(resp.Body)

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != 200 {
		_ = json.Unmarshal(body, &answer)
		return resp.StatusCode, answer.Error
	}

	for _, cookie := range resp.Cookies() {
		http.SetCookie(ctx.Writer, &http.Cookie{
			Name:     cookie.Name,
			Value:    cookie.Value,
			Path:     "/",
			Domain:   config.CFG.Cookie.Domain,
			Expires:  time.Now().Add(time.Duration(config.CFG.Cookie.MaxAge) * time.Second),
			Secure:   false,
			HttpOnly: false,
		})
	}
	return 200, ""
}
