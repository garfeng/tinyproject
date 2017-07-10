package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
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

func copy(src, dst string) error {
	fp2, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer fp2.Close()
	fp1, err := os.Open(src)
	if err != nil {
		return err
	}
	defer fp1.Close()
	_, err = io.Copy(fp2, fp1)
	return err
}

func readConf() {
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		return
	}
	json.Unmarshal(data, config)
}

func writeConf() {
	data, err := json.Marshal(config)
	if err != nil {
		log.Println(err)
	}
	log.Println(config, string(data))
	err = ioutil.WriteFile("./config.json", data, 0644)
	if err != nil {
		log.Println(err)
	}
}

func init() {
	config = new(Config)
	readConf()
}
