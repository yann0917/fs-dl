package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/bogem/id3v2/v2"
)

type ID3Options struct {
	Title  string // 标题
	Artist string // 参与创作的艺术家
	Album  string // 唱片集
	Cover  []byte // 封面
	Year   string
	Genre  string // 流派
}

// DownloadAudio 下载mp3文件
func DownloadAudio(title, fileUrl string, opt ID3Options) error {
	// 请求数据
	resp, err := http.Get(fileUrl)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	fmt.Printf("正在生成文件：【\033[37;1m%s\033[0m】 ", title)

	out, err := os.Create(title)
	if err != nil {
		fmt.Printf("\033[31;1m%s\033[0m\n", "失败"+err.Error())
		panic(err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Printf("\033[31;1m%s\033[0m\n", "失败"+err.Error())
		return err
	}

	tag, err := id3v2.ParseReader(out, id3v2.Options{Parse: true})
	if err != nil {
		log.Fatal("Error while opening mp3 file: ", err)
	}
	defer tag.Close()
	tag.DeleteAllFrames()
	// Set simple text frames.
	tag.SetArtist(opt.Artist)
	tag.SetTitle(opt.Title)
	tag.SetAlbum(opt.Album)

	if len(opt.Cover) > 0 {
		pic := id3v2.PictureFrame{
			Encoding:    id3v2.EncodingUTF8,
			MimeType:    "image/jpeg",
			Picture:     opt.Cover,
			Description: "Cover",
			PictureType: id3v2.PTFrontCover,
		}
		tag.AddAttachedPicture(pic)
	}

	// // Set comment frame.
	// comment := id3v2.CommentFrame{
	// 	Encoding:    id3v2.EncodingUTF8,
	// 	Language:    "eng",
	// 	Description: "My opinion",
	// 	Text:        "Very good song",
	// }
	// tag.AddCommentFrame(comment)

	// Write tag to file.
	if err = tag.Save(); err != nil {
		log.Fatal("Error while saving a tag: ", err)
	}

	fmt.Printf("\033[32;1m%s\033[0m\n", "完成")
	return nil
}

// Download 下载文件
func Download(title, fileUrl string) error {
	// 请求数据
	resp, err := http.Get(fileUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Printf("正在生成文件：【\033[37;1m%s\033[0m】 ", title)

	out, err := os.Create(title)
	if err != nil {
		fmt.Printf("\033[31;1m%s\033[0m\n", "失败"+err.Error())
		panic(err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Printf("\033[31;1m%s\033[0m\n", "失败"+err.Error())
		return err
	}

	fmt.Printf("\033[32;1m%s\033[0m\n", "完成")
	return nil
}

// SaveFile 保存文件
func SaveFile(title, content string) error {

	fmt.Printf("正在生成文件：【\033[37;1m%s\033[0m】 ", title)
	f, err := os.OpenFile(title, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("\033[31;1m%s\033[0m\n", "失败"+err.Error())
		return err
	}
	_, err = f.WriteString(content)
	if err != nil {
		fmt.Printf("\033[31;1m%s\033[0m\n", "失败"+err.Error())
		return err
	}
	if err = f.Close(); err != nil {
		return err
	}
	fmt.Printf("\033[32;1m%s\033[0m\n", "完成")
	return nil
}
