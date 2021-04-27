// Package aleshkin предназначен для выполнения заданий ItUniver
package aleshkin

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/websocket"
)

var database *sql.DB

type Crypta struct {
	//EventType string `json:"e"`
	EventTime            int64  `json:"E"`
	Symbol               string `json:"s"`
	PriceChange          string `json:"p"`
	PriceChangePercent   string `json:"P"`
	WeightedAveragePrice string `json:"w"`
	ClosePrisePrevDay    string `json:"x"`
	ClosePriseCurDay     string `json:"c"`
	QuantityCloseTrade   string `json:"Q"`
	BestBidPrice         string `json:"b"`
	BestBidQuantity      string `json:"B"`
	BestAskPrice         string `json:"a"`
	BestAskQuantity      string `json:"A"`
	OpenPrice            string `json:"o"`
	HighPrice            string `json:"h"`
	LowPrice             string `json:"l"`
	BaseAssetVolume      string `json:"v"`
	QuoteAssetVolume     string `json:"q"`
	OpenTime             int64  `json:"O"`
	CloseTime            int64  `json:"C"`
	FirstTrade           int    `json:"F"`
	LastTrade            int    `json:"L"`
	TotalNumberTrades    int    `json:"n"`
}

//Проверка наличия ошибки
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Вспомогательная функция для завершения считывания данных
func exitCycle(e *bool) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "" {
			*e = false
			break
		}
	}
	checkErr(scanner.Err())
}

// Открытие БД
func databaseOpen() {
	db, err := sql.Open("mysql", "root:password@/gotest")
	checkErr(err)
	database = db
}

// Запись данных в БД
func insertData(db *sql.DB, message *[]byte) {
	var cryptas []Crypta
	json.Unmarshal(*message, &cryptas)
	for _, cryptaOne := range cryptas {
		_, err := db.Exec("insert into gotest.crypto (eventtime, symbol, pricechange, pricechangepercent, weightedaverageprice, closepriseprevday, closeprisecurday, quantityclosetrade, bestbidprice, bestbidquantity, bestaskprice, bestaskquantity, openprice, highprice, lowprice, baseassetvolume, quoteassetvolume, opentime, closetime, firsttrade, lasttrade, totalnumbertrades) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
			time.Unix((cryptaOne.EventTime/1000), (cryptaOne.EventTime%1000)*1000000),
			cryptaOne.Symbol,
			cryptaOne.PriceChange,
			cryptaOne.PriceChangePercent,
			cryptaOne.WeightedAveragePrice,
			cryptaOne.ClosePrisePrevDay,
			cryptaOne.ClosePriseCurDay,
			cryptaOne.QuantityCloseTrade,
			cryptaOne.BestBidPrice,
			cryptaOne.BestBidQuantity,
			cryptaOne.BestAskPrice,
			cryptaOne.BestAskQuantity,
			cryptaOne.OpenPrice,
			cryptaOne.HighPrice,
			cryptaOne.LowPrice,
			cryptaOne.BaseAssetVolume,
			cryptaOne.QuoteAssetVolume,
			time.Unix((cryptaOne.OpenTime/1000), (cryptaOne.OpenTime%1000)*1000000),
			time.Unix((cryptaOne.CloseTime/1000), (cryptaOne.CloseTime%1000)*1000000),
			cryptaOne.FirstTrade,
			cryptaOne.LastTrade,
			cryptaOne.TotalNumberTrades)
		checkErr(err)
	}
}

// Функция StreamCrypto подключается в binance.com через websocket и записывает ticker статистику по всем парам в таблицу БД MySQL
func StreamCrypto() {
	// Подключение к websocket binance.com
	c, _, err := websocket.DefaultDialer.Dial("wss://stream.binance.com:9443/ws/!ticker@arr", nil)
	checkErr(err)
	fmt.Println("Websoket connected. Data loading...")
	defer c.Close()

	//Подключение к БД MySQL
	databaseOpen()
	defer database.Close()

	fmt.Println("Press \"Enter\" to stop data loading and exit from programm")

	exit := true
	go exitCycle(&exit)
	for exit {
		_, message, err := c.ReadMessage()
		checkErr(err)
		insertData(database, &message)
	}
	fmt.Println("The data is written to database")
}
