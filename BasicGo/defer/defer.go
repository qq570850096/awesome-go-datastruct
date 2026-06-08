package _defer

import (
	"errors"
	"fmt"
)

func Error() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from", err)
		}
	}()

	fmt.Println("Start")
	panic(errors.New("Something Wrong!"))
}
