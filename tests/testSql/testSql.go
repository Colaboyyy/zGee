package main

import (
	"fmt"
	"gee"
	"net/http"
)

//func init() {
//	sqlStr := "root:tommyytest1234@tcp(127.0.0.1:3306)/market?charset=utf8&parseTime=true&loc=Local"
//	var err error
//	// 1.打开db
//	SqlDb, err := sql.Open("mysql", sqlStr)
//	if err != nil {
//		fmt.Println("Open Db failed!", err)
//		return
//	}
//	fmt.Println("Open Db Success")
//	// 2.测试db是否连接成功
//	err = SqlDb.Ping()
//	if err != nil {
//		fmt.Println("Connect Db failed!", err)
//		return
//	}
//	fmt.Println("Connect Db Success")
//}

func main() {
	r := gee.New()
	r.Use(gee.Logger())
	r.GET("/", func(ctx *gee.Context) {
		ctx.String(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/sql", gee.FindOneById)
	r.GET("/findchar", gee.FindPronounceByChar)
	r.GET("/findmuldata", gee.FindMulData)
	r.GET("/findmulpro", gee.FindMulPronounceByChar)
	err := r.RUN(":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
}
