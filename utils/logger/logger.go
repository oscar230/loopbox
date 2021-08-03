package logger

import (
	"log"
)

func Info(msg string) {

}

func Warn(msg string) {

}

func Error(msg string) {

}

func Out(s string) {
	log.Println(s)
	return
}
