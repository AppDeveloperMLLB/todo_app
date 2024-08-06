## GET Todo
```zsh
export TOKEN="Authorization: Bearer token"
curl "localhost:8080/v1/todo?page=1&per_page=10&todo_status=completed" -X GET -H $TOKEN | jq
```