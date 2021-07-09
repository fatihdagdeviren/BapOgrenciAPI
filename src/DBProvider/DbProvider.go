package DBProvider

import (
	"BapOgrenciAPI/Models"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
)

type ServiceInterface interface {
	GetStudents() ([]Models.Ogrenci, error)
	GetPersonels() ([]Models.Personel, error)
	OrganizationFromStudent() ([]Models.Organizasyon, error)
	OrganizationFromPersonel() ([]Models.Organizasyon, error)
}

type ServiceManager struct {
	db *gorm.DB
}

var Mgr ServiceInterface

//var MyDB *gorm.DB

func init() {
	dsn := "sqlserver://kullaniciAdi:Sifre@IP?database=DBOgrenci&connection+timeout=30"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	Mgr = &ServiceManager{db: db}
}

func (mgr *ServiceManager) GetStudents() ([]Models.Ogrenci, error) {
	ogrenciModels := make([]Models.Ogrenci, 0)
	sql := "select o.OgrenciNo, o.Ad, o.Soyad,\n (case when op.OgretimDuzeyi=900002 then 1\n       when op.OgretimDuzeyi=900003 then 2\n       " +
		"when op.OgretimDuzeyi=900004 then 3\n       else 0 end) as OgretimDuzeyi, " +
		"kOgretimDuzeyi.Aciklama as OgretimDuzeyiAciklama\n , vw.UstBirimID, vw.BirimID, vw.AltBirimID as ProgramID," +
		" RTRIM(o.OgrenciNo) + '@ogrenci.ege.edu.tr' Email, iCepTel.Aciklama as CepTel\nfrom  Ogrenci o\n" +
		"inner join OrganizasyonProgram op on o.OrganizasyonProgramID = op.OrganizasyonProgramID\n" +
		"inner join Kod kOgretimDuzeyi on op.OgretimDuzeyi = kOgretimDuzeyi.KodID\n" +
		"inner join vwOrganizasyon vw on o.OrganizasyonID = vw.OrganizasyonID\n" +
		"left join Iletisim iCepTel on o.OgrenciID = iCepTel.OgrenciID and iCepTel.IletisimTuru=17002 " +
		"where o.Durum in (905002, 905003) and op.ProgramTuru >= 915002"
	err := mgr.db.Raw(sql).Scan(&ogrenciModels).Error
	return ogrenciModels, err
}

func (mgr *ServiceManager) GetPersonels() ([]Models.Personel, error) {
	personelModels := make([]Models.Personel, 0)
	var sql string = "select * from KBYSPER.DBPersonel.[dbo].[vwUniBapPersonel]"
	err := mgr.db.Raw(sql).Scan(&personelModels).Error
	return personelModels, err
}

func (mgr *ServiceManager) OrganizationFromStudent() ([]Models.Organizasyon, error) {
	organizationList := make([]Models.Organizasyon, 0)
	var sql string = "select OrganizasyonID as Id, Adi as Ad from Organizasyon"
	err := mgr.db.Raw(sql).Scan(&organizationList).Error
	return organizationList, err
}

func (mgr *ServiceManager) OrganizationFromPersonel() ([]Models.Organizasyon, error) {
	organizationList := make([]Models.Organizasyon, 0)
	var sql string = "select OrganizasyonID as Id, Adi as Ad from KBYSPER.DBPersonel.dbo.Organizasyon"
	err := mgr.db.Raw(sql).Scan(&organizationList).Error
	return organizationList, err
}
