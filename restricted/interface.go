package restricted

import (
  "github.com/openrelayxyz/plugeth-utils/core"
  "github.com/openrelayxyz/plugeth-utils/restricted/params"
)

type Backend interface {
  core.Backend
	// General Ethereum API
	ChainDb() Database
  ChainConfig() *params.ChainConfig
}
