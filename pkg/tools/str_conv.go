package tools

import (
	"fmt"
	"strconv"
)

// TransStr 将参数转换为字符串
func TransStr(val interface{}) string {
	switch assertValue := val.(type) {
	case int:
		return strconv.Itoa(assertValue)
	case uint:
		return strconv.Itoa(int(assertValue))
	case uint8:
		return strconv.Itoa(int(assertValue))
	default:
		return fmt.Sprintf("%s", val)
	}
}
