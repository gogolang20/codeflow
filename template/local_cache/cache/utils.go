package cache

import (
	"encoding/json"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// memCache.go
const (
	B = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
)

func ParseSize(size string) (int64, string) {
	re, _ := regexp.Compile("[0-9]+")
	unit := string(re.ReplaceAll([]byte(size), []byte("")))
	num, _ := strconv.ParseInt(strings.Replace(size, unit, "", 1), 10, 64)
	unit = strings.ToUpper(unit)
	var byteNum int64 = 0
	switch unit {
	case "B":
		byteNum = num
	case "KB":
		byteNum = num * KB
	case "MB":
		byteNum = num * MB
	case "GB":
		byteNum = num * GB
	case "TB":
		byteNum = num * TB
	case "PB":
		byteNum = num * PB
	default:
		num = 0
	}

	if num == 0 {
		log.Println("err")
		byteNum = 100 * MB
		unit = "MB"
	}

	return byteNum, strconv.FormatInt(num, 10) + unit
}

func GetValSize(val interface{}) int64 {
	bytes, _ := json.Marshal(val)
	return int64(len(bytes))
}
