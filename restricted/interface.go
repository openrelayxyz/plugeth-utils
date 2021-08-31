package restricted

import (
  "github.com/openrelayxyz/plugeth-utils/core"
)

type Backend interface {
  core.Backend
	// General Ethereum API
	ChainDb() Database
}
