package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/yann0917/fs-dl/services"
	"github.com/yann0917/fs-dl/utils"
)

var (
	businessType int
	bookID       int
	courseID     int
	sort         int
	pageNo       int
	pageSize     int
	publishYear  int
	classifyIds  []int
)

var classCmd = &cobra.Command{
	Use:     "class",
	Short:   "查看全部分类",
	Long:    `使用 fs-dl class 查看分类列表`,
	Example: "fs-dl class",
	Args:    cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return class()
	},
}

var contentCmd = &cobra.Command{
	Use:   "content",
	Short: "查看业务分类下的内容列表",
	Long: "使用 fs-dl content 查看业务分类下的内容列表, " +
		"可查看业务为【1-樊登讲书, 2-非凡精读, 3-李蕾讲经典】下的列表",
	Example: "fs-dl content -b1 -s2 -p1 -l20 -y0",
	Args:    cobra.OnlyValidArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return contentList()
	},
}

var bookInfoCmd = &cobra.Command{
	Use:     "book",
	Short:   "查看某个课程的详情",
	Long:    "使用 fs-dl book 查看某个课程的详情",
	Example: "fs-dl book 123",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("请输入课程ID")
		}
		bookID = utils.String2Int(args[0])
		return bookInfo(bookID)
	},
}

func init() {
	rootCmd.AddCommand(classCmd)
	rootCmd.AddCommand(contentCmd)
	rootCmd.AddCommand(bookInfoCmd)
	contentCmd.Flags().IntVarP(&businessType, "businessType", "b", 1, "业务: 1-樊登讲书, 2-非凡精读, 3-李蕾讲经典")
	contentCmd.Flags().IntVarP(&sort, "sort", "s", 1, "排序: 1-最新, 2-最热")
	contentCmd.Flags().IntVarP(&pageNo, "pageNo", "p", 1, "页码")
	contentCmd.Flags().IntVarP(&pageSize, "pageSize", "l", 10, "每页数量")
	contentCmd.Flags().IntVarP(&publishYear, "publishYear", "y", 0, "发布年份，例如 2024，可不写")
}

func class() (err error) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"业务 ID", "业务名称", "分类 ID", "分类名称", "IDS"})
	table.SetAutoFormatHeaders(true)
	table.SetAutoWrapText(false)

	list, err := Instance.BookClassify()
	if err != nil {
		return
	}
	for _, c := range list {
		for _, cate := range c.Cates {
			var ids []string
			for _, id := range cate.CateIds {
				ids = append(ids, utils.Int2String(id))
			}
			table.Append([]string{
				utils.Int2String(c.BusinessType),
				c.BusinessName,
				utils.Int2String(cate.Id),
				cate.Name,
				strings.Join(ids, ","),
			})
		}
	}
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.Render()
	return
}

func contentList() (err error) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "课程ID", "标题", "简介", "主讲人", "播放量", "上线日期"})
	table.SetAutoFormatHeaders(true)
	table.SetAutoWrapText(false)
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
	for i, data := range list {
		table.Append([]string{strconv.Itoa(i),
			utils.Int2String(data.BookId),
			data.Title, data.Summary,
			data.SpeakerName,
			utils.Int2String(data.PlayCount),
			utils.UnixMilli2DateString(data.PublishTime),
		})
	}

	table.Render()
	return
}

func bookInfo(bookID int) (err error) {
	out := os.Stdout
	table := tablewriter.NewWriter(out)

	detail, err := Instance.BookContent(bookID)
	if err != nil {
		return err
	}

	table.SetHeader([]string{"bookID", "书名", "评分", "主讲人", "上线日期"})
	table.SetAutoFormatHeaders(true)
	table.SetAutoWrapText(false)
	table.Append([]string{
		utils.Int2String(detail.BookInfo.BookId),
		detail.BookInfo.Title,
		detail.BookInfo.Score,
		detail.BookInfo.SpeakerName,
		utils.UnixMilli2DateString(detail.BookInfo.PublishTime),
	})
	table.Render()

	_, _ = fmt.Fprint(out, detail.RecommendVO.RecommendName)
	_, _ = fmt.Fprintln(out)
	_, _ = fmt.Fprint(out, detail.RecommendVO.RecommendInfo)
	_, _ = fmt.Fprintln(out)
	_, _ = fmt.Fprintln(out)
	_, _ = fmt.Fprintln(out, "你将获得\n")

	for _, info := range detail.Acquire.Intros {
		_, _ = fmt.Fprintln(out, info)
	}
	_, _ = fmt.Fprintln(out)
	_, _ = fmt.Fprintln(out)
	_, _ = fmt.Fprintln(out, "作者简介\n")

	for _, info := range detail.Authors {
		_, _ = fmt.Fprintln(out, info.Name)
		_, _ = fmt.Fprintln(out, info.Summary)
	}
	_, _ = fmt.Fprintln(out)
	_, _ = fmt.Fprintln(out)
	_, _ = fmt.Fprintln(out, "精彩选段\n")
	for _, info := range detail.Extract.Infos {
		_, _ = fmt.Fprintln(out, info.Intro)
		_, _ = fmt.Fprintln(out)

	}
	_, _ = fmt.Fprintln(out)

	//
	//// 评论
	//param := services.CommentParam{
	//	ResourceId:    utils.Int2String(bookID),
	//	PageNo:        "1",
	//	PageSize:      "10",
	//	TabType:       1,
	//	BizObjectCode: detail.BizObjectCode,
	//}
	//commentList, err := Instance.CommentList(param)
	//if err != nil {
	//	return err
	//}
	//
	//_, _ = fmt.Fprintln(out, "评论\n")
	//for _, data := range commentList.CommentList {
	//	_, _ = fmt.Fprintln(out, "👤"+data.CommentUserName+
	//		": "+data.CommentTime+
	//		"|📍"+data.IpAddress+
	//		"|👍 "+data.LikeCount+
	//		"|💬 "+data.RepliedTotalCount)
	//	_, _ = fmt.Fprintln(out, data.CommentContent)
	//	_, _ = fmt.Fprintln(out)
	//}
	return
}
