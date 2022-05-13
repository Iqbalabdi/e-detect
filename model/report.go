package model

import (
	"gorm.io/gorm"
)

type Report struct {
	*gorm.Model
	UserID        int
	TipeLaporan   string
	NamaTerlapor  string
	BankID        int
	NoRekening    string
	Platform      string
	KontakPelaku  string
	TotalKerugian string
	FileBukti     string
	Tervalidasi   bool
}

type ReportRepository interface {
	Save(report Report) (Report, error)
	GetReportByUserID() (Report, error)
	GetBankReportByUserID() (Report, error)
	GetPhoneReportByUserID() (Report, error)
	GetReport() ([]Report, error)
	UpdateReport() (Report, error)
	DeleteReport() (Report, error)
}

type ReportUseCase interface {
	SaveRequest(report Report) (Report, error)
	ReadUserReports()
	ReadUserBankReport()
	ReadUserPhoneReport()
	ReadReport()
	ReadReportByID()
	EditReport()
	DeleteReport()
}
