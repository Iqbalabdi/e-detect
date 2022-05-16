package mysql

import (
	"e-detect/model"
	"gorm.io/gorm"
)

type mysqlDisclaimerRepository struct {
	connection *gorm.DB
}

func NewMysqlDisclaimerRepository(db *gorm.DB) model.DisclaimerRepository {
	return &mysqlDisclaimerRepository{
		connection: db,
	}
}

func (m mysqlDisclaimerRepository) Save(disclaimer model.Disclaimer) (err error) {
	//TODO implement me
	if err = m.connection.Save(&disclaimer).Error; err != nil {
		return err
	}
	return
}

func (m mysqlDisclaimerRepository) GetDisclaimerByUserID(id int) (res []model.Disclaimer, err error) {
	//TODO implement me
	if err = m.connection.Table("disclaimers").Where("user_id = ?", id).Find(&res).Error; err != nil {
		return res, err
	}
	return
}

func (m mysqlDisclaimerRepository) UpdateDisclaimer(id int, data model.Disclaimer) (res model.Disclaimer, err error) {
	//TODO implement me
	var NewDisclaimer model.Disclaimer
	m.connection.First(&NewDisclaimer, "id = ?", id)

	if err = m.connection.Model(&NewDisclaimer).Updates(map[string]interface{}{
		//"user_id":        data.UserID,
		"report_id":  data.ReportID,
		"sanggahan":  data.Sanggahan,
		"file_bukti": data.FileBukti,
	}).Error; err != nil {
		return res, err
	}
	return
}

func (m mysqlDisclaimerRepository) DeleteDisclaimer(id int) (err error) {
	//TODO implement me
	var NewDisclaimer model.Disclaimer

	if err = m.connection.First(&NewDisclaimer, id).Error; err != nil {
		return err
	}

	if err = m.connection.Delete(&NewDisclaimer, id).Error; err != nil {
		return err
	}

	return
}

func (m mysqlDisclaimerRepository) Validate(i int) (err error) {
	//TODO implement me
	var NewDisclaimer model.Disclaimer
	if err = m.connection.Model(&NewDisclaimer).Where("id", i).Update("tervalidasi", 1).Error; err != nil {
		return err
	}
	return
}

func (m mysqlDisclaimerRepository) GetAllDisclaimer() (res []model.Disclaimer, err error) {
	//TODO implement me
	if err = m.connection.Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
