package error

import (
	"errors"
	"fmt"
	"testing"
)

func TestName(t *testing.T) {

	err := New(LoginTypeError).Err()
	if err != nil {
		fmt.Println(err.Error())
		s, _ := FromError(err)
		fmt.Println(s.Code)
		fmt.Println(s.Message)
		fmt.Println("==================")
	}

	err2 := errors.New("not")
	if err2 != nil {
		s2, _ := FromError(err2)
		fmt.Println(s2.Code)
		fmt.Println(s2.Message)
		fmt.Println("==================")
	}

	err3 := New(Unknown, "a", "b").Err()
	if err3 != nil {
		fmt.Println(err3.Error())
		s3, _ := FromError(err3)
		fmt.Println(s3.Code)
		fmt.Println(s3.Message)
	}
}
