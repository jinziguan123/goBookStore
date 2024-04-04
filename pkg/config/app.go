/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-04 10:10:37
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-04 10:30:57
 * @FilePath: /goBookStore/pkg/config/app.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:Jinziguan123@/database2?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
