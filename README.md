## Project Structure

```
go-eshop-console/
├── go.mod                # Go module file
├── go.sum                # Dependency checksum file
├── main.go               # Entry point of the application
├── models/               # Folder for entities
│   ├── product.go        # Product struct and related methods
│   ├── order.go          # Order struct and related methods
│   ├── admin.go          # Admin struct (username, password)
│   ├── customer.go       # Customer struct
├── handlers/             # Folder for handling logic
│   ├── admin.go          # Admin-specific functionality (status updates, paging)
│   ├── customer.go       # Customer-specific functionality (invoice generation)
│   ├── order.go          # Order processing logic
│   ├── menu.go           # Menu rendering and navigation logic
├── services/             # Folder for service-related code
│   ├── currency.go       # Currency conversion logic using an external API
│   ├── payment.go        # Payment processing logic
├── utils/                # Folder for utility functions
│   ├── file.go           # File operations (saving invoices, reading config)
│   ├── paginator.go      # Pagination logic for orders
├── configs/              # Configuration files
│   ├── config.json       # Config file for API keys or settings
├── data/                 # Folder for storing data files
│   ├── products.json     # Sample product data
│   ├── orders.json       # Sample order data
│   ├── customers.json    # Sample customer data
├── README.md             # Documentation for the application
└── .gitignore            # Git ignore file for excluding unnecessary files (e.g., compiled binaries)
```

## App Highlights

### Example: Admin Menu:
```
Admin Menu - [yyyy/MM/dd HH:mm]
---------------------------------  
Name: ...
Email: ...

[1] Add Product
[2] Update Product
[3] Delete Product
[4] View All Products
[5] View All Orders (Paginated)
[6] View All Customers
[0] Exit

Select an option: _
```

### Example: Customer Menu

```
Hello [Customer name] - [yyyy/MM/dd HH:mm]
---------------------------------  
Name: ...
Email: ...
Previous Order: [Identifier] [Date]

Customer Menu:
[1] View Products
[2] Add Product to Cart
[3] View Cart
[4] Checkout (Generate Invoice & Currency Conversion)
[5] View Orders
[0] Exit
```

### Example: Pagination Workflow:
```
Admin: View All Orders (Page 1)  
---------------------------------  
Order ID: 1 | Status: Pending ... 
Order ID: 2 | Status: Processing ... 
Order ID: 3 | Status: Completed  ...
...  
[15 orders shown]  

Options:
[N] Next Page   [B] Back

[1] Change Order Status
[2] Display Order Details
[0] Exit orders

Select an option: _

```