package _defer

import (
	"errors"
	"fmt"
)

func Error() {
	// 这里我们使用recover函数进行修复
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from",err)
		}
	}()

	fmt.Println("Start")
	panic(errors.New("Something Wrong!"))
}
