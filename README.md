# API V0.1
## Prepare of MySQL
The next script is for connection database, we are use MySQL
```go
    storage.NewConnection()
	mysqlser := storage.NewMySQL(storage.Pool())
	newmysql := model.NewServices(mysqlser)
```
### Migration
Later, we must copy next script for the migration and creation table of database, also we validated the error
```go
    if err := newmysql.Migrate(); err != nil {
		log.Fatalf("model.Migrate() %v", err)
	}
```
### Insert values
For insert values we must create a model with values, for example title:"value01", body:"body01", also also we validated the error, copy next script
```go
   product := &model.Product{
		Title: "valu01",
		Body:  "body01",
	}
	if err := newmysql.Create(product); err != nil {
		log.Fatalf("model.Create() %v", err)
	}
```
### Update values
For update values we must create a model with values, for example title:"value01", body:"body01", also also we validated the error, copy next script
```go
   product := &model.Product{
		ID:        uint(va),
		Title:     c.FormValue("title"),
		Body:      c.FormValue("body"),
		Update_At: time.Now(),
	}
	if err := newmysql.Update(product); err != nil {
		log.Fatalf("model.Update() %v", err)
	}
```
### Delete values
For Delete values we must add the id
```go
	if err := newmysql.Delete(uint(va)); err != nil {
		log.Fatalf("model.Delete() %v", err)
	}
```

## Routes
We use the framework echo to our api, we must copy the url at Postman or App Frontend.
### POST
Method to create data on our databases
```curl
	http://localhost:999/api/v0/product
```
Result 
```json
	{
	"ID": 8,
	"Title": "example",
	"Body": "sentence of example",
	"Create_At": "0001-01-01T00:00:00Z",
	"Update_At": "0001-01-01T00:00:00Z"
	}
```
### PUT
Method to update a data on our databases,also we must added the id 
```curl
	http://localhost:999/api/v0/product/8
```
Result 
```json
	{
	"ID": 8,
	"Title": "title changed",
	"Body": "sentence changed ",
	"Create_At": "0001-01-01T00:00:00Z",
	"Update_At": "2023-12-29T23:22:09.739206-03:00"
	}
```
### DELETE
Method to delete a data on our databases,also we must added the id 
```curl
	http://localhost:999/api/v0/product/8
```
Result 
```json
	{
		"Data delete ok!"
	}
```
