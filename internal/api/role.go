package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/pkg/curd"
	"github.com/zgwit/iot-master/v4/types"
)

// @Summary 查询角色数量
// @Schemes
// @Description 查询角色数量
// @Tags role
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int64] 返回角色数量
// @Router /role/count [post]
func noopRoleCount() {}

// @Summary 查询角色
// @Schemes
// @Description 这里写描述 get roles
// @Tags role
// @Param search body curd.ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[types.Role] 返回角色信息
// @Router /role/search [post]
func noopRoleSearch() {}

// @Summary 查询角色
// @Schemes
// @Description 查询角色
// @Tags role
// @Param search query curd.ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyList[types.Role] 返回角色信息
// @Router /role/list [get]
func noopRoleList() {}

// @Summary 创建角色
// @Schemes
// @Description 创建角色
// @Tags role
// @Param search body types.Role true "角色信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.Role] 返回角色信息
// @Router /role/create [post]
func noopRoleCreate() {}

// @Summary 修改角色
// @Schemes
// @Description 修改角色
// @Tags role
// @Param id path string true "角色ID"
// @Param role body types.Role true "角色信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.Role] 返回角色信息
// @Router /role/{id} [post]
func noopRoleUpdate() {}

// @Summary 获取角色
// @Schemes
// @Description 获取角色
// @Tags role
// @Param id path string true "角色ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.Role] 返回角色信息
// @Router /role/{id} [get]
func noopRoleGet() {}

// @Summary 删除角色
// @Schemes
// @Description 删除角色
// @Tags role
// @Param id path string true "角色ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[types.Role] 返回角色信息
// @Router /role/{id}/delete [get]
func noopRoleDelete() {}

// @Summary 导出角色
// @Schemes
// @Description 导出角色
// @Tags role
// @Accept json
// @Produce octet-stream
// @Success 200 {object} curd.ReplyList[types.Role] 返回压缩包
// @Router /role/export [get]
func noopRoleExport() {}

// @Summary 导入角色
// @Schemes
// @Description 导入角色
// @Tags role
// @Param file formData file true "压缩包"
// @Accept mpfd
// @Produce json
// @Success 200 {object} curd.ReplyData[int64] 返回角色数量
// @Router /role/import [post]
func noopRoleImport() {}

func roleRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[types.Role]())

	app.POST("/search", curd.ApiSearch[types.Role]())

	app.GET("/list", curd.ApiList[types.Role]())

	app.POST("/create", curd.ParseParamStringId, curd.ApiCreateHook[types.Role](
		curd.GenerateRandomId[types.Role](9), nil))

	app.GET("/:id", curd.ParseParamStringId, curd.ApiGet[types.Role]())

	app.POST("/:id", curd.ParseParamStringId, curd.ApiUpdateHook[types.Role](nil, nil,
		"id", "name", "privileges"))

	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDeleteHook[types.Role](nil, nil))

	app.GET("/export", curd.ApiExport("role", "角色"))

	app.POST("/import", curd.ApiImport("role"))

}
