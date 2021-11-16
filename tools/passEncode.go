package tools

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

var key16 = []byte("ABCDEFGHIJKLMNOP")

// 加密
func AesCtrEncrypt(src []byte) []byte{
	fmt.Printf("明文为 %s\n",src)
	// 1 blockMode
	block,err := aes.NewCipher(key16)
	if err != nil {
		panic(err)
	}
	fmt.Printf("blockSize = %d\n",block.BlockSize())
	// 2 CTR
	iv := bytes.Repeat([]byte("1"),block.BlockSize())
	stream := cipher.NewCTR(block,iv)
	// 加密
	stream.XORKeyStream(src/*密文*/,src/*明文*/)
	return src
}

// 解密
func AesCtrDecrypt(cipherData []byte) []byte {
	block,err := aes.NewCipher(key16)
	if err != nil {
		panic(err)
	}
	iv := bytes.Repeat([]byte("1"),block.BlockSize())
	stream := cipher.NewCTR(block,iv)
	stream.XORKeyStream(cipherData,cipherData)
	return cipherData
}

func EncodeSha256(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	sum := h.Sum(nil)
	s := hex.EncodeToString(sum)
	return string(s)
}

func EncodeMd5(data string) string {
	w := md5.New()
	_, _ = io.WriteString(w, data)
	byData := w.Sum(nil)
	result := fmt.Sprintf("%x", byData)
	return result
}

func Base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func Base64Decode(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

// URL 专用编码
func UrlBase64Encode(info []byte) string {
	encodeInfo := base64.URLEncoding.EncodeToString(info)
	return encodeInfo
}

// URL 专用解码
func UrlBase64Decode(encodeInfo string) ([]byte,error) {
	info,err := base64.URLEncoding.DecodeString(encodeInfo)
	if err != nil {
		return nil,err
	}
	return info,nil
}