package Models

type OrganizasyonSrvcResult struct {
	Message    string
	IsSuccess  bool
	ResultType int32
	Data       *[]Organizasyon
}

type Organizasyon struct {
	Id string
	Ad string
}
