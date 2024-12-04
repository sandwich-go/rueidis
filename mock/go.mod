module github.com/redis/rueidis/mock

go 1.22

toolchain go1.23.3

replace github.com/redis/rueidis => ../

require (
	github.com/redis/rueidis v1.0.49
	go.uber.org/mock v0.4.0
)

require golang.org/x/sys v0.24.0 // indirect
