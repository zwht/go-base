package routersBlog

import (
	. "../../services"
	"errors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"strconv"
)

type BlogDetailController struct{}

// 根据月份和urlen去查询新闻详情
func (c *BlogDetailController) GetBy(ctx iris.Context, year string, month string, urlEn string) mvc.Result {
	month1, _ := strconv.Atoi(month)
	year1, _ := strconv.Atoi(year)
	if month1 == 12 {
		month1 = 1
		year1 = year1 + 1
	} else {
		month1 = month1 + 1
	}
	detail, err := NewsSelectByUrl(urlEn, year+"-"+month+"-01", strconv.Itoa(year1)+"-"+strconv.Itoa(month1)+"-01")
	if err != nil {
		return mvc.Response{Err: errors.New("未找到数据"), Code: 400}
	}
	if detail.State != 1105 {
		return mvc.Response{Err: errors.New("未找到数据!"), Code: 400}
	}
	var news_reviewSearchVo NewsReviewSearchVo
	news_reviewSearchVo.NewId = detail.ID
	reviewList, reviewErr := NewsReviewSelectList(10000, 1, news_reviewSearchVo)
	if reviewErr != nil {
	}

	data := map[string]interface{}{
		"Detail":  detail,
		"Reviews": reviewList,
	}
	var helloView = mvc.View{
		Name: "blog/pages/detail.html",
		Data: data,
	}
	return helloView
}
