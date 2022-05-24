running Makefile:
make create-keypair ENV=test

running create migrate:
make migrate-create NAME=add_users_table

execute migrate up:
make migrate-up
execute  down:
make migrate-down
