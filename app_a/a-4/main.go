package main

import (
	"fmt"
	"syscall"

	"github.com/tmc/keyring"
	"golang.org/x/term"
)

func main() {
	secretValue, err := keyring.Get("progo-keyring-test", "password")
	if err == keyring.ErrNotFound {
		// 未登録だった
		fmt.Println("Secret Value is not found. Please Type:")
		pw, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			panic(err)
		}
		// 登録
		err = keyring.Set("progo-keyring-test", "password", string(pw))
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		// 未知のエラー
		panic(err)
	} else {
		// 登録済みの値を表示
		fmt.Println("Secret Value: ", secretValue)
	}
}
