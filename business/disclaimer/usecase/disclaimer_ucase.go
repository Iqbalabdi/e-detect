package usecase

import "e-detect/model"

type disclaimerUseCase struct {
	disclaimerRepo model.DisclaimerRepository
}

func NewDisclaimerUseCase(r model.DisclaimerRepository) model.DisclaimerUseCase {
	return &disclaimerUseCase{
		disclaimerRepo: r,
	}
}

func (r disclaimerUseCase) SaveRequest(disclaimer model.Disclaimer) (err error) {
	//TODO implement me

	err = r.disclaimerRepo.Save(disclaimer)
	if err != nil {
		return err
	}
	return
}

func (r disclaimerUseCase) ReadUserDisclaimer(id int) (res []model.Disclaimer, err error) {
	//TODO implement me
	res, err = r.disclaimerRepo.GetDisclaimerByUserID(id)
	if err != nil {
		return res, err
	}
	return
}

func (r disclaimerUseCase) ReadDisclaimer() {
	//TODO implement me
	panic("implement me")
}

func (r disclaimerUseCase) ReadDisclaimerByID() {
	//TODO implement me
	panic("implement me")
}

func (r disclaimerUseCase) EditDisclaimer(id int, data model.Disclaimer) (res model.Disclaimer, err error) {
	//TODO implement me
	res, err = r.disclaimerRepo.UpdateDisclaimer(id, data)
	if err != nil {
		return res, err
	}
	return
}

func (r disclaimerUseCase) DeleteDisclaimer(id int) (err error) {
	//TODO implement me
	err = r.disclaimerRepo.DeleteDisclaimer(id)
	if err != nil {
		return err
	}
	return
}

func (r disclaimerUseCase) Validate(i int) error {
	//TODO implement me
	err := r.disclaimerRepo.Validate(i)
	if err != nil {
		return err
	}
	return nil
}

func (r disclaimerUseCase) GetAllDisclaimer() (res []model.Disclaimer, err error) {
	//TODO implement me
	res, err = r.disclaimerRepo.GetAllDisclaimer()
	if err != nil {
		return nil, err
	}
	return
}
