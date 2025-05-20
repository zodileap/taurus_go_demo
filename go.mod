module taurus_go_demo

go 1.24.1

require (
	github.com/google/uuid v1.6.0
	github.com/lib/pq v1.10.9
	github.com/zodileap/taurus_go v0.9.2
)

replace github.com/zodileap/taurus_go v0.9.2 => ../taurus_go

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/redis/go-redis/v9 v9.7.3 // indirect
	golang.org/x/mod v0.24.0 // indirect
	golang.org/x/sync v0.12.0 // indirect
	golang.org/x/tools v0.31.0 // indirect
)
