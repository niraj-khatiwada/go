package main

//import (
//	"bufio"
//	"errors"
//	"fmt"
//	"log"
//	"os"
//)
//
//func main() {
//	// Create file
//	f, err := os.Create("abc.txt")
//	if err != nil {
//		log.Fatal("Oopsie!")
//	}
//	// Close file
//	defer func(f *os.File) {
//		err := f.Close()
//		if err != nil {
//			log.Fatal("Couldn't close file.")
//		}
//	}(f)
//	// Write into created file
//	_, err2 := f.WriteString("Hello Niraj\n")
//	if err2 != nil {
//		log.Fatal("Oopsie!")
//	}
//
//	openFile, err3 := os.Open("abcc.txt")
//	if err3 != nil {
//		if errors.Is(err3, os.ErrNotExist) {
//			log.Fatal("File does not exist.")
//		} else {
//			log.Fatal("Oopsie!")
//		}
//	}
//
//	defer (func(f *os.File) {
//		err := f.Close()
//		if err != nil {
//			log.Fatal("Couldn't close file.")
//		}
//	})(openFile)
//
//	scanner := bufio.NewScanner(openFile)
//	for scanner.Scan() {
//		fmt.Println(scanner.Text())
//	}
//	if err := scanner.Err(); err != nil {
//		log.Fatal("Scanner Error")
//	}
//
//	// Get the status of file
//	_, err4 := os.Stat("abc.txt")
//	if errors.Is(err4, os.ErrNotExist) {
//		log.Fatal("File does not exist.")
//	} else {
//		file, err := os.OpenFile("abc.txt", os.O_WRONLY|os.O_APPEND, 0644)
//		if err != nil {
//			log.Fatal("Oopsie!")
//		}
//		defer (func(f *os.File) {
//			if err := (*f).Close(); err != nil {
//				log.Fatal("Couldn't close file.")
//			}
//
//		})(file)
//		if _, err := file.WriteString("\nHello World"); err != nil {
//			log.Fatal("Error while writing file")
//		}
//	}
//
//}
