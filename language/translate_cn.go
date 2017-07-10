package main

const (
	kCurrentLang = "Chinese"
)

func init() {

	langList[kCurrentLang] = &LangMap{
		"Tiny your Project":            "缩小你的工程",
		"Exit":                         "退出",
		"About":                        "关于",
		"Project Path":                 "工程目录",
		"Output Path":                  "输出目录",
		"select":                       "选择",
		"Start Parse":                  "开始处理",
		"about info":                   "缩小工程 by garfeng\ngithub.com/garfeng/tinyproject",
		"Please select project folder": "请选择工程目录",
		"input is not a dir":           "输入/输出不是一个文件夹",
	}
}
