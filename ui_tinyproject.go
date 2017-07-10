package main

import (
	"fmt"
	"log"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type Config struct {
	project_path string
	output_path  string
}

type Window_TinyProject struct {
	*walk.MainWindow
	ui                    *MainWindow
	lineEdit_project_path *walk.LineEdit
	lineEdit_output_path  *walk.LineEdit

	progressBar_status *walk.ProgressBar
	textEdit_info      *walk.TextEdit
}

func NewWindow_TinyProjec() *Window_TinyProject {
	mw := new(Window_TinyProject)
	mw.initUI()
	icon, err := walk.NewIconFromResource("IDI_ICON1")
	if err == nil {
		mw.ui.Icon = icon
	}
	return mw
}

func (wt *Window_TinyProject) initUI() {
	wt.ui = &MainWindow{
		AssignTo: &wt.MainWindow,
		Title:    tr("Tiny your Project"),
		MenuItems: []MenuItem{
			Action{
				Text: tr("Exit"),
				OnTriggered: func() {
					wt.Close()
				},
			},
			Action{
				Text:        tr("About"),
				OnTriggered: wt.on_about_triggered,
			},
		},
		Size:   Size{640, 480},
		Layout: VBox{},
		Children: []Widget{
			Composite{ //输入输出文件夹选择
				Layout: Grid{MarginsZero: true, Columns: 2},
				Children: []Widget{
					Composite{ // 左侧列表提示
						Layout: VBox{MarginsZero: true},
						Children: []Widget{
							Label{
								Text: tr("Project Path"),
							},
							Label{
								Text: tr("Output Path"),
							},
						},
					},
					Composite{ //  选择框和按钮
						Layout: VBox{MarginsZero: true},
						Children: []Widget{
							Composite{
								Layout: HBox{MarginsZero: true},
								Children: []Widget{
									LineEdit{
										AssignTo:           &wt.lineEdit_project_path,
										AlwaysConsumeSpace: true,
									},
									PushButton{
										OnClicked: wt.on_pushButton_project_path_clicked,
										Text:      tr("select"),
									},
								},
							},
							Composite{
								Layout: HBox{MarginsZero: true},
								Children: []Widget{
									LineEdit{
										AssignTo:           &wt.lineEdit_output_path,
										AlwaysConsumeSpace: true,
									},
									PushButton{
										OnClicked: wt.on_pushButton_output_path_clicked,
										Text:      tr("select"),
									},
								},
							},
						},
					},
				},
			},
			PushButton{
				Text:      tr("Start Parse"),
				OnClicked: wt.startParse,
				MaxSize:   Size{200, 100},
			},
			ProgressBar{
				AssignTo: &wt.progressBar_status,
			},
			TextEdit{
				Enabled:            false,
				AssignTo:           &wt.textEdit_info,
				AlwaysConsumeSpace: true,
				MinSize:            Size{100, 100},
			},
		},
	}
}

func (wt *Window_TinyProject) Setup() {
	_, err := wt.ui.Run()
	if err != nil {
		log.Fatal(err)
	}
}

//	OnTriggered: wt.on_about_triggered,

//	OnClicked: wt.on_pushButton_project_path_clicked,

//	OnClicked: wt.on_pushButton_output_path_clicked,

func (wt *Window_TinyProject) on_about_triggered() {
	walk.MsgBox(wt, tr("About"), tr("about info"), walk.MsgBoxIconInformation)
}

func (wt *Window_TinyProject) selectFolder(input *walk.LineEdit) {
	dlg := new(walk.FileDialog)
	oldPath := input.Text()

	if oldPath != "" {
		dlg.FilePath = oldPath
	}

	dlg.Title = tr("Please select project folder")

	if ok, err := dlg.ShowBrowseFolder(wt); err != nil {
		wt.showLog(err)
	} else if !ok {
		return
	}

	input.SetText(dlg.FilePath)
}

func (wt *Window_TinyProject) on_pushButton_project_path_clicked() {
	wt.selectFolder(wt.lineEdit_project_path)
}

func (wt *Window_TinyProject) on_pushButton_output_path_clicked() {
	wt.selectFolder(wt.lineEdit_output_path)
}

func (w *Window_TinyProject) startParse() {
	w.showLog("hello world")
}

func (w *Window_TinyProject) showLog(data ...interface{}) {
	str := fmt.Sprintln(data...)
	str = str + "\r\n"
	w.textEdit_info.AppendText(str)
}
