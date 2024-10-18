/*
 * @Author: qizk qizk@mail.open.com.cn
 * @Date: 2024-08-29 16:22:14
 * @LastEditors: qizk qizk@mail.open.com.cn
 * @LastEditTime: 2024-09-18 13:17:09
 * @FilePath: \helloworld\zds\question_process.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package import_question

import (
	"errors"
	"io/fs"
	"path/filepath"
	"time"
	"zk/tools/common"
	"zk/tools/logger"
	"zk/tools/models"
)

var (
	CQuestion chan models.Question
	CResult   = make(chan FileData, 1)
)

// 职导狮导题，导题前删除对应科目下所有旧题
type CMD_ZDS_QSP struct {
	Name string
}

func NewCMD_ZDS_QSP(name string) *CMD_ZDS_QSP {
	return &CMD_ZDS_QSP{
		Name: name,
	}
}

func (c CMD_ZDS_QSP) GetName() string {
	return c.Name
}

func (c CMD_ZDS_QSP) Run() {
	workdir := "./runtime/xlsx"

	starttime := time.Now()
	common.Info("===startime %v===", starttime.Format("2006-01-02 15:04:05"))
	go func(path string) {
		defer close(CResult)

		no := 0
		err := filepath.WalkDir(workdir, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			ext := filepath.Ext(path)
			if ext == ".xlsx" {
				no++

				qf := NewQuestionFile(path)
				qf.SetFileNo(no)

				info, err := qf.HandleFile()
				if err != nil {
					if errors.Is(err, ErrSubjectNotExist) {
						logger.ErrorAndShow("%v, file: %v", err, d.Name())
						return nil
					} else {
						return err
					}
				}

				if info != nil {
					CResult <- *info
				}
			}
			return nil
		})

		common.CheckError(err)
	}(workdir)

	for result := range CResult {
		common.Info("%d#, 科目：%s, 文件：%s, 试题：%d", result.XFile.FileNo, result.SubjectKey, result.XFile.Name, len(result.Data))
		// xlsx.BatchInsertData(result.Data, 100)
	}

	endtime := time.Now()
	common.Info("===endime %v===", endtime.Format("2006-01-02 15:04:05"))
	common.Info("===total cost: %v seconds===", endtime.Sub(starttime).Seconds())
}
