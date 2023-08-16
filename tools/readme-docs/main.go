package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"golang.org/x/exp/slices"

	"github.com/lindell/connect-iq-sdk-manager-cli/cmd"
)

const templatePath = "./docs/README.template.md"
const resultingPath = "./README.md"

// These commands should not show up in the readme docs
var bannedCommands = []string{"version"}

type templateData struct {
	MainUsage string
	Commands  []command
}

type command struct {
	Name        string
	Path        string
	Long        string
	Short       string
	Usage       string
	YAMLExample string
}

func main() {
	data := templateData{}

	// Main usage
	data.MainUsage = strings.TrimSpace(cmd.RootCmd().UsageString())

	rootPath := cmd.RootCmd().CommandPath() + " "

	// All commands
	cmds := viewableCommands(cmd.RootCmd())
	for _, cmd := range cmds {
		name := strings.TrimPrefix(cmd.CommandPath(), rootPath)
		data.Commands = append(data.Commands, command{
			Name:        name,
			Path:        strings.ReplaceAll(name, " ", "-"),
			Long:        cmd.Long,
			Short:       cmd.Short,
			Usage:       strings.TrimSpace(cmd.UsageString()),
			YAMLExample: getYAMLExample(cmd),
		})
	}

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Fatal(err)
	}

	tmplBuf := &bytes.Buffer{}
	err = tmpl.Execute(tmplBuf, data)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(resultingPath, tmplBuf.Bytes(), 0600)
	if err != nil {
		log.Fatal(err)
	}
}

// Replace some of the default values in the yaml example with these values
var yamlExamples = map[string]string{}

var listDefaultRegex = regexp.MustCompile(`^\[(.+)\]$`)

func getYAMLExample(cmd *cobra.Command) string {
	if cmd.Flag("config") == nil {
		return ""
	}

	b := strings.Builder{}
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if f.Name == "config" {
			return
		}

		// Determine how to format the example values
		val := f.DefValue
		if val == "-" {
			val = ` "-"`
		} else if val == "[]" {
			val = "\n  - example"
		} else if matches := listDefaultRegex.FindStringSubmatch(val); matches != nil {
			val = "\n  - " + strings.Join(strings.Split(matches[1], ","), "\n  - ")
		} else if val != "" {
			val = " " + val
		}

		if replacement, ok := yamlExamples[f.Name]; ok {
			val = replacement
		}

		usage := strings.Split(strings.TrimSpace(f.Usage), "\n")
		for i := range usage {
			usage[i] = "# " + usage[i]
		}

		b.WriteString(fmt.Sprintf("%s\n%s:%s\n\n", strings.Join(usage, "\n"), f.Name, val))
	})
	return strings.TrimSpace(b.String())
}

func viewableCommands(cmd *cobra.Command) []*cobra.Command {
	cmds := []*cobra.Command{}

	if cmd.Runnable() && !slices.Contains(bannedCommands, cmd.Name()) {
		cmds = append(cmds, cmd)
	}

	for _, command := range cmd.Commands() {
		cmds = append(cmds, viewableCommands(command)...)
	}

	return cmds
}
