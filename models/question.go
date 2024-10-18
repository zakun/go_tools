/*
 * @Author: qizk qizk@mail.open.com.cn
 * @Date: 2024-08-30 15:28:08
 * @LastEditors: qizk qizk@mail.open.com.cn
 * @LastEditTime: 2024-09-02 15:07:53
 * @FilePath: \helloworld\common\question.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package models

type Question struct {
	Id           uint
	SubjectKey   string
	QuestionType int    `gorm:"column:type"`
	Title        string `gorm:"column:question"`
	Content      string
	Option       string
	OptionNum    int
	Options      string
	Answer       string
	Result       string
	Analysis     string
	Analyze      string
	Status       int
	Uid          int
	QType        int
	Addtime      int64
	ComeFrom     string `gorm:"column:from"`
}
