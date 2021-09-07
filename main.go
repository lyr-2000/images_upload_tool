package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func fPrefix() string {
	return time.Now().Format("2006_01_02_15_04_05")
}

// 图片URL的前缀
var prePicURL = "https://cdn.jsdelivr.net/gh/lyr-2000/images_repo_2021_ASUS/"

// git 本地仓库路径
var picPath = "F:\\STATIC_FILE_XXOO_LIN_YANGRUI_USER_CUSTOM_SYS_FILES_WARNING\\region0\\staticFS\\IMAGES\\2021_7_31\\"

var timePrefix = fPrefix()

func main() {
	flag.Parse()
	// \Users\Lenovo\Desktop\piccoding\main.exe "C:\\Users\\Lenovo\\AppData\\Local\\Temp/typora-icon2.png" "C:\\Users\\Lenovo\\AppData\\Local\\Temp/typora-icon.png"

	file := flag.Args()
	//用于协程异步上传
	var (
		wg sync.WaitGroup
	)

	if len(file) == 0 {
		os.Exit(1)
	}
	picURLs := make([]string, len(file))

	for _, v := range file {
		if v == "" {
			continue
		}
		if strings.HasPrefix(v, "https://") || strings.HasPrefix(v, "http://") {
			//网络上的图片，直接不用上传
			picURLs = append(picURLs, v)
			continue
		}
		if !isFileExist(v) {
			if cpFile(v) {
				wg.Add(1)
				//使用协程异步上传，提交响应速度
				// filepath.Base(v) 获取文件名
				picURLs = append(picURLs, prePicURL+timePrefix+filepath.Base(v))
				go func() {
					upload(v)
					wg.Done()
				}()

			}
		} else {
			//如果不存在文件，也是 直接就用原路径
			picURLs = append(picURLs, prePicURL+timePrefix+filepath.Base(v))
		}
	}

	fmt.Println("Upload Success:")
	//打印数组
	for _, v := range picURLs {
		if v == "" {
			continue
		}
		fmt.Println(v)
	}
	//等上传完再退出
	wg.Wait()
}

func isFileExist(fileName string) bool {
	//os.Stat获取文件信息
	_, err := os.Stat(picPath + filepath.Base(fileName))
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true

}

func cpFile(fileName string) bool {

	srcFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return false
	}

	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)
	//创建一个管道 ，写进入
	destFile, err := os.OpenFile(picPath+timePrefix+filepath.Base(fileName), os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)

		return false
	}
	writer := bufio.NewWriter(destFile)
	defer destFile.Close()

	_, _ = io.Copy(writer, reader)

	return true
}

func upload(fileName string) {

	cmd1 := exec.Command("git", "add", ".")
	cmd2 := exec.Command("git", "commit", "-m", filepath.Base(fileName))
	cmd3 := exec.Command("git", "push")

	cmd1.Dir = picPath
	cmd2.Dir = picPath
	cmd3.Dir = picPath

	_ = cmd1.Run()
	_ = cmd2.Run()
	_ = cmd3.Run()

}
