package main

import (
	"go-learning/cipher/cipher"
	"k8s.io/klog/v2"
)

func main() {

	b := []byte("hello,world")

	// 加密a
	encrypt, err := cipher.Encrypt(b)
	if err != nil {
		klog.Error("Encrypt Fail ：", err)
	}
	klog.Info("Encrypt suesses:", encrypt)

	// 解密
	decrypt, err := cipher.Decrypt(encrypt)
	if err != nil {
		klog.Error("Decrypt Fail ：", err)
	}
	klog.Info("Encrypt suesses:", string(decrypt))
}
