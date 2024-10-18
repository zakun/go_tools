package import_question

import (
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"time"
	"zk/tools/common"
	"zk/tools/database"
	"zk/tools/models"
	"zk/tools/xlsx"
)

var (
	ErrSubjectNotExist = errors.New("科目不存在")
	ErrEmptyLine       = errors.New("文件有空行数据")
	ErrInvalidQuestion = errors.New("试题格式不正确")
	CurrentFileData    FileData
)

type QuestionFile struct {
	Name   string
	FileNo int
}

type OptionItme struct {
	Key     string `json:"k"`
	Content string `json:"c"`
}

type FileData struct {
	SubjectKey string
	XFile      *QuestionFile
	Data       []models.Question
}

func NewQuestionFile(name string) *QuestionFile {
	return &QuestionFile{
		Name: name,
	}
}

func (qf *QuestionFile) SetFileNo(num int) {
	qf.FileNo = num
}

func (qf *QuestionFile) HandleFile() (*FileData, error) {
	// 检查科目是否存在
	subjectKey, err := GetSubjectKey(qf.Name)
	if errors.Is(err, ErrSubjectNotExist) {
		return nil, ErrSubjectNotExist
	}

	info := &FileData{
		SubjectKey: subjectKey,
		XFile:      qf,
		Data:       make([]models.Question, 0),
	}

	xlsx := xlsx.NewXlsx(qf.Name, 2)
	xlsx.Process(excelProcess, info)
	// tt, _ := json.Marshal(&xfd)
	// common.Info("info: %+v", string(tt))
	return info, nil
}

func GetSubjectKey(path string) (string, error) {
	basename := filepath.Base(path)
	name, _ := strings.CutSuffix(basename, ".xlsx")

	db := database.Db()
	type suject struct {
		Id   uint64
		Name string
		Skey string
	}
	var sub suject
	db.Table("subject").Select("id", "name", "`key` as skey").Where("name=?", name).First(&sub)

	if sub.Id == 0 {
		return "", fmt.Errorf("%w", ErrSubjectNotExist)
	}

	return sub.Skey, nil
}

func excelProcess(r []string, x *xlsx.Xlsx, params any) error {
	info := params.(*FileData)
	// common.Info("%d#, %v", x.CurrentRowNo, info.SubjectKey)

	if len(r) == 0 {
		common.CheckError(fmt.Errorf("%w, current file: %v, current line: %v", ErrEmptyLine, x.Name, x.CurrentRowNo))
	}

	qs, err := formatSingleQuestion(r)
	if err != nil {
		common.CheckError(fmt.Errorf("%v, current file: %v, current line: %v", err, x.Name, x.CurrentRowNo))
	}
	qs.SubjectKey = info.SubjectKey
	info.Data = append(info.Data, *qs)

	// common.Info("question: %+v", qs)
	return nil
}

func formatSingleQuestion(r []string) (*models.Question, error) {
	qs := &models.Question{}

	qs.ComeFrom = "zk-20240902-imp"
	qs.SubjectKey = ""
	qs.Status = 1
	qs.Uid = 36
	qs.Addtime = time.Now().Unix()

	qs.Title = r[0]
	qs.Content = r[0]

	// 题型处理
	switch r[1] {
	case "单选题":
		qs.QuestionType = 1
		qs.QType = 33
	case "多选题":
		qs.QuestionType = 2
		qs.QType = 34
	case "判断题":
		qs.QuestionType = 3
		qs.QType = 28
	}

	if qs.QuestionType == 0 || len(r) < 15 {
		return nil, fmt.Errorf("%w, raw line: %v, len: %d", ErrInvalidQuestion, r, len(r))
	}

	// 选项处理
	sOptions := make([]string, 0, 4)
	sItem := make([]OptionItme, 0, 4)
	ascii := 65
	for _, op := range r[3:14] {
		if op != "" {
			sOptions = append(sOptions, op)
			sItem = append(sItem, OptionItme{
				Key:     fmt.Sprintf("%c", ascii),
				Content: op,
			})
			ascii++
		}
	}

	if qs.QuestionType == 3 {
		qs.Option = strings.Join(sOptions, "|")
		qs.OptionNum = len(sOptions)
		tmp := []OptionItme{
			{
				Key:     "t",
				Content: "正确",
			},
			{
				Key:     "f",
				Content: "错误",
			},
		}
		soptions, _ := json.Marshal(&tmp)
		qs.Options = string(soptions)

	} else {
		qs.Option = strings.Join(sOptions, "\n")
		qs.OptionNum = len(sOptions)
		soptions, _ := json.Marshal(sItem)
		qs.Options = string(soptions)
	}

	// 答案处理
	answer := r[14]
	qs.Answer = answer
	// 判断题
	if answer != "" && qs.QuestionType == 3 {
		if answer == "正确" {
			qs.Answer = `["t"]`
			qs.Result = "1"
		} else if answer == "错误" {
			qs.Answer = `["f"]`
			qs.Result = "0"
		}
	}
	// 单选多选
	if answer != "" && (qs.QuestionType == 1 || qs.QuestionType == 2) {
		az := "A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P,Q,R,S,T,U,V,W,X,Y,Z"
		sAz := strings.Split(az, ",")
		sAnswer := strings.Split(answer, ",")

		arrIndex := make([]string, 0)
		for _, av := range sAnswer {
			ai := slices.Index(sAz, av)
			if ai != -1 {
				arrIndex = append(arrIndex, strconv.Itoa(ai+1))
			}
		}
		if len(arrIndex) > 0 {
			jAnswer, _ := json.Marshal(&sAnswer)
			qs.Answer = string(jAnswer)
			qs.Result = strings.Join(arrIndex, ",")
		}
	}

	return qs, nil
}

func BatchInsertData(qs []models.Question, size int) {
	db := database.Db()

	db.Table("question").CreateInBatches(qs, size)
}
