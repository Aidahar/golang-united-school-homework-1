package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func main() {

	word := emoji.Sprint("Hello :world_map:!")
	fmt.Println(word)
}
