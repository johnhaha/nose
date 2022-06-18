# 👃 nose

notion api
support basic functions such as create page, db, text block

## 🔽 How to install

```text
go get github.com/johnhaha/nose@v0.0.10
```

## 📃  Page API

### Create Page Client

```go
client := nose.NewPageClient("YOUR-TOKEN","PARENT-PAGE-ID")
```

### Create New Page

```go
client.NewEmptyPage("PAGE-NAME")
```

### Create Database

save Struct data to json, with nose tag

```go
type Sample struct {
	Title     string     `json:"title" nose:"title"`
    Desc      string     `json:"desc" nose:"orange,bold"`
	Count     int        `json:"count"`
	Updated   time.Time  `json:"updated"`
	CreatedAt time.Time  `json:"createdAt"`
} 
client.NewDB("DATABASE-NAME",Sample{})
```

### Append Text Block

only support text

```go
client.AppendTextBlock("TEXT-CONTENT")
```

### Append TODO Block

```go
client.AppendTodoBlock("BLOCK-CONTENT")
```

## 🐬  Database API

### Create Database Client

```go
client := nose.NewDBClient("YOUR-TOKEN","DATABASE-ID")
```

### Insert page to database

save struct data to database

```go
client.SaveData(sample)
```

## 🎎  Client Exchange

### Page Client to DB Client

```go
client := nose.NewPageClient("YOUR-TOKEN","PARENT-PAGE-ID")
dbClient := client.ToDB("DBID")

```

### DB Client to Page Client

```go
client := nose.NewDBClient("YOUR-TOKEN","DB-ID")
dbClient := client.ToPage("PAGE-ID")

```