package output

import (
	"os"

	"github.com/gocarina/gocsv"
)

type Match struct {
	SenderEmail   string
	SenderName    string
	ReceiverEmail string // this is for id purposes for next year
	ReceiverName  string
	Address1      string
	Address2      string
	City          string
	Zip           string
	Province      string
	Country       string
	Phone         string
}

const OutputRow string = "{.SenderEmail},{.SenderName},{.ReceiverEmail},{.ReceiverName},{.Address1},{.Address2},{.City},{.Zip},{.Province},{.Country},{.Phone},{.CreatedAt}"

func WriteCSV(results []*Match, outputFile string) error {
	csvContent, err := gocsv.MarshalString(&results)
	if err != nil {
		return err
	}

	return os.WriteFile(outputFile, []byte(csvContent), 0666)
}
