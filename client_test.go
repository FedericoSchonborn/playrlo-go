package playrlo_test

import (
	_ "embed"
	"testing"

	"github.com/FedericoSchonborn/playrlo"
)

//go:embed testdata/format.rs
var formatCode string

func TestClientFormat(t *testing.T) {
	client := playrlo.NewClient(&playrlo.Config{
		UserAgent: "playrlo-test/0.0.0",
	})

	resp, err := client.Format(playrlo.FormatRequest{
		Code:    formatCode,
		Edition: playrlo.Rust2021,
	})
	if err != nil {
		t.Fatalf("/format call failed: %v", err)
	}

	t.Logf("%#v", resp)
}
