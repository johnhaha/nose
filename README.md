# 👃 nose 

golang api for notion
currently support some basic request

```text
go get github.com/johnhaha/nose@v0.0.3
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

only support title and text

```go
client.NewEmptyDatabase("DATABASE-NAME","TITLE-COLUMN-NAME","COLUMN-NAME1","COLUMN-NAME2")
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

```go
client.NewPage(map[string]{"TITLE":"TITLE"},map[string]{"COLUMN1":"COLUMN1","COLUMN2":"COLUMN2"})
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
