package main

import (
	"fmt"
	"github.com/whatsauth/watoken"
)

func main() {
	privateKey, publicKey := watoken.GenerateKey()
	fmt.Println("Ini Private Key")
	fmt.Println(privateKey)
	fmt.Println("ini public Key")
	fmt.Println(publicKey)
	//generate token for user awangga
	userid := "awangga"
	tokenstring, _ := watoken.Encode(userid, privateKey)
	fmt.Println(tokenstring)
	//decode token to get userid
	useridstring := watoken.DecodeGetId(publicKey, tokenstring)
	if useridstring == "" {
		fmt.Println("expire token")
	}

}
