package sintaxes

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
)

func Int64ParaBytesBigEndian(valor int64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(valor))
	return buf
}

func Int64ParaBytesLittleEndian(valor int64) []byte {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(valor))
	return buf
}

func ParaJSON[T any](valor T) ([]byte, error) {
	return json.Marshal(valor)
}

func DeJSON[T any](dados []byte) (T, error) {
	var valor T
	err := json.NewDecoder(bytes.NewReader(dados)).Decode(&valor)
	return valor, err
}
