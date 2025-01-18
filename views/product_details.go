package views

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Sinet2000/go-eshop-console/internal/entities"
	"github.com/olekukonko/tablewriter"
)

func DisplayProductDetails(product *entities.Product) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"", ""})
	table.SetBorder(true)

	table.Append([]string{"ID", product.ID.Hex()})
	table.Append([]string{"SKU", product.SKU})
	table.Append([]string{"Name", product.Name})
	table.Append([]string{"Description", product.Description})
	table.Append([]string{"Price", fmt.Sprintf("%.2f", product.Price)})
	table.Append([]string{"Stock", strconv.Itoa(product.Stock)})

	table.Render()
}
