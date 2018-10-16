package routers

import (
	"github.com/kataras/iris"
)

func RoutersRestFullInit(app *iris.Application) {
	UserRouter(app)
	UserGroupRouter(app)
	CodeRouter(app)
	NewsTypeRouter(app)
	NewsReviewRouter(app)
	NewsRouter(app)
	VpsRouter(app)
	VpnRouter(app)
	VpnRelationRouter(app)
	FileRouter(app)
}