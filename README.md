# install goose
go install github.com/pressly/goose/v3/cmd/goose@latest

# cd /path/to/your/migrations

# get status of migrations
goose mysql "${RC_USERNAME}:${RC_PASSWORD}@/${RC_DATABASE}?parseTime=true" status

# run all migrations
goose mysql "${RC_USERNAME}:${RC_PASSWORD}@/${RC_DATABASE}?parseTime=true" up

