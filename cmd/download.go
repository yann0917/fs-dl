package cmd

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
	"github.com/yann0917/fs-dl/services"
	"github.com/yann0917/fs-dl/utils"
)

var OutputDir = "output"
var downloadType int

var downloadCmd = &cobra.Command{
	Use:   "dl",
	Short: "下载【樊登讲书, 非凡精读, 李蕾讲经典】的音频或者文稿",
	Long: `使用 fs-dl dl 下载【樊登讲书, 非凡精读, 李蕾讲经典】的音频或者文稿，
可以下载指定ID的内容，或者批量下载`,
	Example: "fs-dl dl 123 -b1 -t1, fs-dl dl -b1 -t1 -p1 -l20",
	Args:    cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			bookID = utils.String2Int(args[0])
			return Download(bookID)
		}
		return DownloadBatch()
	},
}

var downloadCourseCmd = &cobra.Command{
	Use:     "dlc",
	Short:   "下载【课程】的音频",
	Long:    `使用 fs-dl dlc 下载【课程】的音频`,
	Example: "fs-dl dlc 123",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("请填写课程 ID")
		}
		courseID = utils.String2Int(args[0])
		return DownloadCourse(courseID)
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
	rootCmd.AddCommand(downloadCourseCmd)
	downloadCmd.Flags().IntVarP(&businessType, "businessType", "b", 1, "业务: 1-樊登讲书, 2-非凡精读, 3-李蕾讲经典")
	downloadCmd.PersistentFlags().IntVarP(&downloadType, "downloadType", "t", 1, "下载格式, 1-mp3, 2-视频,  3-markdown文档, 4-PDF文档, 5-思维导图jpeg")
	downloadCmd.Flags().IntVarP(&sort, "sort", "s", 1, "排序: 1-最新, 2-最热")
	downloadCmd.Flags().IntVarP(&pageNo, "pageNo", "p", 1, "页码")
	downloadCmd.Flags().IntVarP(&pageSize, "pageSize", "l", 10, "每页数量")
}

func getSubDir(bType int) string {
	list := map[int]string{
		1: "樊登讲书",
		2: "非凡精读",
		3: "李蕾讲经典",
		4: "课程",
	}
	if t, ok := list[bType]; ok {
		return t
	} else {
		return "樊登讲书"
	}
}

func getFileSuffix(dType int) string {
	list := map[int]string{
		1: "mp3",
		2: "mp4",
		3: "md",
		4: "pdf",
		5: "jpeg",
	}
	return list[dType]
}

func DownloadBatch() (err error) {
	param := services.ClassifyBookParam{
		SortType:     sort,
		BusinessType: businessType,
		PageNo:       pageNo,
		PageSize:     pageSize,
		ClassifyIds:  classifyIds,
	}
	list, err := Instance.ClassifyBookList(param)
	if err != nil {
		return err
	}
	for _, book := range list {
		err = Download(book.BookId)
	}
	return nil
}

func Download(bookID int) (err error) {
	detail, err := Instance.BookContent(bookID)
	if err != nil {
		return
	}
	subDir := getSubDir(businessType)
	fileSuffix := getFileSuffix(downloadType)
	filePath, err := utils.Mkdir(OutputDir, utils.FileName(subDir, ""))
	if err != nil {
		return
	}

	bookName := strings.TrimSpace(detail.BookInfo.Title)
	fileName := filepath.Join(filePath, utils.FileName(utils.Int2String(bookID)+"."+bookName, fileSuffix))
	if utils.CheckFileExist(fileName) {
		fmt.Printf("【\033[37;1m%s\033[0m】已存在\n", fileName)
		return
	}

	articleFragmentId, thinkFragmentId := 0, 0
	for _, article := range detail.Articles {
		if article.ModuleCode == "articles" {
			articleFragmentId = article.FragmentId
		}
		if article.ModuleCode == "think" {
			thinkFragmentId = article.FragmentId
		}
	}

	switch downloadType {
	case 1:
		// 获取封面图
		var coverBytes []byte
		if detail.AudioInfo.MediaCoverUrl != "" {
			cover, err1 := http.Get(detail.AudioInfo.MediaCoverUrl)
			if err1 != nil {
				err = err1
				return
			}
			defer cover.Body.Close()
			coverBytes, err = io.ReadAll(cover.Body)
			if err != nil {
				return
			}
		}
		rawURL := detail.AudioInfo.MediaUrl
		if rawURL != "" {
			ext, _ := utils.GetUrlExt(rawURL)
			switch ext {
			case ".mp3":
				var opt utils.ID3Options
				opt.Artist = detail.BookInfo.SpeakerName
				opt.Title = bookName
				opt.Album = getSubDir(detail.BookInfo.BusinessType)
				opt.Cover = coverBytes
				err = utils.DownloadAudio(fileName, rawURL, opt)
			case ".m3u8":
				err = utils.MergeAudioAndVideo([]string{rawURL}, fileName)
				if err != nil {
					fmt.Println(rawURL)
					fmt.Println(err)
				}

			default:
				fmt.Println(rawURL)
			}
		}
	case 2:
		rawURL := detail.VideoInfo.MediaUrl
		err = utils.MergeAudioAndVideo([]string{rawURL}, fileName)
		if err != nil {
			fmt.Println(rawURL)
			fmt.Println(err)
		}
	case 3:
		if articleFragmentId > 0 {
			module, err1 := Instance.BookModuleContent(bookID, articleFragmentId)
			if err1 != nil {
				return err1
			}
			res := utils.Html2Md(module.Content)
			err = utils.SaveFile(fileName, res)
		} else {
			fmt.Printf("【\033[31;1m%s\033[0m】无解读文稿\n", bookName)
		}

	case 4:
		if articleFragmentId > 0 {
			module, err1 := Instance.BookModuleContent(bookID, articleFragmentId)
			if err1 != nil {
				return err1
			}
			doc, _ := goquery.NewDocumentFromReader(strings.NewReader(module.Content))
			text := doc.Find("div.rich_media_content").Map(func(i int, s *goquery.Selection) string {
				s.Find("section").Each(func(i int, item *goquery.Selection) {
					replaceLetterSpacing(item)
				})
				s.Find("p").Each(func(index int, item *goquery.Selection) {
					replaceLetterSpacing(item)
				})
				s.Find("span").Each(func(index int, item *goquery.Selection) {
					replaceLetterSpacing(item)
				})
				res, _ := s.Html()
				return res
			})
			err = utils.Html2Pdf(fileName, bookName, text[0])
			if err != nil {
				return err
			}
		} else {
			fmt.Printf("【\033[31;1m%s\033[0m】无解读文稿\n", bookName)
		}
	case 5:
		if thinkFragmentId > 0 {
			module, err1 := Instance.BookModuleContent(bookID, thinkFragmentId)
			if err1 != nil {
				return err1
			}
			doc, _ := goquery.NewDocumentFromReader(strings.NewReader(module.Content))
			name := utils.Int2String(bookID) + "." + bookName
			doc.Find("img").Each(func(i int, selection *goquery.Selection) {
				if i > 0 {
					fileName = filepath.Join(filePath, utils.FileName(name+"_"+utils.Int2String(i), fileSuffix))
				} else {
					fileName = filepath.Join(filePath, utils.FileName(name, fileSuffix))
				}
				if src, ok := selection.Attr("src"); ok {
					err = utils.Download(fileName, src)
				}
			})
		} else {
			fmt.Printf("【\033[31;1m%s\033[0m】无思维导图\n", bookName)
		}
	}

	return
}

