package manager

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func createTable() *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t")
	table.SetNoWhiteSpace(true)

	return table
}
