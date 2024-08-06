.PHONY: initialize_db
initialize_db:
	PGPASSWORD=password psql -h  127.0.0.1 -U test -d todo_db -f ./db/cleanup.sql
	PGPASSWORD=password psql -h  127.0.0.1 -U test -d todo_db -f ./db/createTable.sql
	PGPASSWORD=password psql -h  127.0.0.1 -U test -d todo_db -f ./db/insertTestData.sql