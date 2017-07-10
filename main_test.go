package main

import (
	"fmt"
	"path/filepath"
	"testing"
)

func Test_isPng(t *testing.T) {
	fmt.Println(isPng("23333.png"))
	fmt.Println(isPng("23333.jpg"))
	fmt.Println(isPng("23333.gif"))
}

func Test_allFilesInDir(t *testing.T) {
	fmt.Println(allFilesInDir("./"))
}

func Test_FilePath(t *testing.T) {
	fmt.Println(filepath.Dir("."))
}
