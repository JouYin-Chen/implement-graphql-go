
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
4. add enum
5. create Mutation function
6. 修改 資料格式 新增interface

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

## mutation
```
mutation {
  reateAuthor {
    name
    age
    books {
      title
      series
    }
  }
}
```

### Enum
新增Enum type, 在Schema type 中 宣告變數為 enum type

查詢

```
query tryEnum($pID: TaskStateEnum, $name: String){
  projectByNameAndState(state: $pID, name: $name){
    id
    name
    state
  }
}

{
  "pID": "inProgress",
  "name": "Learn React Native"
}

```
包含Enum的資料
get author
```
query get($AuthorName: String!) {
  getAuthorByName(name: $AuthorName) {
		name
    age
    gendor #return "male"
    books{
      title
      author
      series
    }
  }
}


{
  "AuthorName": "Brandon Sanderson"
}
```

##interface

修改data 格式
新增interface type
將Schema book type 加入 BookInterface type
query schema 中有需要使用到 interface 變數的 return value 就要加上interface的宣告

```
query implementInterface($BookTitle: String!) {
  getBooksByTitle(title: $BookTitle) {
    ... on getBookByTitleInterface {
      # series
      title
    }
	series
# title
  }
}

{
  "BookTitle": "A Game of Thrones"
}
```

## union
```
  {
    __typename
    getAuthor(name:"Brandon Sanderson"){
      name
      age
      gendor
      books{
        ... on TextBook {
          title
          series
        }
        ... on ShortStory {
          title
          chapters
        }
      }
    }
  }
```
