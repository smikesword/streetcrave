package model


type Business struct {
	Id int `json:"id"`
	OwnerId int `json:"owner_id"`
	Type string `json:"business_type"`
	MediaUrls []string `json:"media_urls"`
}



