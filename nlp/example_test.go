package nlp_test

import (
	"fmt"

	"github.com/janhaans/golang/nlp"
)

func ExampleTokenize() {
	text := "Who's on first?"
	tokens := nlp.Tokenize(text)
	fmt.Println(tokens)

	//Output:
	//[who s on first]
}
