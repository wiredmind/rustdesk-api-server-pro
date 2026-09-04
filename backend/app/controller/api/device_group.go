package api

import (
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/db"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"xorm.io/xorm"
)

type DeviceGroupController struct {
	basicController
}

func (c *DeviceGroupController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/device-group/accessible", "HandleAccessible")
}

func (c *DeviceGroupController) HandleAccessible() mvc.Result {
	current := c.Ctx.URLParamIntDefault("current", 1)
	pageSize := c.Ctx.URLParamIntDefault("pageSize", 10)
	user := c.GetUser()

	query := func() *xorm.Session {
		return c.Db.Table(&model.DeviceGroup{}).Where("user_id = ?", user.Id).Asc("name")
	}
	pagination := db.NewPagination(current, pageSize)
	groups := make([]model.DeviceGroup, 0)
	if err := pagination.Paginate(query, &model.DeviceGroup{}, &groups); err != nil {
		return mvc.Response{Object: iris.Map{"error": err.Error()}}
	}

	data := make([]iris.Map, 0, len(groups))
	for _, group := range groups {
		data = append(data, iris.Map{"name": group.Name})
	}
	return mvc.Response{Object: iris.Map{"data": data, "total": pagination.TotalCount}}
}
