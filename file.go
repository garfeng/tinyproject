package main

import (
	"os"
	"path/filepath"
	"strings"
)

func isDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func isPng(path string) bool {
	if isDir(path) {
		return false
	}
	return strings.HasSuffix(path, ".png")
}

// list dir
// if not exists return empty slice
func scanDir(directory string) []string {
	file, err := os.Open(directory)
	if err != nil {
		return []string{}
	}
	defer file.Close()
	names, err := file.Readdirnames(-1)
	if err != nil {
		return []string{}
	}
	newList := make([]string, 0)
	for _, v := range names {
		newList = append(newList, filepath.Join(directory, v))
	}
	return newList
}

func allFilesInDir(path string) []string {
	result := make([]string, 0)
	if !isDir(path) {
		return result
	}

	tmp := scanDir(path)

	for _, v := range tmp {
		if !isDir(v) {
			result = append(result, v)
		} else {
			tmp_child := allFilesInDir(v)
			result = append(result, tmp_child...)
		}
	}

	return result
}
