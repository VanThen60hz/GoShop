package routers

import (
	"github.com/VanThen60hz/GoShop/internal/routers/manage"
	"github.com/VanThen60hz/GoShop/internal/routers/user"
)

type RouterGroup struct {
	User   user.UserRouterGroup
	Manage manage.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)
