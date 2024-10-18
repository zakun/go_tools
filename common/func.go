/*
 * @Author: qizk qizk@mail.open.com.cn
 * @Date: 2024-08-06 16:56:58
 * @LastEditors: qizk qizk@mail.open.com.cn
 * @LastEditTime: 2024-09-02 17:47:51
 * @FilePath: \helloworld\common\func.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
/*
 * @Author: qizk qizk@mail.open.com.cn
 * @Date: 2024-05-16 14:29:16
 * @LastEditors: qizk qizk@mail.open.com.cn
 * @LastEditTime: 2024-09-02 16:54:23
 * @FilePath: \helloworld\common\func.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package common

import (
	"fmt"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func Info(format string, v ...any) {
	if v != nil {
		fmt.Printf("[Info] "+format+"\n", v...)
	} else {
		fmt.Println(format)
	}
}

func Error(format string, v ...any) {
	if v != nil {
		fmt.Printf("[Error] "+format+"\n", v...)
	} else {
		fmt.Println(format)
	}
}

func UniqueString(s []string) []string {
	m1 := make(map[string]bool)
	result := []string{}

	for _, val := range s {
		if !m1[val] {
			m1[val] = true
			result = append(result, val)
		}
	}
	return result
}
