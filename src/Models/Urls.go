package Models

import (
	"url_shortener/Config"
	_ "github.com/go-sql-driver/mysql"
)


type Url struct {
	ID uint 			`json:"id"`
	Url string 		`json:"url"`
}

func (u *Url) TableName() string {
	return "url"
}

func GetAllUrls(url *[]Url) (err error) {
	if err = Config.DB.Find(url).Error; err != nil {
		return err
	}
	return nil
}

func CreateUrl(url *Url) (err error) {
	if err = Config.DB.Create(url).Error; err != nil {
		return err
	}
	return nil
}

func GetUrl(url *Url, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(url).Error; err != nil {
		return err
	}
	return nil
}

func GetUrlByUrl(url *Url, payloadUrl string) (err error) {
	if err := Config.DB.Where("url = ?", payloadUrl).First(url).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUrl(url *Url, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(url)
	return nil
}