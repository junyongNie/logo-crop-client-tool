package logs

import (
	"github.com/coopsrc/glog"
	"github.com/ying32/govcl/vcl"
)

// 日志
func Console(msg string, err error)  {
	glog.InfoF("info", msg + "==> %s", err)
}

// 日志
func Error(msg string, err error)  {
	glog.ErrorF("error", msg + "==> %s", err)
}

// 错误
func Alert(msg string)  {
	vcl.ThreadSync(func() { // 切换到主线程
		vcl.MessageDlg(msg, 1)
	})
}