package services

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/yann0917/fs-dl/config"
	"github.com/yann0917/fs-dl/utils"
)

var (
	// AesKey AES ECB PKCS5Padding key length 256
	AesKey = config.Conf.AesKey
	Token  = config.Conf.Token
)

const (
	ApiNewBooks         = "/fs-webtool/web/book/v100/newBooks"
	ApiClassify         = "/fs-webtool/web/index/v100/classify"
	ApiLogin            = "/fs-member/user/auth/v100/login"
	ApiSendSms          = "/fs-oneid/verifyCode/api/sendSms"
	ApiListClassifyBook = "/fs-webtool/web/book/v100/listClassifyBook"
	// ApiBookContent book 详情
	ApiBookContent = "/fs-webtool/web/book/v101/content"
	// ApiCommentList book 评论
	ApiCommentList = "/fs-ugc/app/comment/v102/list"
	// ApiUserInfo 用户信息
	ApiUserInfo = "/fs-member/user/profile/v100/detail"
	// ApiBookPortalCategory 门户分类列表
	ApiBookPortalCategory = "/chief-orch/home/bookPortal/v106/category"
	// ApiClassifyBookList book 列表
	ApiClassifyBookList = "/resource-orchestration-system/book/classify/v100/listClassifyBook"

	// ApiKnowledgeList 课程列表
	ApiKnowledgeList = "/smart-orch/knowledge/v101/list"
	// ApiCourseInfo 课程详情
	ApiCourseInfo = "/smart-orch/course/v100/info"
	// ApiProgramList 课程下的节目列表
	ApiProgramList = "/smart-orch/program/v100/list"
	// ApiGiftCardConfig 礼品卡配置
	ApiGiftCardConfig = "fdtalk-orch/app/giftCard/v100/config"
	// ApiUserGiftCards 我的礼品卡列表
	ApiUserGiftCards = "/fdtalk-orch/app/giftCard/v100/userGiftCards"

	// ApiGiftCardDetail 礼品卡详情
	ApiGiftCardDetail = "fdtalk-orch/app/giftCard/v100/detail"
)

type Response struct {
	Data       respData `json:"data"`
	Msg        string   `json:"msg"`
	Status     string   `json:"status"` // "0000"-成功
	SystemMsg  string   `json:"systemMsg"`
	SystemTime int64    `json:"systemTime"`
}

type T2 struct {
	AreaCode               string `json:"areaCode"`
	Mobile                 string `json:"mobile"`
	VerificationCode       string `json:"verificationCode"`
	AuthType               int    `json:"authType"`
	VerificationCodeSource string `json:"verificationCodeSource"`
	PromotionInfo          struct {
		TrackingSourceId int `json:"trackingSourceId"`
	} `json:"promotionInfo"`
}
type T3 struct {
	Data []struct {
		BusinessName string `json:"businessName"`
		BusinessType int    `json:"businessType"`
		Cates        []struct {
			CateIds []int  `json:"cateIds"`
			Id      int    `json:"id"`
			Name    string `json:"name"`
		} `json:"cates"`
		Sort  int `json:"sort"`
		Years []struct {
			Year     int    `json:"year"`
			YearName string `json:"yearName"`
		} `json:"years"`
	} `json:"data"`
	Msg        string `json:"msg"`
	Status     string `json:"status"`
	SystemMsg  string `json:"systemMsg"`
	SystemTime int64  `json:"systemTime"`
}

// ClassifyBookParam 请求Book列表param
type ClassifyBookParam struct {
	SortType     int   `json:"sortType"`     // 1-最新, 2-最热
	BusinessType int   `json:"businessType"` // 1-樊登讲书, 2-非凡精读, 3-李蕾讲经典, 4-课程, 5-训练营, 6-电子书
	PageNo       int   `json:"pageNo"`       // 页码
	PageSize     int   `json:"pageSize"`     // 每页数量
	ClassifyIds  []int `json:"classifyIds"`  // 分类 Ids
	PublishYear  int   `json:"publishYear"`
}
type BookContentParam struct {
	BookId     int    `json:"bookId,omitempty"`
	Token      string `json:"token"`
	FragmentId int    `json:"fragmentId,omitempty"`
}

