package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	number := 128.98112
	fmt.Printf(fmt.Sprintf("%.2f\n", number))

	data := []byte{0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xA, 0xB, 0xC, 0xD, 0xE, 0xF}
	hexString := fmt.Sprintf("%x\n", data)
	hexString2 := hex.EncodeToString(data)
	fmt.Println(hexString)
	fmt.Println(hexString2)
}
