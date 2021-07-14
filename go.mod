module github.com/jgarland/bitly_backend_test

go 1.15

require (
	github.com/bitly/openapi v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.8.1
	golang.org/x/oauth2 v0.0.0-20210628180205-a41e5a781914 // indirect
)

replace github.com/bitly/openapi => ./openapi
