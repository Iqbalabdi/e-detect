package model

import (
	"database/sql"
	"gorm.io/gorm"
)

type Disclaimer struct {
	*gorm.Model
	UserID      int          `json:"user_id" form:"user_id"`
	ReportID    int          `json:"report_id" form:"report_id"`
	Sanggahan   string       `json:"sanggahan" form:"sanggahan"`
	FileBukti   string       `json:"file_bukti" form:"file_bukti"`
	Tervalidasi sql.NullBool `gorm:"default:false"`
}

type DisclaimerRepository interface {
	Save(Disclaimer Disclaimer) error
	GetDisclaimerByUserID(int) ([]Disclaimer, error)
	UpdateDisclaimer(int, Disclaimer) (Disclaimer, error)
	DeleteDisclaimer(int) error
	Validate(int) error
	GetAllDisclaimer() ([]Disclaimer, error)
}

type DisclaimerUseCase interface {
	SaveRequest(Disclaimer Disclaimer) error
	ReadUserDisclaimer(int) ([]Disclaimer, error)
	EditDisclaimer(int, Disclaimer) (Disclaimer, error)
	DeleteDisclaimer(int) error
	Validate(int) error
	GetAllDisclaimer() ([]Disclaimer, error)
}
