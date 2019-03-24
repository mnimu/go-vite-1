package chain_utils

import (
	"encoding/binary"
	"github.com/vitelabs/go-vite/chain/block"
	"github.com/vitelabs/go-vite/common/types"
)

func SerializeAccountIdHeight(accountId, height uint64) []byte {
	return nil
}

func DeserializeAccountIdHeight(buf []byte) (uint64, uint64) {
	return 0, 0
}

func DeserializeHashList(buf []byte) []*types.Hash {
	return nil
}

func SerializeAccountId(accountId uint64) []byte {
	return nil
}

func DeserializeAccountId(buf []byte) uint64 {
	return 0
}

func SerializeHeight(height uint64) []byte {
	return nil
}

func SerializeLocation(location *chain_block.Location) []byte {
	bytes := make([]byte, 12)
	binary.BigEndian.PutUint64(bytes, location.FileId())

	binary.BigEndian.PutUint32(bytes[8:], uint32(location.Offset()))

	return bytes
}

func DeserializeLocation(bytes []byte) *chain_block.Location {
	return chain_block.NewLocation(FixedBytesToUint64(bytes[:8]), int64(binary.BigEndian.Uint32(bytes[8:])))
}

func SerializeUint64(number uint64) []byte {
	return nil
}
func DeserializeUint64(buf []byte) uint64 {
	return 0
}

func Uint64ToFixedBytes(height uint64) []byte {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, height)
	return bytes
}

func FixedBytesToUint64(bytes []byte) uint64 {
	return binary.BigEndian.Uint64(bytes)
}
