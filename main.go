package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// 定义要清理的 Mac 系统文件
var junkDir = []string{
	".fseventsd",
	".Spotlight-V100",
	".TemporaryItems",
	".Trashes",
}

func main() {
	// 获取当前工作目录
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("无法获取当前目录: %v\n", err)
		return
	}

	fmt.Println()
	fmt.Println("开始清理 Mac 系统生成的垃圾文件...")
	fmt.Printf("当前工作目录为: %s\n", currentDir)
	fmt.Println()

	// 检测工作目录是否为根目录并清理垃圾文件夹
	if filepath.Clean(currentDir) == filepath.Dir(currentDir) {
		for _, junkDirName := range junkDir {
			cleanJunkDir(junkDirName)
		}
	}

	// 清理工作目录及子目录中的垃圾文件
	cleanJunkFile(currentDir)

	fmt.Println()
	fmt.Println("垃圾文件清理完成，按回车键退出...")
	fmt.Scanln()
}

// cleanJunkDir 清理根目录下的垃圾文件夹
func cleanJunkDir(junkDirName string) {
	_, err := os.Stat(junkDirName)
	if err != nil {
		return
	}
	if err := os.RemoveAll(junkDirName); err != nil {
		fmt.Printf("无法删除垃圾文件夹: %s - %v\n", junkDirName, err)
	} else {
		fmt.Printf("成功删除垃圾文件夹: %s\n", junkDirName)
	}
}

// cleanJunkFile 清理工作目录及子目录中的垃圾文件
func cleanJunkFile(currentDir string) {
	filepath.Walk(currentDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if !info.IsDir() && (strings.HasPrefix(info.Name(), "._") || info.Name() == ".DS_Store") {
			if err := os.Remove(path); err != nil {
				fmt.Printf("无法删除垃圾文件: %s - %v\n", path, err)
			} else {
				fmt.Printf("成功删除垃圾文件: %s\n", path)
			}
		}
		return nil
	})
}
