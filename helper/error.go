package helper

// trata a mensagem de erro, tratando como uma condicao critica
func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}
