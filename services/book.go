package services

// Category 全部分类
type Category struct {
	BusinessName string `json:"businessName"`
	BusinessType int    `json:"businessType"`
	BusinessZone int    `json:"businessZone"`
	Sort         int    `json:"sort"`
	Cates        []Cate `json:"cates"`
	Years        []Year `json:"years"`
}

type Cate struct {
	CateIds []int  `json:"cateIds"`
	Id      int    `json:"id"`
	Name    string `json:"name"`
}

type Year struct {
	Year     int    `json:"year"`
	YearName string `json:"yearName"`
}
type BookTag struct {
	Id      int    `json:"id"`
	Key     string `json:"key,omitempty"`
	Name    string `json:"name"`
	TagType int    `json:"tagType"`
}

type ClassifyBook struct {
	BookId         int    `json:"bookId"`
	BusinessType   int    `json:"businessType"`
	CoverImg       string `json:"coverImg"`
	PlayCount      int    `json:"playCount"`
	PublishTime    int64  `json:"publishTime"`
	Score          string `json:"score"`
	SpeakerName    string `json:"speakerName"`
	Summary        string `json:"summary"`
	Title          string `json:"title"`
	WatermarkImage string `json:"watermarkImage"`
}

type Book struct {
	BookId         int    `json:"bookId"`
	BusinessType   int    `json:"businessType"`
	CoverImg       string `json:"coverImg"`
	PlayCount      int    `json:"playCount"`
	PublishTime    int64  `json:"publishTime"`
	Score          string `json:"score"`
	SpeakerName    string `json:"speakerName"`
	Summary        string `json:"summary"`
	Title          string `json:"title"`
	WatermarkImage string `json:"watermarkImage"`
}

type Rank struct {
	Num      int    `json:"num"`
	RankType int    `json:"rankType"`
	TagId    int    `json:"tagId"`
	TagName  string `json:"tagName"`
}

type Author struct {
	Name    string   `json:"name"`
	Summary string   `json:"summary"`
	Tags    []string `json:"tags"`
}

