module github.com/redis/rueidis/om

go 1.22

toolchain go1.23.3

replace github.com/redis/rueidis => ../

require (
	github.com/oklog/ulid/v2 v2.1.0
	github.com/redis/rueidis v1.0.49
)

require golang.org/x/sys v0.24.0 // indirect
