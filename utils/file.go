package utils

import (
	"archive/zip"
	"errors"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// 可执行程序所在文件夹
var CurrentPath string

// web静态目录
var WebDir string
var Webs []string

func init() {
	CurrentPath, _ = GetCurrentPath()
	WebDir = CurrentPath + "webs"
	Webs = GetFolders("")
}

func Zip(srcFile string, destZip string) error {
	zipfile, err := os.Create(destZip)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	filepath.Walk(srcFile, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = strings.TrimPrefix(path, filepath.Dir(srcFile)+"/")
		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
		}
		return err
	})

	return err
}

func Unzip(zipFile string, destDir string) error {
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	for _, f := range zipReader.File {
		fpath := filepath.Join(destDir, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return err
			}

			inFile, err := f.Open()
			if err != nil {
				return err
			}
			defer inFile.Close()

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, inFile)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func GetCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\". `)
	}
	return string(path[0 : i+1]), nil
}

// GetNginxFiles 获取目录下所有文件
func GetNginxFiles(path string) []Json {
	fs, _ := os.ReadDir(path)

	list := make([]Json, len(fs))
	for index, file := range fs {
		list[index] = Json{}
		if file.IsDir() {
			list[index]["name"] = file.Name()
			list[index]["type"] = "directory"
		} else {
			list[index]["name"] = file.Name()
			list[index]["type"] = "file"
		}
	}

	return list
}

func GetFolders(path string) []string {
	list := make([]string, 0)
	fs, _ := os.ReadDir(WebDir + "/" + path)
	for _, file := range fs {
		if file.IsDir() {
			list = append(list, file.Name())
		}
	}
	return list
}

// GetFiles 获取指定目录及所有子目录下的所有文件
func GetFiles(path string) []string {
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	list := make([]string, 0)
	fs, _ := os.ReadDir(WebDir + "/" + path)
	for _, file := range fs {
		if file.IsDir() {
			list = append(list, GetFiles(path+file.Name())...)
		} else {
			list = append(list, path+file.Name())
		}
	}

	return list
}

func CamelToSnake(text string) string {
	re := regexp.MustCompile("([a-z])([A-Z])")
	snakeCase := re.ReplaceAllString(text, "${1}_${2}")
	return strings.ToLower(snakeCase)
}
