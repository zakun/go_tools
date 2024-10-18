/*
 * @Author: qizk qizk@mail.open.com.cn
 * @Date: 2024-05-23 14:17:02
 * @LastEditors: qizk qizk@mail.open.com.cn
 * @LastEditTime: 2024-05-27 17:09:40
 * @FilePath: \helloworld\models\user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package models

import "gorm.io/gorm"

type Profile struct {
	UserId         uint
	IdentifyNumber string
	User           *User
	*gorm.Model
}

func (Profile) TableName() string {
	return "profile"
}
