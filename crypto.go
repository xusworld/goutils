package goutils

import (
	"crypto/md5"
	"crypto/sha1"
	"io"
)

// TODO 封装一些常见的加密/解密方法

// WARNING: MD5 is cryptographically broken and should not be used for secure applications.
func MD5(data []byte) []byte {
	hash := md5.New()
	hash.Write(data)
	return hash.Sum(nil)
}

func SHA1(data string) []byte {
	hash := sha1.New()
	_, _ = io.WriteString(hash, data)
	return hash.Sum(nil)
}