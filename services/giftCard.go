package services

type UserGiftCards struct {
	ExpiredText   string     `json:"expiredText"`
	GiftCards     []GiftCard `json:"giftCards"`
	ReceivedText  string     `json:"receivedText"`
	SendableCount int        `json:"sendableCount"`
}

type GiftCard struct {
	CardId          int64  `json:"cardId"`
	CardTypeName    string `json:"cardTypeName"`
	Code            string `json:"code"`
	CoverId         string `json:"coverId"`
	ExpireDate      int64  `json:"expireDate"`
	ImgUrl          string `json:"imgUrl"`
	IsRepeat        bool   `json:"isRepeat"`
	ListenMoreFlag  bool   `json:"listenMoreFlag"`
	Name            string `json:"name"`
	OwnerRewardText string `json:"ownerRewardText"`
	PlatformSended  bool   `json:"platformSended"`
	SelfUsable      bool   `json:"selfUsable"`
	Status          int    `json:"status"`
	StyleId         string `json:"styleId"`
	Tid             string `json:"tid"`
	TimeUnit        int    `json:"timeUnit"`
	TimeValue       int    `json:"timeValue"`
	UseStatus       int    `json:"useStatus"`
}

type GiftCardConfig struct {
	Name    string `json:"name"`
	SkuType string `json:"skuType"`
	Tid     int    `json:"tid"`
}

type GiftCardDetail struct {
	CardId                  int64  `json:"cardId"`
	CardNumber              string `json:"cardNumber"`
	Code                    string `json:"code"`
	CollectingTitle         string `json:"collectingTitle"`
	CompletedBooksCountText string `json:"completedBooksCountText"`
	CoverId                 string `json:"coverId"`
	CoverImg                string `json:"coverImg"`
	CreateDate              int64  `json:"createDate"`
	ExpireDate              int64  `json:"expireDate"`
	JoinedDaysText          string `json:"joinedDaysText"`
	ListenMoreFlag          bool   `json:"listenMoreFlag"`
	Name                    string `json:"name"`
	OwnerRewardText         string `json:"ownerRewardText"`
	PlatformSended          bool   `json:"platformSended"`
	PurchaseDate            int64  `json:"purchaseDate"`
	Repeat                  bool   `json:"repeat"`
	SelfUsable              bool   `json:"selfUsable"`
	SenderName              string `json:"senderName"`
	ShareContent            string `json:"shareContent"`
	ShareImg                string `json:"shareImg"`
	ShareLink               string `json:"shareLink"`
	ShareTitle              string `json:"shareTitle"`
	Status                  int    `json:"status"`
	StyleId                 string `json:"styleId"`
	TimeUnit                int    `json:"timeUnit"`
	TimeValue               int    `json:"timeValue"`
	Tips                    string `json:"tips"`
	UseStatus               int    `json:"useStatus"`
	UserId                  int    `json:"userId"`
	WishMessage             string `json:"wishMessage"`
}
