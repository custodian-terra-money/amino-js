package src

import (
	"github.com/terra-project/amino-js/go/lib/tendermint/tendermint/libs/bech32"
)

var EncodeBech32 = bech32.ConvertAndEncode
var DecodeBech32 = bech32.DecodeAndConvert
