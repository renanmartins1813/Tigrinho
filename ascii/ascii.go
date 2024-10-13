package ascii

import (
	"io/ioutil"
)

func Ganhou() string {
	b, err := ioutil.ReadFile("./ascii/ascii.txt")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func Perdeu() string {
	b, err := ioutil.ReadFile("./ascii/perdeu.txt")
	if err != nil {
		panic(err)
	}
	return string(b)
}
