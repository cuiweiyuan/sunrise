package testfactory

import (
	"bytes"
	"sort"

	tmrand "github.com/cometbft/cometbft/libs/rand"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sunriselayer/sunrise/pkg/appconsts"
	apprand "github.com/sunriselayer/sunrise/pkg/random"
)

const (
	// nolint:lll
	TestAccName  = "test-account"
	TestAccAddr  = "sunrise1g39egf59z8tud3lcyjg5a83m20df4kcc9kspal"
	TestAccMnemo = `ramp soldier connect gadget domain mutual staff unusual first midnight iron good deputy wage vehicle mutual spike unlock rocket delay hundred script tumble choose`
	ChainID      = "test-app"
)

func Repeat[T any](s T, count int) []T {
	ss := make([]T, count)
	for i := 0; i < count; i++ {
		ss[i] = s
	}
	return ss
}

// GenerateRandNamespacedRawData returns random data of length count. Each chunk
// of random data is of size shareSize and is prefixed with a random blob
// namespace.
func GenerateRandNamespacedRawData(count int) (result [][]byte) {
	for i := 0; i < count; i++ {
		rawData := tmrand.Bytes(appconsts.ShareSize)
		namespace := apprand.RandomBlobNamespace().Bytes()
		copy(rawData, namespace)
		result = append(result, rawData)
	}

	sortByteArrays(result)
	return result
}

func sortByteArrays(src [][]byte) {
	sort.Slice(src, func(i, j int) bool { return bytes.Compare(src[i], src[j]) < 0 })
}

func RandomAccountNames(count int) []string {
	accounts := make([]string, 0, count)
	for i := 0; i < count; i++ {
		accounts = append(accounts, tmrand.Str(10))
	}
	return accounts
}

func GenerateAccounts(count int) []string {
	accs := make([]string, count)
	for i := 0; i < count; i++ {
		accs[i] = tmrand.Str(20)
	}
	return accs
}

func GetAddresses(keys keyring.Keyring) []sdk.AccAddress {
	recs, err := keys.List()
	if err != nil {
		panic(err)
	}
	addresses := make([]sdk.AccAddress, 0, len(recs))
	for idx, rec := range recs {
		addresses[idx], err = rec.GetAddress()
		if err != nil {
			panic(err)
		}
	}
	return addresses
}

func GetAddress(keys keyring.Keyring, account string) sdk.AccAddress {
	rec, err := keys.Key(account)
	if err != nil {
		panic(err)
	}
	addr, err := rec.GetAddress()
	if err != nil {
		panic(err)
	}
	return addr
}

// func RandomEVMAddress() gethcommon.Address {
// 	return gethcommon.BytesToAddress(tmrand.Bytes(gethcommon.AddressLength))
// }

func TestKeyring(cdc codec.Codec, accounts ...string) keyring.Keyring {
	kb := keyring.NewInMemory(cdc)

	for _, acc := range accounts {
		_, _, err := kb.NewMnemonic(acc, keyring.English, "", "", hd.Secp256k1)
		if err != nil {
			panic(err)
		}
	}

	_, err := kb.NewAccount(TestAccName, TestAccMnemo, "", "", hd.Secp256k1)
	if err != nil {
		panic(err)
	}

	return kb
}
