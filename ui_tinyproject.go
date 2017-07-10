package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type Config struct {
	Project_path string
	Output_path  string
}

type Window_TinyProject struct {
	*walk.MainWindow
	ui                    *MainWindow
	lineEdit_project_path *walk.LineEdit
	lineEdit_output_path  *walk.LineEdit

	pushButton_start *walk.PushButton

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
										Text:               config.Project_path,
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
										Text:               config.Output_path,
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
				AssignTo:  &wt.pushButton_start,
			},
			ProgressBar{
				AssignTo: &wt.progressBar_status,
				MaxValue: 100,
				MinValue: 0,
			},
			TextEdit{
				AssignTo:           &wt.textEdit_info,
				AlwaysConsumeSpace: true,
				MinSize:            Size{100, 100},
				VScroll:            true,
				MaxLength:          65535,
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

	config.Project_path = wt.lineEdit_project_path.Text()
	config.Output_path = wt.lineEdit_output_path.Text()

	log.Println(config)

	writeConf()
}

func (wt *Window_TinyProject) on_pushButton_project_path_clicked() {
	wt.selectFolder(wt.lineEdit_project_path)
}

func (wt *Window_TinyProject) on_pushButton_output_path_clicked() {
	wt.selectFolder(wt.lineEdit_output_path)
}

func (wt *Window_TinyProject) initProgressBar() {
	wt.progressBar_status.SetValue(0)
}

func (wt *Window_TinyProject) setValueOfProgressBar(v float64) {
	wt.progressBar_status.SetValue(int(v))
}

func (w *Window_TinyProject) parseCore() {
	defer w.pushButton_start.SetEnabled(true)

	input := w.lineEdit_project_path.Text()
	output := w.lineEdit_output_path.Text()
	if !isDir(input) {
		w.showLog(tr("input is not a dir"))
		return
	}

	err := os.RemoveAll(output)
	if err != nil {
		w.showLog(err)
		return
	}
	err = os.MkdirAll(output, 0755)
	if err != nil {
		w.showLog(err)
		return
	}

	fileList := allFilesInDir(input)
	one_step := 100.0 / float64(len(fileList))
	current_data := float64(0)

	w.initProgressBar()

	for _, v := range fileList {
		w.showLog(v)
		newPath := strings.Replace(v, input, output, 1)
		newDir, _ := filepath.Split(newPath)
		os.MkdirAll(newDir, 0755)
		if isPng(v) {

			//i := strings.Replace(v, " ", "\\ ", -1)
			//o := strings.Replace(newPath, " ", "\\ ", -1)
			func() {
				cmd := exec.Command("./pngquant.exe", "--skip-if-larger",
					"--output", newPath, "--speed", "1", "255", "--", v)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				<-time.After(1e7)
				err := cmd.Run()
				if err != nil {
					w.showLog("Error:", err, "copy the original file")
					copy(v, newPath)
				}
			}()
		} else {
			copy(v, newPath)
		}
		current_data += one_step
		w.setValueOfProgressBar(current_data)
	}
}

func (w *Window_TinyProject) startParse() {
	w.pushButton_start.SetEnabled(false)
	go w.parseCore()
}

func (w *Window_TinyProject) showLog(data ...interface{}) {
	if w.textEdit_info.TextLength() > w.textEdit_info.MaxLength()/4 {
		w.textEdit_info.SetText("")
	}

	str := fmt.Sprintln(data...)
	str = str + "\r\n"
	w.textEdit_info.AppendText(str)
	log.Println(data...)
}
