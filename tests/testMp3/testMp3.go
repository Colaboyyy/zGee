package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.Use(gee.Logger(), gee.DbFindProByChar())
	r.LoadHTMLGlob("./gee/templates/index.tmpl")
	r.Static("/assets", "./gee/static")
	r.GET("/listen", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", gee.H{
			"charCN":    ctx.Query("charCN"),
			"pronounce": ctx.SqlRes.Data,
			"path":      "assets/mp3/" + fmt.Sprintf("%s", ctx.SqlRes.Data) + ".mp3",
		})
	})
	r.RUN(":9999")
}
