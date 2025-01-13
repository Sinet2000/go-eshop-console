package tables

import (
	"fmt"
	"os"

	"github.com/Sinet2000/go-eshop-console/internal/entities"
	"github.com/olekukonko/tablewriter"
)

func ListProducts(products []entities.Product) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetCaption(true, "WSC - Products Stock 📦")
	table.SetHeader([]string{"ID", "Name", "SKU", "Price"})

	for _, product := range products {
		var idString string
		if product.ID.IsZero() {
			idString = "N/A"
		} else {
			idString = product.ID.Hex()
		}

		table.Append([]string{
			idString,
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
