package model

import (
	"database/sql"
	"e-detect/business/response"
	"gorm.io/gorm"
)

type Report struct {
	*gorm.Model
	UserID        int          `json:"user_id" form:"user_id"`
	TipeLaporan   string       `json:"tipe_laporan" form:"tipe_laporan"`
	NamaTerlapor  string       `json:"nama_terlapor" form:"nama_terlapor"`
	BankID        *int         `json:"bank_id" form:"bank_id"`
	NoRekening    string       `json:"no_rekening" form:"no_rekening"`
	Platform      string       `json:"platform" form:"platform"`
	KontakPelaku  string       `json:"kontak_pelaku" form:"kontak_pelaku"`
	TotalKerugian string       `json:"total_kerugian" form:"total_kerugian"`
	FileBukti     string       `json:"file_bukti" form:"file_bukti"`
	Tervalidasi   sql.NullBool `gorm:"default:false"`
	Disclaimer    Disclaimer
}

type ReportRepository interface {
	Save(report Report) error
	GetReportByUserID(int) ([]Report, error)
	GetBankReportByUserID() (Report, error)
	GetPhoneReportByUserID() (Report, error)
	GetReport() ([]Report, error)
	UpdateReport(int, Report) (Report, error)
	DeleteReport(int) error
	Statistic() (response.StatisticResponse, error)
	DetectBank(string) ([]Report, error)
	DetectPhone(string) ([]Report, error)
	Validate(int) error
	GetAllReport() ([]Report, error)
}

type ReportUseCase interface {
	SaveRequest(report Report) error
	ReadUserReports(int) ([]Report, error)
	ReadUserBankReport()
	ReadUserPhoneReport()
	ReadReport()
	ReadReportByID()
	EditReport(int, Report) (Report, error)
	DeleteReport(int) error
	Statistic() (response.StatisticResponse, error)
	DetectBank(string) ([]Report, error)
	DetectPhone(string) ([]Report, error)
	Validate(int) error
	GetAllReport() ([]Report, error)
}
