package starter

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
)

var sysTab []IComp

// RegisterComp 注册要输出的配置信息
func RegisterComp(comp IComp) {
	sysTab = append(sysTab, comp)
}

// Print 控制台输出配置信息
func Print() {
	t := table.NewWriter()

	// 输出到控制台
	t.SetOutputMirror(os.Stdout)
	t.AppendSeparator()

	// 输出项目信息
	t.AppendHeader(table.Row{"COMP_NAME", "KEY", "VALUE"})
	for _, val := range sysTab {
		if !val.IsDev() {
			continue
		}

		rows := make([]table.Row, 0)
		rows = append(rows, []interface{}{val.CompName()})

		for _, v := range val.GetCompItem() {
			rows = append(rows, []interface{}{"", v.Key, v.Value})
		}
		t.AppendRows(rows)
		t.AppendSeparator()
	}
	t.Render()
}
