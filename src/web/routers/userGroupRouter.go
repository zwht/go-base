package routers

import (
	. "../../tools/http"
	. "../controllers"
	"github.com/kataras/iris"
)

//code model
func UserGroupRouter(app *iris.Application) {
	routes := app.Party("/v1/user/group", func(ctx iris.Context) {
		ctx.Next()
	})
	{
		routes.Post("/add", Permission(UserGroupCreate, "1001"))
		routes.Post("/update", Permission(UserGroupUpdate, "1001"))
		routes.Get("/del/{id:string}", Permission(UserGroupDeleteById, "1001"))
		routes.Get("/getById/{id:string}", Permission(UserGroupGetById, "1001"))
		routes.Post("/list/{pageNum:int}/{pageSize:int}", Permission(UserGroupGetList, "1001"))
	}

}
