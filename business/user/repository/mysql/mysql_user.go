package mysql

import (
	"e-detect/config"
	"e-detect/model"
	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	connection *gorm.DB
}

func NewMysqlUserRepository(db *gorm.DB) model.UserRepository {
	return &mysqlUserRepository{
		connection: db,
	}
}

func (m mysqlUserRepository) GetAll() (res []model.User, err error) {
	//TODO implement me
	if err = m.connection.Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (m mysqlUserRepository) GetByQuery(user model.User) (res model.User, err error) {
	//TODO implement me
	var userAuth model.User
	if err = m.connection.Where("email = ?", user.Email).First(&userAuth).Error; err != nil {
		return
	}
	return userAuth, nil

}

func (m mysqlUserRepository) Create(user model.User) (res model.User, err error) {
	//TODO implement me
	if err = m.connection.Save(&user).Error; err != nil {
		return res, err
	}
	return user, nil
}

func (m mysqlUserRepository) Update(id int, data model.User) (res model.User, err error) {
	//TODO implement me
	var user model.User
	config.DB.First(&user, "id = ?", id)

	if err = config.DB.Model(&user).Updates(map[string]interface{}{"name": data.Name, "email": data.Email, "password": data.Password, "phone": data.Phone}).Error; err != nil {
		return user, err
	}
	return user, err
}
