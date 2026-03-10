module github.com/ava-labs/ac-staking-contracts

go 1.24.11

toolchain go1.24.12

replace github.com/ava-labs/avalanchego/graft/coreth => github.com/ava-labs/avalanchego/graft/coreth v0.0.0-20251203215505-70148edc6eca

replace github.com/ava-labs/avalanchego/graft/subnet-evm => github.com/ava-labs/avalanchego/graft/subnet-evm v0.0.0-20260105172535-1a59a6f646ef

replace github.com/ava-labs/avalanchego/graft/evm => github.com/ava-labs/avalanchego/graft/evm v0.0.0-20260105172535-1a59a6f646ef

require (
	github.com/ava-labs/avalanchego/graft/subnet-evm v0.0.0-20260105172535-1a59a6f646ef
	github.com/ava-labs/libevm v1.13.15-0.20251210210615-b8e76562a300
)

require (
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/bits-and-blooms/bitset v1.20.0 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.3.5 // indirect
	github.com/consensys/gnark-crypto v0.18.1 // indirect
	github.com/crate-crypto/go-kzg-4844 v1.1.0 // indirect
	github.com/deckarep/golang-set/v2 v2.1.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.4.0 // indirect
	github.com/ethereum/c-kzg-4844 v1.0.0 // indirect
	github.com/fsnotify/fsnotify v1.9.0 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/holiman/uint256 v1.2.4 // indirect
	github.com/shirou/gopsutil v3.21.11+incompatible // indirect
	github.com/supranational/blst v0.3.14 // indirect
	github.com/tklauser/go-sysconf v0.3.15 // indirect
	github.com/tklauser/numcpus v0.10.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	golang.org/x/crypto v0.45.0 // indirect
	golang.org/x/exp v0.0.0-20241215155358-4a5509556b9e // indirect
	golang.org/x/mod v0.29.0 // indirect
	golang.org/x/sync v0.18.0 // indirect
	golang.org/x/sys v0.38.0 // indirect
	golang.org/x/time v0.12.0 // indirect
	golang.org/x/tools v0.38.0 // indirect
)
