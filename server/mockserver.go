package server

import (
	"github.com/0xPolygonHermez/zkevm-bridge-service/bridgectrl"
	"github.com/0xPolygonHermez/zkevm-bridge-service/db/pgstorage"
)

// RunMockServer runs mock server
func RunMockServer() (*bridgectrl.BridgeController, error) {
	dbCfg := pgstorage.NewConfigFromEnv()
	err := pgstorage.InitOrReset(dbCfg)
	if err != nil {
		return nil, err
	}

	store, err := pgstorage.NewPostgresStorage(dbCfg)
	if err != nil {
		return nil, err
	}

	bt, err := bridgectrl.MockBridgeCtrl(store)
	if err != nil {
		return nil, err
	}

	cfg := Config{
		GRPCPort: "9090",
		HTTPPort: "8080",
	}

	return bt, RunServer(store, bt, cfg)
}
