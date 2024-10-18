/*
 * @Author: qizk qizk@mail.open.com.cn
 * @Date: 2024-08-29 16:20:06
 * @LastEditors: qizk qizk@mail.open.com.cn
 * @LastEditTime: 2024-08-30 09:25:09
 * @FilePath: \helloworld\command\common.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package import_question

type Command interface {
	GetName() string
	Run()
}
