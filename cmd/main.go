package main

import (
	"github.com/coretech/newProject/internal/api"
	"github.com/coretech/newProject/internal/config"
	"github.com/coretech/newProject/internal/service"
	"strings"
	"sync"
)

func main()  {
	const (
		yamlFile = "config.yml"
		separator = ","
	)

	appConfig := config.New(yamlFile)
	countries := strings.Split(appConfig.Countries,separator)
	capChannel := len(appConfig.Countries)

	c := make(chan api.University,capChannel)
	done := make(chan bool)

	var wg sync.WaitGroup

	for _, country:= range countries {
		wg.Add(1)

		go func(country string, c chan api.University) {
			defer  wg.Done()
			api.GetRequest(country, c)
		}(country, c)
	}
	
	go service.Service(c,done)

	wg.Wait()
	close(c)

	<- done
}