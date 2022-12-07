package views

import (
	"fmt"
	"time"
)

func Add(firstNum, secondNum int) string {
	return fmt.Sprintf("%d", firstNum+secondNum)
}

func Sub(firstNum, secondNum int) string {
	return fmt.Sprintf("%d", firstNum-secondNum)
}

func PrintDate(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}
