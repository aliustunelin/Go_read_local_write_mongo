//csv okuyacak
package coreLib

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"time"

	"../shared"
)

//csv içinden saedce ad ve soyad alcak str döncek
func ReadNamesTableList() [][]string { //go'da 2 array döncem

	//dosya yolunu direk yazma, shared config den gelsin onlar
	fileCsv := shared.Config.FILEURL

	//okuduğum datanın temizi
	personData := make([][]string, 0)

	//read csv
	f, err := os.Open(fileCsv)
	simpleErrorHandle(err)

	defer f.Close()
	csvReader := csv.NewReader(f)
	// ; ile split edecek
	csvReader.Comma = shared.Config.SPERATOR

	dataPerson, err := csvReader.ReadAll()
	simpleErrorHandle(err)

	//gelen dataperson'ı personData içine temize çekelim
	for i, line := range dataPerson {

		//header varsa ve split 0'dan büyükse alalım içine
		if (shared.Config.ISHEADER && i > 0) || !shared.Config.ISHEADER {
			//alınacak isim soyisim appedn edildi
			personData = append(personData, []string{line[1], line[2]})
		}
	}
	return personData
}

//nick atamayı ayrı yapalım, ileride logic vs değişirse 1 yerden değiştirelim
func CreateNick(person []string) string {
	//nick oluşturma ad'ın 3 harfi, soyadın 3 harfinden

	return person[0][0:3] + person[1][0:3]
}

//modeldeki No random oluşcak
func CreateNo(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	//random algom
	num := rand.Intn(max-min) + min
	return num
}

func simpleErrorHandle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
