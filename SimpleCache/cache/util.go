package cache

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	B = 1 << (iota * 10)
	KB
	MB
	GB
)

func ParseSize(size string) (int64, string) {
	re, _ := regexp.Compile("[\\d]+")
	unit := string(re.ReplaceAll([]byte(size), []byte("")))
	num, _ := strconv.ParseInt(strings.Replace(size, unit, "", 1), 10, 64)
	unit = strings.ToUpper(unit)
	fmt.Println(unit)
	var bytenum int64 = 0
	switch unit {
	case "B":
		bytenum = num
		break
	case "K":
	case "KB":
		bytenum = num * KB
		unit = "KB"
		break
	case "MB":
	case "M":
		bytenum = num * MB
		unit = "MB"
		break
	case "G":
	case "GB":
		bytenum = num * GB
		unit = "GB"
		break
	default:
		bytenum = 0
		panic("不支持的大小")
	}

	fmt.Println(bytenum, unit)
	return bytenum, unit
}

func FormatKey(key string) string {
	return string([]byte(key))
}
