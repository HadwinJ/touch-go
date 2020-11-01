package main

import (
	"log"
	"os"
)

type messageType int

const (
	INFO messageType = 0 + iota
	WARNING
	ERROR
	FATAL
)

var (
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
	FatalLogger   *log.Logger
)

func init() {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.LUTC|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.LUTC|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.LUTC|log.Ltime|log.Lshortfile)
	FatalLogger = log.New(file, "FATAL: ", log.LUTC|log.Ltime|log.Lshortfile)
}

func main() {

	InfoLogger.Println("UTC - an information message!")
	WarningLogger.Println("UTC - a warning")
	ErrorLogger.Println("UTC - an error")
	FatalLogger.Fatal("we crash!")
}
