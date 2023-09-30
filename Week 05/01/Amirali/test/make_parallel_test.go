package main

import (
	"make/mymake"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var pm *mymake.Make

func TestParaExecSimple(t *testing.T) {
	out, err := pm.ExecuteParallel("compile")
	assert.NoError(t, err)
	assert.Equal(t, "compile\n", out)
}

func TestParaExecEmpty(t *testing.T) {
	out, err := pm.ExecuteParallel("empty")
	assert.NoError(t, err)
	assert.Equal(t, "", out)
}

func TestParaExecWithDep1(t *testing.T) {
	out, err := pm.ExecuteParallel("all2")
	assert.NoError(t, err)

	assert.Contains(t, out, "clean")
	assert.Contains(t, out, "compile")
	assert.Contains(t, out, `all
echo "done"
done
`)
	assert.True(t, strings.HasSuffix(out, "done\n"))
}

func TestParaExecWithDep2(t *testing.T) {
	out, err := pm.ExecuteParallel("all")
	assert.NoError(t, err)

	assert.Contains(t, out, "compiling 1")
	assert.Contains(t, out, "compiling 3")
	assert.Contains(t, out, "compiling 2")

	assert.Less(t, strings.Index(out, "3"), strings.Index(out, "2"))

	assert.True(t, strings.HasSuffix(out, "linking\ndone\n"))
}

func TestParaExecMultipleTarget(t *testing.T) {
	out, err := pm.ExecuteParallel("all", "all2")
	assert.NoError(t, err)

	assert.Contains(t, out, "compiling 1")
	assert.Contains(t, out, "compiling 3")
	assert.Contains(t, out, "compiling 2")

	assert.Contains(t, out, "linking")
	assert.Contains(t, out, "done")
	assert.Contains(t, out, "clean")
	assert.Contains(t, out, "compile")
	assert.Contains(t, out, "all")
	assert.Contains(t, out, `echo "done"`)

	assert.Less(t, strings.Index(out, "3"), strings.Index(out, "2"))
	assert.Less(t, strings.Index(out, "all"), strings.LastIndex(out, "done"))
}

func TestParaExecDefaultTarget(t *testing.T) {
	out, err := pm.ExecuteParallel()
	assert.NoError(t, err)

	assert.Contains(t, out, "compiling 1")
	assert.Contains(t, out, "compiling 3")
	assert.Contains(t, out, "compiling 2")

	assert.Less(t, strings.Index(out, "3"), strings.Index(out, "2"))

	assert.True(t, strings.HasSuffix(out, "linking\ndone\n"))
}

func TestParaExecCircle1(t *testing.T) {
	out, err := pm.ExecuteParallel("circle1")

	assert.NoError(t, err)

	assert.Contains(t, out, "compiling 1")
	assert.Contains(t, out, "compiling 3")
	assert.Contains(t, out, "compiling 2")

	assert.Less(t, strings.Index(out, "3"), strings.Index(out, "2"))

	assert.Equal(t, len("compiling 1\n")*3, len(out))
}

func TestParaExecCircle2(t *testing.T) {
	out, err := pm.ExecuteParallel("circle2")
	assert.NoError(t, err)

	assert.Contains(t, out, "compiling 1")
	assert.Contains(t, out, "compiling 3")
	assert.Contains(t, out, "compiling 2")

	assert.Less(t, strings.Index(out, "3"), strings.Index(out, "2"))

	assert.Equal(t, len("compiling 1\n")*3, len(out))
}

func TestParaExecInvalidSyntax(t *testing.T) {
	_, err := pm.ExecuteParallel("invalid_syntax")
	assert.Error(t, err)
}

func TestParaExecInvalidCommand(t *testing.T) {
	_, err := pm.ExecuteParallel("invalid_command")
	assert.Error(t, err)
}

func TestParaExecInvalidTarget(t *testing.T) {
	_, err := pm.ExecuteParallel("invalid_target")
	assert.Error(t, err)
}
