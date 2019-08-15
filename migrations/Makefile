ENV=dev
COMMAND=sql-migrate -env="${ENV}"
new:
	sql-migrate -env="${ENV}" create_user_table
up:
	sql-migrate up -env="${ENV}"
down:
	sql-migrate down -env="${ENV}" down 