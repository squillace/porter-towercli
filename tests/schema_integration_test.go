// +build integration

package tests

import (
	"bytes"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

// Add a test that checked that the schema was packed into the binary
// properly. Requires a make clean xbuild-all first.
func TestSchema(t *testing.T) {
	schemaBackup := "../pkg/towercli/schema/schema.json.bak"
	schemaPath := "../pkg/towercli/schema/schema.json"

	defer os.Rename(schemaBackup, schemaPath)
	err := os.Rename(schemaPath, schemaBackup)
	require.NoError(t, err, "failed to sabotage the schema.json")

	output := &bytes.Buffer{}
	cmd := exec.Command("../bin/mixins/towercli/towercli", "schema")
	cmd.Stdout = output
	cmd.Stderr = output

	err = cmd.Start()
	require.NoError(t, err, "failed to start the towercli schema command")

	err = cmd.Wait()
	t.Log(output)
	require.NoError(t, err, "towercli schema failed")
}
