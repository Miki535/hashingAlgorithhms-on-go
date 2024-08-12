package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
	"time"

	"golang.org/x/crypto/sha3"
)

func main() {
	var te string
	var ts int8
	var info string
	fmt.Println("Hello in our hashing program! (start or exit)")
	fmt.Scan(&te)
	for {
		if te == "start" {
			fmt.Println("\nEnter information:")
			reader := bufio.NewReader(os.Stdin)
			info, _ = reader.ReadString('\n')
			info = info[:len(info)-1]

			fmt.Println("Now choose hashing algoritm:")
			time.Sleep(100 * time.Millisecond)
			fmt.Println("1) md5")
			time.Sleep(100 * time.Millisecond)
			fmt.Println("2) sha1")
			time.Sleep(100 * time.Millisecond)
			fmt.Println("3) sha256")
			time.Sleep(100 * time.Millisecond)
			fmt.Println("4) sha512")
			time.Sleep(100 * time.Millisecond)
			fmt.Println("5) sha512_224")
			time.Sleep(100 * time.Millisecond)
			fmt.Println("6) sha512_256")
			time.Sleep(100 * time.Millisecond)
			fmt.Println("7) sha512_384")
			time.Sleep(100 * time.Millisecond)
			fmt.Println("8) sha512_512")
			time.Sleep(100 * time.Millisecond)
			fmt.Println("9) sha3")
			time.Sleep(100 * time.Millisecond)
			fmt.Println("10) sha3_224")
			time.Sleep(100 * time.Millisecond)
			fmt.Println("11) sha3_256")
			fmt.Scan(&ts)
			switch ts {
			case 1:
				hashing(info, "md5")
			case 2:
				hashing(info, "sha1")
			case 3:
				hashing(info, "sha256")
			case 4:
				hashing(info, "sha512")
			case 5:
				hashing(info, "sha512_224")
			case 6:
				hashing(info, "sha512_256")
			case 7:
				hashing(info, "sha512_384")
			case 8:
				hashing(info, "sha512_512")
			case 9:
				hashing(info, "sha3")
			case 10:
				hashing(info, "sha3_224")
			case 11:
				hashing(info, "sha3_256")
			default:
				fmt.Println("Unknown hashing algorithm")
				os.Exit(0)
			}
		} else if te == "exit" {
			fmt.Println("Bye!")
			os.Exit(0)
		} else {
			fmt.Println("Unknown command!")
			os.Exit(0)
		}
	}
}

func hashing(info string, test string) {
	data := []byte(info)
	switch test {
	case "md5":
		fmt.Printf("%x", md5.Sum(data))
	case "sha1":
		fmt.Printf("%x", sha1.Sum(data))
	case "sha256":
		fmt.Printf("%x", sha256.Sum256(data))
	case "sha512":
		fmt.Printf("%x", sha512.Sum512(data))
	case "sha512_224":
		fmt.Printf("%x", sha256.Sum224(data))
	case "sha512_256":
		fmt.Printf("%x", sha256.Sum256(data))
	case "sha512_384":
		fmt.Printf("%x", sha512.Sum384(data))
	case "sha512_512":
		fmt.Printf("%x", sha512.Sum512(data))
	case "sha3":
		fmt.Printf("%x", sha3.Sum256(data))
	case "sha3_224":
		fmt.Printf("%x", sha3.Sum224(data))
	case "sha3_256":
		fmt.Printf("%x", sha3.Sum256(data))
	}
}
