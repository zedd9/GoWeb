package main

import (
	"fmt"

	"github.com/zedd9/goweb/web9/cipher"
	"github.com/zedd9/goweb/web9/lzw"
)

type Component interface {
	Operator(string)
}

var sentData string
var recvData string

type SendCompoent struct {
}

func (self *SendCompoent) Operator(data string) {
	// Send Data
	sentData = data
}

type ZipComponet struct {
	com Component
}

func (self *ZipComponet) Operator(data string) {
	zipData, err := lzw.Write([]byte(data))
	if err != nil {
		panic(err)
	}

	self.com.Operator(string(zipData))
}

type EncryptComponent struct {
	key string
	com Component
}

func (self *EncryptComponent) Operator(data string) {
	encryptData, err := cipher.Encrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}

	self.com.Operator(string(encryptData))
}

type DecryptComponent struct {
	key string
	com Component
}

func (self *DecryptComponent) Operator(data string) {
	decryptData, err := cipher.Decrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}

	self.com.Operator(string(decryptData))
}

type UnzipComponent struct {
	com Component
}

func (self *UnzipComponent) Operator(data string) {
	unzipdata, err := lzw.Read([]byte(data))
	if err != nil {
		panic(err)
	}

	self.com.Operator(string(unzipdata))
}

type ReadComponent struct {
}

func (self *ReadComponent) Operator(data string) {
	recvData = data
}

func main() {
	sender := &EncryptComponent{
		key: "abcde",
		com: &ZipComponet{
			com: &SendCompoent{}}}

	sender.Operator("Hello World")

	fmt.Println(sentData)

	receiver := &UnzipComponent{
		com: &DecryptComponent{key: "abcde",
			com: &ReadComponent{}}}

	receiver.Operator(sentData)
	fmt.Println(recvData)
}
