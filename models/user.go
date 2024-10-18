/*
 * @Author: qizk qizk@mail.open.com.cn
 * @Date: 2024-05-23 14:17:02
 * @LastEditors: qizk qizk@mail.open.com.cn
 * @LastEditTime: 2024-05-27 17:09:16
 * @FilePath: \helloworld\models\user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package models

import "gorm.io/gorm"

type User struct {
	Name  string
	Email string

	Profile *Profile
	*gorm.Model
}

func (User) TableName() string {
	return "user"
}
