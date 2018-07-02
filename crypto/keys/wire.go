package keys

import (
	amino "github.com/tendermint/go-amino"
	tcrypto "github.com/tendermint/tendermint/crypto"
)

var cdc = amino.NewCodec()

func init() {
	tcrypto.RegisterAmino(cdc)
	cdc.RegisterInterface((*Info)(nil), nil)
	//	cdc.RegisterConcrete(ccrypto.PrivKeyLedgerSecp256k1{},
	//		"tendermint/PrivKeyLedgerSecp256k1", nil)
	cdc.RegisterConcrete(localInfo{}, "crypto/keys/localInfo", nil)
	cdc.RegisterConcrete(ledgerInfo{}, "crypto/keys/ledgerInfo", nil)
	cdc.RegisterConcrete(offlineInfo{}, "crypto/keys/offlineInfo", nil)
}