type T4 struct {
	Data struct {
		Acquire struct {
			Intros []string `json:"intros"`
		} `json:"acquire"`
		Articles []struct {
			FragmentId int    `json:"fragmentId"`
			ModuleCode string `json:"moduleCode"`
			ModuleName string `json:"moduleName"`
		} `json:"articles"`
		AudioInfo struct {
			Duration                  int `json:"duration"`
			FragmentId                int `json:"fragmentId"`
			LoudnessNormalizationInfo struct {
			} `json:"loudnessNormalizationInfo"`
			MediaCoverUrl            string `json:"mediaCoverUrl"`
			MediaFilesize            int    `json:"mediaFilesize"`
			MediaUrl                 string `json:"mediaUrl"`
			TrialCompletedButtonText string `json:"trialCompletedButtonText"`
			TrialCompletedText       string `json:"trialCompletedText"`
			TrialDuration            int    `json:"trialDuration"`
		} `json:"audioInfo"`
		Authors []struct {
			Name    string   `json:"name"`
			Summary string   `json:"summary"`
			Tags    []string `json:"tags"`
		} `json:"authors"`
		BizObjectCode string `json:"bizObjectCode"`

		BookRights struct {
			Free      bool `json:"free"`
			RightType int  `json:"rightType"`
			Trial     bool `json:"trial"`
		} `json:"bookRights"`
		Extract struct {
			Infos []struct {
				Intro string `json:"intro"`
				Page  string `json:"page"`
			} `json:"infos"`
		} `json:"extract"`
		VideoInfo struct {
			Duration                  int `json:"duration"`
			FragmentId                int `json:"fragmentId"`
			LoudnessNormalizationInfo struct {
			} `json:"loudnessNormalizationInfo"`
			MediaCoverUrl            string `json:"mediaCoverUrl"`
			MediaFilesize            int    `json:"mediaFilesize"`
			MediaUrl                 string `json:"mediaUrl"`
			TrialCompletedButtonText string `json:"trialCompletedButtonText"`
			TrialCompletedText       string `json:"trialCompletedText"`
			TrialDuration            int    `json:"trialDuration"`
		} `json:"videoInfo"`
	} `json:"data"`
	Msg        string `json:"msg"`
	Status     string `json:"status"`
	SystemMsg  string `json:"systemMsg"`
	SystemTime int64  `json:"systemTime"`
}
type BookContent struct {
	Acquire struct {
		Intros []string `json:"intros"`
		Title  string   `json:"title"`
	} `json:"acquire"`
	BizObjectCode string `json:"bizObjectCode"`
	BookRights    struct {
		Free      bool `json:"free"`
		RightType int  `json:"rightType"`
		Trial     bool `json:"trial"`
	} `json:"bookRights"`

	Authors           []Author      `json:"authors"`
	Articles          []Article     `json:"articles"`
	AudioInfo         AudioInfo     `json:"audioInfo"`
	VideoInfo         AudioInfo     `json:"videoInfo"`
	BarPoints         []BarPoint    `json:"barPoints"`
	BookComponent     BookComponent `json:"bookComponent"`
	BookInfo          BookInfo      `json:"bookInfo"`
	BusinessIntroduce string        `json:"businessIntroduce"`
	BusinessType      int           `json:"businessType"`
	ButtonCloseDay    int           `json:"buttonCloseDay"`
	ButtonSharingFlag bool          `json:"buttonSharingFlag"`
	ButtonSharingText string        `json:"buttonSharingText"`
	CommentCount      int           `json:"commentCount"`
	CpsTitle          string        `json:"cpsTitle"`
	Content           string        `json:"content"` // 富文本内容
	Ebook             Ebook         `json:"ebook"`
	Extract           struct {
		Infos []struct {
			Intro string `json:"intro"`
		} `json:"infos"`
	} `json:"extract"`
	FloatingLayerText  string      `json:"floatingLayerText"`
	FragmentId         int         `json:"fragmentId"`
	Free               bool        `json:"free"`
	HasBought          bool        `json:"hasBought"`
	HasDownloaded      bool        `json:"hasDownloaded"`
	IsFavorite         bool        `json:"isFavorite"`
	Label              int         `json:"label"`
	ListenMinute       int         `json:"listenMinute"`
	MemberOnly         bool        `json:"memberOnly"`
	ModuleList         []Module    `json:"moduleList"`
	PaperBook          PaperBook   `json:"paperBook"`
	PolishFlow         PolishFlow  `json:"polishFlow"`
	RankVO             RankVO      `json:"rankVO"`
	ReceiveFlag        bool        `json:"receiveFlag"`
	RecommendBook      []BookInfo  `json:"recommendBook"`
	RecommendVO        RecommendVO `json:"recommendVO"`
	RightEndTime       int64       `json:"rightEndTime"`
	RightType          int         `json:"rightType"`
	ShowFrequency      int         `json:"showFrequency"`
	ShowVideo          bool        `json:"showVideo"`
	Speakers           []Speaker   `json:"speakers"`
	TopicVO            TopicVO     `json:"topicVO"`
	Trial              bool        `json:"trial"`
	Type               int         `json:"type"`
	UnlockedText       string      `json:"unlockedText"`
	UserCompletedCount int         `json:"userCompletedCount"`
}

type Article struct {
	FragmentId       int    `json:"fragmentId"`
	FragmentIdEncode string `json:"fragmentIdEncode"`
	ModuleCode       string `json:"moduleCode"`
	ModuleImg1       string `json:"moduleImg1"`
	ModuleName       string `json:"moduleName"`
	Op               string `json:"op"`
	ShowFlag         int    `json:"showFlag"`
	Tid              int    `json:"tid"`
	TransId          int    `json:"transId"`
	Type             int    `json:"type"`
	View             string `json:"view"`
}

type AudioInfo struct {
	Duration                  int `json:"duration"`
	FragmentId                int `json:"fragmentId"`
	LoudnessNormalizationInfo struct {
	} `json:"loudnessNormalizationInfo"`
	MediaCoverUrl            string `json:"mediaCoverUrl"`
	MediaFilesize            int    `json:"mediaFilesize"`
	MediaUrl                 string `json:"mediaUrl"`
	TrialCompletedButtonText string `json:"trialCompletedButtonText"`
	TrialCompletedText       string `json:"trialCompletedText"`
	TrialDuration            int    `json:"trialDuration"`
}

type Speaker struct {
	BookNum         int           `json:"bookNum"`
	Crypto          string        `json:"crypto"`
	HeadImageRawUrl string        `json:"headImageRawUrl"`
	Id              string        `json:"id"`
	Name            string        `json:"name"`
	OldId           string        `json:"oldId"`
	ReadCount       int           `json:"readCount"`
	SpeakerBooks    []SpeakerBook `json:"speakerBooks"`
	Summary         string        `json:"summary"`
	Type            int           `json:"type"`
	UserId          int           `json:"userId"`
	HeaderImageUrl  string        `json:"headerImageUrl"`
}

