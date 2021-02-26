package ui

import (
	"errors"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types/colors"
	"os"
	"os/exec"
	"path/filepath"
	Logs "workspace/govcl_img/logs"
)

type TMainForm struct {
	*vcl.TForm
}

var (
	mainForm *TMainForm
)

var chAddr *chan string
var ProgressBar *vcl.TGauge

func Init(ch chan string) {
	chAddr = &ch

	vcl.RunApp(&mainForm)
}

// 创建窗口
func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	// 创建菜单
	itemSupport := vcl.NewMenuItem(f)
	itemSupport.SetCaption("支持的分辨率")
	itemSupport.SetOnClick(func(sender vcl.IObject) {
		vcl.ShowMessage("8, 16, 20, 28, 29, 32, 40, 48, 58, 60, 64, 72, 80, 87, 96, 100, 108, 120, 128, 144, 180, 192, 256, 320, 512, 1024")
	})

	itemAbout := vcl.NewMenuItem(f)
	itemAbout.SetCaption("关于我们")
	itemAbout.SetOnClick(func(sender vcl.IObject) {
		vcl.ShowMessage("©️黑豹科技有限公司 408811093@qq.com")
	})

	item := vcl.NewMenuItem(f)
	item.SetCaption("帮助")
	item.Add(itemSupport)
	item.Add(itemAbout)

	menu := vcl.NewMainMenu(f)
	menu.Items().Add(item)



	// 设置窗口位置
	f.ScreenCenter() // 窗口居中
	f.SetColor(colors.ClWhite)
	f.SetWidth(240) // 窗口宽度
	f.SetHeight(240) // 窗口高度
	f.SetAutoSize(false)
	f.EnabledSystemMenu(false) // ?
	f.SetOnConstrainedResize(func(sender vcl.IObject, minWidth, minHeight, maxWidth, maxHeight *int32) {
		*minWidth = 240
		*minHeight = 240
		*maxWidth = 240
		*maxHeight = 240
	}) // 约束窗口大小
	f.SetCaption("LOGO裁剪工具") // 设置窗口标题

	// 拖入文件事件
	f.SetAllowDropFiles(true) // 允许拖拽入文件
	f.SetOnDropFiles(func(sender vcl.IObject, aFileNames []string) {
		for i := 0; i < len(aFileNames); i++ {
			*chAddr <- aFileNames[i]
		}
	})

	// 创建背景
	icns, error := staticPath("govcl_img.icns")
	img := vcl.NewImage(f)
	img.SetParent(f)
	img.SetStretch(true) // 拉伸
	img.SetProportional(true) // 成比例
	img.SetBounds(70, 40, 100, 100) // 设置文件位置
	if error == nil {
		img.Picture().LoadFromFile(icns) // 从本地读取文件
	} else {
		Logs.Console("\n 读取 logo 文件失败 ", error)
	}

	// 创建提示文字
	font := vcl.NewFont()
	font.SetSize(12) // 字体大小
	font.SetColor(0x6C6C6C) // bgr 颜色

	text := vcl.NewStaticText(f)
	text.SetParent(f)
	text.SetFont(font)
	text.SetTop(165)
	text.SetLeft(60)
	text.SetAutoSize(true) // 自动调整大小
	text.SetCaption("拖拽 logo 到这里")

	// 创建进度条
	//ProgressBar = vcl.NewProgressBar(f)
	//ProgressBar.SetParent(f)
	//ProgressBar.SetTop(220)
	//ProgressBar.SetLeft(5)
	//ProgressBar.SetWidth(230)
	//ProgressBar.SetPosition(10)

	ProgressBar = vcl.NewGauge(f)
	ProgressBar.SetParent(f)
	ProgressBar.SetProgress(50)
	ProgressBar.SetBounds(70, 40, 100, 100) // 设置文件位置
	ProgressBar.Hide()
}

// 获取静态文件
func staticPath(filename string) (string, error) {
	var file, path string
	var err error

	// 取运行绝对路径
	file, err = exec.LookPath(os.Args[0])
	if err == nil {
		path = filepath.Dir(file)
		path1 := path + "/static/" + filename
		_, err = os.Stat(path1)
		if err == nil {
			return path1, nil
		} else {
			Logs.Console("", err)
		}
	}

	// 打包后的目录
	path2 := path + "/../Resources/" + filename
	_, err = os.Stat(path2)
	if err == nil {
		return path2, nil
	} else {
		Logs.Console("", err)
	}

	// 相对路径
	path, err = os.Getwd()
	if err == nil {
		path3 := path + "/static/" + filename
		_, err = os.Stat(path3)
		if err == nil {
			return path3, nil
		} else {
			Logs.Console("", err)
		}
	}

	return "", errors.New("未找到文件目录")
}
