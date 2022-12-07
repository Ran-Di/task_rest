package main

import (
	"fmt"
	"task_rest/crypting"
)

func main() {
	fmt.Println("Test GiT")

	fmt.Println("Crypting 'AAABBACACACACDACDCDDDD' ")
	fmt.Println("Result: ", crypting.Crypt("AAABBACACACACDACDCDDDD")) //AAABBACACACACD ABBFBBFBBFBBDBBFBBFBBFBBDE
	fmt.Println("Decrypting '3A2B4(AC)DA2(CD)3D' ")
	fmt.Println("Result: ", crypting.Decrypt("3A2B4(AC)DA2(CD)3D"))

}
