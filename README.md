
1. [Git Guidelines](#git-guidelines)
2. [Commands](#commands)
3. [Project Structure](#project-structure)
4. [App Highlights](#app-highlights)
5. [Go Concurrency & Context - Quick Overview](#go-concurrency--context---quick-overview)
6. [MongoDB Docs](#mongodb-docs)

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
WSC - Admin Menu 🛠️
Hello [AdminName] - [yyyy/MM/dd HH:mm]
---------------------------------  
Name: ...
Email: ...

[1] Manage Products
[2] Manage Orders
[3] Manage Customers
[4] Analytics
[5] System Settings
[0] Quit

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
[0] Quit

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

### Example: Products List Pagination:
```bash
Admin: Products (Page 1)  
---------------------------------  
Order ID: 67854f6b298f5112b3ce1a87 | Name: ... 
Order ID: 67854f6b298f5112b3ce1a89 | Name: ... 
Order ID: 67854f6b298f5112b3ce1a8b | Name:  ...
...  
[15 products shown]  

Options:
[N] Next Page   [B] Back

[1] Update product
[2] Display Product Details
[3] Delete Product
[0] Quit

Select an option: _

```


### Example: Orders List Pagination:
```bash
Admin: Orders (Page 1)  
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

## Go Concurrency & Context - Quick Overview
> Go is designed with **concurrency** in mind. It uses a lightweight, powerful concurrency model based on **goroutines** and **channels**. Additionally, the standard library offers **context** to manage deadlines, cancellation, and request-scoped data across function boundaries.
> 
> These features are especially relevant when building **network** or **database-driven** applications because you need to efficiently handle multiple simultaneous operations and gracefully cancel or time out long-running tasks.

### Goroutines

A **goroutine** is a lightweight thread of execution in Go. You can create one using the `go` keyword before a function call.

#### Key Points
1. **Lightweight**: Thousands of goroutines can run on a few OS threads.
2. **Non-blocking**: Other goroutines keep running even if one is blocked (e.g., waiting on I/O).
3. **Share Memory by Communicating**: The recommended practice is to avoid sharing memory across threads if possible; use **channels** to coordinate instead.

##### Example

```go
func main() {
    // This runs in the main goroutine
    go sayHello("Alice") // This starts a new goroutine
    go sayHello("Bob")

    // Wait for input to prevent main() from exiting immediately
    fmt.Scanln()
}

func sayHello(name string) {
    fmt.Printf("Hello, %s!\n", name)
}
```

- Each call to sayHello with go runs concurrently.
- Scanln() is used so the program doesn’t exit immediately (in real code, you often use sync.WaitGroup or a more advanced approach).

### Channels
A `channel` in Go is a typed conduit to send and receive values between goroutines. Channels help synchronize goroutines and enable safe communication without extensive locking.

**Key Points**
* **Typed**: A channel has a specific type (e.g., chan int).
* **Unbuffered** vs. **Buffered**:
  * Unbuffered channels block until the receiver is ready.
  * Buffered channels allow a limited queue of messages.

##### Example
```go
func main() {
    ch := make(chan string)

    // Producer goroutine
    go func() {
        ch <- "ping" // Send a message to the channel
    }()

    // Consumer goroutine
    msg := <-ch // Receive the message (blocks until producer sends)
    fmt.Println(msg)
}

```

- The producer goroutine sends `"ping"` into the channel.
- The main goroutine blocks on `<-ch` until it receives a value.

#### Closing a Channel
When you’re done sending values, you can close(ch). Receivers get a zero value if they continue to read from a closed channel, and can check if the channel is closed by the two-value receive form: v, ok := <-ch.

### Context
The `context` package helps manage `cancellation`, `deadlines`, and other `request-scoped` values. This is very useful in long-running or I/O-bound processes where you might want to cancel or timeout the operation.

#### Key Points
1. **Inheritance**: You create a context from a parent (e.g., context.Background()) and pass it down call chains.
2. **Cancellation**: If a context is canceled, all functions using it should stop as soon as possible and clean up.
3. **Timeouts**: context.WithTimeout sets a deadline after which the context is automatically canceled.

##### Example
```go
func main() {
    // Create a context that times out after 3 seconds
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    // Pass ctx to a function that might block
    err := doLongRunningTask(ctx)
    if err != nil {
        fmt.Println("Task failed or timed out:", err)
    } else {
        fmt.Println("Task completed successfully")
    }
}

func doLongRunningTask(ctx context.Context) error {
    select {
    case <-time.After(5 * time.Second):
        // Simulate a long task
        return nil
    case <-ctx.Done():
        // The context was canceled or expired
        return ctx.Err()
    }
}

```

- `time.After(5s)` simulates a 5-second task, but the context times out at 3 seconds, so the task will cancel early.
- When `ctx.Done()` is signaled, we handle cleanup (if needed) and return.

### Error Handling Patterns
Although `error handling` isn’t strictly part of concurrency, it’s crucial when dealing with asynchronous operations that can fail. In Go:

1. `Check Errors Early`: After each function call that might fail, check the error. This avoids silent failures or confusion later.
2. `Wrap Errors`: Use fmt.Errorf("... %w", err) or similar to provide more context about where an error occurred.
3. `Use Sentinel Errors`: Sometimes define package-level var ErrSomething = errors.New("...") for repeated checks.

#### Example
```go
if err := doSomething(); err != nil {
    return fmt.Errorf("could not do something: %w", err)
}
```

### Summary
- `Goroutines`: Let you run functions concurrently. They are cheap and numerous compared to OS threads.
- `Channels`: Provide a way to safely communicate between goroutines without excessive locking.
- `Context`: Manages cancellation and timeouts, essential in network/database interactions.
- `Error Handling`: Go uses explicit error returns; checking and wrapping errors keeps code understandable.

- These features work together to build highly concurrent yet maintainable applications in Go. Start simple—launch goroutines, pass data via channels, and wrap calls in contexts for robust cancellation and timeouts. Over time, incorporate more advanced patterns like worker pools, sync.WaitGroup, or pipeline concurrency as your needs grow.
---

## MongoDB Docs
MongoDB is a NoSQL document-oriented database. Instead of tables and rows, data is stored in collections and documents. This allows for flexible schemas and the ability to store complex data structures within a single record.

If you are transitioning from SQL:
- **Database** in MongoDB is analogous to a database in SQL.
- **Collection** in MongoDB is analogous to a table in SQL.
- **Document** in MongoDB is analogous to a row in SQL, but stored in BSON (binary JSON) format.

### Setting Up
1. **Install the MongoDB Go Driver:**
   ```bash
   go get go.mongodb.org/mongo-driver/mongo
   go get go.mongodb.org/mongo-driver/bson
   ```
2. **Connect to MongoDB:**
   ```go
   package main
   
   
   func NewMongoService(dbName string) (*MongoDbContext, error) {
       config.LoadConfig()
   
       mongoURI := config.GetEnv("MONGO_URI")
       mongoUser := config.GetEnv("MONGO_USER")
       mongoPassword := config.GetEnv("MONGO_PASSWORD")
       mongoAuthSource := config.GetEnv("MONGO_AUTH_SOURCE")
   
       credential := options.Credential{
           AuthSource: mongoAuthSource,
           Username:   mongoUser,
           Password:   mongoPassword,
       }
   
       ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
       defer cancel()
   
       serverAPI := options.ServerAPI(options.ServerAPIVersion1)
       opts := options.Client().ApplyURI(mongoURI).SetAuth(credential).SetServerAPIOptions(serverAPI)
       client, err := mongo.Connect(ctx, opts)
   
       if err != nil {
           log.Fatalf("Error connecting to MongoDB: %v", err)
           return nil, err
       }
   
       if err := ensureHealthy(client, ctx); err != nil {
           // If ping fails, disconnect to avoid leaving the client in an inconsistent state
           _ = client.Disconnect(context.Background())
           return nil, err
       }
   
       logger.PrintlnColoredText("Successfully connected to MongoDB!", logger.GreenTxtColorCode)
       return &MongoDbContext{
           Client: client,
           DB:     client.Database(dbName),
       }, nil
   }
   ```

### Basic Concepts
#### Collections vs. Tables
- In MongoDB, you do not need to define a schema before creating a collection. You can simply start inserting documents into a collection that does not yet exist, and MongoDB will create it on the fly.

#### Documents vs. Rows
- A MongoDB document is a JSON-like structure (BSON under the hood). This means each document can have different fields.

#### Flexible Schema
- In an SQL database, all rows in a table must have the same columns (though columns can be NULL). In MongoDB, each document in a collection can have different or additional fields if needed.

### Data Types
MongoDB supports various data types, some of which directly map to Go data types:

1. **String** (string in Go)
2. **Boolean** (bool in Go)
3. **Integer** (int, int32, int64 in Go)
4. **Double** (float64 in Go)
5. **Date** (represented as time.Time in Go)
6. **Array** (represented as slices in Go)
7. **Object** / Embedded Document (represented as bson.M or structs in Go)
8. **ObjectId** (represented by primitive.ObjectID in Go from the MongoDB driver)

### Quick Reference for BSON Types
#### 1. `bson.M`
- **Definition**: `bson.M` is a Go map: `map[string]interface{}`.
- **Usage**: Quick, unordered BSON documents (filters, updates, inserts).
- **Example**:
  ```go
  filter := bson.M{"age": 30}
  update := bson.M{"$set": bson.M{"active": true}}
  ```
**When to Use:**
Most CRUD operations (Find, Update, Insert) when the order of fields does not matter.

#### 2. `bson.D`
- **Definition**: `bson.D` is a slice of bson.E structs, which preserve order of fields
```go
type D []E
type E struct {
    Key   string
    Value interface{}
}
```

- **Usage**: Primarily for aggregation pipelines or any scenario where you want fields in a specific order.
Example:
```go
filter := bson.D{
    {Key: "age", Value: bson.D{{Key: "$gt", Value: 25}}},
    {Key: "name", Value: "John"},
}
```

**When to Use:**
- For aggregation stages ($match, $project, $group, etc.)** or if you explicitly need ordered fields.

#### 3. `primitive.D`
- **Definition**: Identical concept to bson.D, but located in the primitive package.
- **Usage**: Same as bson.D. The driver uses bson.D and primitive.D somewhat interchangeably.
Example:
```go
pipeline := mongo.Pipeline{
    {{"$match", primitive.D{{"age", primitive.D{{"$gt", 25}}}}}},
}

```

#### 4. Which One Should You Use?
- `bson.M`: Most common, unordered. Use it for everyday queries and updates if you don’t care about field order.
- `bson.D / primitive.D`: For ordered documents, especially in aggregation pipelines or queries that rely on specific field order.

### CRUD Operations
#### Create
To insert data into a MongoDB collection, you can use:

* `InsertOne()`: to insert a single document
* `InsertMany()`: to insert multiple documents

**Example** (`InsertOne`):
```go
package main

import (
"context"
"fmt"
"log"
"time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
    Name  string `bson:"name"`
    Email string `bson:"email"`
    Age   int    `bson:"age"`
}

func main() {
   ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
   defer cancel()

    client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        log.Fatal(err)
    }
    defer client.Disconnect(ctx)

    // Choose the database and collection
    collection := client.Database("testdb").Collection("users")

    // Create a user
    newUser := User{
        Name:  "John Doe",
        Email: "john@example.com",
        Age:   30,
    }

    // Insert the user document
    res, err := collection.InsertOne(ctx, newUser)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Inserted document ID:", res.InsertedID)
}
```

#### Read
To retrieve documents from a collection, you use `Find()` or `FindOne()`.

```go
func main() {
   // ...connection code omitted for brevity...

    collection := client.Database("testdb").Collection("users")

    // Find all users
    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        log.Fatal(err)
    }
    defer cursor.Close(ctx)

    var users []User
    if err = cursor.All(ctx, &users); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Found users: %+v\n", users)

    // Find a single user
    var singleUser User
    err = collection.FindOne(ctx, bson.M{"name": "John Doe"}).Decode(&singleUser)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Found one user: %+v\n", singleUser)
}
```

#### Update
To update documents:

* `UpdateOne()`: update a single document
* `UpdateMany()`: update multiple documents
* `ReplaceOne()`: replace a single document completely
Example `(UpdateOne)`:
```go
update := bson.M{
   "$set": bson.M{
      "age": 31,
   },
}

filter := bson.M{"email": "john@example.com"}

res, err := collection.UpdateOne(ctx, filter, update)
if err != nil {
   log.Fatal(err)
}

fmt.Printf("Matched %v documents and updated %v documents.\n", res.MatchedCount, res.ModifiedCount)
```

#### Delete
To delete documents:

* `DeleteOne()`
* `DeleteMany()`
Example (`DeleteOne`):

```go
filter := bson.M{"name": "John Doe"}
res, err := collection.DeleteOne(ctx, filter)
if err != nil {
   log.Fatal(err)
}

fmt.Printf("Deleted %v documents.\n", res.DeletedCount)
```

### Queries and Filters
#### Basic Filters
Queries (filters) in MongoDB are specified as bson.M objects (maps in Go) or by using typed APIs. Common filters resemble SQL WHERE clauses but in JSON-like form.

```go
// Example Filter to find all documents where age is 30
filter := bson.M{"age": 30}
cursor, err := collection.Find(ctx, filter)
```

#### Operators
MongoDB provides a variety of query operators, such as:

* `$gt`, `$gte`: Greater than, greater or equal
* `$lt`, `$lte`: Less than, less or equal
* `$eq`, `$ne`: Equal, not equal
* `$in`, `$nin`: In, not in
* `$and`, `$or`: AND, OR

Example using `$gt`:
```go
filter := bson.M{
   "age": bson.M{
      "$gt": 25,
   },
}
cursor, err := collection.Find(ctx, filter)
```

### Projections (Select)
In SQL, you might specify columns to return (e.g. `SELECT name, age FROM users`). In MongoDB, use a projection to limit which fields are returned.

```go
options := options.Find().SetProjection(bson.M{"email": 1, "_id": 0})
// Only return the "email" field and exclude the "_id" field

cursor, err := collection.Find(ctx, bson.M{}, options)
if err != nil {
   log.Fatal(err)
}
```

### Limiting Results
Like `LIMIT` in SQL, use SetLimit:
```go
options := options.Find().SetLimit(5)
cursor, err := collection.Find(ctx, bson.M{}, options)
if err != nil {
   log.Fatal(err)
}
```

### Aggregation Pipelines
An aggregation pipeline is a sequence of stages (like $match, $project, $group, etc.) that transform your documents. In Go, you often build pipelines using mongo.Pipeline—a slice of documents that each contain one stage.

#### Basic Pipeline with $match and $project
```go
pipeline := mongo.Pipeline{
    // 1) $match stage: filter documents where age > 25
    {{"$match", bson.M{"age": bson.M{"$gt": 25}}}},

    // 2) $project stage: select only "name" and "email" fields
    {{"$project", bson.M{"name": 1, "email": 1, "_id": 0}}},
}

cursor, err := collection.Aggregate(ctx, pipeline)
```

### Pagination Example
A common way to paginate with MongoDB is to use skip and limit. For instance, if you want 10 items per page:

* page (the current page number)
* limit (items per page)
* skip = (page - 1) * limit

```go

page := 2
limit := int64(10)
skip := (page - 1) * limit

opts := options.Find().
    SetSkip(skip).
    SetLimit(limit)

cursor, err := collection.Find(ctx, bson.M{}, opts)
if err != nil {
    // handle error
}
defer cursor.Close(ctx)

// process results...

```
### Transactions
MongoDB supports multi-document transactions in replica set or sharded deployments. This is somewhat analogous to transactions in SQL databases. You start a session, begin a transaction, perform operations, and then commit or abort.

```go
session, err := client.StartSession()
if err != nil {
    log.Fatal(err)
}
defer session.EndSession(ctx)

// Transaction function
callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
    // All operations in this function will be part of a transaction
    collection := client.Database("testdb").Collection("users")

    _, err := collection.InsertOne(sessCtx, bson.M{"name": "Jane", "age": 28})
    if err != nil {
        return nil, err
    }

    // ... more operations ...

    return nil, nil
}

_, err = session.WithTransaction(ctx, callback)
if err != nil {
    log.Fatal(err)
}

fmt.Println("Transaction complete.")

```
---