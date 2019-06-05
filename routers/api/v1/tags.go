package v1

// 定义tags表的相关操作的控制器方法

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/imanner/gin-xinshulaila/models"
	"github.com/imanner/gin-xinshulaila/pkg/e"
	"github.com/imanner/gin-xinshulaila/pkg/setting"
	"github.com/imanner/gin-xinshulaila/pkg/util"
	"net/http"
)

// 获取多个文章标签
// curl 127.0.0.1:8000/api/v1/tags
func GetTags(c *gin.Context) {
	// 获取参数 name
	name := c.Query("name")

	// 定义变量
	maps := make(map[string]interface{})	// 接受参数的条件
	data := make(map[string]interface{}) 	// 返回真实的数据

	// 若参数name不为空
	if name != "" {
		maps["name"] = name
	}

	// 声明状态变量
	var state int = -1
	// c.Query可用于获取?name=test&state=1这类URL参数，而c.DefaultQuery则支持设置一个默认值
	// 获取参数 state
	if arg := c.Query("state"); arg != "" {
		// arg变为int赋给state
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	// 获取code  这里是不是一上来就应该获取code？？？？
	code := e.SUCCESS
	msg := e.GetMsg(code)

	// 调用模型中的方法获取列表和总数
	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	// 以json的形式返回
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : msg,
		"data" : data,
	})
}


//新增文章标签
// postman: post: http://127.0.0.1:8000/api/v1/tags?name=1&state=1&created_by=test
func AddTag(c *gin.Context) {
	// name 参数
	name := c.Query("name")
	// 划重点
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")

	// beego 的验证部分
	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")		// 必要参数
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")  //字符串长度
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")	// 限定值

	// 初始化返回code
	code := e.INVALID_PARAMS
	// 如果验证没有问题
	if ! valid.HasErrors() {
		// 如果不存在标签
		if ! models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			// 如果已经存在改标签
			code = e.ERROR_EXIST_TAG
		}
	}
	msg := e.GetMsg(code)

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : msg,
		"data" : make(map[string]string),  // 这样保证了返回数据的统一
	})
}

//修改文章标签
// PUT访问http://127.0.0.1:8000/api/v1/tags/1?name=edit1&state=0&modified_by=edit1，查看code是否返回200
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}

			models.EditTag(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]string),
	})
}

//删除文章标签
// delete : http://127.0.0.1:8000/api/v1/tags/1，查看code是否返回200
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			models.DeleteTag(id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]string),
	})
}
