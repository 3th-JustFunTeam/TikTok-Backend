package main

type errorCode struct {
	code    string
	message string
}

var Success = errorCode{
	code:    "00000",
	message: "success",
}

var RuntimeError = errorCode{
	code:    "10001",
	message: "runtime error",
}
