package handle

import (
	"github.com/nfnt/resize"
	"github.com/ying32/govcl/vcl"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"
	Logs "workspace/govcl_img/logs"
	UI "workspace/govcl_img/ui"
)

var size []int

func Cropping(img chan image.Image, sizeConfig []int)  {
	size = sizeConfig

	for  {
		tailoring(<-img)
	}
}

/** 裁剪逻辑
 */
func tailoring(img image.Image) {
	// 获取图片宽高
	var imageWeight = img.Bounds()
	var width = strconv.Itoa(imageWeight.Max.X)
	var height = strconv.Itoa(imageWeight.Max.Y)

	// 取图片可用面积最小正方形
	if width != height {
		Logs.Alert("图片分辨率: " + width + "*" + height + ", 请选择正方形图片")
		return
	}

	// 创建文件夹
	dir, err := PathExists()
	if err!= nil {
		Logs.Error("创建文件夹失败", err)
		return
	}

	// 裁剪图像
	go images(img, dir)
}

// 异步裁剪图像
func images(img image.Image, dir string) {
	i := 1
	num := len(size)
	for _,w := range size{
		// 裁剪图片
		cropping(img, dir, w)

		// 进度条渲染
		fraction := int32(float32(i)/float32(num)*100) // 计算进度条
		vcl.ThreadSync(func() { // 切换到主线程渲染UI
			//UI.ProgressBar.SetPosition(fraction)
			if fraction != 100 {
				UI.ProgressBar.Show()
			} else {
				UI.ProgressBar.Hide()
				vcl.ShowMessage("转换文件到: " + dir)
			}
			UI.ProgressBar.SetProgress(fraction)
		})

		// 睡 1/10 秒
		time.Sleep(time.Second/10)
		i++
	}
}

/** 创建文件夹
 * @param 文件夹路径
 * @return 新建文件夹路径、error
 */
func PathExists() (string, error) {
	t := time.Now()
	paths, fileName := filepath.Split(Filepath) // 获取文件路径、文件名
	ext := filepath.Ext(Filepath) // 获取文件扩展名
	reg := regexp.MustCompile(ext) // 创建正则
	filename := reg.ReplaceAllString(fileName, "") // 正则去掉路径中的扩展名
	dir := paths + "/" + filename + "_" + t.Format("20060102150405") + "/" //保存文件夹

	_, err := os.Stat(dir)
	if err != nil {
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			return dir, err
		}
	}
	return dir, nil
}

/** 裁剪图片
 * @param 图片资源、裁剪宽度、保存位置
 */
func cropping(file image.Image, dir string, w int) {
	var (
		widthUint = uint(w)
		heightUint = uint(w)
		widthString = strconv.Itoa(w)
		heightString = strconv.Itoa(w)
	)

	// 调用resize库进行图片缩放(如果高度填0，resize.Resize函数中会自动计算缩放图片的宽高比)
	m1 := resize.Resize(widthUint, heightUint, file, resize.Lanczos3)

	// 将两个图片合成一张
	newImg := image.NewNRGBA(image.Rect(0, 0, w, w)) //创建一个新RGBA图像
	draw.Draw(newImg, newImg.Bounds(), m1, m1.Bounds().Min, draw.Over) // 画上第一张缩放后的图片

	// 裁剪PNG
	pngFile, _ := os.Create(dir + widthString + ".png")
	defer pngFile.Close()
	png.Encode(pngFile, newImg)
	Logs.Console("\n 裁剪完成 (PNG): 宽度 " + widthString + ", 高度 "+ heightString, nil)

	// 裁剪JPG
	jpgFile, _ := os.Create(dir + widthString + ".jpg")
	defer jpgFile.Close()
	jpeg.Encode(jpgFile, newImg, &jpeg.Options{100})
	Logs.Console("\n 裁剪完成 (JPG): 宽度 " + widthString + ", 高度 "+ heightString, nil)

	return
}