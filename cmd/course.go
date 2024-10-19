package cmd

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/yann0917/fs-dl/services"
	"github.com/yann0917/fs-dl/utils"
)

var courseCmd = &cobra.Command{
	Use:     "course",
	Short:   "查看课程列表",
	Long:    "使用 fs-dl course 查看课程列表, 加上课程 ID 可查看具体课程的节目列表",
	Example: "fs-dl course -s1 -p1 -l20 , fs-dl course 400003477",
	Args:    cobra.OnlyValidArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return courseList()
		} else {
			courseID = utils.String2Int(args[0])
			return programList(courseID)
		}
	},
}

func init() {
	rootCmd.AddCommand(courseCmd)
	courseCmd.Flags().IntVarP(&sort, "sort", "s", 1, "排序: 1-最新, 2-最热")
	courseCmd.Flags().IntVarP(&pageNo, "pageNo", "p", 1, "页码")
	courseCmd.Flags().IntVarP(&pageSize, "pageSize", "l", 10, "每页数量")
	courseCmd.Flags().IntSliceVarP(&classifyIds, "classifyIds", "c", []int{}, "分类 ID，全部可不写")
}

func courseList() (err error) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "课程ID", "标题", "简介", "讲书人", "讲数", "播放量"})
	table.SetAutoFormatHeaders(true)
	table.SetAutoWrapText(false)
	param := services.KnowledgeListParam{
		ViewType:    sort,
		PageNo:      pageNo,
		PageSize:    pageSize,
		CategoryIds: classifyIds,
	}
	list, err := Instance.KnowledgeList(param)
	if err != nil {
		return err
	}
	for i, data := range list {
		table.Append([]string{strconv.Itoa(i),
			data.Id,
			data.Title, data.Introduct,
			data.Author,
			utils.Int2String(data.TotalPublishNo),
			utils.Int2String(data.PlayCount),
		})
	}
	table.Render()
	return
}

func programList(courseID int) (err error) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "ID", "名称", "时长", "播放次数"})
	table.SetAutoFormatHeaders(true)
	table.SetAutoWrapText(false)
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
	for i, data := range list {
		table.Append([]string{strconv.Itoa(i),
			utils.Int2String(data.Id),
			data.Title,
			utils.FormatSeconds(data.Duration),
			utils.Int2String(data.ReadCount),
		})
	}
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.Render()
	return
}
