package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/360EntSecGroup-Skylar/excelize/v2"

)

var frd *os.File
var fed *os.File

type fileDir struct {
	SrcDir string
	DestDir string
}

type fileDirs struct {
	FileDir []fileDir
}

func main() {
	frd,_=os.OpenFile("record"+time.Now().Format("20060102150405")+".txt",os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	fed,_=os.OpenFile("error"+time.Now().Format("20060102150405")+".txt",os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	file_xlsx := "excel.xlsx"
	f, err := excelize.OpenFile(file_xlsx)
	if err != nil {
		println(err.Error())
		return
	}
	rows, _ := f.GetRows("Sheet1")
	rData :=  make([]string, 0)
	dData :=  make([]string, 0)
	for _, row := range rows {
		for k1, colCell := range row {
			if k1 == 0{
				rData = append(rData,colCell)
			}
			if k1 == 1{
				dData = append(dData,colCell)
			}
		}
	}
	var config fileDirs
	if _,err := toml.DecodeFile("file_dir.toml",&config);err !=nil{
		fmt.Println(err)
		fmt.Println("配置文件错误")
		return
	}
	fmt.Println(rData)
	fmt.Println(dData)
	for _,v := range config.FileDir{
		for kk,vv := range rData{
			fmt.Println("开始复制源目录"+v.SrcDir+vv)
			tracefile("开始复制源目录"+v.SrcDir+vv,frd)
			fmt.Println("开始复制到目标目录"+v.DestDir+dData[kk])
			tracefile("开始复制到目标目录"+v.DestDir+dData[kk],frd)
			CopyDir(v.SrcDir+vv,v.DestDir+dData[kk],vv)
		}
	}
	frd.Close()
	fed.Close()
}

func CopyDir(srcPath string, destPath string,dirname string) error {
	moveFlag := true
	d, _ := pathExists(destPath)
	if d == false {
		fmt.Println("创建目录:" + destPath)
		tracefile("创建目录:" + destPath,frd)
		//创建目录
		err := os.Mkdir(destPath, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			tracefile(err.Error(),fed)
		}
		err = os.Mkdir(destPath+"/"+dirname, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			tracefile(err.Error(),fed)
		}
	}
	srcPath = strings.Replace(srcPath,"\\","/",-1)
	destPath = strings.Replace(destPath+"/"+dirname,"\\","/",-1)
	//检测目录正确性
	if srcInfo, err := os.Stat(srcPath); err != nil {
		fmt.Println(err.Error())
		tracefile(err.Error(),fed)
		return err
	} else {
		if !srcInfo.IsDir() {
			e := errors.New("srcPath不是一个正确的目录！")
			fmt.Println(e.Error())
			tracefile(e.Error(),fed)
			return e
		}
	}
	if destInfo, err := os.Stat(destPath); err != nil {
		fmt.Println(err.Error())
		tracefile(err.Error(),fed)
		return err
	} else {
		if !destInfo.IsDir() {
			e := errors.New("destInfo不是一个正确的目录！")
			fmt.Println(e.Error())
			tracefile(e.Error(),fed)
			return e
		}
	}
	//加上拷贝时间:不用可以去掉
	//destPath = destPath + "_" + time.Now().Format("20060102150405")
	destPath = destPath
	err := filepath.Walk(srcPath, func(path string, f os.FileInfo, err error) error{
		if f == nil {
			return err
		}
		if !f.IsDir() {
			path := strings.Replace(path, "\\", "/", -1)
			destNewPath := strings.Replace(path, srcPath, destPath, -1)
			fmt.Println("复制文件:" + path + " 到 " + destNewPath)
			tracefile("复制文件:" + path + " 到 " + destNewPath,frd)
			copyFile(path, destNewPath)
			moveFlag = false
		}
		return nil
	})
	if moveFlag {
		fmt.Println(srcPath+"{{{不存在需要移动的文件}}}")
		tracefile(srcPath+"{{{不存在需要移动的文件}}}",fed)
	}
	if err != nil {
		fmt.Printf(err.Error())
		tracefile(err.Error(),fed)
	}
	return err
}

//生成目录并拷贝文件
func copyFile(src, dest string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err.Error())
		tracefile(err.Error(),fed)
		return
	}
	defer srcFile.Close()
	//分割path目录
	destSplitPathDirs := strings.Split(dest, "/")

	//检测时候存在目录
	destSplitPath := ""
	for index, dir := range destSplitPathDirs {
		if index < len(destSplitPathDirs)-1 {
			destSplitPath = destSplitPath + dir + "/"
			b, _ := pathExists(destSplitPath)
			if b == false {
				fmt.Println("创建目录:" + destSplitPath)
				tracefile("创建目录:" + destSplitPath,frd)
				//创建目录
				err := os.Mkdir(destSplitPath, os.ModePerm)
				if err != nil {
					fmt.Println(err)
					tracefile(err.Error(),fed)
				}
			}
		}
	}
	dstFile, err := os.Create(dest)
	if err != nil {
		fmt.Println(err.Error())
		tracefile(err.Error(),fed)
		return
	}
	defer dstFile.Close()
	return io.Copy(dstFile, srcFile)
}

//检测文件夹路径时候存在
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func tracefile(str_content string,fd *os.File)  {
	fd_time:=time.Now().Format("2006-01-02 15:04:05")
	fd_content:=strings.Join([]string{"======",fd_time,"=====","\n",str_content,"\n"},"")
	buf:=[]byte(fd_content)
	fd.Write(buf)
}
