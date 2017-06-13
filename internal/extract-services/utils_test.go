package internal

import (
	"fmt"
	"testing"
)

func TestGetPDF(t *testing.T) {
	ret, err := getPDF("6760-1315010-118", "FR")
	if err != nil {
		fmt.Printf("getPDF: %v\n", err)
	}
	fmt.Println(ret)
}
