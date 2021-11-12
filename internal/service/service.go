package service

import (
	"fmt"
	"github.com/coretech/newProject/internal/api"
	"github.com/coretech/newProject/internal/repository"
	"log"
	"sync"
)

func Service(c chan api.University, done chan bool) {
	console := repository.Console{}
	csvF := repository.CSVfile{}

	_, err := repository.CreateCSV()
	if err != nil {
		fmt.Println("Unable to create file:", err)
		log.Fatalln(err)
	}

	var wg sync.WaitGroup

	for element := range c {
		wg.Add(3)
		go func(element api.University) {
			defer wg.Done()
			WriteAllData(console,element)
		}(element)

		go func(element api.University) {
			defer wg.Done()
			WriteAllData(csvF, element)
		}(element)

		go func(element api.University) {
			defer wg.Done()
			repository.HealthCheck(element)
		}(element)

	}

	wg.Wait()
	done <- true
}

type WriteData interface {
	Write(u api.University)
}
func WriteAllData(w WriteData, u api.University){
	w.Write(u)
}


