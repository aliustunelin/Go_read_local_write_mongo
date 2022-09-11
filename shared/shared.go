//asset yolu tutar
package shared

//goda başharfi büyük public,
var Config = configuration{
	FILEURL:  "/home/alistnel/Müzik/calis/proje/namesTable.csv",
	MONGOURL: "mongodb://localhost/squid",
	SPERATOR: ';',
	ISHEADER: true,
}

// private
type configuration struct {
	//dosyamın yolu
	FILEURL string
	//mongo ip'den veya localden gelcek
	MONGOURL string
	//başlık varmı?
	ISHEADER bool
	//csv , iile böler bu ayırıcıyı beliticez
	//goda char yok ama rune var ' ile belirtilir
	SPERATOR rune
}
