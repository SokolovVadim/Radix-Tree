package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"unicode"
)

const max_size = 2000
const dna_size = 1000000

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
	//var counter = 0
	for sc.Scan() {
		if len(result) > max_size {
			break
		}
		line = sc.Text()  // GET the line string
		byteStream := make(map[string] interface{})

		err = json.Unmarshal([]byte(line), &byteStream)
		if err != nil {
			fmt.Println(err)
		}
		text, ok := byteStream["reviewText"].(string)
		if ok {
			result += preprocessData(text) + " "
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return "", err
	}
	return result, nil
}

func writeToFile(text string, filename string) error {
	file, err := os.Create(filename)
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

func preprocessData(text string) string {
	byteSeq := []byte(text)
	for pos, char := range text {
		if !unicode.IsLetter(char) {
			byteSeq[pos] = ' '
		}
	}
	return string(byteSeq)
}

/*func TestParseJson(t *testing.T) {
	text, err := parseJson("C:\\Users\\Vadim\\GolandProjects\\Gift_Cards\\Gift_Cards.json")
	if err != nil {
		t.Errorf("ParseJson failed! Error: %v", err)
	}
	// add EOF
	text += "$"
	err = writeToFile(text[:max_size], "data.txt")
	if err != nil {
		t.Errorf("writeToFile failed! Error: %v", err)
	}
}*/

func parseDNA(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string
	text := string(content)
	fmt.Println("len:", len(text))
	return text, nil
}

func TestParseDNA(t *testing.T) {
	text, err := parseDNA("C:\\Users\\Vadim\\GolandProjects\\TestCorpora\\dna.50MB\\dna.50MB")
	if err != nil {
		t.Errorf("ParseJson failed! Error: %v", err)
	}
	// add EOF
	text += "$"
	err = writeToFile(text[:dna_size], "dna.txt")
	if err != nil {
		t.Errorf("writeToFile failed! Error: %v", err)
	}
}
