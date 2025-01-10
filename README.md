## Project Structure
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