## Commands
### Go commands
```bash
# Initialize a new Go module in the current directory
go mod init <module-name>

# Add a dependency to the module
go get <package-name>
go get -u <package-name>@none

# Remove unused dependencies and tidy up go.mod and go.sum
go mod tidy

# Build the Go project
go build

# Force-update all dependencies to their latest versions
go get -u ./...

# View a graph of module dependencies
go mod graph

# Verify dependencies against go.sum
go mod verify

# Run a Go program
go run <file>.go

# Run tests in the current directory
go test ./...

# Clean up the build cache
go clean -cache

# Clean up the module cache
go clean -modcache

# List all installed Go packages
go list all

# Format Go source code
go fmt ./...

# View documentation for a package
go doc <package>

# Download all dependencies
go mod download

# Install a Go tool (e.g., golangci-lint, staticcheck)
go install <tool>@latest

# Check for any linting issues
golangci-lint run

# Run static analysis on your code
staticcheck ./...

```
---

## Project Structure

```bash
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
```bash
Admin Menu - [yyyy/MM/dd HH:mm]
---------------------------------  
Name: ...
Email: ...

[1] Manage Products
[2] Manage Orders
[3] Manage Customers
[4] Analytics
[5] System Settings
[0] Exit

Select an option: _
```

### Example: Admin Menu (Products Management):
```bash
WSC - Product Management 🛠️ - [yyyy/MM/dd HH:mm]
---------------------------------

[1] List All Products
[2] List Products (Paginated)
[3] Create Product
[4] Update Product
[5] Delete Product
[0] Exit

Select an option: _
```

---

### Example: Customer Menu

```bash
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
```bash
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

---

## Git Guidelines

This document outlines the naming conventions for branches and commits to ensure clarity, traceability, and consistency in the project.

### **Branch Naming Convention**
Branch names should follow this structure: `<phase>/(WSC-<ticket-number>)-<feature-or-task-name>`

### **Phases**

- `init`: Setting up the initial project structure.
- `feat`: For new features.
- `enhance`: Feature enhancements and improvements.
- `fix`: For bug fixes.
- `docs`: For documentation changes.
- `refactor`: For code restructuring.
- `test`: For adding or improving tests.
- `infra`: infrastructure management, ci/cd, including cloud setup, networking, servers

### **Examples**
- `init/(WSC-<ticket-number>)-project-setup`: Setting up the initial project structure.
- `feat/(WSC-4)-product-listing`: Adding a new feature for listing products.
- `enhance/(WSC-8)-filter-products`: Enhancing filtering options for products.
- `fix/(WSC-12)-error-handling`: Fixing bugs in error handling.
- `docs/(WSC-15)-update-readme`: Updating project documentation.
- `refactor/(WSC-11)-catalog-module`: Restructuring the catalog module for efficiency.
- `test/(WSC-14)-add-unit-tests`: Adding unit tests for core functionality.
- `infra/(WSC-15)-add-azure-logging`: Adding Logging to Azure

### **Commit Message Rules**

Commit messages should follow this format: `type(WSC-<ticket-number>):<short description>`

### **Best Practices**
1. **Write Descriptive Commit Messages**:
   - Keep them concise yet informative.
   - Example: `fix(WSC-2): handle invalid payment input gracefully`.

2. **Group Related Changes in Branches**:
   - Focus each branch on a single task or feature.
   - Avoid unrelated changes in a single branch.

3. **Keep Commits Atomic**:
   - Each commit should represent a logical change.
   - Example: Separate commits for adding functionality and fixing bugs.

4. **Use Pull Requests (PRs)**:
   - Always create a pull request for branch merges.
   - Include detailed descriptions of changes in the PR.

### **Workflow Example**

1. **Branch Creation**:
   - Create a branch for the task: `core/product-listing`.

2. **Commit Example**:
   - `feat(PAT): implement product listing with stock status`

3. **Pull Request**:
   - PR Title: `Implement product listing feature`
   - PR Description:
     - Adds functionality to display product list with stock status.
     - Handles edge cases for empty product list.

---