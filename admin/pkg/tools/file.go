package tools

import (
	"os"
	"path/filepath"
)

// GetPath 通过文件名称获取到文件的绝对路径
func GetPath(fileName string) (string, error) {
	if filepath.IsAbs(fileName) {
		return fileName, nil
	}

	var path string
	cwd, err := os.Getwd()
	if err != nil {
		return path, err
	}

	path = filepath.Clean(filepath.Join(cwd, fileName))
	return path, nil
}

// IsExists 判断文件是否存在
func IsExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
