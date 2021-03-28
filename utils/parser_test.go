package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
)

func parseJson(filename string) (string, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	// close jsonFile on exit and check for its returned error
	defer func() {
		if err := jsonFile.Close(); err != nil {
			panic(err)
		}
	}()

	// Start reading from the file with a scanner.
	sc := bufio.NewScanner(jsonFile)
	var line string
	var result string
	for sc.Scan() {
		line = sc.Text()  // GET the line string
		// println(line)
		byteStream := make(map[string] interface{})
		err = json.Unmarshal([]byte(line), &byteStream)
		if err != nil {
			fmt.Println(err)
		}
		// println("review:")
		if byteStream["reviewText"] != nil{
			result += byteStream["reviewText"].(string) + " "
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return "", err
	}
	return result, nil
}

func writeToFile(text string) error {
	file, err := os.Create("data.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	if _, err := file.WriteString(text); err != nil {
		panic(err)
	}
	return err
}

func TestParseJson(t *testing.T) {
	text, err := parseJson("C:\\Users\\Vadim\\GolandProjects\\Magazine_Subscriptions\\Magazine_Subscriptions.json")
	if err != nil {
		t.Errorf("ParseJson failed! Error: %v", err)
	}
	err = writeToFile(text)
	if err != nil {
		t.Errorf("writeToFile failed! Error: %v", err)
	}
}
