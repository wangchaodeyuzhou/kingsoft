package log

import (
	"io"
	"log"
	"os"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	file, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	Trace = log.New(io.Discard, "Trace:", log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(os.Stdout, "Info: ", log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(os.Stdout, "Warning:", log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(io.MultiWriter(file, os.Stderr), "Error:", log.Ldate|log.Ltime|log.Lshortfile)
}
