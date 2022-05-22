package main

type errorCode struct {
	code    string
	message string
}

var Success = errorCode{code: "00000", message: "success"}

var RuntimeError = errorCode{code: "10001", message: "runtime error"}

var ParamError = errorCode{code: "10002", message: "parameters error"}

var UserNameAlreadyExistsError = errorCode{code: "20001", message: "user name already exists"}

var UserNotExistsError = errorCode{code: "20002", message: "user doesn't exist"}

var UserCreditsError = errorCode{code: "20003", message: "user name or password incorrect"}

var UserNotLoginError = errorCode{code: "20004", message: "user doesn't login or token expired"}
