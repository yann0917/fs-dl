package services

// Knowledge 课程
type Knowledge struct {
	Author         string `json:"author"`
	BizType        int    `json:"bizType"`
	HasBuy         int    `json:"hasBuy"`
	Id             string `json:"id"`
	Introduct      string `json:"introduct"`
	PicUrl         string `json:"picUrl"`
	PlayCount      int    `json:"playCount"`
	Title          string `json:"title"`
	TotalPublishNo int    `json:"totalPublishNo"`
	WatermarkImage string `json:"watermarkImage"`
}

// CourseInfo 课程详情
type CourseInfo struct {
	ActualSaleScene []interface{} `json:"actualSaleScene"`
	AdditionCount   int           `json:"additionCount"`
	AlbumCoverUrl   string        `json:"albumCoverUrl"`
	Author          string        `json:"author"`
	AuthorIntro     string        `json:"authorIntro"`
	BizType         int           `json:"bizType"`
	BuyCount        int           `json:"buyCount"`
	CategoryName    string        `json:"categoryName"`
	CategoryNo      int           `json:"categoryNo"`
	CategoryNoList  []int         `json:"categoryNoList"`
	CounselorFlag   int           `json:"counselorFlag"`
	CouponList      []interface{} `json:"couponList"`
	CoverImage      string        `json:"coverImage"`
	Disabled        bool          `json:"disabled"`
	HasBought       bool          `json:"hasBought"`
	Id              int           `json:"id"`
	ImageText       string        `json:"imageText"`
	LearningInfo    LearningInfo  `json:"learningInfo"`
	OriginalPrice   string        `json:"originalPrice"`
	PackageList     []interface{} `json:"packageList"`
	Poster          bool          `json:"poster"`
	PublishedCount  int           `json:"publishedCount"`
	PurchaseImage   string        `json:"purchaseImage"`
	ReadCount       int           `json:"readCount"`
	ReadCountTotal  int           `json:"readCountTotal"`
	RelateChapters  int           `json:"relateChapters"`
	SellPrice       string        `json:"sellPrice"`
	Sharable        bool          `json:"sharable"`
	ShareImageUrl   string        `json:"shareImageUrl"`
	ShareLink       string        `json:"shareLink"`
	ShowType        int           `json:"showType"`
	SourceInfo      SourceInfo    `json:"sourceInfo"`
	SubTitle        string        `json:"subTitle"`
	SuitedPeople    string        `json:"suitedPeople"`
	Title           string        `json:"title"`
	TotalPublishNo  int           `json:"totalPublishNo"`
	TrialPrograms   int           `json:"trialPrograms"`
	Type            int           `json:"type"`
	UnlockNum       int           `json:"unlockNum"`
}

type LearningInfo struct {
	Finished      int    `json:"finished"`
	Learning      int    `json:"learning"`
	ShareText     string `json:"shareText"`
	TotalDuration int    `json:"totalDuration"`
}

type SourceInfo struct {
	Duration int `json:"duration"`
	Type     int `json:"type"`
}

// Program 节目
type Program struct {
	AlbumId           int          `json:"albumId"`
	AlbumName         string       `json:"albumName"`
	AudioUrl          string       `json:"audioUrl"` // 音频链接
	ChapterInfo       *ChapterInfo `json:"chapterInfo,omitempty"`
	Duration          int          `json:"duration"`
	Finished          int          `json:"finished"`
	FragmentId        int          `json:"fragmentId"`
	Free              bool         `json:"free"`
	Id                int          `json:"id"`
	IsLimitedTimeFree bool         `json:"isLimitedTimeFree"`
	MediaFilesize     int          `json:"mediaFilesize"`
	PublishTime       int64        `json:"publishTime"`
	ReadCount         int          `json:"readCount"`
	Seq               string       `json:"seq"` // 排序
	ShowType          int          `json:"showType"`
	Title             string       `json:"title"`
	TitleImageUrl     string       `json:"titleImageUrl"`
	TopFlag           int          `json:"topFlag"`
	Trial             bool         `json:"trial"`
	Unlock            bool         `json:"unlock"`
	UnlockType        int          `json:"unlockType,omitempty"`
	VideoFragmentId   int          `json:"videoFragmentId"`
}

type ChapterInfo struct {
	ChapterId   int    `json:"chapterId"`
	ChapterName string `json:"chapterName"`
	ChapterSeq  int    `json:"chapterSeq"`
	ProgramNum  int    `json:"programNum"`
}
