package Models

type PersonelSrvcResult struct {
	Message    string
	IsSuccess  bool
	ResultType int32
	Data       *[]Personel
}

type Personel struct {
	TcKimlikNo           string
	SicilNo              string
	PersonelTipi         string
	Adi                  string
	Soyadi               string
	Unvan                string
	Kadrounvani          string
	KadroBirim           string
	KadroBolum           string
	KadroAnaBilimDali    string
	GorevBirim           string
	GorevBolumAdi        string
	GorevAnaBilimDaliAdi string
	Eposta               string
	Telefon              string
	DogumYeri            string
	DogumYili            string
	Cinsiyet             string
	GirisTarihi          string
	Derece               string
	EkGosterge           string
	Aktif                string
	KadroBosalmaTarihi   string
}
