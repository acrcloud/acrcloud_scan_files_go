package checker

import (
	"acrcloud-scan-tool-golang/logger"
	"fmt"
	"os"
)

const componentName = "checker"

func isFileExists(filename string) (bool, error) {
	_, err := os.Stat(filename)

	if err == nil {
		return true, nil

	}
	return false, err
}

func CheckFilesExists() {
	files := []string{
		"config.yaml",
		"libacrcloud_extr_tool.dylib",
	}

	for _, f := range files {
		if flag, err := isFileExists(f); !flag {
			logger.LogFatal(componentName, fmt.Sprintf("%s doesn't exist or occur some error, "+
				"if you have questions please check or contact support@acrcloud.com", f), err)
		}
	}

}
