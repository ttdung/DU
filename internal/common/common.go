package common

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ethcommon "github.com/ethereum/go-ethereum/common"
	//internalCommon "github.com/ttdung/du/internal/clients/common"
	"strings"
	"time"
)

const (
	AsaDecimals      = 18
	DefaultSleepTime = 2 * time.Second
	ZeroAddress      = "0000000000000000000000000000000000000000"
	DecimalDigits    = 5
)

// FormatMDLink creates a ref-text message following the MarkDown standard.
func FormatMDLink(msg string, link string) string {
	return fmt.Sprintf("[%v](%v)", msg, link)
}

func AccountAddressToHex(addr string) (string, error) {
	ethAddr, err := AccountAddressToEthAddr(addr)
	if err != nil {
		return "", err
	}

	return ethAddr.Hex(), nil
}

// MustAccountAddressToHex is the same as AccountAddressToHex except that it will panic upon errors.
func MustAccountAddressToHex(addr string) string {
	addr, err := AccountAddressToHex(addr)
	if err != nil {
		panic(err)
	}

	return addr
}

// AccountAddressToEthAddr parses the given address to an ETH address.
func AccountAddressToEthAddr(addr string) (ethcommon.Address, error) {
	zeroAddr := ethcommon.HexToAddress(addr)
	if addr == ZeroAddress {
		return zeroAddr, nil
	}
	if strings.HasPrefix(addr, sdk.GetConfig().GetBech32AccountAddrPrefix()) {
		// Check to see if address is Cosmos bech32 formatted
		toAddr, err := sdk.AccAddressFromBech32(addr)
		if err != nil {
			return zeroAddr, fmt.Errorf("%v is not a valid Bech32 address", addr)
		}
		ethAddr := ethcommon.BytesToAddress(toAddr.Bytes())
		return ethAddr, nil
	}

	if !strings.HasPrefix(addr, "0x") {
		addr = "0x" + addr
	}

	valid := ethcommon.IsHexAddress(addr)
	if !valid {
		return zeroAddr, fmt.Errorf("%s is not a valid Ethereum or Cosmos address", addr)
	}

	ethAddr := ethcommon.HexToAddress(addr)

	return ethAddr, nil
}
