package cache

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unsafe"
)

const (
	B = 1 << (iota * 10)
	KB
	MB
	GB
)

func ParseSize(size string) (int64, string, error) {
	if size == "" {
		return 0, "", fmt.Errorf("empty size string")
	}

	re, err := regexp.Compile("[\\d]+")
	if err != nil {
		return 0, "", fmt.Errorf("invalid regex: %v", err)
	}

	unit := string(re.ReplaceAll([]byte(size), []byte("")))
	numStr := strings.Replace(size, unit, "", 1)
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		return 0, "", fmt.Errorf("invalid number: %v", err)
	}

	if num < 0 {
		return 0, "", fmt.Errorf("negative size not allowed")
	}

	unit = strings.ToUpper(unit)
	var bytenum int64

	switch unit {
	case "B":
		bytenum = num
		unit = "B"
	case "K", "KB":
		bytenum = num * KB
		unit = "KB"
	case "M", "MB":
		bytenum = num * MB
		unit = "MB"
	case "G", "GB":
		bytenum = num * GB
		unit = "GB"
	default:
		return 0, "", fmt.Errorf("unsupported unit: %s", unit)
	}

	return bytenum, unit, nil
}

func FormatKey(key string) string {
	return string([]byte(key))
}

//func CalSize(key string, value interface{}) int64 {
//	s := unsafe.Sizeof(key) + unsafe.Sizeof(value)
//	return int64(s)
//}

func CalSize(key string, value *mValue) int64 {
	// Calculate the size of the key
	keySize := int64(len(key))

	// Calculate the size of the value
	valueSize := int64(0)
	switch v := value.val.(type) {
	case string:
		valueSize = int64(len(v))
	case int, int32, int64, float32, float64:
		valueSize = 8 // Assuming 8 bytes for numeric types
	case bool:
		valueSize = 1 // 1 byte for boolean
	default:
		valueSize = int64(unsafe.Sizeof(v)) // Use unsafe.Sizeof for unknown types
	}

	// Add the size of the expiration time (time.Time)
	expSize := int64(8) // Assuming 8 bytes for time.Time

	// Total size is the sum of key size, value size, and expiration size
	return keySize + valueSize + expSize
}
