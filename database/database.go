/*
 * @Author: qizk qizk@mail.open.com.cn
 * @Date: 2024-05-23 11:29:00
 * @LastEditors: qizk qizk@mail.open.com.cn
 * @LastEditTime: 2024-09-18 11:23:47
 * @FilePath: \helloworld\database\database.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package database

import (
	"fmt"
	"zk/tools/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Db() *gorm.DB {
	db, err := connection("db_default")
	if err != nil {
		panic(err)
	}

	return db
}

func DbName(name string) *gorm.DB {
	db, err := connection(name)
	if err != nil {
		panic(err)
	}

	return db
}

func connection(name string) (*gorm.DB, error) {
	if db != nil {
		return db, nil
	}

	port := config.Get("Port", name).String()
	host := config.Get("Host", name).String()
	username := config.Get("Username", name).String()
	password := config.Get("Password", name).String()
	database := config.Get("Database", name).String()

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", username, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		// Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}
