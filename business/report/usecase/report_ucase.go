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

func (r reportUseCase) ReadUserReports() {
	//TODO implement me

	res, err = r.reportRepo.GetReportByUserID()
	if err != nil {
		return nil, err
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

func (r reportUseCase) EditReport() {
	//TODO implement me
	panic("implement me")
}

func (r reportUseCase) DeleteReport() {
	//TODO implement me
	panic("implement me")
}
