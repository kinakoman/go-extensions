package caller

import (
	"fmt"
	"testing"
)

func TestHere(t *testing.T) {
	fmt.Println(Here(0).Format())
	fmt.Println(Here(0).FormatShort())
}
