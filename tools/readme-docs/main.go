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

	"github.com/lindell/connect-iq-manager/cmd"
)

const templatePath = "./docs/README.template.md"
const resultingPath = "./README.md"

type templateData struct {
	MainUsage string
	Commands  []command
}

type command struct {
	Name        string
	Long        string
	Short       string
	Usage       string
	YAMLExample string
}

func main() {
	data := templateData{}

	// Main usage
	data.MainUsage = strings.TrimSpace(cmd.RootCmd().UsageString())

	subCommands := cmd.RootCmd().Commands()

	// All commands
	cmds := []struct {
		cmd *cobra.Command
	}{
		{
			cmd: commandByName(subCommands, "version"), // TODO: Update to real command
		},
	}
	for _, c := range cmds {
		data.Commands = append(data.Commands, command{
			Name:        c.cmd.Name(),
			Long:        c.cmd.Long,
			Short:       c.cmd.Short,
			Usage:       strings.TrimSpace(c.cmd.UsageString()),
			YAMLExample: getYAMLExample(c.cmd),
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

func commandByName(cmds []*cobra.Command, name string) *cobra.Command {
	for _, command := range cmds {
		if command.Name() == name {
			return command
		}
	}
	panic(fmt.Sprintf(`could not find command "%s"`, name))
}
