
執行
`$GOPATH/bin/air -c .air.conf`

web界面
URL: http://localhost:8080/graphql

GET
```
URL
http://localhost:8080/graphql?query={fetchTasks{id, title}}
```


POST
```
URL
http://localhost:8080/graphql
body
{
	"query": "query {fetchTasks{id, title}}"
}
```


0. http server
1. create data
2. create Query function
    1. create query type return type
    2. set return type & function Args
    3. create function Resolve return value
3. add function for variables

## variables
```
query get($TaskID: ID!) {
  getTask(id: $TaskID) {
    id
    title
  }
}

QUERY VARIABLES
{
  "TaskID": 1
}
```
