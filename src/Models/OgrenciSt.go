package Models

type OgrenciSrvcResult struct {
	Message    string
	IsSuccess  bool
	ResultType int32
	Data       *[]Ogrenci
}

type Ogrenci struct {
	OgrenciNo             string
	Ad                    string
	Soyad                 string
	OgretimDuzeyi         string
	OgretimDuzeyiAciklama string
	UstBirimID            string
	BirimID               string
	ProgramID             string
	Email                 string
	CepTel                string
}
