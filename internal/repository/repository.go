package repository

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/coretech/newProject/internal/api"
	"log"
	"net/http"
	"os"
)

type CSVfile struct {}
type Console struct {}

func (cons Console) Write( u api.University)  {
	fmt.Printf("\nSending data to the console: %s",u)
}

func (csvF CSVfile) Write( u api.University)  {
	var site string
	for _,sites:= range u.WebPages{
		site += ","+sites
	}
	strUniversity:= u.Name +","+ u.Country+ site

	var message []string
	message  = append(message, strUniversity)

	file, err := OpenCSV()
	defer close(file)
	if err != nil {
		log.Fatalf("We are not able to process row to CSV file: %s",err)
	}

	w := csv.NewWriter(file)

	if err := w.Write(message); err != nil {
		log.Fatalln("error writing record to csv:", err)
	}
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

//Open CSV file
func OpenCSV() (*os.File, error) {
	file, err := os.OpenFile("file.csv", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	return file,nil
}


// Write in Health Check with validation and verification of sites for performance
func HealthCheck(u api.University) {

	for _,v := range u.WebPages {

		_, err := http.Get(v)
		if err != nil {
			fmt.Printf("\nThis site is not working - %s",err)
			return
		}
		fmt.Println("This site is good working : ",v)
	}
}


// Create CSV file
func CreateCSV() (*os.File, error) {
	file, err := os.Create("file.csv")
	defer close(file)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		return nil, errors.New("this is error")
	}

	return file, err
}

// Close file
func close(f *os.File) {
	err := f.Close()
	if err != nil {
		fmt.Println("Unable to create file:", err)
	}
}

