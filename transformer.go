package main

import (
	"bitbucket.org/camilo_crespo/camilocrespo/processing"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

const LINE_LENGTH = 65

// Transaction represents data in a row from the original file
type Transaction struct {
	Id         string `json:"id"`
	RecordDate string `json:"recorddate"`
	CardNumber string `json:"cardnumber"`
	Name       string `json:"name"`
	Amount     string `json:"amount"`
}

// TransactionWrap represents the top level document
type TransactionsWrap struct {
	Transactions []Transaction `json:"transactions"`
}

// Processes a flat file into Json
// Syntax to run: go run transformer.go input_path output_path
func main() {
	// open file
	// take input path from the command line
	f, error := os.Open(os.Args[1])
	check(error)
	defer f.Close()
	// process the file
	doc := processFile(f)
	// write the JSON to a file
	// take output path from the command line
	writeOutput(os.Args[2], doc)
}

func processFile(f io.Reader) TransactionsWrap {
	index := 1
	txs := make([]Transaction, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		tx := processLine(line, index)
		txs = append(txs, tx)
		index++
	}
	if err := scanner.Err(); err != nil {
		check(err)
	}
	return TransactionsWrap{Transactions: txs}
}

func processLine(line string, index int) Transaction {
	// validate line is expected length, no less, no more
	if len(line) != LINE_LENGTH {
		panic(fmt.Sprintf("Line %d in the input file doesn't have the valid length", index))
	}
	// call processors (and indirectly validate)
	indexString := strconv.Itoa(index)
	dateToken := processing.ProcessDate(line[0:10])
	cardToken := processing.ProcessCard(line[10:25])
	nameToken := processing.ProcessName(line[25:55])
	amountToken := processing.ProcessAmount(line[55:65])
	tx := Transaction{Id: indexString, RecordDate: dateToken, CardNumber: cardToken, Name: nameToken, Amount: amountToken}
	return tx
}

func writeOutput(outPath string, doc TransactionsWrap) {
	b, _ := json.Marshal(doc)
	var out bytes.Buffer
	json.Indent(&out, b, "", "\t")
	err := ioutil.WriteFile(outPath, out.Bytes(), 0644)
	fmt.Println(out.String())
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
