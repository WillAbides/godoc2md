package goreadme_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/WillAbides/godoc2md/goreadme"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testPkg = "github.com/WillAbides/godoc2md/goreadme/testdata/pkgs/pkg1"

func updateTestdata() error {
	err := goreadme.WriteReadme(testPkg, "testdata/want/pkgs/pkg1/README.md")
	return err
}

func TestReadmeMD(t *testing.T) {
	want, err := ioutil.ReadFile(filepath.FromSlash("testdata/want/pkgs/pkg1/README.md"))
	require.NoError(t, err)
	var buf bytes.Buffer
	err = goreadme.ReadmeMD(testPkg, &buf)
	assert.NoError(t, err)
	assert.Equal(t, string(want), buf.String())
}

func TestWriteReadme(t *testing.T) {
	want, err := ioutil.ReadFile(filepath.FromSlash("testdata/want/pkgs/pkg1/README.md"))
	require.NoError(t, err)
	tmpDir, err := ioutil.TempDir("", "")
	require.NoError(t, err)
	readmePath := filepath.Join(tmpDir, "README.md")
	err = goreadme.WriteReadme(testPkg, readmePath)
	assert.NoError(t, err)
	got, err := ioutil.ReadFile(readmePath) //nolint:gosec
	require.NoError(t, err)
	assert.Equal(t, string(want), string(got))
	require.NoError(t, os.RemoveAll(tmpDir))
}

func TestVerifyReadme(t *testing.T) {
	t.Run("readme doesn't exist", func(t *testing.T) {
		ok, err := goreadme.VerifyReadme(testPkg, "this/should/not/exists")
		assert.NoError(t, err)
		assert.False(t, ok)
	})

	t.Run("matches", func(t *testing.T) {
		ok, err := goreadme.VerifyReadme(testPkg, filepath.FromSlash("testdata/want/pkgs/pkg1/README.md"))
		assert.NoError(t, err)
		assert.True(t, ok)
	})

	t.Run("doesn't match", func(t *testing.T) {
		ok, err := goreadme.VerifyReadme(testPkg, filepath.FromSlash("testdata/want/pkgs/pkg1/badREADME.md"))
		assert.NoError(t, err)
		assert.False(t, ok)
	})
}
