<<<<<<<< HEAD:tables/product_table.go
package tables
========
package views
>>>>>>>> 4295e61 (#WSC-3_1: Improved archtiecture, added db support for products):views/product_table.go

import (
	"fmt"
	"os"

	"github.com/Sinet2000/go-eshop-console/internal/entities"
	"github.com/olekukonko/tablewriter"
)

<<<<<<<< HEAD:tables/product_table.go
func ListProducts(products []entities.Product) {
========
func ShowProductTable(products []entities.Product) {
>>>>>>>> 4295e61 (#WSC-3_1: Improved archtiecture, added db support for products):views/product_table.go
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
	table.SetFooter([]string{"", "", "Count", fmt.Sprintf("%d", len(products))})
	table.Render()
}

func limitStringLength(str string, length int) string {
	if len(str) > length {
		return str[:length] + "..."
	}
	return str
}
