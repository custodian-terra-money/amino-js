package rootmulti

import (
	"github.com/terra-project/amino-js/go/lib/cosmos/cosmos-sdk/store/types"
)

type storeInfo struct {
	Name string
	Core storeCore
}

type storeCore struct {
	CommitID types.CommitID
}
