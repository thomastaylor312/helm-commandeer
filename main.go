package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const globalUsage = `
Commandeer cluster resources into a new chart

By default, commandeer expects to match all resources given. When using a label
query, it expects to match at least one resource in the cluster. It will error
if these conditions are not met. To change this behavior, use the --ignore-missing
flag. 

You can combine the list of specific resources with a label query to match additional
resources. The label query functions exactly like label queries in kubectl.

Basic example:

	$ helm commandeer deployment:my-deployment service:my-service persistentvolumeclaim:my-pvc
`

var (
	labelQuery    string
	ignoreMissing bool
	debug         bool
	dryRun        bool
	releaseName   string
	outputPath    string
)

var version = "DEV"

func main() {
	cmd := &cobra.Command{
		Use:   "commandeer [flags] []<kind>:<name>",
		Short: fmt.Sprintf("commandeer cluster resources into a new chart (helm-commandeer %s)", version),
		Long:  globalUsage,
		RunE:  commandeer,
	}

	f := cmd.Flags()
	f.BoolVarP(&debug, "debug", "v", false, "show the generated manifests on STDOUT")
	f.BoolVar(&dryRun, "dry-run", false, "show what resources will be commandeered")
	f.BoolVarP(&debug, "ignore-missing", "i", false, "if a specifed resource does not exist, ignore it instead of erroring."+
		"There must still be at least 1 matching resource between the list of <kind>:<name> and label queries")
	f.StringVarP(&releaseName, "name", "n", "", "the name you want to use for the generated release. If not provided, one will be generated for you")
	// We won't actually get the working directory yet to print here to avoid the error handling unless we absolutely need to
	f.StringVarP(&outputPath, "output", "o", "", "where to output the generated chart. Defaults to the current working directory")
	f.StringVarP(&labelQuery, "label", "l", "", "a valid label query to select resources")

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func commandeer(cmd *cobra.Command, args []string) error {
	fmt.Printf("Commandeer called with args: %v\n", args)
	return nil
}
