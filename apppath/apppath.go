package apppath

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetAppPath() (string, error) {
	execFilePath, e := exec.LookPath(os.Args[0])
	if e != nil {
		fmt.Println("获取当前执行文件路径失败, %v", e)
		return "", e
	} else {
		//解决linux下相对路径的问题
		execFilePath, _ = filepath.Abs(execFilePath)
		fmt.Println("当前执行文件路径:", execFilePath)
	}

	binDirPath := filepath.Dir(execFilePath)
	var appPath string
	if strings.HasSuffix(binDirPath, "bin") {
		appPath = filepath.Dir(binDirPath)
	} else if strings.HasSuffix(binDirPath, "src") { //just for developer
		appPath = filepath.Dir(binDirPath)
	} else {
		appPath = binDirPath
	}

	fmt.Println("应用程序所在路径:", appPath)
	/*
		workpath, e := os.Getwd()
		if e != nil {
			fmt.Println("获取当前工作路径失败", e)
			return
		} else {
			fmt.Println("当前工作路径", workpath)
		}
	*/

	return appPath, nil
}
