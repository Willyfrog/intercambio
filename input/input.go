package input

import (
	"os"

	"github.com/gocarina/gocsv"
)

type Row struct {
	Title     string `csv:"Title"`
	Amount    string `csv:"Amount"`
	Currency  string `csv:"Currency"`
	Name      string `csv:"Name"`
	Email     string `csv:"Email"`
	Address1  string `csv:"Address1"`
	Address2  string `csv:"Address2"`
	City      string `csv:"City"`
	Zip       string `csv:"Zip"`
	Province  string `csv:"State/Province"`
	Country   string `csv:"Country"`
	Phone     string `csv:"Phone"`
	CreatedAt string `csv:"Created at"`
}

func LoadCSV(input string) ([]*Row, error) {
	in, err := os.Open("clients.csv")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	rows := []*Row{}
	if err := gocsv.UnmarshalFile(in, &rows); err != nil {
		return nil, err
	}
	return rows, nil
}
