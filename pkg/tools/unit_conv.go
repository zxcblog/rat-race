package tools

import (
	"errors"
	"strconv"
	"strings"
)

const (
	_        = iota
	KB int64 = 1 << (10 * iota)
	MB int64 = 1 << (10 * iota)
	GB int64 = 1 << (10 * iota)
	TB int64 = 1 << (10 * iota)
)

// UnitConvInt64 计算单位转换int64
func UnitConvInt64(size, unit string) (int64, error) {
	s, err := strconv.ParseInt(size, 10, 64)
	if err != nil {
		return 0, err
	}

	switch strings.ToLower(unit) {
	case "kb":
		return s * KB, nil
	case "mb":
		return s * MB, nil
	case "gb":
		return s * GB, nil
	case "tb":
		return s * TB, nil
	}
	return s, errors.New("请输入正确单位信息： KB, MB, GB, TB")
}
