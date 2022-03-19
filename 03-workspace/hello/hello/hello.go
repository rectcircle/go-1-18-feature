package hello

import (
	"fmt"

	stringutil "github.com/rectcircle/go-1-18-feature/03-workspace/util/string"
)

func Say() {
	fmt.Println(stringutil.Reverse("Hello"))
	fmt.Println(stringutil.ToUpper("Hello"))
	fmt.Println(stringutil.ToLower("Hello"))
}
