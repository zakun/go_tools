/*
 * @Author: qizk qizk@mail.open.com.cn
 * @Date: 2024-09-18 10:52:13
 * @LastEditors: qizk qizk@mail.open.com.cn
 * @LastEditTime: 2024-09-18 13:27:50
 * @FilePath: \zk_tools\cmd\impQuestion.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"zk/tools/import_question"

	"github.com/spf13/cobra"
)

// impQuestionCmd represents the impQuestion command
var impQuestionCmd = &cobra.Command{
	Use:   "impQuestion",
	Short: "职导狮导题",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		imp := import_question.NewCMD_ZDS_QSP("ZDS_IMP_QUESTION")
		imp.Run()
	},
}

func init() {
	rootCmd.AddCommand(impQuestionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// impQuestionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// impQuestionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
