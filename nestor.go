package Nestor

import (
	"fmt"
	"os"
	"time"
)

const logDirectory = "Logs"

type Nestor struct {
	Level string
	File  string
	Print bool
	Write bool
}

func (nestor Nestor) Info(name string, message string) {

	if nestor.Level == "DEBUG" || nestor.Level == "TRACE" || nestor.Level == "INFO" {
		var msg = fmt.Sprintf("%s  INFO %s %s", time.Now().Format("02-01-2006 15:04:05"), name, message)
		//if nestor.Print {
		fmt.Println(msg)
		//}
		if nestor.Write {
			nestor.WriteToFile(msg, nestor.File)
		}
	}
}

func (nestor Nestor) Debug(name string, message string) {

	if nestor.Level == "DEBUG" || nestor.Level == "TRACE" {
		var msg = fmt.Sprintf("%s DEBUG %s %s", time.Now().Format("2006-01-02 15:04:05"), name, message)
		if nestor.Print {
			fmt.Println(msg)
		}
		if nestor.Write {
			nestor.WriteToFile(msg, nestor.File)
		}
	}
}

func (nestor Nestor) Trace(name string, message string) {

	if nestor.Level == "TRACE" {
		var msg = fmt.Sprintf("%s TRACE %s %s", time.Now().Format("2006-01-02 15:04:05"), name, message)
		if nestor.Print {
			fmt.Println(msg)
		}
		if nestor.Write {
			nestor.WriteToFile(msg, nestor.File)
		}
	}
}

func (nestor Nestor) Error(name string, message string) {

	var msg = fmt.Sprintf("%s ERROR %s %s", time.Now().Format("2006-01-02 15:04:05"), name, message)
	if nestor.Print {
		fmt.Println(msg)
	}
	if nestor.Write {
		nestor.WriteToFile(msg, nestor.File)
	}
}

func (nestor Nestor) WriteToFile(msg string, fileLog string) {
	_ = os.Mkdir(fmt.Sprintf(logDirectory), 0666)
	_ = os.Mkdir(fmt.Sprintf("%s/%s (%d)", logDirectory, time.Now().Month(), time.Now().Year()), 0666)
	_ = os.Mkdir(fmt.Sprintf("%s/%s (%d)/%02d-%02d-%d/", logDirectory, time.Now().Month(), time.Now().Year(),
		time.Now().Day(), time.Now().Month(), time.Now().Year()), 0666)
	path := fmt.Sprintf("%s/%s (%d)/%02d-%02d-%d", logDirectory, time.Now().Month(), time.Now().Year(),
		time.Now().Day(), time.Now().Month(), time.Now().Year())
	//	fmt.Println(path)

	file, err := os.OpenFile(fmt.Sprintf("%s/%s", path, fileLog), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	_, err = file.WriteString(fmt.Sprintf("%s\n", msg))
	if err != nil {
		return
	}

	err = file.Close()
	if err != nil {
		return
	}
}
