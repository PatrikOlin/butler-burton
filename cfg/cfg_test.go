package cfg

import (
	"os"
	"testing"
)

func TestGetConfigPath(t *testing.T) {
	got := GetConfigPath()
	want := os.Getenv("HOME") + "/.config/butlerburton/config.yml"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
