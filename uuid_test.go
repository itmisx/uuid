package uuid

import (
	"fmt"
	"testing"
)

func TestUUid(t *testing.T) {
	Init(1, 1)
	fmt.Println(UUID{}.String())
}
