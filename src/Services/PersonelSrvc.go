package Services

import (
	cons "BapOgrenciAPI/Consts"
	dbProvider "BapOgrenciAPI/DBProvider"
	"BapOgrenciAPI/Models"
)

func GetPersonels() Models.PersonelSrvcResult {
	personelResult := Models.PersonelSrvcResult{}
	personelList, err := dbProvider.Mgr.GetPersonels()
	personelResult.IsSuccess = (err == nil)
	personelResult.ResultType = 1
	if err != nil {
		personelResult.Message = err.Error()
	} else {
		personelResult.Message = cons.SrvcSuccess
	}
	personelResult.Data = &personelList
	return personelResult
}
