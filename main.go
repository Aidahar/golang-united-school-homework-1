package solution

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func GetMessage() {
	word := emoji.Sprint("Hello :world_map:!")
	fmt.Println(word)
}
