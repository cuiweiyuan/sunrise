package testnode

import (
	"context"
	"encoding/hex"

	"github.com/sunriselayer/sunrise/app/consts"
	"github.com/sunriselayer/sunrise/app/encoding"
	testencoding "github.com/sunriselayer/sunrise/test/util/encoding"
	"github.com/sunriselayer/sunrise/test/util/testfactory"

	sdkmath "cosmossdk.io/math"
	tmrand "github.com/cometbft/cometbft/libs/rand"
	rpctypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func TestAddress() sdk.AccAddress {
	bz, err := sdk.GetFromBech32(testfactory.TestAccAddr, "sunrise")
	if err != nil {
		panic(err)
	}
	return sdk.AccAddress(bz)
}

func QueryWithoutProof(clientCtx client.Context, hashHexStr string) (*rpctypes.ResultTx, error) {
	hash, err := hex.DecodeString(hashHexStr)
	if err != nil {
		return nil, err
	}

	node, err := clientCtx.GetNode()
	if err != nil {
		return nil, err
	}

	return node.Tx(context.Background(), hash, false)
}

func NewKeyring(accounts ...string) (keyring.Keyring, []sdk.AccAddress) {
	cdc := encoding.MakeConfig(testencoding.ModuleEncodingRegisters...).Codec
	kb := keyring.NewInMemory(cdc)

	addresses := make([]sdk.AccAddress, len(accounts))
	for idx, acc := range accounts {
		rec, _, err := kb.NewMnemonic(acc, keyring.English, "", "", hd.Secp256k1)
		if err != nil {
			panic(err)
		}
		addr, err := rec.GetAddress()
		if err != nil {
			panic(err)
		}
		addresses[idx] = addr
	}
	return kb, addresses
}

func RandomAddress() sdk.Address {
	name := tmrand.Str(6)
	_, addresses := NewKeyring(name)
	return addresses[0]
}

func FundKeyringAccounts(accounts ...string) (keyring.Keyring, []banktypes.Balance, []authtypes.GenesisAccount) {
	kr, addresses := NewKeyring(accounts...)
	genAccounts := make([]authtypes.GenesisAccount, len(accounts))
	genBalances := make([]banktypes.Balance, len(accounts))

	for i, addr := range addresses {
		balances := sdk.NewCoins(
			sdk.NewCoin(consts.BondDenom, sdkmath.NewInt(99999999999999999)),
		)

		genBalances[i] = banktypes.Balance{Address: addr.String(), Coins: balances.Sort()}
		genAccounts[i] = authtypes.NewBaseAccount(addr, nil, uint64(i), 0)
	}
	return kr, genBalances, genAccounts
}

func GenerateAccounts(count int) []string {
	accs := make([]string, count)
	for i := 0; i < count; i++ {
		accs[i] = tmrand.Str(20)
	}
	return accs
}
