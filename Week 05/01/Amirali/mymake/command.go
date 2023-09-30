package mymake

import (
	"bytes"
	"errors"
	"io"
	"make/os_util"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Make struct {
	Path       string
	TargetList []string
	Compiled   []string
}

type Target struct {
	Name     string
	Deps     []string
	Commands []string
}

func NewMake(dir string) *Make {
	return &Make{
		Path: dir,
	}
}

func (m *Make) Parse() (map[string]Target, error) {
	targets := map[string]Target{}
	makefile := filepath.Join(m.Path, "makefile")
	makefileContent, _ := os.ReadFile(makefile)
	lines := strings.Split(string(makefileContent), "\n")
	lastTarget := ""
	for _, line := range lines {
		if len(line) == 0 {
			lastTarget = ""
			continue
		}

		re := regexp.MustCompile(`(\w+)(:\s?)([\w\ ]*)`)
		matches := re.FindStringSubmatch(line)
		if len(matches) > 0 {
			lastTarget = matches[1]
			m.TargetList = append(m.TargetList, lastTarget)

			if matches[3] != "" {
				targets[lastTarget] = Target{
					Name:     lastTarget,
					Deps:     strings.Split(string(matches[3]), " "),
					Commands: []string{},
				}
			} else {
				targets[lastTarget] = Target{
					Name:     lastTarget,
					Commands: []string{},
				}
			}
			continue
		}

		re = regexp.MustCompile(`^([\t\ ]+)(.*)`)
		matches = re.FindStringSubmatch(line)
		if len(matches) > 0 && lastTarget != "" && matches[2] != "" {
			targets[lastTarget] = Target{
				Name:     lastTarget,
				Deps:     targets[lastTarget].Deps,
				Commands: append(targets[lastTarget].Commands, matches[2]),
			}
		}
	}
	return targets, nil
}

func (m *Make) execute(target, commands string, outputWriter io.Writer) error {
	m.Compiled = append(m.Compiled, target)
	if strings.HasPrefix(commands, "@") {
		commands = commands[1:]
	} else {
		outputWriter.Write([]byte(commands + "\n"))
	}
	return os_util.Run(commands, outputWriter)
}

func (m *Make) executeDeps(allTargets map[string]Target, dep string, outputWriter io.Writer) error {
	targetDep, isExit := allTargets[dep]
	if !isExit {
		return errors.New("invalid target")
	}
	for _, c := range m.Compiled {
		if c == dep {
			return nil
		}
	}
	for _, deps := range targetDep.Deps {
		m.executeDeps(allTargets, deps, outputWriter)
	}
	for _, commands := range targetDep.Commands {
		err := m.execute(dep, commands, outputWriter)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Make) Execute(targets ...string) (string, error) {
	var result bytes.Buffer
	allTargets, err := m.Parse()
	if len(targets) == 0 && len(allTargets) > 0 {
		targets = append(targets, m.TargetList[0])
	}

	m.Compiled = []string{}

	for _, target := range targets {
		t, isExit := allTargets[target]
		if !isExit {
			return "", errors.New("invalid target")
		}
		for _, deps := range t.Deps {
			m.executeDeps(allTargets, deps, &result)
		}
		for _, commands := range t.Commands {
			err = m.execute(target, commands, &result)
			if err != nil {
				break
			}
		}
		if err != nil {
			break
		}
	}
	return result.String(), err
}

func (m *Make) ExecuteParallel(targets ...string) (string, error) {
	var result bytes.Buffer
	allTargets, err := m.Parse()
	if len(targets) == 0 && len(allTargets) > 0 {
		targets = append(targets, m.TargetList[0])
	}

	m.Compiled = []string{}

	for _, target := range targets {
		t, isExit := allTargets[target]
		if !isExit {
			return "", errors.New("invalid target")
		}
		for _, deps := range t.Deps {
			m.executeDeps(allTargets, deps, &result)
		}
		for _, commands := range t.Commands {
			err = m.execute(target, commands, &result)
			if err != nil {
				break
			}
		}
		if err != nil {
			break
		}
	}
	return result.String(), err
}
