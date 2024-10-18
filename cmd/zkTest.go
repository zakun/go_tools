/*
 * @Author: qizk qizk@mail.open.com.cn
 * @Date: 2024-09-18 16:11:12
 * @LastEditors: qizk qizk@mail.open.com.cn
 * @LastEditTime: 2024-10-17 10:13:34
 * @FilePath: \zk_tools\cmd\zkTest.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"sync"
	"zk/tools/logger"

	"github.com/spf13/cobra"
)

var isTest bool
var name string

var echo string

// zkTestCmd represents the zkTest command
var zkTestCmd = &cobra.Command{
	Use:   "zk-test [--name yourname | -t]",
	Short: "test cmd for zk",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("zkTest called")

		test()

		logger.Info("final echo: %v", echo)
		return nil
	},
}

func init() {

	zkTestCmd.Flags().StringVar(&name, "name", "zk", "name from cmd")
	zkTestCmd.Flags().BoolVarP(&isTest, "test", "t", false, "test from cmd")

	rootCmd.AddCommand(zkTestCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// zkTestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// zkTestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func test() {

	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}
