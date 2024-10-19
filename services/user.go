package services

type T5 struct {
	AreaCode      string `json:"areaCode"`
	Mobile        string `json:"mobile"`
	AuthToken     string `json:"authToken"`
	AuthType      int    `json:"authType"` // 5
	PromotionInfo struct {
		TrackingSourceId int `json:"trackingSourceId"`
	} `json:"promotionInfo"`
}

type T6 struct {
	Mobile   string `json:"mobile"`   // 176xxxx0917
	AreaCode string `json:"areaCode"` // +86
	DeviceId string `json:"deviceId"` // sbeqam9nb2g
}
type UserInfo struct {
	AccountBalance     float64 `json:"accountBalance"`
	AreaCode           string  `json:"areaCode"`
	AvatarUrl          string  `json:"avatarUrl"`
	Belong             string  `json:"belong"`
	BindWeChat         bool    `json:"bindWeChat"`
	BirthDate          string  `json:"birthDate"`
	CityBlong          string  `json:"cityBlong"`
	Counselor          bool    `json:"counselor"`
	CurrentBonus       int     `json:"currentBonus"`
	CurrentDateTime    int64   `json:"currentDateTime"`
	Education          string  `json:"education"`
	Email              string  `json:"email"`
	EncryptedUid       string  `json:"encryptedUid"`
	ExpireTime         int64   `json:"expire_time"`
	FirstLoginTime     int64   `json:"firstLoginTime"`
	Gender             int     `json:"gender"`
	Interest           string  `json:"interest"`
	IsTrial            bool    `json:"isTrial"`
	IsTemp             bool    `json:"is_temp"`
	IsVip              bool    `json:"is_vip"`
	MaritalStatus      int     `json:"maritalStatus"`
	Mobile             string  `json:"mobile"`
	Occupation         string  `json:"occupation"`
	Point              int     `json:"point"`
	ProBlong           string  `json:"proBlong"`
	RegisterSourceName string  `json:"registerSourceName"`
	RegisterTime       int64   `json:"registerTime"`
	StatusInVipCode    int     `json:"statusInVipCode"`
	StatusInVipSwitch  int     `json:"statusInVipSwitch"`
	Uid                int     `json:"uid"`
	UserPromoType      int     `json:"userPromoType"`
	UserRoleCode       string  `json:"userRoleCode"`
	UserStatus         int     `json:"userStatus"`
	Username           string  `json:"username"`
	VipBeginTime       int64   `json:"vipBeginTime"`
	WxSubscribedStatus int     `json:"wxSubscribedStatus"`
}

type LearningRank struct {
	EncodeUserId           string `json:"encodeUserId"`
	HeadImg                string `json:"headImg"`
	LikeCount              int    `json:"likeCount"`
	Liked                  bool   `json:"liked"`
	Rank                   int    `json:"rank"`
	UserName               string `json:"userName"`
	WeeklyLearningDuration int    `json:"weeklyLearningDuration"`
}
