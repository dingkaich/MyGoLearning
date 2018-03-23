package mycrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func mybase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func myhex() {
	fmt.Println(hex.EncodeToString([]byte("1234567890abcdef")))
	fmt.Println(hex.EncodeToString([]byte("ding")))
}

func MycryptMain() {
	fmt.Println(mybase64("12345678"))
	myhex()
	encryptaes("dingkaich", hex.EncodeToString([]byte("ding123")), hex.EncodeToString([]byte("12347890abcdef")))
	decryptaes("6n/pc08zcp+F7D7tF2XdcQ==", hex.EncodeToString([]byte("ding123")), hex.EncodeToString([]byte("12347890abcdef")))

}

//encryptaes
//cryptcontent  string 原始要加密的内容
// cipher hex 格式,密码短语，不足blocksize，按PKCS5 PADDING
// 初始向量  hex格式 不足blocksize，按PKCS5 PADDING
// reuturn 值为base64后的密文
func encryptaes(cryptcontent, key, iv string) {
	byte_cryptcontent := []byte(cryptcontent)
	byte_cipher, _ := hex.DecodeString(key)
	byte_iv, _ := hex.DecodeString(iv)

	if len(byte_cipher)%aes.BlockSize != 0 {
		fmt.Println(byte_cipher)
		byte_cipher = PKCS5Padding(byte_cipher, aes.BlockSize)
	}

	switch len(byte_cipher) {
	case 16:
		fmt.Println("aes-128")
	case 32:
		fmt.Println("aes-256")
	default:
		fmt.Println("error")
		return
	}

	if len(byte_iv)%aes.BlockSize != 0 {
		byte_iv = PKCS5Padding(byte_iv, aes.BlockSize)
	}

	byte_cryptcontent = PKCS5Padding(byte_cryptcontent, aes.BlockSize)

	block, err := aes.NewCipher(byte_cipher)
	if err != nil {
		fmt.Println(err)
		return
	}
	// The IV needs to be unique, but not secure. Therefore it's common to

	mode := cipher.NewCBCEncrypter(block, byte_iv)

	mode.CryptBlocks(byte_cryptcontent, byte_cryptcontent)
	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.
	final_str := base64.StdEncoding.EncodeToString(byte_cryptcontent)
	fmt.Printf("%s\n", final_str)
}

//decryptaes
//cryptcontent  string 原始要加密的内容
// cipher hex 格式,密码短语，不足blocksize，按PKCS5 PADDING
// 初始向量  hex格式 不足blocksize，按PKCS5 PADDING
// reuturn 值为base64后的密文
func decryptaes(cryptcontent, hexkey, hexiv string) {
	byte_cryptcontent, _ := base64.StdEncoding.DecodeString(cryptcontent)
	byte_key, _ := hex.DecodeString(hexkey)
	byte_iv, _ := hex.DecodeString(hexiv)

	if len(byte_key)%aes.BlockSize != 0 {
		byte_key = PKCS5Padding(byte_key, aes.BlockSize)
	}

	switch len(byte_key) {
	case 16:
		fmt.Println("aes-128")
	case 32:
		fmt.Println("aes-256")
	default:
		fmt.Println("error")
		return
	}

	if len(byte_iv)%aes.BlockSize != 0 {
		byte_iv = PKCS5Padding(byte_iv, aes.BlockSize)
	}

	block, err := aes.NewCipher(byte_key)
	if err != nil {
		panic(err)
	}

	mode := cipher.NewCBCDecrypter(block, byte_iv)
	mode.CryptBlocks(byte_cryptcontent, byte_cryptcontent)
	byte_cryptcontent = PKCS5UNPadding(byte_cryptcontent, aes.BlockSize)
	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	fmt.Printf("%s\n", string(byte_cryptcontent))
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	paddlen := blockSize - len(src)%blockSize
	tailblock := bytes.Repeat([]byte{byte(paddlen)}, paddlen)
	return append(src, tailblock...)
}

func PKCS5UNPadding(src []byte, blockSize int) []byte {
	by := src[len(src)-1]
	intby := len(src) - int(by)
	return src[:intby]
}
