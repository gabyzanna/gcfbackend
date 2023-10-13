package gcfbackend

import (
	"fmt"
	"testing"
)

func TestGCHandlerFunc(t *testing.T) {
	data := GCHandlerFunc("string", "gisjson", "gisjson")

	fmt.Printf("%+v", data)
}