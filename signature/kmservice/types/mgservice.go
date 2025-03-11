package types

type CoinType uint32

// https://github.com/satoshilabs/slips/blob/master/slip-0044.md
// erc20 token代币参考未占用地址新增
const (
	CoinTypeBTC      CoinType = 0x80000000
	CoinTypeLTC      CoinType = 0x80000002
	CoinTypeETH      CoinType = 0x8000003c
	CoinTypeEOS      CoinType = 0x800000c2
	CoinTypeMATIC    CoinType = 0x800003c6
	CoinTypeTRON     CoinType = 0x800000c3
	CoinTypeBNB      CoinType = 0x800002ca
	CoinTypeDASH     CoinType = 0x80000005
	CoinTypeBCH      CoinType = 0x80000091
	CoinTypeFTM      CoinType = 0x800003ef
	CoinTypeAVAXC    CoinType = 0x8000232d
	CoinTypeOp       CoinType = 0x80000266
	CoinTypeArbitrum CoinType = 0x80002329
	CoinTypeFilCoin  CoinType = 0x800001cd
)

type ArgsGetAddress struct {
	Wid          uint32
	CoinType     CoinType
	AddressIndex uint32
}

type ReplyGetAddress struct {
	Address string
	Pubkey  string
	Coin    string
}

type ArgsSign struct {
	Wid          uint32
	CoinType     CoinType
	AddressIndex uint32
	Hash         []byte
	LenVin       int
}

type ReplySign struct {
	Result []byte
	Err    string
}
