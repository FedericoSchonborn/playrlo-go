package playrlo_test

import (
	_ "embed"
	"testing"

	"github.com/FedericoSchonborn/playrlo"
)

var (
	//go:embed testdata/compile.rs
	compileCode string
	//go:embed testdata/format.rs
	formatCode string
)

var client = playrlo.NewClient(&playrlo.Config{
	UserAgent: "playrlo-test/0.0.0",
})

func TestClientCompile(t *testing.T) {
	resp, err := client.Compile(playrlo.CompileRequest{
		Target:    playrlo.ASM,
		Channel:   playrlo.Nightly,
		Mode:      playrlo.Release,
		Edition:   playrlo.Some(playrlo.Rust2021),
		CrateType: playrlo.Binary,
		Code:      compileCode,
		Tests:     false,
	})
	if err != nil {
		t.Fatalf("/compile call failed: %v", err)
	}

	if !resp.Success {
		t.Fatal("!success")
	}

	t.Logf("%v", resp)
}

func TestClientFormat(t *testing.T) {
	resp, err := client.Format(playrlo.FormatRequest{
		Code:    formatCode,
		Edition: playrlo.Some(playrlo.Rust2021),
	})
	if err != nil {
		t.Fatalf("/format call failed: %v", err)
	}

	if !resp.Success {
		t.Fatal("!success")
	}

	t.Logf("%v", resp)
}
