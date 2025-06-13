package get_conf

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Read() {
	file, err := os.Open("./config.txt")
	if err != nil {
		log.Fatal("Error opening config file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			if key == "" {
				continue
			}
			os.Setenv(key, value)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading config file:", err)
	}
}