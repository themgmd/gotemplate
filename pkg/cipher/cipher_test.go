package cipher

import (
	"fmt"
	"testing"
)

func TestCipher_Encode(t *testing.T) {
	testcases := []struct {
		Text string
	}{
		{Text: "test long text here test long text here test long text here test long text here test long text here"},
		{Text: "test long text here"},
		{Text: "test"},
		{Text: "t"},
		{Text: ""},
	}

	for i := range testcases {
		fmt.Printf("test #%d\n", i+1)

		key, err := NewKey()
		if err != nil {
			t.Errorf("NewKey: %s", err.Error())
			return
		}
		fmt.Printf("key: %x\n", key)

		cipher := New(key)

		ciphertext, err := cipher.Encode(testcases[i].Text)
		if err != nil {
			t.Errorf("Error occurred while encode text: %s", err.Error())
			return
		}

		fmt.Printf("encoded: %s\n\n", ciphertext)
	}
}

func TestCipher_Decode(t *testing.T) {
	testcases := []struct {
		Key        []byte
		CipherText string
		PlaintText string
	}{
		{
			Key:        []byte("55c2b23b3d73b9238958ff19e885bda4"),
			CipherText: "4279424eed6abcae6782b82bd73b4595c96a5087538bea23d0e8e7d3b40941cd546f15444a6a107db6b785b78f82dd",
			PlaintText: "test long text here",
		},
		{
			Key:        []byte("ee9be81b59b4297228b3cfab36cfc819"),
			CipherText: "c6c61a3ae60a0a9f407a439f5772c6a35b7d96378dfa2289d5d83143d50165b6",
			PlaintText: "test",
		},
	}

	for i := range testcases {
		fmt.Printf("test #%d\n", i+1)

		cipher := New(testcases[i].Key)

		decoded, err := cipher.Decode(testcases[i].CipherText)
		if err != nil {
			t.Errorf("Error occurred while decode text: %s", err.Error())
			return
		}

		if decoded != testcases[i].PlaintText {
			t.Errorf("Expected %s, got: %s", testcases[i].PlaintText, decoded)
		}
	}
}
