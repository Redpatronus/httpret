package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/redpatronus/httpret/service"
)

func ParseConfig(filepath string) *service.Svc {
	s := &service.Svc{
		Browser:    &service.SvcBr{},
		VirusTotal: &service.SvcVt{},
		IpInfo:     &service.SvcIi{},
	}
	// Open our jsonFile
	jsonFile, err := os.Open(filepath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	log.Println("Successfully Opened: " + filepath)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return s
}
