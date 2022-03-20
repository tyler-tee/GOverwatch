package main

import (
		"bufio"
		"encoding/json"
		"flag"
		"fmt"
		"log"
		"os"
		"os/exec"
		"strings"
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

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func masscanParser(path string) (string, string) {
	lines, err := readLines(path)

	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	var addresses_raw []string
	var ports_raw []string

	for _, line := range lines {
		if len(strings.Split(line, " ")) > 2{
			addresses_raw = append(addresses_raw, strings.Split(line, " ")[3])
			ports_raw = append(ports_raw, strings.Split(line, " ")[2])
		}
	}

	addresses := strings.Join(addresses_raw, " ")
	ports := strings.Join(ports_raw, ",")

	return addresses, ports
}

// func scan_handler(target string, timestamp string, portQuant string, portType string, osDetection bool)
func main() {
	config := getConfig().Scan
	massCmd := fmt.Sprintf("masscan %s", config.Targets)
	massCmd += fmt.Sprintf(" -p%s", config.Ports)
	massCmd += fmt.Sprintf(" --rate %s -oL mass_results", config.Rate)

	err := exec.Command(massCmd)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Masscan complete.")

	addresses, ports := masscanParser("./mass_results")

	nmapCmd := "nmap "

	if config.osDetection {
		nmapCmd += "-O "
	}

	nmapCmd += "-p " + ports + " -Pn " + addresses + " -oX nmap_results.txt"

	err = exec.Command(nmapCmd)

	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	fmt.Println("Nmap complete.")
}
