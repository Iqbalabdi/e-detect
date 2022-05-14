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

func (r reportUseCase) SaveRequest(report model.Report) (res model.Report, err error) {
	//TODO implement me

	res, err = r.reportRepo.Save(report)
	if err != nil {
		return res, err
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

func (r reportUseCase) Statistic() (model.Report, error) {
	//TODO implement me
	panic("implement me")
}
