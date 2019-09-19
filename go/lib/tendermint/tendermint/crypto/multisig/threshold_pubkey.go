package multisig

import (
	"github.com/terra-project/amino-js/go/lib/tendermint/tendermint/crypto"
)

type PubKeyMultisigThreshold struct {
	K       uint            `json:"threshold"`
	PubKeys []crypto.PubKey `json:"pubkeys"`
}

var _ crypto.PubKey = PubKeyMultisigThreshold{}
