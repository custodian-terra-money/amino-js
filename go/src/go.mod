module github.com/terra-project/amino-js/go/src

go 1.12

require (
	github.com/terra-project/amino-js/go/extensions v0.0.0
	github.com/cosmos/cosmos-sdk v0.37.0
	github.com/tendermint/go-amino v0.15.0
	github.com/terra-project/amino-js/go/lib v0.0.0
)

replace github.com/terra-project/amino-js/go/extensions => ../extensions

replace github.com/terra-project/amino-js/go/lib => ../lib
