package category

import (
	"github.com/astaxie/beego/validation"
	"github.com/gaogep/EchoPlay/models"
	"github.com/gaogep/EchoPlay/settings"
	"github.com/labstack/echo"
	"github.com/unknwon/com"
	"net/http"
	"strconv"
)

// 获取分类
func GetCategory(c echo.Context) error {
	code := http.StatusOK
	maps := make(map[string]interface{})
	resp := make(map[string]interface{})

	totalItems := models.GetTagTotal(maps)
	page := com.StrTo(c.QueryParam("page")).MustInt()
	pageSize := settings.GlobalConf["PAGESIZE"].(int)
	pageOffset := (page - 1) * pageSize
	totalPage := (totalItems / pageSize) + 1

	if page <= 0 || page > totalPage {
		page = 1
		code = http.StatusBadRequest
		resp["message"] = "页码错误"
		return c.JSON(code, &resp)
	}

	nextPage := strconv.Itoa(page + 1)
	prevPage := "1"
	if page > 1 {
		prevPage = strconv.Itoa(page - 1)
	}

	resp["list"] = models.GetCategoryList(pageOffset, pageSize, maps)
	resp["cur_page"] = page
	resp["total_page"] = totalPage
	resp["prev_page"] = c.Request().Host + c.Echo().Reverse("GetCategorys") + "?page=" + prevPage
	resp["next_page"] = c.Request().Host + c.Echo().Reverse("GetCategorys") + "?page=" + nextPage

	return c.JSON(code, &resp)
}

// 新增分类
func AddCategory(c echo.Context) error {
	code := http.StatusOK
	name := c.FormValue("name")
	resp := make(map[string]interface{})

	valid := validation.Validation{}
	valid.Required(name, "name").Message("类别名称不能为空")
	valid.MaxSize(name, 20, "name").Message("类别名不能超过20个字符")

	if !valid.HasErrors() {
		if models.ExistedCategoryByName(name) {
			resp["message"] = "分类已存在"
			code = http.StatusForbidden
		} else {
			models.CreateCategory(name)
			resp["message"] = "分类创建成功"
		}
	} else {
		for _, err := range valid.Errors {
			resp[err.Key] = err.Message
		}
		code = http.StatusBadRequest
	}

	return c.JSON(code, &resp)
}

// 更新分类
func UpdateCategory(c echo.Context) error {
	code := http.StatusOK
	name := c.FormValue("name")
	id := com.StrTo(c.Param("id")).MustInt()
	resp := make(map[string]interface{})

	valid := validation.Validation{}
	valid.Required(name, "name").Message("类别名称不能为空")
	valid.MaxSize(name, 20, "name").Message("类别名不能超过20个字符")

	if !valid.HasErrors() {
		if models.ExistedCategoryById(id) {
			data := make(map[string]interface{})
			data["name"] = name
			models.UpdateCategory(id, data)
			resp["message"] = "更新成功"
		} else {
			resp["message"] = "分类不存在"
			code = http.StatusForbidden
		}
	} else {
		for _, err := range valid.Errors {
			resp[err.Key] = err.Message
		}
		code = http.StatusBadRequest
	}

	return c.JSON(code, &resp)
}

// 删除分类
func DeleteCategory(c echo.Context) error {
	code := http.StatusOK
	id := com.StrTo(c.Param("id")).MustInt()
	resp := make(map[string]interface{})

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if !valid.HasErrors() {
		if models.ExistedCategoryById(id) {
			models.DeleteCategory(id)
			resp["message"] = "删除成功"
		} else {
			resp["message"] = "分类不存在"
			code = http.StatusForbidden
		}
	} else {
		for _, err := range valid.Errors {
			resp[err.Key] = err.Message
		}
		code = http.StatusBadRequest
	}

	return c.JSON(code, &resp)
}

func RegisterCategoryHandler(g *echo.Group) {
	g.GET("/categorys", GetCategory).Name = "GetCategorys"
	g.POST("/categorys", AddCategory).Name = "AddCategorys"
	g.PUT("/categorys/:id", UpdateCategory).Name = "UpdateCategory"
	g.DELETE("/categorys/:id", DeleteCategory).Name = "DeleteCategory"
}
