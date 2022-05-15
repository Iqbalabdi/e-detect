package model

type Disclaimer struct {
	UserID      int
	ReportID    int
	Sanggahan   string
	FileBukti   string
	Tervalidasi bool
}
