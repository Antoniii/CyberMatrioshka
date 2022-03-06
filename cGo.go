package main

// #include <Python.h>
// void callC() {
//     Py_Initialize();
//     PyRun_SimpleString("import sys");
//     PyRun_SimpleString("sys.path.append('.')");
//     PyRun_SimpleString("import pyenjs");
//     Py_Finalize();
// }
import "C"

import "fmt"

func main() {
	fmt.Println("A Go statement!")
	C.callC()
	fmt.Println("Another Go statement!")
}

// CGO_CFLAGS="-I/usr/include/python3.10" CGO_LDFLAGS="-lpython3.10" go run cGo.go
