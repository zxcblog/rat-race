package tools

import "testing"

func TestGetPath(t *testing.T) {
	cases := []struct {
		FileName, AbsPath string
	}{
		{"../conf.ini", "E:\\rat-race\\pkg\\conf.ini"},
		{"../config.ini", "E:\\rat-race\\pkg\\conf.ini"},
	}

	for _, c := range cases {
		t.Run(c.FileName, func(t *testing.T) {
			path, err := GetPath(c.FileName)
			if err != nil {
				t.Fatalf("通过文件名获取文件绝对路径失败：%s", err.Error())
			}
			if path != c.AbsPath {
				t.Fatalf("通过文件名获取到的文件绝对路径错误：文件名：%s， 文件绝对路径：%s, 预设绝对路径：%s", c.FileName, path, c.AbsPath)
			}
		})
	}
}
