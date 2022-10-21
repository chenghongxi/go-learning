package main

import (
	"go-learning/cipher/cipher"
	"k8s.io/klog/v2"
)

func main() {

	// 定义变量b 类型为：byte数组
	b := []byte("hello,world")

	// 加密b
	encrypt, err := cipher.Encrypt(b)
	if err != nil {
		klog.Error("Encrypt Fail ：", err)
	}
	klog.Info("Encrypt suesses:", encrypt)

	// 对上一步加密变量解密
	decrypt, err := cipher.Decrypt(encrypt)
	if err != nil {
		klog.Error("Decrypt Fail ：", err)
	}
	klog.Info("Encrypt suesses:", string(decrypt))
}
