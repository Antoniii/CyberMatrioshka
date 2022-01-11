package main
// #include <stdio.h>
// void callC() {
//    printf("Calling C code! $> fatal error: Python.h: No such file or directory #include <Python.h>\n");
// }
import "C"

import "fmt"
func main() {
	fmt.Println("A Go statement!")
    C.callC()
    fmt.Println("Another Go statement!")
}

// go run cGo.go 
// go build cGo.go && ./cGo