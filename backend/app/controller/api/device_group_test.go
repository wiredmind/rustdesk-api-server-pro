package api

import (
	"net/http"
	"net/http/httptest"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"
	"testing"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func TestAccessibleDeviceGroupsReturnsCurrentUsersGroups(t *testing.T) {
	engine, err := db.NewEngine(&config.DbConfig{
		Driver:   "sqlite",
		Dsn:      "file::memory:?cache=shared",
		TimeZone: "UTC",
	})
	if err != nil {
		t.Fatal(err)
	}
	defer engine.Close()
	if err := engine.Sync(new(model.DeviceGroup)); err != nil {
		t.Fatal(err)
	}
	if _, err := engine.Insert(
		&model.DeviceGroup{UserId: 7, Name: "Family"},
		&model.DeviceGroup{UserId: 8, Name: "Other"},
	); err != nil {
		t.Fatal(err)
	}

	app := iris.New()
	ctx := app.ContextPool.Acquire(
		httptest.NewRecorder(),
		httptest.NewRequest(http.MethodGet, "/api/device-group/accessible?current=1&pageSize=10", nil),
	)
	defer app.ContextPool.Release(ctx)
	ctx.Values().Set(config.CurrentUserKey, &model.User{Id: 7})

	controller := DeviceGroupController{
		basicController: basicController{Ctx: ctx, Db: engine},
	}
	response, ok := controller.HandleAccessible().(mvc.Response)
	if !ok {
		t.Fatalf("response type = %T, want mvc.Response", controller.HandleAccessible())
	}
	payload, ok := response.Object.(iris.Map)
	if !ok {
		t.Fatalf("response object type = %T, want iris.Map", response.Object)
	}
	if total := payload["total"]; total != int64(1) {
		t.Fatalf("total = %v, want 1", total)
	}
	groups, ok := payload["data"].([]iris.Map)
	if !ok || len(groups) != 1 || groups[0]["name"] != "Family" {
		t.Fatalf("groups = %#v, want one Family group", payload["data"])
	}
}