type Module struct {
	KeyWord  string `json:"keyWord"`
	Name     string `json:"name"`
	ShowFlag bool   `json:"showFlag"`
	Sort     int    `json:"sort"`
}

type PaperBook struct {
	ModuleCode string `json:"moduleCode"`
	ModuleName string `json:"moduleName"`
	Op         string `json:"op"`
	Share      bool   `json:"share"`
	Spuid      string `json:"spuid"`
	TransId    int    `json:"transId"`
	Url        string `json:"url"`
}

type PolishFlow struct {
	Persons []Person `json:"persons"`
	Title   string   `json:"title"`
}

type Person struct {
	Flow  string `json:"flow"`
	Image string `json:"image"`
	Name  string `json:"name"`
}

type SpeakerBook struct {
	BookId         int    `json:"bookId"`
	BookName       string `json:"bookName"`
	BookNum        int    `json:"bookNum"`
	CoverIcon      string `json:"coverIcon"`
	ReadCount      int    `json:"readCount"`
	SpeakerId      int    `json:"speakerId"`
	WatermarkImage string `json:"watermarkImage"`
}

type BarPoint struct {
	Description string `json:"description"`
	Time        int    `json:"time"`
}

type BookComponent struct {
	CompAbstract string `json:"compAbstract"`
	CompBanner   string `json:"compBanner"`
	CompComment  string `json:"compComment"` // 获取评论列表需要的参数
}

type BookInfo struct {
	BookId         int    `json:"bookId"`
	BusinessType   int    `json:"businessType"`
	CoverImg       string `json:"coverImg"`
	PlayCount      int    `json:"playCount"`
	PublishTime    int64  `json:"publishTime"`
	Score          string `json:"score"`
	SpeakerName    string `json:"speakerName"`
	Summary        string `json:"summary"`
	Title          string `json:"title"`
	WatermarkImage string `json:"watermarkImage"`
}

type Ebook struct {
	EbookId    string `json:"ebookId"`
	ModuleCode string `json:"moduleCode"`
	ModuleImg1 string `json:"moduleImg1"`
	ModuleName string `json:"moduleName"`
	Op         string `json:"op"`
	ShowFlag   int    `json:"showFlag"`
	Tid        int    `json:"tid"`
	View       string `json:"view"`
}

type RankVO struct {
	DimensionType  int    `json:"dimensionType"`
	DimensionValue string `json:"dimensionValue"`
	IsAll          bool   `json:"isAll"`
	RankCode       string `json:"rankCode"`
	RankDesc       string `json:"rankDesc"`
	RankName       string `json:"rankName"`
	Ranking        int    `json:"ranking"`
}

type RecommendVO struct {
	RecommendInfo string `json:"recommendInfo"`
	RecommendName string `json:"recommendName"`
}

type TopicVO struct {
	TopicContent string `json:"topicContent"`
}

type ClassifySpeaker struct {
	Icon string `json:"icon"`
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CommentList struct {
	CommentList []Comment `json:"commentList"`
	Page        Page      `json:"page"`
}

type Comment struct {
	ActivityName      string `json:"activityName"`
	BizObjectCode     string `json:"bizObjectCode"`
	CommentContent    string `json:"commentContent"`
	CommentId         int    `json:"commentId"`
	CommentTime       string `json:"commentTime"`
	CommentUserId     int    `json:"commentUserId"`
	CommentUserImg    string `json:"commentUserImg,omitempty"`
	CommentUserName   string `json:"commentUserName"`
	CreateTime        int64  `json:"createTime"`
	Crypto            string `json:"crypto"`
	HasSpeaker        bool   `json:"hasSpeaker"`
	LikeCount         string `json:"likeCount"`
	MarkLike          bool   `json:"markLike"`
	Marked            bool   `json:"marked"`
	RepliedTotalCount string `json:"repliedTotalCount,omitempty"`
	SpeakerMarkLike   bool   `json:"speakerMarkLike"`
	CommentIP         string `json:"commentIP,omitempty"`
	IpAddress         string `json:"ipAddress,omitempty"`
	Status            int    `json:"status"`
	TopicId           int    `json:"topicId,omitempty"`
	TopicTitle        string `json:"topicTitle,omitempty"`
}

type Page struct {
	PageCount  string `json:"pageCount"`
	PageNo     string `json:"pageNo"`
	PageSize   string `json:"pageSize"`
	TotalCount string `json:"totalCount"`
}
