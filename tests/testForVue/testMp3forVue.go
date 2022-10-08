package main

import (
	"gee"
)

func main() {
	r := gee.New()
	r.Use(gee.Logger())
	r.Static("/assets", "./gee/static")
	r.GET("/findchar", gee.FindPronounceByChar)
	//r.GET("/listen", func(ctx *gee.Context) {
	//	gee.FindPronounceByChar(ctx)
	//	ctx.JSON(http.StatusOK, gee.H{
	//		"charCN":    ctx.Query("charCN"),
	//		"pronounce": ctx.SqlRes.Data,
	//		"path":      "assets/mp3/" + fmt.Sprintf("%s", ctx.SqlRes.Data) + ".mp3",
	//	})
	//})
	r.GET("/listen", gee.FindMp3ForPronounce)
	r.RUN(":9999")
}
