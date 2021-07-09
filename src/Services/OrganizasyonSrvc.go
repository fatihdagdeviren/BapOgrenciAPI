package Services

import (
	cons "BapOgrenciAPI/Consts"
	dbProvider "BapOgrenciAPI/DBProvider"
	"BapOgrenciAPI/Models"
)

func GetOrganizationFromStudent() Models.OrganizasyonSrvcResult {
	studentResult := Models.OrganizasyonSrvcResult{}
	res, err := dbProvider.Mgr.OrganizationFromStudent()
	studentResult.IsSuccess = (err == nil)
	studentResult.ResultType = 1
	if err != nil {
		studentResult.Message = err.Error()
	} else {
		studentResult.Message = cons.SrvcSuccess
	}
	studentResult.Data = &res
	return studentResult
}

func GetOrganizationFromPersonel() Models.OrganizasyonSrvcResult {
	personelStudent := Models.OrganizasyonSrvcResult{}
	res, err := dbProvider.Mgr.OrganizationFromPersonel()
	personelStudent.IsSuccess = (err == nil)
	personelStudent.ResultType = 1
	if err != nil {
		personelStudent.Message = err.Error()
	} else {
		personelStudent.Message = cons.SrvcSuccess
	}
	personelStudent.Data = &res
	return personelStudent
}
