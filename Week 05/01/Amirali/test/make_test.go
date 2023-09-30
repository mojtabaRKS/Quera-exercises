package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	assert.NotNil(t, m)
}

func TestParseTargets(t *testing.T) {
	targets, err := m.Parse()
	assert.NoError(t, err)
	assert.NotEmpty(t, targets)
	names := make([]string, 0, len(targets))
	for _, target := range targets {
		names = append(names, target.Name)
	}
	assert.ElementsMatch(t,
		[]string{
			"all", "compile1", "compile2", "compile3",
			"all2", "clean", "compile", "empty",
			"circle1", "circle2",
			"invalid_syntax", "invalid_command",
		},
		names)
}

func TestParseDeps(t *testing.T) {
	targets, err := m.Parse()
	assert.NoError(t, err)
	assert.ElementsMatch(t, []string{"clean", "compile"}, targets["all2"].Deps)
	assert.Empty(t, targets["clean"].Deps)
	assert.Empty(t, targets["compile"].Deps)
}

func TestParseCommands(t *testing.T) {
	targets, err := m.Parse()
	assert.NoError(t, err)
	assert.Equal(t, []string{
		`@echo "all"`,
		`echo "done"`,
	}, targets["all2"].Commands)
	assert.Equal(t, []string{
		`@echo "compile"`,
	}, targets["compile"].Commands)
	assert.Equal(t, []string{
		`@echo "clean"`,
	}, targets["clean"].Commands)
	assert.Equal(t, []string{}, targets["empty"].Commands)
}

func TestExecSimple(t *testing.T) {
	out, err := m.Execute("compile")
	assert.NoError(t, err)
	assert.Equal(t, "compile\n", out)
}

func TestExecEmpty(t *testing.T) {
	out, err := m.Execute("empty")
	assert.NoError(t, err)
	assert.Equal(t, "", out)
}

func TestExecWithDep1(t *testing.T) {
	out, err := m.Execute("all2")
	assert.NoError(t, err)
	assert.Equal(t,
		`clean
compile
all
echo "done"
done
`, out)
}

func TestExecWithDep2(t *testing.T) {
	out, err := m.Execute("all")
	assert.NoError(t, err)
	assert.Equal(t,
		`compiling 1
compiling 3
compiling 2
linking
done
`, out)
}

func TestExecMultipleTarget(t *testing.T) {
	out, err := m.Execute("all", "all2")
	assert.NoError(t, err)
	assert.Equal(t,
		`compiling 1
compiling 3
compiling 2
linking
done
clean
compile
all
echo "done"
done
`, out)
}

func TestExecDefaultTarget(t *testing.T) {
	out, err := m.Execute()
	assert.NoError(t, err)
	assert.Equal(t,
		`compiling 1
compiling 3
compiling 2
linking
done
`, out)
}

func TestExecCircle1(t *testing.T) {

	out, err := m.Execute("circle1")
	assert.NoError(t, err)
	assert.Equal(t,
		`compiling 3
compiling 2
compiling 1
`, out)
}
func TestExecCircle2(t *testing.T) {

	out, err := m.Execute("circle2")
	assert.NoError(t, err)
	assert.Equal(t,
		`compiling 1
compiling 3
compiling 2
`, out)
}

func TestExecInvalidSyntax(t *testing.T) {
	_, err := m.Execute("invalid_syntax")
	assert.Error(t, err)
}

func TestExecInvalidCommand(t *testing.T) {
	_, err := m.Execute("invalid_command")
	assert.Error(t, err)
}

func TestExecInvalidTarget(t *testing.T) {
	_, err := m.Execute("invalid_target")
	assert.Error(t, err)
}
