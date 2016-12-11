package simplelog

import "testing"

func TestOutput(t *testing.T) {
	output(DEBUG, "It's %s", "OK")
}

func TestDebug(t *testing.T) {
	Debug("It's %s", "OK")
}
