package utilities

import (
	"bytes"
	"encoding/hex"
	"log"
)

func init() {
	// Enable line numbers in logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// GetBytes - Returns the corresponding byte stream, which is passed as a string stream into the attribute hexValues with hex values
func GetBytes(hexValues string) []byte {
	decStr, _ := hex.DecodeString(hexValues)
	byteReader := bytes.NewReader(decStr)
	var b bytes.Buffer
	byteReader.WriteTo(&b)
	return b.Bytes()

}

func RlpCtrl(rlpBytes []byte) {
	if rlpBytes[0] == RlpStr {
		DecodeRlpMsg(rlpBytes[1:], "string")
		return
	}
	if rlpBytes[0] == RlpNr {
		DecodeRlpMsg(rlpBytes[1:], "number")
		return
	}
	// throw error if none
	log.Fatal("Could not find mapping for command byte in rlp stream")

}

func RlpDecoder(list1 string) {
	bytes1 := GetBytes(list1)
	RlpCtrl(bytes1)
}

// EncodeStrings - The rules for the encoder can be found here: https://eth.wiki/en/fundamentals/rlp
func DecodeRlpMsg(rlpBytes []byte, rlpType string) {
	lenStr := rlpBytes[0] - 0x80
	rlpString := rlpBytes[1 : lenStr+1]
	log.Printf("RLP %s: %s", rlpType, string(rlpString))
	remaining := rlpBytes[lenStr+1:]
	if len(remaining) != 0 {
		DecodeRlpMsg(remaining, rlpType)
	}
}
