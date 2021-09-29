package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	// "os/exec"
)

type Configuration struct {
	Scan struct {
		Title       string
		Targets     string
		Rate        string
		Ports       string
		PortTypes   string
		osDetection bool
	}
}

func getConfig() Configuration {
	// Try to open config file "scan.conf" in same directory - Prompt if not found.
	configLoc := flag.String("configLoc", "./scan.conf", "Specify the configuration file.")
	flag.Parse()
	file, err := os.Open(*configLoc)
	if err != nil {
		log.Fatal("Failed to open config file: ", err)
	}
	defer file.Close()

	// Read in JSON and return Configuration struct
	decoder := json.NewDecoder(file)
	Config := Configuration{}
	err = decoder.Decode(&Config)
	if err != nil {
		log.Fatal("can't decode config JSON: ", err)
	}
	log.Println(Config.Scan.osDetection)

	return Config
}

// func scan_handler(target string, timestamp string, portQuant string, portType string, osDetection bool)

func main() {
	config := getConfig().Scan
	massCmd := fmt.Sprintf("masscan %s", config.Targets)
	massCmd += fmt.Sprintf(" -p%s", config.Ports)
	massCmd += fmt.Sprintf(" --rate %s -oL mass.txt", config.Rate)

	fmt.Printf(massCmd)
	/*
		err := exec.Command(massCmd)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Scan complete.")
	*/

	/*
		m := masscan.New()
		m.SetPorts(config.Scan.Ports)

		m.SetRanges(config.Scan.Targets)

		m.SetRate(config.Scan.Rate)

		err := m.Run()
		if err != nil {
			fmt.Println("Scan failed: ", err)
			return
		}

		results, err := m.Parse()
		if err != nil {
			fmt.Println("Parsed scan results: ", err)
			return
		}

		for _, result := range results {
			fmt.Println(result)
		}
	*/

}
