package utils

import (
	"bufio"
	"os"
	"github.com/charmbracelet/log"
	"fmt"
)

func AppendIfNotExist(filename string, newline string) string {
	if exist, _ := LineExists(filename, newline); exist {
		return	"IP seems already exist."
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Errorf("Failed to open file: %s", err)
	}
	defer file.Close()

	if _, err := fmt.Fprintln(file, newline); err != nil {
		log.Errorf("Failed to write to file: %s", err)
	}
	
	return fmt.Sprintf("Successfully added IP: %s", newline)
}


func LineExists(filename string, lineToFind string) (bool, error) {
	file, err := os.Open(filename)
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == lineToFind {
			return true, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return false, err
	}

	return false, nil
}