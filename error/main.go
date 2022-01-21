package main

import (
	"errors"
	"fmt"
	"strings"
)

// TODO error messageleri json reponcelere nasil koyalarim ?
// Paketin verecegi errorlar boyle tanimlanabilir
var (
	ErrInternal  = errors.New("errorex: internal error")
	ErrBilmemene = errors.New("errorex: bilmemne")
)

// https://youtu.be/29LLRKIL_TI?t=1201
type Error struct {
	Code    int
	Message string
	Detail  interface{}
}

// implements Error interface
func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", strings.ToLower(e.Message), fmt.Sprint(e.Code))
}

func main() {
	// constant errorler kucuk-orta projelerde mantili olabilir
	// customError()
	// constantError()

	// stdlib ornekleri
	fmt.Printf("")

}

func customError() {
	myErr := Error{Code: 1, Message: "Deneme error"}
	panic(myErr)
}

func myFunc() error {
	return ErrInternal
}

func constantError() {
	err := myFunc()
	fmt.Printf("error %+v\n", err)

	if err != nil {
		panic(err)
	}
}
