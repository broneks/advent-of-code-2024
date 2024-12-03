package shared

import (
	"bufio"
	"fmt"
	"os"
)

func ScanFile(filepath string, onScanLine func(line string)) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		onScanLine(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading lines:", err)
	}
}
