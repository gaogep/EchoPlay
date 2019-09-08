package post

import (
	"github.com/astaxie/beego/validation"
	"github.com/gaogep/EchoPlay/models"
	"github.com/labstack/echo"
	"github.com/unknwon/com"
	"net/http"
)

func GetPosts(c echo.Context) error {

	return nil
}

func GetPost(c echo.Context) error {
	code := http.StatusOK
	resp := make(map[string]interface{})
	postID := com.StrTo(c.Param("id")).MustInt()

	if models.ExistedPostById(postID) {
		post := models.GetPost(postID)
		resp["title"] = post.Title
		resp["content"] = post.Content
		userInfo := make(map[string]interface{})
		userInfo["user_id"] = post.UserID
		userInfo["user_name"] = post.User.NickName
		userInfo["user_email"] = post.User.Email
		resp["user_info"] = userInfo
		resp["category_id"] = post.CategoryID
	} else {
		code = http.StatusNotFound
		resp["message"] = "文章不存在"
	}

	return c.JSON(code, &resp)
}

func AddPost(c echo.Context) error {
	resp := make(map[string]interface{})
	data := make(map[string]interface{})

	code := http.StatusOK
	title := c.FormValue("title")
	content := c.FormValue("content")
	uid := com.StrTo(c.FormValue("user_id")).MustInt()
	cid := com.StrTo(c.FormValue("category_id")).MustInt()

	valid := validation.Validation{}
	valid.Required(title, "title").Message("标题不能为空")
	valid.MaxSize(title, 50, "title").Message("标题不能超过50个字")
	valid.Required(content, "content").Message("内容不能为空")

	if !valid.HasErrors() {
		if models.ExistedCategoryById(cid) {
			valid := validation.Validation{}
			valid.Required(title, "title").Message("标题不能为空")
			valid.Required(content, "content").Message("内容不能为空")

			if !valid.HasErrors() {
				data["title"] = title
				data["content"] = content
				data["user_id"] = uid
				data["category_id"] = cid
				models.CreatePost(data)
				resp["message"] = "文章创建成功"
			} else {
				for _, err := range valid.Errors {
					resp[err.Key] = err.Message
				}
				code = http.StatusBadRequest
				resp["message"] = "文章创建失败"
			}
		} else {
			code = http.StatusBadRequest
			resp["message"] = "分类不存在"
		}
	} else {
		for _, err := range valid.Errors {
			resp[err.Key] = err.Message
		}
		code = http.StatusBadRequest
	}

	return c.JSON(code, &resp)
}

func UpdatePost(c echo.Context) error {
	resp := make(map[string]interface{})

	code := http.StatusOK
	title := c.FormValue("title")
	content := c.FormValue("content")
	pid := com.StrTo(c.FormValue("post_id")).MustInt()
	cid := com.StrTo(c.FormValue("category_id")).MustInt()

	valid := validation.Validation{}
	valid.Required(title, "title").Message("标题不能为空")
	valid.MaxSize(title, 50, "title").Message("标题不能超过50个字")
	valid.Required(content, "content").Message("内容不能为空")

	if !valid.HasErrors() {
		if models.ExistedPostById(pid) && models.ExistedCategoryById(cid) {
			data := make(map[string]interface{})
			data["title"] = title
			data["content"] = content
			data["cid"] = cid
			models.UpdatePost(pid, data)
			resp["message"] = "文章更新成功"
		} else {
			resp["message"] = "文章或者分类不存在"
		}
	} else {
		for _, err := range valid.Errors {
			resp[err.Key] = err.Message
		}
		code = http.StatusBadRequest
	}

	return c.JSON(code, &resp)
}

func DeletePost(c echo.Context) error {
	code := http.StatusOK
	resp := make(map[string]interface{})
	pid := com.StrTo(c.FormValue("post_id")).MustInt()

	if models.ExistedPostById(pid) {
		models.DeletePost(pid)
		resp["message"] = "文章删除成功"
	} else {
		resp["message"] = "文章不存在"
		code = http.StatusBadRequest
	}

	return c.JSON(code, &resp)
}

func RegisterPostHandler(g *echo.Group) {
	g.GET("/posts", GetPosts)
	g.GET("/posts/:id", GetPost)
	g.POST("/posts", AddPost)
	g.PUT("/posts/:id", UpdatePost)
	g.DELETE("/posts/:id", DeletePost)
}
