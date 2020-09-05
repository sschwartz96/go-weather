package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// load config
	cfg, err := loadConfig()
	if err != nil {
		fmt.Printf("error loading config:\n\t%v\n", err)
		return
	}

	// check command line flags

	// lookup data
	//temp, err := getCurrentTempurature(cfg)
	cw, err := getCurrentWeather(cfg)
	if err != nil {
		fmt.Printf("error receiving current temp: \n\t%v\n", err)
		return
	}

	fmt.Printf("temp: %0.f\n", cw.Main.Temp)
	fmt.Printf("clouds: %d%%\n", cw.Clouds.All)

	res, _ := http.Get("https://api.ipify.org")
	ip, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(ip))
}
