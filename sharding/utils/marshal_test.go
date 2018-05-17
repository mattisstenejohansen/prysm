package utils

import (
	"math/rand"
	"reflect"
	"testing"
)

func buildRawBlob(size int64) []RawBlob {
	tempbody := make([]RawBlob, size)
	for i := int64(0); i < size; i++ {
		var rawblob RawBlob
		rawblob.data = buildBlob(size)
		flagset := byte(rand.Int()) >> 7
		if flagset == byte(1) {
			rawblob.flags.skipEvmExecution = true

		}
		tempbody[i] = rawblob
	}

	return tempbody
}

func buildBlob(size int64) []byte {
	tempbody := make([]byte, size)
	for i := int64(0); i < size; i++ {
		tempbody[i] = byte(rand.Int())
	}

	return tempbody
}

func TestSize(t *testing.T) {
	for i := 0; i < 300; i++ {
		size := int64(i)
		blob := buildRawBlob(size)
		chunksafterSerialize := size / chunkDataSize
		terminalchunk := size % chunkDataSize
		if terminalchunk != 0 {
			chunksafterSerialize = chunksafterSerialize + 1
		}
		chunksafterSerialize = chunksafterSerialize * size
		sizeafterSerialize := chunksafterSerialize * chunkSize

		drefbody := make([]*RawBlob, len(blob))
		for s := 0; s < len(blob); s++ {
			drefbody[s] = &(blob[s])

		}
		serializedblob, err := Serialize(drefbody)
		if err != nil {
			t.Errorf("Error Serializing blob:%v\n %v", err, serializedblob)
		}

		if int64(len(serializedblob)) != sizeafterSerialize {
			t.Errorf("Error Serializing blobs the lengths are not the same:\n %d \n %d", int64(len(serializedblob)), sizeafterSerialize)
		}
	}

}
func TestSerializeAndDeserializeblob(t *testing.T) {

	for i := 1; i < 300; i++ {

		blob := buildRawBlob(int64(i))

		drefbody := make([]*RawBlob, len(blob))
		for s := 0; s < len(blob); s++ {
			drefbody[s] = &(blob[s])
		}

		serializedblob, err := Serialize(drefbody)

		if err != nil {
			t.Errorf("Error Serializing blob at index %d:\n%v\n%v", i, err, serializedblob)
		}
		raw, err2 := Deserialize(serializedblob)
		if err2 != nil {
			t.Errorf("Error Serializing blob at index %d:\n%v due to \n%v", i, raw, err2)
		}

		if !reflect.DeepEqual(blob, raw) {
			t.Errorf("Error Serializing blobs at index %d, the serialized and deserialized versions are not the same:\n\n %v \n\n %v \n\n %v", i, blob, serializedblob, raw)
		}
	}

}
