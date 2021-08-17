package apple

import (
	"fmt"
	"testing"
)

func TestSayHello(t *testing.T) {
	msg := sayHello("apple")
	fmt.Println(msg)
}
