package client

import (
	"testing"

	"github.com/loqutus/O-1/pkg/types"
)

func TestInfo(t *testing.T) {
	types.Client.HostName = "localhost"
	types.Client.Port = "6969"
	err := Info()
	if err != nil {
		t.Fatal(err)
	}
}
