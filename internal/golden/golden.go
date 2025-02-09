// Package golden asserts golden files contents.
package golden

import (
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

var update = flag.Bool("update", false, "update .golden files")

const golden = ".golden"

func RequireEqual(tb testing.TB, out []byte) {
	tb.Helper()
	RequireEqualExt(tb, out, "")
}

func RequireEqualExt(tb testing.TB, out []byte, ext string) {
	tb.Helper()
	doRequireEqual(tb, out, ext, golden)
}

func RequireEqualTxt(tb testing.TB, out []byte) {
	tb.Helper()
	doRequireEqual(tb, out, ".txt", golden)
}

func RequireEqualJSON(tb testing.TB, out []byte) {
	tb.Helper()
	doRequireEqual(tb, out, ".json", golden)
}

func RequireEqualRb(tb testing.TB, out []byte) {
	tb.Helper()
	doRequireEqual(tb, out, ".rb", golden)
}

func RequireEqualYaml(tb testing.TB, out []byte) {
	tb.Helper()
	doRequireEqual(tb, out, ".yaml", "")
}

func RequireReadFile(tb testing.TB, path string) []byte {
	tb.Helper()
	bts, err := os.ReadFile(path)
	require.NoError(tb, err)
	return bts
}

func doRequireEqual(tb testing.TB, out []byte, ext, suffix string) {
	tb.Helper()

	golden := "testdata/" + tb.Name() + ext + suffix
	if *update {
		require.NoError(tb, os.MkdirAll(filepath.Dir(golden), 0o755))
		require.NoError(tb, os.WriteFile(golden, out, 0o655))
	}

	gbts, err := os.ReadFile(golden)
	require.NoError(tb, err)

	require.Equal(tb, string(gbts), string(out))
}
