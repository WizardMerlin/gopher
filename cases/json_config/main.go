package main

import (
_	"io/ioutil"
	"os"
	"fmt"
	"encoding/json"
)

type Config struct {
	Database struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"database"`
	Host string `json:"host"`
	Port string `json:"port"`
}

func LoadConfiguration(filename string) (Config, error) {
	var config Config
	//configFile, err := ioutil.ReadFile(filename) //slice
	configFile, err:= os.Open(filename) // * File
	defer configFile.Close()

	if err != nil {
		return config, err
	}
	//json.Marshal --- slice
	//json.Unmarshal  --- slice
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config) //store it in struct config
	return config, err
}

func (config Config) String() string {
	return fmt.Sprintf("Config{database{host:%s, port:%s}, host:%s, port:%s}", 
		config.Database.Host, config.Database.Port, config.Host, config.Port)
}

func main() {
	fmt.Println("starting the application...")
	if config, err := LoadConfiguration("config.json"); err != nil {
		fmt.Println(err);
		os.Exit(1)
	} else {
		fmt.Printf("%v\n", config)
	}

}