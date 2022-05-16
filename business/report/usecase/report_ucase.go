package usecase

import "e-detect/model"

type reportUseCase struct {
	reportRepo model.ReportRepository
}

func NewReportUseCase(r model.ReportRepository) model.ReportUseCase {
	return &reportUseCase{
		reportRepo: r,
	}
}

func (r reportUseCase) SaveRequest(report model.Report) (err error) {
	//TODO implement me

	err = r.reportRepo.Save(report)
	if err != nil {
		return err
	}
	return
}

func (r reportUseCase) ReadUserReports() (res []model.Report, err error) {
	//TODO implement me

	res, err = r.reportRepo.GetReportByUserID()
	if err != nil {
		return res, err
	}
	return
}

func (r reportUseCase) ReadUserBankReport() {
	//TODO implement me
	panic("implement me")
}

func (r reportUseCase) ReadUserPhoneReport() {
	//TODO implement me
	panic("implement me")
}

func (r reportUseCase) ReadReport() {
	//TODO implement me
	panic("implement me")
}

func (r reportUseCase) ReadReportByID() {
	//TODO implement me
	panic("implement me")
}

func (r reportUseCase) EditReport(id int, data model.Report) (res model.Report, err error) {
	//TODO implement me
	res, err = r.reportRepo.UpdateReport(id, data)
	if err != nil {
		return res, err
	}
	return
}

func (r reportUseCase) DeleteReport(id int) (err error) {
	//TODO implement me
	err = r.reportRepo.DeleteReport(id)
	if err != nil {
		return err
	}
	return
}

func (r reportUseCase) Statistic() (totalReport int64, totalBank int64, totalPhone int64, totalCost int64, err error) {
	//TODO implement me
	totalReport, totalBank, totalPhone, totalCost, err = r.reportRepo.Statistic()

	return
}

func (r reportUseCase) DetectBank(number string) (report []model.Report, err error) {
	//TODO implement me
	report, err = r.reportRepo.DetectBank(number)
	if err != nil {
		return
	}
	return
}

func (r reportUseCase) DetectPhone(s string) (phone []model.Report, err error) {
	//TODO implement me
	phone, err = r.reportRepo.DetectPhone(s)
	if err != nil {
		return
	}
	return
}

func (r reportUseCase) Validate(i int) error {
	//TODO implement me
	err := r.reportRepo.Validate(i)
	if err != nil {
		return err
	}
	return nil
}

func (r reportUseCase) GetAllReport() (res []model.Report, err error) {
	//TODO implement me
	res, err = r.reportRepo.GetAllReport()
	if err != nil {
		return nil, err
	}
	return
}
