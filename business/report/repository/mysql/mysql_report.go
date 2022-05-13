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

func (m mysqlReportRepository) Save(report model.Report) (res model.Report, err error) {
	//TODO implement me

	if err = m.connection.Save(&report).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (m mysqlReportRepository) GetReportByUserID() (model.Report, error) {
	//TODO implement me
	if err = m.connection.Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
	panic("implement me")
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

func (m mysqlReportRepository) UpdateReport() (model.Report, error) {
	//TODO implement me
	panic("implement me")
}

func (m mysqlReportRepository) DeleteReport() (model.Report, error) {
	//TODO implement me
	panic("implement me")
}