type CommentParam struct {
	ObjectSource      string `json:"objectSource,omitempty"`
	PageCount         string `json:"pageCount,omitempty"`
	ResourceId        string `json:"resourceId,omitempty"`
	PageNo            string `json:"pageNo,omitempty"`
	PageSize          string `json:"pageSize,omitempty"`
	TabType           int    `json:"tabType,omitempty"`       // 1-推荐, 2-最热, 3-最新
	BizObjectCode     string `json:"bizObjectCode,omitempty"` // fd_book_comment, ff_book_comment, ll_book
	TotalCount        string `json:"totalCount,omitempty"`
	SubResourceId     string `json:"subResourceId,omitempty"`
	ObjectSourceValue string `json:"objectSourceValue,omitempty"`
}

type T struct {
	BizObjectCode string `json:"bizObjectCode"`
	ResourceId    int    `json:"resourceId"`
	TabType       int    `json:"tabType"`
	PageSize      int    `json:"pageSize"`
}
type KnowledgeListParam struct {
	CategoryIds    []int `json:"categoryIds"`
	PageNo         int   `json:"pageNo"`
	ViewType       int   `json:"viewType"`
	PageSize       int   `json:"pageSize"`
	BookReadStatus int   `json:"bookReadStatus"`
}

type CourseInfoParam struct {
	ProgramCardId int `json:"programCardId"`
	CourseId      int `json:"courseId"`
}
type ProgramListParam struct {
	Page     ProgramPage `json:"page"`
	CourseId int         `json:"courseId"`
}

type ProgramPage struct {
	PageNo     int `json:"pageNo"`
	PageSize   int `json:"pageSize"`
	TotalCount int `json:"totalCount"`
}

type UserInfoParam struct {
	Token                string `json:"token"`
	IncludeBusinessTypes []int  `json:"includeBusinessTypes"`
}

type respData []byte

func (r *respData) UnmarshalJSON(data []byte) error {
	*r = data
	return nil
}

func (r *respData) String() string {
	return string(*r)
}

func (s *Service) GetUserInfo() (user UserInfo, err error) {
	param := UserInfoParam{
		IncludeBusinessTypes: []int{1, 2, 3},
		Token:                Token,
	}
	cipher, err := handleEncryptParam(param)
	if err != nil {
		return
	}
	resp, err := s.client.R().
		SetBody(cipher).
		Post(ApiUserInfo)
	reader, err := handleHTTPResponse(resp, err)
	if err != nil {
		return
	}
	err = handleJSONParse(reader, &user)
	return
}

func (s *Service) BookClassify() (list []Category, err error) {
	resp, err := s.client.R().
		Post(ApiClassify)
	reader, err := handleHTTPResponse(resp, err)
	if err != nil {
		return
	}
	err = handleJSONParse(reader, &list)
	return
}

//func (s *Service) BookPortalCategory() (list []Category, err error) {
//	resp, err := s.client.R().
//		Post(ApiBookPortalCategory)
//	reader, err := handleHTTPResponse(resp, err)
//	if err != nil {
//		return
//	}
//	err = handleJSONParse(reader, &list)
//	return
//}

// ClassifyBookList BookList
func (s *Service) ClassifyBookList(param ClassifyBookParam) (list []ClassifyBook, err error) {
	cipher, err := handleEncryptParam(param)
	if err != nil {
		return
	}
	resp, err := s.client.R().
		SetBody(cipher).
		Post(ApiListClassifyBook)
	reader, err := handleHTTPResponse(resp, err)
	if err != nil {
		return
	}
	err = handleJSONParse(reader, &list)
	return
}

