package util

import (
	"strconv"
)

func StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
func IntTostring(i int) string {
	str1 := strconv.Itoa(i)
	return str1

}

func GetTypeMap(value interface{}) map[string]interface{} {

	switch inst := value.(type) {
	case map[string]interface{}:
		return inst
	default:
		return nil

	}

}

func GetTypeString(value interface{}) string {

	switch inst := value.(type) {
	case string:
		return inst
	default:
		return "unknow"

	}

}

func GetTypeInt(value interface{}) int {

	switch inst := value.(type) {
	case int:
		return inst
	default:
		return 0

	}

}
