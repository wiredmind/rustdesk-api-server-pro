package admin

import (
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"xorm.io/xorm"
)

type DevicesController struct {
	basicController
}

func (c *DevicesController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/devices/list", "HandleList")
	b.Handle("POST", "/devices/connect", "HandleConnect")
}

var deviceSortableColumns = map[string]bool{
	"":            true,
	"id":          true,
	"hostname":    true,
	"rustdesk_id": true,
	"username":    true,
	"version":     true,
	"os":          true,
	"memory":      true,
	"last_user":   true,
	"last_online": true,
	"is_online":   true,
	"created_at":  true,
}

func (c *DevicesController) HandleList() mvc.Result {
	currentPage := c.Ctx.URLParamIntDefault("current", 1)
	pageSize := c.Ctx.URLParamIntDefault("size", 15)
	hostname := c.Ctx.URLParamDefault("hostname", "")
	username := c.Ctx.URLParamDefault("username", "")
	rustdesk_id := c.Ctx.URLParamDefault("rustdesk_id", "")
	last_user := c.Ctx.URLParamDefault("last_user", "")
	version := c.Ctx.URLParamDefault("version", "")
	os := c.Ctx.URLParamDefault("os", "")
	online := c.Ctx.URLParamDefault("online", "")
	keyword := c.Ctx.URLParamDefault("keyword", "")
	sortBy := c.Ctx.URLParamDefault("sort_by", "")
	sortOrder := c.Ctx.URLParamDefault("sort_order", "")

	if !deviceSortableColumns[sortBy] {
		sortBy = ""
	}
	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = ""
	}
	if pageSize <= 0 || pageSize > 200 {
		pageSize = 15
	}

	query := func() *xorm.Session {
		q := c.Db.Table(&model.Device{})

		if hostname != "" {
			q.Where("hostname LIKE ?", "%"+hostname+"%")
		}
		if username != "" {
			q.Where("username LIKE ?", "%"+username+"%")
		}
		if rustdesk_id != "" {
			q.Where("rustdesk_id LIKE ?", "%"+rustdesk_id+"%")
		}
		if last_user != "" {
			q.Where("last_user LIKE ?", "%"+last_user+"%")
		}
		if version != "" {
			q.Where("version = ?", version)
		}
		if os != "" {
			q.Where("os LIKE ?", "%"+os+"%")
		}
		switch online {
		case "online":
			q.Where("is_online = ?", 1)
		case "offline":
			q.Where("is_online = ?", 0)
		}
		if keyword != "" {
			like := "%" + keyword + "%"
			q.Where(
				"hostname LIKE ? OR rustdesk_id LIKE ? OR username LIKE ? OR last_user LIKE ? OR os LIKE ?",
				like, like, like, like, like,
			)
		}

		order := sortBy
		if order == "" {
			order = "is_online DESC, last_online DESC, id DESC"
		} else {
			switch sortOrder {
			case "asc":
				order = sortBy + " ASC"
			default:
				order = sortBy + " DESC"
			}
		}
		q.OrderBy(order)
		return q
	}

	pagination := db.NewPagination(currentPage, pageSize)
	deviceList := make([]model.Device, 0)
	err := pagination.Paginate(query, &model.Device{}, &deviceList)
	if err != nil {
		return c.Error(nil, err.Error())
	}

	list := make([]iris.Map, 0, len(deviceList))
	for _, d := range deviceList {
		list = append(list, iris.Map{
			"id":          d.Id,
			"rustdesk_id": d.RustdeskId,
			"hostname":    d.Hostname,
			"username":    d.Username,
			"uuid":        d.Uuid,
			"version":     d.Version,
			"os":          d.Os,
			"memory":      d.Memory,
			"cpu":         d.Cpu,
			"last_user":   d.LastUser,
			"last_online": formatTime(d.LastOnline),
			"is_online":   d.IsOnline,
			"conns":       d.Conns,
			"created_at":  d.CreatedAt.Format(config.TimeFormat),
		})
	}
	return c.Success(iris.Map{
		"total":   pagination.TotalCount,
		"records": list,
		"current": currentPage,
		"size":    pageSize,
	}, "ok")
}

type connectRequest struct {
	RustdeskId string `json:"rustdesk_id"`
	Action     string `json:"action"`
}

// deviceActionAuthority maps a portal action to the RustDesk client's own
// rustdesk://<authority>/<id> uni-link scheme. These authorities are read
// directly from the RustDesk client source (core_main.rs): --connect,
// --play, --file-transfer map to "connect", "play", "file-transfer"; the
// built-in terminal session uses the "terminal" authority (confirmed by the
// RustDesk maintainers). There is no supported way to inject or execute an
// arbitrary shell command through this link - the terminal authority only
// opens RustDesk's own terminal UI, gated by the remote device's
// "enable-terminal" permission.
var deviceActionAuthority = map[string]string{
	"remote":   "connect",
	"file":     "file-transfer",
	"mirror":   "play",
	"terminal": "terminal",
}

func (c *DevicesController) HandleConnect() mvc.Result {
	var req connectRequest
	if err := c.Ctx.ReadJSON(&req); err != nil {
		return c.Error(nil, err.Error())
	}
	if req.RustdeskId == "" {
		return c.Error(nil, "RustdeskIdEmpty")
	}
	authority, ok := deviceActionAuthority[req.Action]
	if !ok {
		authority = "connect"
		req.Action = "remote"
	}
	return c.Success(iris.Map{
		"rustdesk_id": req.RustdeskId,
		"action":      req.Action,
		"client_url":  "rustdesk://" + authority + "/" + req.RustdeskId,
		"scheme":      "rustdesk",
	}, "ok")
}

func formatTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(config.TimeFormat)
}
