package main

// 32
import (
	"fmt"
	"time"

	"./model"

	"./coreLib"
)

func main() {
	//main kaç sn sürmüş calc etsin
	defer calcTime(time.Now(), "Get All Transaction") //defer fonk kapanırken calısır

	// TODO: verilerimi alayım bi, model, readData'm olsun
	// call model
	transactionArray := make([]model.TransactionLayer, 0)

	//call clean value
	dataPerson := coreLib.ReadNamesTableList()

	// TODO: ham verimizi okuduk, json dosyamızı oluşturalım, fin

	//TODO: json ok, channel ile her kayıt gelince senkron dinleyelim
	personChannel := make(chan model.TransactionLayer, 0)
	//channel fonk go ile çağır
	go printPerPlayer(personChannel, dataPerson)

	//TODO: kayıtları dinleyelim simdi, kayıt gelince buraya gelcek
	for person := range personChannel {
		//direk channel range ile dinleniyor, yenisi gelince artar
		transactionArray = append(transactionArray, person)
		fmt.Printf("Nick:%s No:%d\n", person.Nick, person.No)
		time.Sleep(300 * time.Millisecond)
	}
}

//channel burası publisher
func printPerPlayer(personChannel chan model.TransactionLayer, personList [][]string) {
	for _, person := range personList {
		modelUser := model.TransactionLayer{
			Nick: coreLib.CreateNick(person),
			No:   coreLib.CreateNo(1, 471),
		}
		personChannel <- modelUser
	}
}

//main kaç sn sürmüş calc etsin
func calcTime(start time.Time, description string) {
	elapsed := time.Since(start)
	fmt.Printf("%s Total Time:%s", description, elapsed)

}
