package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

var (
	// This variable should be set by the linker when compiling
	// go build -ldflags "-X main.secretKey=newsecretkey" main.go
	// or manually right here before build
	secretKey = "secretkey"
	alphabet  = []string{
		"0123456789",
		"abcdefghijklmnopqrstuvwxyz",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"\"!#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"}
	power       = len(alphabet)
	passwordLen = 18
)

func getAlphabet(abc []string, pwr int) string {
	var res string
	if len(abc) == 0 {
		abc = alphabet
	}
	if pwr < 1 || pwr > len(abc) {
		pwr = len(abc)
	}
	for _, c := range abc[:pwr] {
		res += c
	}
	return res
}

func generatePassword(secKey, key, abc string, ln int) string {
	var password string
	counter := 0
	zero := big.NewInt(0)
	i := new(big.Int)
	d := big.NewInt(int64(len(abc)))
	m := new(big.Int)
	for {
		h := sha256.Sum256([]byte(secKey + key + strconv.Itoa(ln) + strconv.Itoa(counter)))
		i.SetBytes(h[:])
		for i.Cmp(zero) > 0 {
			i.DivMod(i, d, m)
			password += string(abc[m.Uint64()])
			if len(password) >= ln {
				return password
			}
		}
		counter++
	}
}

func init() {
	flag.IntVar(&passwordLen, "l", passwordLen, "password length")
	flag.IntVar(&power, "p", power, fmt.Sprintf("password strength 1-%s", strconv.Itoa(power)))
	flag.Parse()
}

func main() {
	nargs := len(os.Args)
	if nargs < 2 || nargs%2 != 0 {
		fmt.Printf("Usage:\n%s [-l] [-p] <keyword>\n\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
	key := os.Args[nargs-1]
	abc := getAlphabet(alphabet, power)
	pass := generatePassword(secretKey, key, abc, passwordLen)
	fmt.Println(pass)
}
