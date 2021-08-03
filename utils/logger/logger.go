package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

func Info(msg string) {
	if fmt.Sprint(os.Getenv("LOG_INFO")) == "true" {
		flag := "[info]"
		if fmt.Sprint(os.Getenv("COLORED_OUTPUT")) == "true" {
			msg = color.New(color.FgHiCyan).Sprint(msg)
		}
		out(flag + msg)
	}
}

func Warn(msg string) {
	if fmt.Sprint(os.Getenv("LOG_WARN")) == "true" {
		flag := "[warn]"
		if fmt.Sprint(os.Getenv("COLORED_OUTPUT")) == "true" {
			msg = color.New(color.FgHiYellow).Sprint(msg)
		}
		out(flag + msg)
	}
}

func Error(msg string) {
	if fmt.Sprint(os.Getenv("LOG_ERROR")) == "true" {
		flag := "[error]"
		if fmt.Sprint(os.Getenv("COLORED_OUTPUT")) == "true" {
			msg = color.New(color.FgHiRed).Sprint(msg)
		}
		out(flag + msg)
	}
}

func out(in string) {
	pid := fmt.Sprint(os.Getpid())
	appName := os.Getenv("APP_NAME")
	log.Println("(" + pid + ")[" + appName + "]" + in)
}
