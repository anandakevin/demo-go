package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// struct for structure of
type DataLine struct {
	Created []string `json:"created"`
	Deleted []string `json:"deleted"`
}

func printContent(data map[string]DataLine) {
	// print content of map
	for key, value := range data {
		fmt.Println("Key: ", key)
		fmt.Println("Created: ", value.Created)
		fmt.Println("Deleted: ", value.Deleted)
	}
}

func writeJSON(data map[string]DataLine) {
	// print to JSON file from logData
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	// write to file
	f, err := os.Create("output.json")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	_, err = f.Write(jsonData)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
}

func main() {

	lineLength := 83

	f, err := os.Open("app.log")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	byteData := make([]byte, lineLength)

	// declare array of string
	var logData map[string]DataLine = make(map[string]DataLine)
	var lines []string

	// read line by line
	for {
		_, err := f.Read(byteData)
		if err != nil {
			fmt.Printf(err.Error())
			break
		}

		// append line to array
		lines = append(lines, string(byteData))

		// fetch timestamp, action keyword, and UUID from the line
		timestamp := string(byteData[1:17])
		action := string(byteData[37:44])
		uuid := string(byteData[46:82])

		// insert timestamp if not exist in map
		if _, ok := logData[timestamp]; !ok {
			logData[timestamp] = DataLine{Created: []string{}, Deleted: []string{}}
		}

		if action == "created" {
			tempCreated := logData[timestamp].Created
			tempCreated = append(tempCreated, uuid)
			logData[timestamp] = DataLine{Created: tempCreated, Deleted: logData[timestamp].Deleted}
		} else if action == "deleted" {
			tempDeleted := logData[timestamp].Created
			tempDeleted = append(tempDeleted, uuid)
			logData[timestamp] = DataLine{Created: logData[timestamp].Created, Deleted: tempDeleted}
		} else {
			// fmt.Println("Invalid action")
			return
		}
	}

	writeJSON(logData)

	f.Close()
}
