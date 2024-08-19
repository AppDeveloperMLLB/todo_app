## GET TodoList
```zsh
export TOKEN="Authorization: Bearer token"
curl "localhost:8080/v1/todo?page=1&per_page=10&todo_status=completed" -X GET -H $TOKEN | jq
```

## GET Todo
```zsh
export TOKEN="Authorization: Bearer token"
curl "localhost:8080/v1/todo/6" -X GET -H $TOKEN | jq
```

## Create Todo
```zsh
curl "localhost:8080/v1/todo" -X POST -d '{"title": "test", "description": "This is test", "status": "todo"}' -H $TOKEN | jq
```

## Update Todo
```zsh
curl "localhost:8080/v1/todo" -X PUT -d '{"id": 6,"title": "update", "description": "This is update test", "status": "todo"}' -H $TOKEN | jq
```

## Delete Todo
```zsh
curl "localhost:8080/v1/todo/6" -X DELETE -H $TOKEN | jq
```