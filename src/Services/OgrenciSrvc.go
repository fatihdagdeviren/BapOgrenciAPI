package Services

import (
	cons "BapOgrenciAPI/Consts"
	dbProvider "BapOgrenciAPI/DBProvider"
	"BapOgrenciAPI/Models"
)

func GetStudents() Models.OgrenciSrvcResult {
	ogrenciResult := Models.OgrenciSrvcResult{}
	ogrenciList, err := dbProvider.Mgr.GetStudents()
	ogrenciResult.IsSuccess = (err == nil)
	ogrenciResult.ResultType = 1
	if err != nil {
		ogrenciResult.Message = err.Error()
	} else {
		ogrenciResult.Message = cons.SrvcSuccess
	}
	ogrenciResult.Data = &ogrenciList
	return ogrenciResult
}
