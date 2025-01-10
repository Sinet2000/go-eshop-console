<<<<<<<< HEAD:internal/entities/order.go
package entities
========
package order_scope
>>>>>>>> ca972b1 (#WSC-3: Product seeding & menu output & improved structure):modules/order/order.go

import (
	"fmt"
	"time"

	"github.com/Sinet2000/go-eshop-console/exceptions"
<<<<<<<< HEAD:internal/entities/order.go
========
	product_scope "github.com/Sinet2000/go-eshop-console/modules/product"
>>>>>>>> ca972b1 (#WSC-3: Product seeding & menu output & improved structure):modules/order/order.go
)

type OrderStatus int

// iota is a predefined identifier that is used to simplify the definition of incrementing constant values in const blocks.
const (
	Pending           OrderStatus = iota // 0 - Order placed, but not yet processed
	Shipped                              // 1 - Order shipped to the customer
	Delivered                            // 2 - Order delivered to the customer
	Cancelled                            // 3 - Order cancelled
	WaitingForPayment                    // 4 - Order placed, awaiting payment (after invoice sent)
	Paid                                 // 5 - Order paid (cash, credit, etc.)
	OrderConfirmed                       // 6 - Order placed and confirmed (for private clients, no invoice yet)
	RefundRequested                      // 7 - Customer has requested a refund
	Refunded                             // 8 - Refund processed and completed
)

type Order struct {
	ID           int
	CustomerID   int
	Status       OrderStatus
	TotalPrice   float64
	Currency     string
	OrderDate    time.Time
	ShippingDate time.Time
	Products     []product_scope.Product
}

func CreateOrder(customerID int, currency string) (*Order, error) {
	if customerID < 1 {
		return nil, &exceptions.DomainException{Message: "Invalid customer ID"}
	}

	order := &Order{
		CustomerID: customerID,
		Status:     Pending,
		Currency:   currency,
		OrderDate:  time.Now(),
		Products:   []product_scope.Product{},
	}

	return order, nil
}

func (o *Order) ChangeStatus(newStatus OrderStatus) error {
	if !o.CanChangeStatusTo(newStatus) {
		return &exceptions.DomainException{
			Message: fmt.Sprintf("Cannot transition order from %s to %s", o.Status.String(), newStatus.String()),
		}
	}

	o.Status = newStatus
	return nil
}

func (o *Order) CanChangeStatusTo(newStatus OrderStatus) bool {
	switch o.Status {
	case Pending:
		return newStatus == WaitingForPayment || newStatus == Shipped || newStatus == Cancelled
	case Shipped:
		return newStatus == Delivered || newStatus == Cancelled
	case WaitingForPayment:
		return newStatus == Paid || newStatus == Cancelled
	case Paid:
		return newStatus == Shipped || newStatus == Cancelled
	case Delivered, Refunded, Cancelled:
		return false // No more status changes after these
	}
	return false
}

func (o *Order) CalculateTotalPrice() float64 {
	var total float64
	for _, product := range o.Products {
		total += product.Price
	}
	// Add business logic for taxes, discounts, etc. here
	return total
}

func (o *Order) GetOrderSummary() string {
	return fmt.Sprintf("Order ID: %d\nCustomer ID: %d\nStatus: %s\nTotal Price: %.2f %s\nOrder Date: %s\nShipping Date: %s\n",
		o.ID, o.CustomerID, o.Status.String(), o.TotalPrice, o.Currency, o.OrderDate.Format("2006-01-02"), o.ShippingDate.Format("2006-01-02"))
}

// String method to convert OrderStatus to string (for easy printing)
func (os OrderStatus) String() string {
	switch os {
	case Pending:
		return "Pending"
	case Shipped:
		return "Shipped"
	case Delivered:
		return "Delivered"
	case Cancelled:
		return "Cancelled"
	case WaitingForPayment:
		return "WaitingForPayment"
	case Paid:
		return "Paid"
	case OrderConfirmed:
		return "OrderConfirmed"
	case RefundRequested:
		return "RefundRequested"
	case Refunded:
		return "Refunded"
	default:
		return "Unknown"
	}
}
