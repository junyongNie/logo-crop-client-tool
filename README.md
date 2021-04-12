# logo-Crop
一键生成国内各大应用市场、微信、支付宝所需要的 logo 尺寸裁剪客户端工具

不会编译的可以直接下载 Releases 版本

<b>本项目 UI部分 基于 <a href="https://github.com/ying32/govcl">github.com/ying32/govcl</a>，特别鸣谢</b>

<br/>
<br/>
<br/>
<br/>
使用说明 

<p>
    <b>go build . <b style="color: red;">(build前一定要将 liblcl.dylib 放在 $GOPATH/src 目录下， 其他系统 lib 文件请自行下载)</b></b>
    <br/>
    <img src="https://status.psd1403.com/uploads/govcl_img_001.png" style="max-width: 640px; max-height: 480px;"/>
</p>

<p>
    <b>build 完成后在目录下会生成 govcl_img 文件，双击打开</b>
    <br/>
    <img src="https://status.psd1403.com/uploads/govcl_img_002.png" style="max-width: 640px; max-height: 480px;"/>
</p>

<p>
    <b>打开 govcl_img 文件后会生成 govcl_img.app 文件</b>
    <br/>
    <img src="https://status.psd1403.com/uploads/govcl_img_003.png" style="max-width: 640px; max-height: 480px;"/>
</p>

<p>
    <b>到这里 APP 已经可以使用了，直接拖拽想要裁剪的图片到 APP 内，进度条开始</b>
    <br/>
    <img src="https://status.psd1403.com/uploads/govcl_img_006.png" style="max-width: 640px; max-height: 480px;"/>
</p>

<p>
    <b>进度条百分百后，会在当前目录下创建一个前缀为 logo 后缀为 当天日期的文件夹，文件夹内为裁剪好的图片</b>
    <br/>
    <img src="https://status.psd1403.com/uploads/govcl_img_007.png" style="max-width: 640px; max-height: 480px;"/>
</p>

<p>
    <b>如果你还想更换 APP 图标，可以在 APP 内右键=》显示包内容 </b>
    <br/>
    <img src="https://status.psd1403.com/uploads/govcl_img_004.png" style="max-width: 640px; max-height: 480px;"/>
</p>

<p>
    <b>进入 Contents=》Resources 文件夹，替换掉 govcl_img.icns</b>
    <br/>
    <img src="https://status.psd1403.com/uploads/govcl_img_005.png" style="max-width: 640px; max-height: 480px;"/>
</p>

<p>
    <b>更换图标后的效果</b>
    <br/>
    <img src="https://status.psd1403.com/uploads/govcl_img_008.png" style="max-width: 640px; max-height: 480px;"/>
</p>


