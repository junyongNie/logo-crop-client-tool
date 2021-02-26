package main

import (
	"image"
	Hand "workspace/govcl_img/handle"
	UI "workspace/govcl_img/ui"
)

/*  !!!!!!!!!!!!!!!!!!!!!!!!!!!!
	!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	没有改成异步
	第一次裁剪未完成，直接拖入下一张图片报错
	选择小于1024的图片是否提示？
*/


var filepath string
var ch_file = make(chan string)
var ch_image = make(chan image.Image)

func main() {
	// 读取文件
	go Hand.Read(ch_file, ch_image)

	// 输出文件
	go Hand.Cropping(ch_image, Size)

	// 初始化UI
	UI.Init(ch_file)
}
