/*
 * @Author: qizk qizk@mail.open.com.cn
 * @Date: 2024-10-15 14:59:24
 * @LastEditors: qizk qizk@mail.open.com.cn
 * @LastEditTime: 2024-10-17 09:29:12
 * @FilePath: \zk_tools\xsearch\question.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package xsearch

import (
	"encoding/json"
	"regexp"
	"slices"
	"strings"
	"zk/tools/common"
	"zk/tools/database"

	"gorm.io/gorm"
)

type Question struct {
	Id          int
	QuestionId  int
	Title       string
	OriginTitle string
	Answer      string
}

var db *gorm.DB

func init() {
	db = database.DbName("db_kaopei")
}

func Process(isAnswer bool) {

	var results []Question
	var sum int
	result := db.Table("x_question").FindInBatches(&results, 1000, func(tx *gorm.DB, batch int) error {

		sum = sum + len(results)
		for _, question := range results {
			id := question.Id
			if isAnswer {
				rightAnser := FindAnswer(question)
				// logger.Info("==rightAnswer==: %v", rightAnser)

				tx.Table("x_question").Where("id=?", id).Update("answer", rightAnser)
			} else {
				title := question.OriginTitle
				id := question.Id
				title = TrimHtmlTags(title)
				title = Replace(title)
				title = strings.ToUpper(title)

				tx.Table("x_question").Where("id=?", id).Update("title", title)
			}

		}

		common.Info("%v#, total record: %v ", batch, sum)
		return nil
	})

	common.Info("affacted rows: %v", result.RowsAffected)

}

func Replace(title string) string {

	sPattern := []string{
		`\r?\n`,
		`^[一|二|三|四|五|六|七|八|九|十]+、*`,
		`（`,
		`）`,
		`[“|”]`,
		`[。|．]`,
		`：`,
		`；`,
		`，`,
		`？`,
		`[\s| |　]`,
		`[_]+`,
		`[—]+`,
		`^[（|\(][\d|A-Z|a-z]+[）|\)]`,
		`^\d+\.`,
		`^\.`,
		`\(每小题\d+分,共\d+分\)`,
		`\(每小题\d+分\)`,
		`\(每题\d+分\)`,
		`^阅读理解:[选择|填空|判断|多选|简答|问答|简述]+题`,
		`[\.\(\)|\.|\(\)\.|\?|_|-|:]+$`,
	}

	patterns := map[string]string{
		`\r?\n`: "",
		`^[一|二|三|四|五|六|七|八|九|十]+、*`: "",
		`（`:                          "(",
		`）`:                          ")",
		`[“|”]`:                      "\"",
		`[。|．]`:                      ".",
		`：`:                          ":",
		`；`:                          ";",
		`，`:                          ",",
		`？`:                          "?",
		`[\s| |　]`:                   "",
		`[_]+`:                       "_",
		`[—]+`:                       "-",
		`^[（|\(][\d|A-Z|a-z]+[）|\)]`: "",
		`^\d+\.`:                     "",
		`^\.`:                        "",
		`\(每小题\d+分,共\d+分\)`:          "",
		`\(每小题\d+分\)`:                "",
		`\(每题\d+分\)`:                 "",
		`^阅读理解:[选择|填空|判断|多选|简答|问答|简述]+题`: "阅读理解:",
		`[\.\(\)|\.|\(\)\.|\?|_|-|:]+$`:  "",
	}

	for _, key := range sPattern {
		reg := regexp.MustCompile(key)

		// common.Info("reg: %v", key)
		// common.Info("==o: %v", title)
		title = reg.ReplaceAllString(title, patterns[key])
		// common.Info("==n: %v", title)
		// common.Info(">>>>>>>>>>")
	}

	return title
}

func TrimHtmlTags(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile(`<[\S\s]+?>`)
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile(`<style[\S\s]+?</style>`)
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile(`<script[\S\s]+?</script>`)
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile(`<[\S\s]+?\>`)
	src = re.ReplaceAllString(src, "")
	//去除连续的换行符
	re, _ = regexp.Compile(`\s{2,}`)
	src = re.ReplaceAllString(src, "")
	return strings.TrimSpace(src)
}

func FindAnswer(xquestion Question) string {
	qid := xquestion.QuestionId

	type QuestionInfo struct {
		Id       int
		Type     int
		Option   string
		Options  string
		Answer   string
		Result   string
		Analysis string
	}
	var qi QuestionInfo

	db.Table("question").Where("id=?", qid).First(&qi)

	type OptionsItem struct {
		K string
		C string
	}

	var rightAnser = ""
	switch qi.Type {
	case 1:
		fallthrough
	case 2:
		var oi []OptionsItem
		var ai []string

		json.Unmarshal([]byte(qi.Options), &oi)
		json.Unmarshal([]byte(qi.Answer), &ai)
		// logger.Info("info: %v", qi)
		// logger.Info("==oi==: %+v", oi)
		// logger.Info("==ai==: %+v", ai)
		s1 := make([]string, 0)
		for _, ov := range oi {
			if slices.Contains(ai, ov.K) {
				s1 = append(s1, ov.C)
			}
		}

		rightAnser = strings.Join(s1, ",")
	case 3:
		// logger.Info("判断题")
		rightAnser = "错误"
		if qi.Result == "1" {
			rightAnser = "正确"
		}
	case 5:
		// logger.Info("填空题")
		rightAnser = strings.ReplaceAll(qi.Option, "\n", ",")
	default:
		// logger.Info("简答题")
		rightAnser = qi.Analysis
	}

	return rightAnser
}
