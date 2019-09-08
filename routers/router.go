package routers

import (
	"github.com/gaogep/EchoPlay/routers/apis/v1/category"
	"github.com/gaogep/EchoPlay/routers/apis/v1/post"
	"github.com/gaogep/EchoPlay/routers/apis/v1/user"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//e.Pre(middleware.AddTrailingSlash())

	// 注册路由组
	apiv1 := e.Group("/api/v1")

	// 注册分类路由
	category.RegisterCategoryHandler(apiv1)
	post.RegisterPostHandler(apiv1)
	user.RegisterUserHandler(apiv1)

	return e
}
