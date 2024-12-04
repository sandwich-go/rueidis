module github.com/redis/rueidis/rueidisprob

go 1.22

toolchain go1.23.3

replace github.com/redis/rueidis => ../

require (
	github.com/redis/rueidis v1.0.49
	github.com/twmb/murmur3 v1.1.8
)

require golang.org/x/sys v0.24.0 // indirect
