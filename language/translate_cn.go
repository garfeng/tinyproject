package main

const (
	kCurrentLang = "Chinese"
)

func init() {

	langList[kCurrentLang] = &LangMap{
		"UVSS Update Tool":   "UVSS 升级工具",
		"Exit":               "退出",
		"About":              "关于",
		"Language":           "语言",
		"The Update package": "升级包",
		"Your computer IP":   "你的电脑IP",
		"Please select":      "请选择",
		"Select":             "选择",
		"All device IP to update, one IP each line.": "填入所有要升级的设备IP，每个IP一行",
		"One device IP a line":                       "每个IP一行",
		"Click to update":                            "点击以升级",
		"Update":                                     "升级",
		"Click the button above to update the camera": "点击上方按钮升级相机",
		"finished in update":                          "完成升级",
		"is not a valid device ip":                    "不是一个有效的IP",
		"Blue Stream UVSS Camera Update Tool":         "蓝溪科技UVSS相机升级工具",
		"Succeed": "成功",
		"New config will take effect at next setup":                             "配置将会在下次启动时生效",
		"Setup the connection":                                                  "创建链接中……",
		"Fail to setup connection, please check the log.txt and send it to us.": "连接创建失败，请检查log.txt，将它发送给我们。",
		"fail to update":                          "升级失败：",
		"wait for reboot":                         "等待相机重启",
		"fail to read the version info of device": "读取设备当前版本信息失败",
		"fail to read local version info":         "读取本地包版本信息失败",
		"versions not match":                      "版本不匹配，请重新尝试升级",
		"fail to parse version of device":         "分析相机当前版本信息失败",
		"fail to parse version package":           "分析本地包版本信息失败",
		"Package not valid":                       "无效的升级包",
		"reconnect":                               "尝试重连",
		"fail to copy update_manager":             "升级管理器复制失败",
		"id":                "序号",
		"ip":                "IP",
		"current version":   "当前版本",
		"version to update": "将要升级的版本",
		"Add a camera":      "添加一个相机IP",
		"add":               "添加",
		"click an item to delete it":                          "鼠标左键单击一个列表条目，可以删除它",
		"please select update package":                        "请选择升级包",
		"is not a valid device ip or not a valid uvss camera": "不是一个有效的UVSS相机",
	}
}
