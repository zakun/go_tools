/*
 * @Author: qizk qizk@mail.open.com.cn
 * @Date: 2024-10-15 14:57:33
 * @LastEditors: qizk qizk@mail.open.com.cn
 * @LastEditTime: 2024-10-15 15:00:07
 * @FilePath: \zk_tools\cmd\xquestion.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
/*
 * @Author: qizk qizk@mail.open.com.cn
 * @Date: 2024-10-15 14:57:33
 * @LastEditors: qizk qizk@mail.open.com.cn
 * @LastEditTime: 2024-10-15 14:58:15
 * @FilePath: \zk_tools\cmd\xquestion.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"zk/tools/xsearch"

	"github.com/spf13/cobra"
)

// xquestionCmd represents the xquestion command
var answer bool

var xquestionCmd = &cobra.Command{
	Use:   "xquestion",
	Short: "x-question",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		xsearch.Process(answer)
		// logger.Info("answer: %v", answer)
	},
}

func init() {

	xquestionCmd.Flags().BoolVarP(&answer, "answer", "t", false, "查找试题答案")

	rootCmd.AddCommand(xquestionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// xquestionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// xquestionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
