package main

import (
	"github.com/AnuchitO/say/v2"
	oldsay "github.com/AnuchitO/say"
)

func main() {
	say.Something("new before", "my talk")
	oldsay.Something("Old say")
}
