# images_upload_tool
github 图床上传工具  【喜欢的欢迎 点击 star  】

##  typora 代码配置

![image](https://user-images.githubusercontent.com/46613910/127744213-a712db07-42d1-4723-a74f-a61515461b66.png)


![image](https://user-images.githubusercontent.com/46613910/127744226-ec110750-cac7-4b6e-8003-5e0a416c0968.png)

打包后 的 exe ,的文件路径 复制到typora 设置 那里就可以了

上传图片后的效果：
https://cdn.jsdelivr.net/gh/lyr-2000/images_repo_2021_ASUS/2021_07_31_22__59_01a0.png
![](https://cdn.jsdelivr.net/gh/lyr-2000/images_repo_2021_ASUS/2021_07_31_22__59_01a0.png)


##  食用方法
你需要下载 go的运行环境



```go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)


func fPrefix() string {
	return time.Now().Format("2006_01_02_15__04_05")
}

// 改成  https://cdn.jsdelivr.net/gh/你的账号/仓库名字/
// 图片URL的前缀
var prePicURL = "https://cdn.jsdelivr.net/gh/lyr-2000/images_repo_2021_ASUS/"

// 肯定要 先 git init 初始化 的，这个目录，然后懂得都懂
// git 本地仓库路径
var picPath = "F:\\STATIC_FILE_XXOO_LIN_YANGRUI_USER_CUSTOM_SYS_FILES_WARNING\\region0\\staticFS\\IMAGES\\2021_7_31\\"
// 生成时间前缀，防止文件覆盖
var timePrefix = fPrefix()



func main() {
	flag.Parse()
	// \Users\Lenovo\Desktop\piccoding\main.exe "C:\\Users\\Lenovo\\AppData\\Local\\Temp/typora-icon2.png" "C:\\Users\\Lenovo\\AppData\\Local\\Temp/typora-icon.png"

	file := flag.Args()

	if len(file) == 0 {
		os.Exit(1)
	}
	picURLs := make([]string, len(file))

	for _, v := range file {
		if !isFileExist(v) {
			if cpFile(v) {
				upload(v)
				// filepath.Base(v) 获取文件名
				picURLs = append(picURLs, prePicURL+timePrefix+filepath.Base(v))
			}
		} else {
			picURLs = append(picURLs, prePicURL+timePrefix+filepath.Base(v))
		}
	}

	fmt.Println("Upload Success:")
	//打印数组
	for _, v := range picURLs {
		fmt.Println(v)
	}

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


```




改一下 代码，然后 go build 就可以运行了

