package mysql

import (
	"e-detect/config"
	"e-detect/model"
	"gorm.io/gorm"
)

type mysqlReportRepository struct {
	connection *gorm.DB
}

func NewMysqlReportRepository(db *gorm.DB) model.ReportRepository {
	return &mysqlReportRepository{
		connection: db,
	}
}

func (m mysqlReportRepository) Save(report model.Report) (res model.Report, err error) {
	//TODO implement me

	if err = m.connection.Save(&report).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (m mysqlReportRepository) GetReportByUserID() (res []model.Report, err error) {
	//TODO implement me
	if err = m.connection.Where("user_id = ?", 1).Find(&res).Error; err != nil {
		return res, err
	}
	return
}

func (m mysqlReportRepository) GetBankReportByUserID() (model.Report, error) {
	//TODO implement me
	panic("implement me")
}

func (m mysqlReportRepository) GetPhoneReportByUserID() (model.Report, error) {
	//TODO implement me
	panic("implement me")
}

func (m mysqlReportRepository) GetReport() ([]model.Report, error) {
	//TODO implement me
	panic("implement me")
}

func (m mysqlReportRepository) UpdateReport(id int, data model.Report) (res model.Report, err error) {
	//TODO implement me
	var NewReport model.Report
	config.DB.First(&NewReport, "id = ?", id)

	if err = config.DB.Model(&NewReport).Updates(map[string]interface{}{
		//"user_id":        data.UserID,
		"nama_terlapor":  data.NamaTerlapor,
		"bank_id":        data.BankID,
		"no_rekening":    data.NoRekening,
		"platform":       data.Platform,
		"kontak_pelaku":  data.KontakPelaku,
		"total_kerugian": data.TotalKerugian,
		"file_bukti":     data.FileBukti,
	}).Error; err != nil {
		return res, err
	}
	return
}

func (m mysqlReportRepository) DeleteReport(id int) (err error) {
	//TODO implement me
	var report model.Report

	if err := config.DB.First(&report, id).Error; err != nil {
		return err
	}

	if err := config.DB.Delete(&report, id).Error; err != nil {
		return err
	}

	return
}
