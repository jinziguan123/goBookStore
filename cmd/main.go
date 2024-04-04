/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-04 10:09:10
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-04 13:03:01
 * @FilePath: /goBookStore/cmd/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinziguan123/goBookStore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
