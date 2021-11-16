package tools

import (
	"github.com/thedevsaddam/gojsonq"
	"os"
)

// 读取JSON文件
func ReadJsonFile(filePath string) interface{} {
	wdPath,_ := os.Getwd()
	fileName := wdPath + filePath
	jsonData := gojsonq.New().File(fileName).Get()
	return jsonData
}
