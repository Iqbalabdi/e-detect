package usecase

import (
	"e-detect/model"
	validator "github.com/go-playground/validator/v10"
)

type userUsecase struct {
	userRepo model.UserRepository
	validate *validator.Validate
}

func NewUserUseCase(u model.UserRepository) model.UserUseCase {
	return &userUsecase{
		userRepo: u,
		validate: validator.New(),
	}
}

func (u *userUsecase) GetAll() (res []model.User, err error) {
	res, err = u.userRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return
}

func (u *userUsecase) Create(user model.User) (res model.User, err error) {
	var ok bool
	if ok, err = u.Validate(&user); !ok {
		return res, err
	}

	res, err = u.userRepo.Create(user)
	if err != nil {
		return res, err
	}
	return
}

func (u *userUsecase) Validate(user *model.User) (bool, error) {
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *userUsecase) Login(data model.LoginRequest) (res model.User, ok bool, err error) {

	if err = u.validate.Struct(data); err != nil {
		return
	}

	res.Email = data.Email
	res, err = u.userRepo.GetByQuery(res)

	if err != nil || res.Password != data.Password {
		return res, false, err
	}

	return res, true, nil
}

func (u *userUsecase) Update(id int, user model.User) (res model.User, err error) {
	//TODO implement me
	res, err = u.userRepo.Update(id, user)
	if err != nil {
		return res, err
	}
	return
}
