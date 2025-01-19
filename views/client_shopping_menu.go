package views

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"time"
)

func DisplayClientShoppingMenu() {
	fmt.Println("WSC - Shopping 🛒")
	currentTime := time.Now().Format("2006-01-02 15:04")
	fmt.Printf("Current time: %s\n", currentTime)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"", ""})

	table.Append([]string{"[1]", "List All Products"})
	table.Append([]string{"[2]", "List Products (Paginated)"})
	table.Append([]string{"[3]", "Get Product details"})
	table.Append([]string{"[0]", "Quit 🚪"})

	table.Render()

}
