package test

import (
	"fmt"
	"testing"

	"../utils"
)

func TestGenerateUUID(t *testing.T) {
	uuids := make(map[string]int)
	for i := 1; i <= 100; i++ {
		uuids[utils.GenerateUUID()]++
	}
	for k, v := range uuids {
		fmt.Printf("%s:%d\n", k, v)
		if v > 1 {
			t.Error("Duplicate UUID found")
		}
	}
}
