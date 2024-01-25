# nfon-crud

Create a CRUD service with Golang:

Your task is to develop a simple CRUD service in Golang for managing items. Each item should have a unique ID and a name. The service should support the following operations:

Create Item: Add a new item to the system.

Get Item by ID: Retrieve an item by its ID.

Get All Items: Retrieve all items. Additionally, implement a query parameter to filter items by name.

Update Item: Modify an existing item by its ID.

Delete Item: Remove an item from the system.

Feel free to implement it as you see fit, but simplicity and readability will be highly valued.

## Kick Off
[Kick-Off](https://github.com/DhirenB94/nfon-crud/issues/1) with a task breakdown can be found here

## Important Decisions Made
1. As the task mentions simplicity, I am only using an in memory database (maps) rather than persisting the data in an actual database
2. The updateItem operation will only be a PATCH and not a PUT because we only have 2 fields (ID and Name) and only the name should be modifiable
3. All operations involving writes (Create Item & Update Item) are done so via JSON in the request body

## Next Steps
1. Persist the data by implementing a file system or a database
3. Implement a basic frontend and so users can input items, and parse these forms

## Run Locally

Clone the project
```bash
  git clone https://github.com/DhirenB94/nfon-crud.git
```
Go to the project directory
```bash
  cd nfon-crud
```
Start the server
```bash
  go run cmd/web/main.go
```

