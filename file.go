package goutils

import (
	"bufio"
	"errors"
	"io"
	"io/ioutil"
	"os"
)

func ReadLines(path string) ([]string, int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, 0, err
	}
	defer file.Close()

	var lines []string
	lineCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		lineCount++
	}

	if scanner.Err() == bufio.ErrTooLong {
		panic(scanner.Err())
	}
	return lines, lineCount, scanner.Err()
}


func ReadLinesV2(path string) ([]string, int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, 0, err
	}
	defer file.Close()

	var lines []string
	lineCount := 0
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		lines = append(lines, line)
		lineCount++
		if err == io.EOF {
			return lines, lineCount, nil
		}
		if err != nil {
			return lines, lineCount, err
		}
	}
}

func ListDir(path string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	filenames := make([]string, len(files))
	for idx, file := range files {
		filenames[idx] = file.Name()
	}
	return filenames
}

// path exists or not
func IsFileExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

//  path is a directory or not
func IsDir(path string) bool {
	if info, err := os.Stat(path); err == nil && info.IsDir() {
		return true
	}
	return false
}

// path is a file or not
func IsFile(path string) bool {
	if info, err := os.Stat(path); err == nil && info.Mode().IsRegular() {
		return true
	}
	return false
}

// remove the named file or empty directory
// https://gist.github.com/novalagung/13c5c8f4d30e0c4bff27
func RemoveFile(path string) error {
	err := os.Remove(path)
	return err
}

// create a new file with the specified name
func CreateFile(path string) error {
	if IsFileExist(path) == false {
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()
		return nil
	}
	return errors.New("goutils: file " + path + " already exists.")
}


// TODO
// 1 并发读取
// 2 写入文件

