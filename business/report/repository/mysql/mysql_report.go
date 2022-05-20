package mysql

import (
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

func (m mysqlReportRepository) Save(report model.Report) (err error) {
	//TODO implement me

	if err = m.connection.Save(&report).Error; err != nil {
		return err
	}
	return
}

func (m mysqlReportRepository) GetReportByUserID(user_id int) (res []model.Report, err error) {
	//TODO implement me
	if err = m.connection.Table("reports").Where("user_id = ?", user_id).Find(&res).Error; err != nil {
		return res, err
	}
	return res, err
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
	m.connection.Table("reports").First(&NewReport, "id = ?", id)

	if err = m.connection.Model(&NewReport).Updates(map[string]interface{}{
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
	return NewReport, err
}

func (m mysqlReportRepository) DeleteReport(id int) (err error) {
	//TODO implement me
	var report model.Report

	if err := m.connection.First(&report, id).Error; err != nil {
		return err
	}

	if err := m.connection.Delete(&report, id).Error; err != nil {
		return err
	}

	return
}

func (m mysqlReportRepository) Statistic() (totalReport int64, totalBank int64, totalPhone int64, totalCost int64, err error) {

	m.connection.Table("reports").Count(&totalReport)
	m.connection.Table("reports").Where("tipe_laporan = ?", "phone").Count(&totalPhone)
	m.connection.Table("reports").Where("tipe_laporan = ?", "rekening").Count(&totalBank)
	m.connection.Table("reports").Select("sum(total_kerugian)").Row().Scan(&totalCost)

	return
}

func (m mysqlReportRepository) DetectBank(i string) (report []model.Report, err error) {
	//TODO implement me
	if err = m.connection.Table("reports").Where("no_rekening = ?", i).Find(&report).Error; err != nil {
		return
	}

	return
}

func (m mysqlReportRepository) DetectPhone(s string) (phone []model.Report, err error) {
	//TODO implement me
	if err = m.connection.Table("reports").Where("kontak_pelaku = ?", s).Find(&phone).Error; err != nil {
		return
	}

	return
}

func (m mysqlReportRepository) Validate(i int) (err error) {
	//TODO implement me
	var report model.Report
	if err = m.connection.Model(&report).Where("id", i).Update("tervalidasi", 1).Error; err != nil {
		return err
	}
	return
}

func (m mysqlReportRepository) GetAllReport() (res []model.Report, err error) {
	//TODO implement me
	if err = m.connection.Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
