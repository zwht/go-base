package routersBlog

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func RoutersBlogInit(app *iris.Application) {
	mvc.New(app.Party("/")).Handle(new(BlogIndexController))
	mvc.New(app.Party("/detail/")).Handle(new(BlogDetailController))
}
