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

---

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
- `chore`: For maintenance tasks like setup.
- `docs`: For documentation changes.
- `refactor`: For code restructuring.
- `test`: For adding or improving tests.

### **Examples**
- `init/(WSC-<ticket-number>)-project-setup`: Setting up the initial project structure.
- `feat/(WSC-4)-product-listing`: Adding a new feature for listing products.
- `enhance/(WSC-8)-filter-products`: Enhancing filtering options for products.
- `fix/(WSC-12)-error-handling`: Fixing bugs in error handling.
- `chore/(WSC-1)-project-setup`: Setting up the initial project structure.
- `docs/(WSC-15)-update-readme`: Updating project documentation.
- `refactor/(WSC-11)-catalog-module`: Restructuring the catalog module for efficiency.
- `test/(WSC-14)-add-unit-tests`: Adding unit tests for core functionality.

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