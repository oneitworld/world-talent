package models

type Audit struct {
	ID           int64  `json:"id" gorm:"primary_key;auto_increment"`
	Datetime     string `json:"dateTime" gorm:"size:50;not null;`
	IPAddress    string `json:"idAddress" gorm:"size:40;not null"`
	APIName      string `json:apiName gorm:"size:100;not null;`
	URL          string `json:"url" gorm:"size:1024;not null;`
	HTTPMethod   string `json:method gorm:"size:10;not null;`
	HTTPRequest  string `json:"request" gorm:"type:text"`
	HTTPResponse string `json:"response" gorm:"type:text"`
	Success      bool   `json:"success"`
	Status       int    `json:"status" gorm:"default:true"`
	Channel      string `json:"channel"`
}