// BookContent Book
func (s *Service) BookContent(bookId int) (detail BookContent, err error) {
	param := BookContentParam{
		BookId: bookId,
		Token:  Token,
	}
	cipher, err := handleEncryptParam(param)
	if err != nil {
		return
	}

	resp, err := s.client.R().
		SetBody(cipher).
		Post(ApiBookContent)
	reader, err := handleHTTPResponse(resp, err)
	if err != nil {
		return
	}
	err = handleJSONParse(reader, &detail)
	return
}

// BookModuleContent get articles:思维导图、文字稿
func (s *Service) BookModuleContent(bookId, fragmentId int) (detail BookContent, err error) {
	param := BookContentParam{
		BookId:     bookId,
		FragmentId: fragmentId,
		Token:      Token,
	}
	cipher, err := handleEncryptParam(param)
	if err != nil {
		return
	}
	resp, err := s.client.R().
		SetBody(cipher).
		Post(ApiBookContent)
	reader, err := handleHTTPResponse(resp, err)
	if err != nil {
		return
	}
	err = handleJSONParse(reader, &detail)
	return
}

func (s *Service) KnowledgeList(param KnowledgeListParam) (list []Knowledge, err error) {
	cipher, err := handleEncryptParam(param)
	if err != nil {
		return
	}
	resp, err := s.client.R().
		SetBody(cipher).
		Post(ApiKnowledgeList)
	reader, err := handleHTTPResponse(resp, err)
	if err != nil {
		return
	}
	err = handleJSONParse(reader, &list)
	return
}

func (s *Service) CourseInfo(param CourseInfoParam) (detail CourseInfo, err error) {
	cipher, err := handleEncryptParam(param)
	if err != nil {
		return
	}
	resp, err := s.client.R().
		SetBody(cipher).
		Post(ApiCourseInfo)
	reader, err := handleHTTPResponse(resp, err)
	if err != nil {
		return
	}
	err = handleJSONParse(reader, &detail)
	return
}

func (s *Service) ProgramList(param ProgramListParam) (list []Program, err error) {
	cipher, err := handleEncryptParam(param)
	if err != nil {
		return
	}
	resp, err := s.client.R().
		SetBody(cipher).
		Post(ApiProgramList)
	reader, err := handleHTTPResponse(resp, err)
	if err != nil {
		return
	}
	err = handleJSONParse(reader, &list)
	return
}

// handleEncryptParam 加密参数
func handleEncryptParam(param interface{}) (cipher string, err error) {
	data, err := utils.MarshalJSON(param)
	if err != nil {
		return
	}
	cipher, err = utils.AesEcbEncryptBase64(data, []byte(AesKey))
	if err != nil {
		return
	}
	return
}

func handleHTTPResponse(resp *resty.Response, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, errors.New("404 NotFound")
	}
	if resp.StatusCode() == http.StatusBadRequest {
		return nil, errors.New("400 BadRequest")
	}
	if resp.StatusCode() == http.StatusUnauthorized {
		return nil, errors.New("401 Unauthorized")
	}
	if resp.StatusCode() == 496 {
		return nil, errors.New("496 NoCertificate")
	}

	data := resp.Body()
	return data, nil
}

func handleJSONParse(reader []byte, v interface{}) error {
	result := new(Response)

	plainText, err := utils.AesEcbDecryptByBase64(string(reader), []byte(AesKey))
	if err != nil {
		fmt.Printf("err1: %s \n", err.Error())
		return err
	}

	err = utils.UnmarshalJSON(plainText, &result)
	if err != nil {
		fmt.Printf("err2: %s \n", err.Error())
		return err
	}
	if !result.isSuccess() {
		// 未登录或者登录凭证无效
		err = errors.New("服务异常，请稍后重试。errMsg:" + result.Msg + ", " + result.SystemMsg)
		return err
	}
	err = utils.UnmarshalJSON(result.Data, v)
	if err != nil {
		fmt.Printf("Unmarshal Data err: %s", err.Error())
		return err
	}

	return nil
}

func (r *Response) isSuccess() bool {
	return r.Status == "0000"
}
