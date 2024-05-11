package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Serverni ochishda xatolik:", err)
		return
	}
	defer ln.Close()
	fmt.Println("Malumot kutilyabti...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Klient bilan bog'lanishda xatolik:", err)
			continue
		}

		fmt.Println("Yangi klient ulandi:", conn.RemoteAddr())
		receiveFile(conn)
	}
}

func receiveFile(conn net.Conn) {
	defer conn.Close()

	receivedFile, err := os.Create("receiveFile.txt")
	if err != nil {
		fmt.Println("Faylni yaratishda xatolik:", err)
		return
	}
	defer receivedFile.Close()
	_, err = io.Copy(receivedFile, conn)
	if err != nil {
		fmt.Println("Faylni yozishda xatolik:", err)
		return
	}

	fmt.Println("Fayl muvaffaqiyatli qabul qilindi.")
}
