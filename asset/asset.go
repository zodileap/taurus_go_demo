package asset

import (
	"fmt"

	"github.com/yohobala/taurus_go/asset"
)

func main() {
	var assets asset.Assets
	assets.Add("test.txt", []byte("Hellow, World!"))
	if err := assets.Write(); err != nil {
		fmt.Print(err)
	}
}
