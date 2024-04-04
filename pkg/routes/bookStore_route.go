/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-04 10:11:54
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-04 12:32:27
 * @FilePath: /goBookStore/pkg/routes/bookStore_route.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package routes

import (
	"github.com/gorilla/mux"
	"github.com/jinziguan123/goBookStore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
