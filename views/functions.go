package views

import "fmt"

func Add(firstNum, secondNum int) string {
	return fmt.Sprintf("%d", firstNum+secondNum)
}
func Sub(firstNum, secondNum int) string {
	return fmt.Sprintf("%d", firstNum-secondNum)
}
