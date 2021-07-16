package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"sort"
	"strings"
)



type KeyValue struct {
	Key string
	Value int
}

type Output struct {
	Words string `json:"words"`
	Count int	`json:"count"`
}

func sortMapByValue(m map[string]int) []Output {
	var tempObjArray []KeyValue
	var output []Output
	for k, v := range m {
		tempObjArray = append(tempObjArray, KeyValue{k, v})
	}

	sort.Slice(tempObjArray, func(i, j int) bool {
		return tempObjArray[i].Value > tempObjArray[j].Value
	})

	for _, kv := range tempObjArray {
		output = append(output, Output{Words:kv.Key, Count:kv.Value})
	}
	return output
}

func FindTop10(c *fiber.Ctx) error {
	readFile, err := c.FormFile("textFile")
	if err != nil {
		fmt.Println("Send file with form key `textFile`")
		return c.Status(http.StatusBadRequest).SendString("Send file with form key `textFile`")
	}
	file, fileOpenError := readFile.Open()
	if fileOpenError != nil {
		fmt.Println("Unable to open file")
		return c.Status(http.StatusInternalServerError).SendString("Unable to open file")
	}
	defer file.Close()
	//create byte stream to read the file into
	fileData := make([]byte, readFile.Size)
	_, readError := file.Read(fileData)
	if readError != nil {
		fmt.Println("Unable to read file data")
		return c.Status(http.StatusInternalServerError).SendString("Unable to read file data")
	}

	// removes the dots and special characters
	replacer := strings.NewReplacer(".", "", ",","",":", "", ";","", "!", "", "`","", "\"", "", "?","")

	newFileData := replacer.Replace(string(fileData))

	// splitting the file string to words
	words := strings.Fields(string(newFileData))

	countMap := make(map[string]int)

	// iterating through the words to count the occurrence
	for _, word := range words {
		countMap[word] = countMap[word] + 1
	}

	sortedArr := sortMapByValue(countMap)
	top10 := sortedArr

	if len(sortedArr) >= 10 {
		top10 = sortedArr[:10]
	}

	return c.JSON(top10)
}
