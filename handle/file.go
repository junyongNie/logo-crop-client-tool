package handle

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	Logs "workspace/govcl_img/logs"
)

var Filepath string

func Read(ch chan string, img chan image.Image) {
	for {
		// 读取文件
		read(<-ch, img)
	}
}

/** 读取图片
 * @param 图片路径
 * @return 图片资源、真实文件类型、error
 */
func read(filename string, img chan image.Image) {
	// 复制 filename 真实 value
	i := make([]byte,len(filename))
	copy(i, filename)
	Filepath = string(i)

	file, err := os.Open(filename) //打开图片
	defer file.Close()

	if err != nil {
		Logs.Error("解析图片失败", err)
	}

	// image.Decode 图片
	imgSource,_,err :=image.Decode(file)
	if err == nil {
		img <- imgSource
	} else {
		Logs.Alert("解析图片失败")
	}
}