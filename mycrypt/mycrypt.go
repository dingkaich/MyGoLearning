package mycrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"

	"golang.org/x/crypto/pbkdf2"
)

// 传入Base64的rootkey
func mypbkdf2(rootkey string) []byte {
	pwd, _ := base64.StdEncoding.DecodeString(rootkey)
	// pwd := []byte(rootkey)                                                                                         // 用户设置的原始密码
	salt := []byte{1, 2, 3, 3, 4, 5, 6, 6, 7, 8, 9, 0, 1, 1, 2, 3, 5, 6} // 盐，是一个随机字符串，每一个用户都不一样，在这里我们随机选择 "I1lrI7wqJOJZ" 作为盐
	// iterations := 1000                                                                                             // 1000 次
	// digest := sha1.New
	log.Println(pwd)  // digest 算法，使用 sha256
	log.Println(salt) // digest 算法，使用 sha256

	// 第一步：使用 pbkdf2 算法加密
	dk := pbkdf2.Key(pwd, salt, 1000, 16, sha1.New)
	log.Println(dk)

	// 第二步：Base64 编码
	str := base64.StdEncoding.EncodeToString(dk)
	log.Println(str)

	// 第三步：组合加密算法、迭代次数、盐、密码和分割符号 "$"
	// log.Println("pbkdf2_sha256" + "$" + strconv.FormatInt(int64(iterations), 10) + "$" + string(salt) + "$" + str)
	return dk
}

func mybase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func myhex() {
	fmt.Println(hex.EncodeToString([]byte("1234567890abcdef")))
	fmt.Println(hex.EncodeToString([]byte("ding")))
}

func MycryptMain() {
	// rootkey, _ := base64.StdEncoding.DecodeString(string("VQ6cFVjNCGXUW049fB1QnQ=="))
	mypbkdf2("YXNkZmFzZmFmCg==")

	// aaa, _ := hex.DecodeString("ba49a299d69ed6e3414d1df0ff07ed80")
	// fmt.Printf("len=%d,", len(aaa))
	// PKCS5UNPadding(secondkey, 16)
	// fmt.Println(len(secondkey), hex.EncodeToString(secondkey))
	// fmt.Println("==", string(mingwen))

	// fmt.Println(mybase64("12345678"))
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
	// a := make([]byte, 16)
	// copy(a, byte_cryptcontent[:16])
	final_str := base64.StdEncoding.EncodeToString(byte_cryptcontent)

	fmt.Printf("%s\n", final_str)
}

func decryptaes1(cryptcontent string, hexkey string, hexiv string) string {
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
		return ""
	}

	if len(byte_iv)%aes.BlockSize != 0 {
		byte_iv = PKCS5Padding(byte_iv, aes.BlockSize)
	}

	block, err := aes.NewCipher(byte_key)
	if err != nil {
		panic(err)
	}

	mode := cipher.NewCBCDecrypter(block, byte_iv)
	// mode.CryptBlocks(byte_cryptcontent, byte_cryptcontent[:16])
	fmt.Println("lencyp=", len(byte_cryptcontent))
	// byte_cryptcontent_real := make([]byte, 0, 16)
	// copy(byte_cryptcontent_real, byte_cryptcontent)
	// byte_cryptcontent_real = append(byte_cryptcontent_real, byte_cryptcontent[:16]...)

	mode.CryptBlocks(byte_cryptcontent, byte_cryptcontent)

	// byte_cryptcontent = PKCS5UNPadding(byte_cryptcontent, aes.BlockSize)
	fmt.Printf(" decrypt len=%d value=%s\n", len(byte_cryptcontent), hex.EncodeToString(byte_cryptcontent))
	return hex.EncodeToString(byte_cryptcontent)
	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	// fmt.Printf("%s\n", string(byte_cryptcontent))

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
	fmt.Println("len=", len(byte_cryptcontent))
	// return byte_cryptcontent
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
