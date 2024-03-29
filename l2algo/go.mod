module l2algo

go 1.21

toolchain go1.22.0

require (
	bc v0.0.0
	github.com/algorand/go-algorand-sdk/v2 v2.4.0
	merkle v0.0.0
)

require (
	github.com/algorand/go-codec/codec v1.1.10 // indirect
	github.com/cometbft/cometbft v0.38.5 // indirect
	golang.org/x/crypto v0.20.0 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
)

replace (
	bc => ../../quartz/bc
	merkle => ../../quartz/merkle
	qz => ../../quartz/qz
)
