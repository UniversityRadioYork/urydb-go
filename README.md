# URYdb-go

A tiny package that looks for URY database connection URLs in a few places and returns an open sql.DB instance.

## Usage
```go
db, err := urydb.GetDB()
```
That's it.
