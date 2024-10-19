package utils

import (
	"bytes"
	"fmt"
)

func Html2Pdf(fileName, title, content string) (err error) {

	buf := new(bytes.Buffer)

	article := genHeadHtml() + content + `
	</div>
</body>
</html>`
	buf.Write([]byte(article))
	pdf := PdfOption{
		FileName: fileName,
		Title:    title,
		PageSize: "A4",
		Toc:      false,
	}
	fmt.Printf("正在生成文件：【\033[37;1m%s\033[0m】 ", fileName)
	err = pdf.GenPdf(buf)
	return
}

func genHeadHtml() (result string) {
	result = `<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
   <meta name="viewport" content="width=device-width, initial-scale=1.0">
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
	<style>
		@font-face { font-family: "FZFangSong-Z02"; src:local("FZFangSong-Z02"), url("https://imgcdn.umiwi.com/ttf/fangzhengfangsong_gbk.ttf"); }
		@font-face { font-family: "FZKai-Z03"; src:local("FZFangSong-Z02S"), url("https://imgcdn.umiwi.com/ttf/0315911813008928624065681028886857980055.ttf"); }
		@font-face { font-family: "FZKai-Z03"; src:local("FZKai-Z03"), url("https://imgcdn.umiwi.com/ttf/fangzhengkaiti_gbk.ttf"); }
		@font-face { font-family: "PingFang SC"; src:local("PingFang SC"); }
		table, tr, td, th, tbody, thead, tfoot {page-break-inside: avoid !important;}
		img { page-break-inside: avoid; border-style: none;max-width: 100% !important;}
		img.epub-footnote { padding-right:5px;}
		img-info {page-break-inside: avoid; border-style: none;display: block;margin-top: -6px;margin-bottom: 18px;color: #A1A8AD;font-size: 12pt;text-align: center;line-height: 1.85;padding: 0 7px;}
		body {font-family: "-apple-system-font", "BlinkMacSystemFont", "Microsoft YaHei UI", "Microsoft YaHei", "Source Han Serif SC", "Spectral", "Pingfang SC", "SFUI Text", "Source Han Serif SC", "Noto Sans CJK SC", "Roboto", "lucida grande", "lucida sans unicode", lucida, helvetica, "Hiragino Sans GB", "WenQuanYi Micro Hei";color:#4C4948;text-align:left;line-height:1.8;}
		em {font-style: normal;}
		h2>code { background-color: rgb(255, 96, 2);padding: 0.5%;border-radius: 10%;color: white;}
		p>em {color: rgb(255, 96, 2);}
		.timestamp {
           font-weight: 500;
           color: #2E4E6F;
           text-decoration: underline;
       }
	</style>
</head>
<body>
<div class="course-content">
`
	return
}
