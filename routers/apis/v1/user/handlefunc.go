package user

import (
	"github.com/astaxie/beego/validation"
	"github.com/dgrijalva/jwt-go"
	"github.com/gaogep/EchoPlay/models"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"time"
)

func JwtLogin(e echo.Context) error {
	code := http.StatusOK
	resp := make(map[string]interface{})

	nickname := e.FormValue("nickname")
	passwd := e.FormValue("passwd")

	valid := validation.Validation{}
	valid.Required(nickname, "nickname").Message("用户名不能为空")
	valid.Required(passwd, "passwd").Message("密码不能为空")

	if !valid.HasErrors() {
		if models.ValidateUserByName(nickname, passwd) {
			token := jwt.New(jwt.SigningMethodES256)
			claims := jwt.MapClaims{
				"nick_name": nickname,
				"exp_time":  time.Now().Add(time.Hour * 72).Unix(),
			}
			token.Claims = claims
			tokenStr, err := token.SigningString()
			if err != nil {
				log.Fatal("jwt signed error: ", err)
			}
			resp["token"] = tokenStr
		} else {
			code = http.StatusForbidden
			resp["message"] = "用户名或密码错误"
		}
	} else {
		for _, err := range valid.Errors {
			resp[err.Key] = err.Message
		}
		code = http.StatusBadRequest
	}

	return e.JSON(code, &resp)
}

func RegisterUserHandler(g *echo.Group) {
	g.POST("/login", JwtLogin)
}
