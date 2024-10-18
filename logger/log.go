/*
 * @Author: qizk qizk@mail.open.com.cn
 * @Date: 2024-05-21 17:22:03
 * @LastEditors: qizk qizk@mail.open.com.cn
 * @LastEditTime: 2024-09-23 12:10:39
 * @FilePath: \helloworld\logger\log.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AElog
 */
package logger

import (
	"log"
	"os"
	"time"
)

var applogger *log.Logger
var apploggerfile *os.File

func init() {
	if applogger == nil {
		applogger = log.New(os.Stderr, "", log.LstdFlags)
	}

	fileName := "runtime/" + time.Now().Format("2006_01_02") + "_app.log"
	apploggerfile, _ = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
}

func Info(msg string, v ...any) {
	applogger.SetPrefix("[Info] ")
	applogger.Printf(msg, v...)
}

func Error(msg string, v ...any) {
	applogger.SetPrefix("[Error] ")

	applogger.Printf(msg, v...)
}

func InfoAndShow(msg string, v ...any) {
	Info(msg, v...)

	applogger.SetOutput(apploggerfile)
	Info(msg, v...)

	applogger.SetOutput(os.Stderr)
}

func ErrorAndShow(msg string, v ...any) {
	Error(msg, v...)

	applogger.SetOutput(apploggerfile)
	Error(msg, v...)

	applogger.SetOutput(os.Stderr)
}
