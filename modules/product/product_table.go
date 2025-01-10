package product_scope

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func PrintProductTable(products []Product) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetCaption(true, "WSC - Products Stock 📦")
	table.SetHeader([]string{"ID", "Name", "SKU", "Price"})

	for _, product := range products {
		table.Append([]string{
			fmt.Sprintf("%d", product.ID),
			limitStringLength(product.Name, 45),
			product.SKU,
			fmt.Sprintf("%.2f", product.Price),
		})
	}
	table.Render()
}

func limitStringLength(str string, length int) string {
	if len(str) > length {
		return str[:length] + "..."
	}
	return str
}
