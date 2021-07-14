
## Requirements
https://gist.github.com/bitlytanner/0ea563f6a68decaefb2debe91e303d75#requirements

## Design
I chose to write this in Golang so that I could use the compiled binary to build a very compact docker image, and I chose to use the OpenAPI becuase it appears to be the most official golang library option for interfacing with the bitly API.

https://dev.bitly.com/docs/sdks/openapi-30

## Dependencies
 Docker
 curl

## Build
`docker build -t bitly_backend_test .`

## Run
`docker run --rm -p 127.0.0.1:8080:8080/tcp -it bitly_backend_test:latest`

`curl -v -H 'Authorization: Bearer {TOKEN}' localhost:8080`
