package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func splitPerElf() func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	searchBytes := []byte("\n\n")
	searchLength := len("\n\n")
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		dataLen := len(data)

		if atEOF && dataLen == 0 {
			return 0,nil,nil
		}

		if i := bytes.Index(data, searchBytes); i >= 0 {
			return i + searchLength, data[0:i], nil
		}

		if atEOF {
			return dataLen, data, nil
		}

		return 0, nil, nil
	}
}

func main() {
	file, err := os.Open("./calories")

	check(err)

	defer func(file *os.File) {
		err := file.Close()
		check(err)
	}(file)

	scanner := bufio.NewScanner(file)

	scanner.Split(splitPerElf())
	check(scanner.Err())

	var entities []int

	for scanner.Scan() {
		lines := strings.Fields(scanner.Text())
		total := 0

		for _, line := range lines {
			entry, err := strconv.Atoi(line)
			check(err)
			total += entry
		}

		entities = append(entities, total)
	}

	sort.Ints(entities)
	fmt.Println(entities[len(entities)-1])
	fmt.Println(entities[len(entities)-1] + entities[len(entities)-2] + entities[len(entities)-3])
}