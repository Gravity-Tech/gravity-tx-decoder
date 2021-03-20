package main

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
)

type ID [32]byte
type TxFunc string
type Type string
type Arg struct {
	Type  Type
	Value []byte
}
type Value interface{}
type StringValue struct {
	Value string
}
type IntValue struct {
	Value int64
}
type BytesValue struct {
	Value []byte
}

type Transaction struct {
	Id           ID
	SenderPubKey [32]byte
	Signature    [72]byte
	Func         TxFunc
	Timestamp    uint64
	Args         []Arg
}

type ChainType byte

const (
	Ethereum ChainType = iota
	Waves
	Binance
)

func main() {

	s0 := "eyJJZCI6WzQwLDY4LDcyLDIwOSwxMTIsMjM3LDMwLDE2MiwyMywxMzMsMTY3LDQsMTI5LDE3OSwyMjYsOTgsMTQwLDk1LDExLDE2MywzOSw0NywzNCwxNCwxODAsODgsMjMsNTcsMTYyLDYyLDM5LDE2Nl0sIlNlbmRlclB1YktleSI6WzIwMiwyMDMsMTEzLDY5LDE4MywxODMsMiwxNywyMzcsNjcsMTcyLDIxNCw3MiwxMzUsMTMxLDU0LDE0NSw5MywxNTIsMjYsMTkwLDE2NiwxODksMTEsNTksMjIxLDEwNiw3OSwyNDUsMjE4LDIxMywyMDddLCJTaWduYXR1cmUiOlsxMzYsMTksMjIsMTksMTQ0LDczLDU4LDE4NywyMjEsMTk3LDYsNjYsMTAxLDE1MCw1LDExLDIxOSwxNjAsNDUsNiwyMTAsMjQxLDExMSwyMjcsMTQ4LDg1LDE0NCwxMzMsMjMsMTE3LDkyLDIwLDE2MSwxLDgsMTIxLDI0MiwxNDQsMjQ3LDU1LDcxLDE2OSwxMjIsMTQ2LDIzNCwxNDcsMTY4LDIzNSwyMjMsODMsMjEyLDMzLDE2OCw4OSw0Miw3MCw3MywyMDYsMTQzLDI0NywzMSw3MywyMTcsMiwwLDAsMCwwLDAsMCwwLDBdLCJGdW5jIjoic2lnbk5ld0NvbnN1bHMiLCJUaW1lc3RhbXAiOjE2MTQ2MjIwOTgsIkFyZ3MiOlt7IlR5cGUiOiJieXRlcyIsIlZhbHVlIjoiQWc9PSJ9LHsiVHlwZSI6ImludCIsIlZhbHVlIjoiQUFBQUFBQUFBQVU9In0seyJUeXBlIjoiYnl0ZXMiLCJWYWx1ZSI6IjNxallmNXVkbnBJMHpFODVIS1c5aHcvN3IvbUVTcnFFOU1xblI0aU0xeXh6Wm9YMFVzZVhIUzRiNmViN2dsSU9vSE13RlFUbllBdDh4ODU2K1UrZWJ3RT0ifV19"
	s1 := "eyJJZCI6WzIwMCw5MCwxODksMjA1LDkyLDg5LDEzMiwxMjUsMjUxLDExNSw1NiwxMSwyMzcsMjI4LDE5NSwxMDMsMTk5LDk3LDE4MCwxMiw0NiwyMTcsMTM1LDEwOSw5NCwzMywyMDMsMTA4LDE3NSwxMTgsMTAyLDExM10sIlNlbmRlclB1YktleSI6WzgwLDI1MywyNCwxNjUsMTkzLDE1MCwxNTQsMzUsMTA1LDI0NywxMjAsMjE5LDE4OCwxNjgsMjQ3LDE2NywyMzMsMTU4LDE4LDU0LDIxNiwyNywyMjcsMTIwLDIsODUsMjE2LDE2MywyMTgsMTM3LDI0OSwxOTZdLCJTaWduYXR1cmUiOls5MiwxNjAsNTksMTYwLDE3MiwxNTQsNDEsNzQsMTExLDExLDU4LDI1MSwxMDgsNjgsMjAsMTQ0LDU0LDIxNywxNDMsMTQsMTUwLDE4Myw3NSwxMTUsNzgsMjQ1LDg5LDI0MiwxNTYsMTkzLDEyNSwyNDAsMTI3LDE5MiwzNCwyMDksMjAxLDIyMSwxNTEsMjIyLDYsMjI1LDUsNyw5MCwyMzYsOTMsODMsMjA2LDY1LDIwNCwyMTEsMjM5LDMxLDIwMCwxMDEsMjQ3LDExOSwxMzgsMjE5LDI1MSw5MywyMTYsOCwwLDAsMCwwLDAsMCwwLDBdLCJGdW5jIjoic2lnbk5ld0NvbnN1bHMiLCJUaW1lc3RhbXAiOjE2MTQ2MjIwOTgsIkFyZ3MiOlt7IlR5cGUiOiJieXRlcyIsIlZhbHVlIjoiQWc9PSJ9LHsiVHlwZSI6ImludCIsIlZhbHVlIjoiQUFBQUFBQUFBQVU9In0seyJUeXBlIjoiYnl0ZXMiLCJWYWx1ZSI6ImpIL2NDdW5ocVlNQldzVDdIYjBtR2ZnT005QkVDS0tzdzBPd2hhQzllaWtKQTAzUGVjZ1NMQVNCa0VKVmQrL0l1WXAxTi9iblJlOTU1YjlDZm13eWV3QT0ifV19"

	tx0 := parseTx(s0)
	tx1 := parseTx(s1)

	printTx(tx0)
	printTx(tx1)

}

func printTx(tx Transaction) {

	ch := ChainType(tx.Args[0].Value[0])

	id := binary.BigEndian.Uint64(tx.Args[1].Value)

	fmt.Println("type:", tx.Func, "chain:", ch.String(), "roundId:", id, "sign:", tx.Args[2].Value)
}

func parseTx(s string) Transaction {

	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		fmt.Println("error:", err)
	}

	var tx Transaction
	err = json.Unmarshal(b, &tx)
	if err != nil {
		fmt.Println("error:", err)
	}

	return tx
}

func (ch ChainType) String() string {
	switch ch {
	case Ethereum:
		return "ethereum"
	case Waves:
		return "waves"
	case Binance:
		return "bsc"
	default:
		return "ethereum"
	}
}
