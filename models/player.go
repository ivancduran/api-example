package models

import (
	"github.com/jinzhu/gorm"
)

type Player struct {
	gorm.Model
	Name      string `json:"name, omitempty"`                           // Player name
	AccountId int    `json:"account_id, omitempty"`                     // Easycast: account id
	Active    int    `json:"active, omitempty" sql:"default:1"`         // status
	Style     string `json:"style, omitempty"`                          // Url css
	StyleJson string `json:"style_json, omitempty" sql:"type:longtext"` // json default configuration
	Autoplay  bool   `json:"autoplay, omitempty" sql:"default:false" `  // Autoplay content
	Controls  bool   `json:"controls, omitempty" sql:"default:true"`    // video players controls
	Entry     string `json:"entry, omitempty"`                          // unique string entry id
	Token     bool   `json:"token"`                                     // player token

	TitleLabel bool   `json:"title_label, omitempty"` // Title true false
	Adblock    bool   `json:"adblock, omitempty"`     // Code to detect adblocks
	Watermark  string `json:"watermark, omitempty"`   // url

	BrandActive bool   `json:"brand_active, omitempty"` // true false
	BrandImg    string `json:"brand_img, omitempty"`    // image url
	BrandLink   string `json:"brand_link, omitempty"`   // link

	AdsActive bool   `json:"ads_active, omitempty"` // true false
	AdsUrl    string `json:"ads_url, omitempty"`    // url vast

	QualitySelector bool `json:"quality_selector"`         // quality selector default true
	SocialActive    bool `json:"social_active, omitempty"` // true false
}

type PlayerStyle struct {
	Data string `json:"data"`
}

func (this *Player) FindAllById(Db *gorm.DB, ID int, p *[]Player) error {

	if err := Db.Where("account_id = ? && active = 1", ID).Find(&p).Error; err != nil {
		return err
	}

	return nil

}

func (this *Player) FindById(Db *gorm.DB, entry string, accountID int) error {

	if err := Db.Where("entry = ? AND account_id = ?", entry, accountID).First(&this).Error; err != nil {
		return err
	}

	return nil
}

func (this *Player) Update(Db *gorm.DB, entry string, accountID int, b map[string]interface{}) error {
	if err := Db.Model(&this).Where("entry = ? AND account_id = ?", entry, accountID).Updates(b).Error; err != nil {
		return err
	}

	return nil
}

// UpdateJsonStyle
func (p *Player) UpdateJSONStyle(Db *gorm.DB, entry string, accountID int, styleJson string) {
	Db.Where("entry = ? and account_id = ?", entry, accountID).First(&p)
	p.StyleJson = styleJson
	Db.Where("entry = ? and account_id = ?", entry, accountID).Save(&p)
}

func (this *Player) Delete(Db *gorm.DB, entry string, accountID int) error {
	if err := Db.Model(&this).Where("entry = ? AND account_id = ? ", entry, accountID).Update("active", 0).Error; err != nil {
		return err
	}
	return nil
}

func (p *Player) DeletePlayers(Db *gorm.DB, id uint) error {

	if result := Db.Where("account_id = ?", id).Delete(Player{}); result.Error != nil {
		if result.RecordNotFound() {
			return nil
		}
		return result.Error
	}

	return nil
}
