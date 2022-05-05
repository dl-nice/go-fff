package utils

import (
	"log"
	"testing"
)

func TestFFFAddressEncode(t *testing.T) {

	log.Println(FFFAddressEncode("0xBbC473B491D84ad69954da075efd3cf142D54EEE"))

	log.Println(FFFAddressDecode("FFF5tXp4dJJWemxB4RPMmnVAcanbnsz8MH5Vi3Uf1caR8wtiKtW5nUsGmr"))
}
