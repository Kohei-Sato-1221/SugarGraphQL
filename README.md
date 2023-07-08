# SugarGraphQL
Sample fullstack GraphQL application


## SQLBoilder
```
go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest
```

## memo
https://zenn.dev/hsaki/books/golang-graphql/viewer/resolverbasic

curl \
  -X POST \
  -H "Content-Type: application/json" \
  -d '{ "query": "query { user(name: \"kohekohe\") { id name projectV2(number: 1) { title }} }" }' \
  http://localhost:8080/query