// DownloadCourse 下载课程音频
func DownloadCourse(courseID int) (err error) {
	cParam := services.CourseInfoParam{
		ProgramCardId: 0,
		CourseId:      courseID,
	}
	detail, err := Instance.CourseInfo(cParam)
	if err != nil {
		return err
	}
	albumName, titleImageUrl := detail.Title, detail.AlbumCoverUrl

	param := services.ProgramListParam{
		Page: services.ProgramPage{
			PageNo: 1, PageSize: 1000,
		},
		CourseId: courseID,
	}
	list, err := Instance.ProgramList(param)
	if err != nil {
		return err
	}

	// 获取封面图
	var coverBytes []byte
	if titleImageUrl != "" {
		cover, err1 := http.Get(titleImageUrl)
		if err1 != nil {
			return err1
		}
		defer cover.Body.Close()
		coverBytes, err = io.ReadAll(cover.Body)
		if err != nil {
			return err
		}
	}

	subDir := getSubDir(4)
	fileSuffix := getFileSuffix(1)
	filePath, err := utils.Mkdir(OutputDir, utils.FileName(subDir, ""), utils.FileName(albumName, ""))
	if err != nil {
		return err
	}

	for _, program := range list {
		seq := program.Seq
		if program.ChapterInfo != nil {
			seq = utils.Int2String(program.ChapterInfo.ChapterSeq) + "-" + program.Seq
		}
		title := strings.TrimSpace(program.Title)
		fileName := filepath.Join(filePath, utils.FileName(seq+"."+title, fileSuffix))
		if utils.CheckFileExist(fileName) {
			fmt.Printf("【\033[37;1m%s\033[0m】已存在\n", fileName)
			continue
		}
		rawURL := program.AudioUrl
		if rawURL != "" {
			// 获取文件的扩展名
			ext, _ := utils.GetUrlExt(rawURL)
			switch ext {
			case ".mp3":
				var opt utils.ID3Options
				opt.Artist = detail.Author
				opt.Title = title
				opt.Album = albumName
				opt.Cover = coverBytes
				err = utils.DownloadAudio(fileName, rawURL, opt)
			case ".m3u8":
				err = utils.MergeAudioAndVideo([]string{rawURL}, fileName)
				if err != nil {
					fmt.Println(rawURL)
					fmt.Println(err)
				}

			default:
				fmt.Println(rawURL)
			}
		}
	}
	return
}

// replaceLetterSpacing
func replaceLetterSpacing(s *goquery.Selection) {
	if style, exists := s.Attr("style"); exists {
		styles := strings.Split(style, ";")
		// letter-spacing 导致 wkhtmltopdf 文字被截断
		for i, s := range styles {
			keyValue := strings.Split(s, ":")
			if len(keyValue) == 2 {
				key := strings.TrimSpace(keyValue[0])
				value := strings.TrimSpace(keyValue[1])
				if key == "letter-spacing" {
					value = "0.3px"
					styles[i] = fmt.Sprintf("%s: %s", key, value)
					break
				}
			}
		}
		// 重建style字符串
		newStyle := strings.Join(styles, ";")
		if !strings.HasSuffix(newStyle, ";") {
			newStyle += ";" // 确保以分号结尾
		}
		s.SetAttr("style", newStyle)
	}
}
