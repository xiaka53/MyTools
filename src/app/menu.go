package app

import (
	"MyTools/src/utils"
	"strings"
)

//标签类型
type Operation struct {
	Key     string
	Label   string
	End     bool
	Process func(args []string) error
}

var menuMap [][]Operation

var operations = make(map[string]Operation)

func init() {
	menuMap = [][]Operation{
		{
			{Key: "new", Label: "创建账号"},
			{Key: "import", Label: "导入账号"},
			{Key: "export", Label: "导出账号"},
		},
		{
			{Key: "del", Label: "删除账号"},
			{Key: "edit", Label: "退出", End: true},
		},
	}
}

func showMenu() {
	var columnsMaxWidths = make(map[int]int)

	for _, v := range menuMap {
		for key, val := range v {
			maxLen := int(utils.ZhLen(operationFormat(val)))
			if _, exists := columnsMaxWidths[key]; !exists {
				columnsMaxWidths[key] = maxLen
			}
			if columnsMaxWidths[key] < maxLen {
				columnsMaxWidths[key] = maxLen
			}
			operations[val.Key] = val
		}
	}
	for _, v := range menuMap {
		str := ""
		for key, val := range v {
			str += stringPadding(operationFormat(val), columnsMaxWidths[key]) + "\t"
		}
		utils.InfoLn(strings.TrimSpace(str))
		str = ""
	}

}

func operationFormat(operation Operation) (str string) {
	str = "[" + operation.Key + "] " + operation.Label
	return
}

func stringPadding(str string, paddingLen int) string {
	if len(str) < paddingLen {
		return stringPadding(str+" ", paddingLen)
	} else {
		return str
	}
}
