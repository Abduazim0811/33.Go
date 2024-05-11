package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	serverAddr := "127.0.0.1:8080"

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Serverga ulanishda xatolik:", err)
		return
	}
	defer conn.Close()

	filePaths := getFilePathsFromTerminal()
	if len(filePaths) == 0 {
		fmt.Println("Fayl manzillari kiritilmagan.")
		return
	}

	for _, filePath := range filePaths {
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("Faylni ochishda xatolik (%s): %v\n", filePath, err)
			continue
		}
		defer file.Close()

		_, err = io.Copy(conn, file)
		if err != nil {
			fmt.Printf("Faylni serverga yuborishda xatolik (%s): %v\n", filePath, err)
			continue
		}

		fmt.Printf("Fayl %s muvaffaqiyatli serverga yuborildi.\n", filePath)
	}
}


func getFilePathsFromTerminal() []string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Yuborish uchun fayllarning manzillarini vergul bilan ajrating:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "" {
		return []string{}
	}
	filePaths := strings.Split(input, ",")
	for i, filePath := range filePaths {
		filePaths[i] = strings.TrimSpace(filePath)
	}
	return filePaths
}
