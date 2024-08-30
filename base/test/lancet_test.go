package test

import (
	"fmt"
	"testing"

	"github.com/duke-git/lancet/v2/cryptor"
)

func TestExampleAesEcbEncrypt(t *testing.T) {
	data := "hello"
	key := "abcdefghijklmnop"

	encrypted := cryptor.AesEcbEncrypt([]byte(data), []byte(key))

	decrypted := cryptor.AesEcbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func TestExampleAesEcbDecrypt(t *testing.T) {
	data := "hello"
	key := "abcdefghijklmnop"

	encrypted := cryptor.AesEcbEncrypt([]byte(data), []byte(key))

	decrypted := cryptor.AesEcbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleAesCbcEncrypt() {
	data := "hello"
	key := "abcdefghijklmnop"

	encrypted := cryptor.AesCbcEncrypt([]byte(data), []byte(key))

	decrypted := cryptor.AesCbcDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleAesCbcDecrypt() {
	data := "hello"
	key := "abcdefghijklmnop"

	encrypted := cryptor.AesCbcEncrypt([]byte(data), []byte(key))

	decrypted := cryptor.AesCbcDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleAesCtrCrypt() {
	data := "hello"
	key := "abcdefghijklmnop"

	encrypted := cryptor.AesCtrCrypt([]byte(data), []byte(key))
	decrypted := cryptor.AesCtrCrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleAesCfbEncrypt() {
	data := "hello"
	key := "abcdefghijklmnop"

	encrypted := cryptor.AesCfbEncrypt([]byte(data), []byte(key))

	decrypted := cryptor.AesCfbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleAesCfbDecrypt() {
	data := "hello"
	key := "abcdefghijklmnop"

	encrypted := cryptor.AesCfbEncrypt([]byte(data), []byte(key))

	decrypted := cryptor.AesCfbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleAesOfbEncrypt() {
	data := "hello"
	key := "abcdefghijklmnop"

	encrypted := cryptor.AesOfbEncrypt([]byte(data), []byte(key))

	decrypted := cryptor.AesOfbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleAesOfbDecrypt() {
	data := "hello"
	key := "abcdefghijklmnop"

	encrypted := cryptor.AesOfbEncrypt([]byte(data), []byte(key))

	decrypted := cryptor.AesOfbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleDesEcbEncrypt() {
	data := "hello"
	key := "abcdefgh"

	encrypted := cryptor.DesEcbEncrypt([]byte(data), []byte(key))

	decrypted := cryptor.DesEcbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleDesEcbDecrypt() {
	data := "hello"
	key := "abcdefgh"

	encrypted := cryptor.DesEcbEncrypt([]byte(data), []byte(key))

	decrypted := cryptor.DesEcbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleDesCbcEncrypt() {
	data := "hello"
	key := "abcdefgh"

	encrypted := cryptor.DesCbcEncrypt([]byte(data), []byte(key))

	decrypted := cryptor.DesCbcDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleDesCbcDecrypt() {
	data := "hello"
	key := "abcdefgh"

	encrypted := cryptor.DesCbcEncrypt([]byte(data), []byte(key))

	decrypted := cryptor.DesCbcDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleDesCtrCrypt() {
	data := "hello"
	key := "abcdefgh"

	encrypted := cryptor.DesCtrCrypt([]byte(data), []byte(key))
	decrypted := cryptor.DesCtrCrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleDesCfbEncrypt() {
	data := "hello"
	key := "abcdefgh"

	encrypted := cryptor.DesCfbEncrypt([]byte(data), []byte(key))

	decrypted := cryptor.DesCfbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleDesCfbDecrypt() {
	data := "hello"
	key := "abcdefgh"

	encrypted := cryptor.DesCfbEncrypt([]byte(data), []byte(key))

	decrypted := cryptor.DesCfbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleDesOfbEncrypt() {
	data := "hello"
	key := "abcdefgh"

	encrypted := cryptor.DesOfbEncrypt([]byte(data), []byte(key))

	decrypted := cryptor.DesOfbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleDesOfbDecrypt() {
	data := "hello"
	key := "abcdefgh"

	encrypted := cryptor.DesOfbEncrypt([]byte(data), []byte(key))

	decrypted := cryptor.DesOfbDecrypt(encrypted, []byte(key))

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleGenerateRsaKey() {
	// Create ras private and public pem file
	err := cryptor.GenerateRsaKey(4096, "rsa_private.pem", "rsa_public.pem")
	if err != nil {
		return
	}

	fmt.Println("foo")

	// Output:
	// foo
}

func ExampleRsaEncrypt() {
	// Create ras private and public pem file
	err := cryptor.GenerateRsaKey(4096, "rsa_private.pem", "rsa_public.pem")
	if err != nil {
		return
	}

	data := []byte("hello")
	encrypted := cryptor.RsaEncrypt(data, "rsa_public.pem")
	decrypted := cryptor.RsaDecrypt(encrypted, "rsa_private.pem")

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleRsaDecrypt() {
	// Create ras private and public pem file
	err := cryptor.GenerateRsaKey(4096, "rsa_private.pem", "rsa_public.pem")
	if err != nil {
		return
	}

	data := []byte("hello")
	encrypted := cryptor.RsaEncrypt(data, "rsa_public.pem")
	decrypted := cryptor.RsaDecrypt(encrypted, "rsa_private.pem")

	fmt.Println(string(decrypted))

	// Output:
	// hello
}

func ExampleBase64StdEncode() {
	base64Str := cryptor.Base64StdEncode("hello")

	fmt.Println(base64Str)

	// Output:
	// aGVsbG8=
}

func ExampleBase64StdDecode() {
	str := cryptor.Base64StdDecode("aGVsbG8=")

	fmt.Println(str)

	// Output:
	// hello
}

func ExampleHmacMd5() {
	str := "hello"
	key := "12345"

	hms := cryptor.HmacMd5(str, key)
	fmt.Println(hms)

	// Output:
	// e834306eab892d872525d4918a7a639a
}

func ExampleHmacMd5WithBase64() {
	str := "hello"
	key := "12345"

	hms := cryptor.HmacMd5WithBase64(str, key)
	fmt.Println(hms)

	// Output:
	// 6DQwbquJLYclJdSRinpjmg==
}

func ExampleHmacSha1() {
	str := "hello"
	key := "12345"

	hms := cryptor.HmacSha1(str, key)
	fmt.Println(hms)

	// Output:
	// 5c6a9db0cccb92e36ed0323fd09b7f936de9ace0
}

func ExampleHmacSha1WithBase64() {
	str := "hello"
	key := "12345"

	hms := cryptor.HmacSha1WithBase64(str, key)
	fmt.Println(hms)

	// Output:
	// XGqdsMzLkuNu0DI/0Jt/k23prOA=
}

func ExampleHmacSha256() {
	str := "hello"
	key := "12345"

	hms := cryptor.HmacSha256(str, key)
	fmt.Println(hms)

	// Output:
	// 315bb93c4e989862ba09cb62e05d73a5f376cb36f0d786edab0c320d059fde75
}

func ExampleHmacSha256WithBase64() {
	str := "hello"
	key := "12345"

	hms := cryptor.HmacSha256WithBase64(str, key)
	fmt.Println(hms)

	// Output:
	// MVu5PE6YmGK6Ccti4F1zpfN2yzbw14btqwwyDQWf3nU=
}

func ExampleHmacSha512() {
	str := "hello"
	key := "12345"

	hms := cryptor.HmacSha512(str, key)
	fmt.Println(hms)

	// Output:
	// dd8f1290a9dd23d354e2526d9a2e9ce8cffffdd37cb320800d1c6c13d2efc363288376a196c5458daf53f8e1aa6b45a6d856303d5c0a2064bff9785861d48cfc
}

func ExampleHmacSha512WithBase64() {
	str := "hello"
	key := "12345"

	hms := cryptor.HmacSha512WithBase64(str, key)
	fmt.Println(hms)

	// Output:
	// 3Y8SkKndI9NU4lJtmi6c6M///dN8syCADRxsE9Lvw2Mog3ahlsVFja9T+OGqa0Wm2FYwPVwKIGS/+XhYYdSM/A==
}

func ExampleMd5String() {
	md5Str := cryptor.Md5String("hello")
	fmt.Println(md5Str)

	// Output:
	// 5d41402abc4b2a76b9719d911017c592
}

func ExampleMd5StringWithBase64() {
	md5Str := cryptor.Md5StringWithBase64("hello")
	fmt.Println(md5Str)

	// Output:
	// XUFAKrxLKna5cZ2REBfFkg==
}

func ExampleMd5Byte() {
	md5Str := cryptor.Md5Byte([]byte{'a'})
	fmt.Println(md5Str)

	// Output:
	// 0cc175b9c0f1b6a831c399e269772661
}

func ExampleMd5ByteWithBase64() {
	md5Str := cryptor.Md5ByteWithBase64([]byte("hello"))
	fmt.Println(md5Str)

	// Output:
	// XUFAKrxLKna5cZ2REBfFkg==
}

func ExampleSha1() {
	result := cryptor.Sha1("hello")
	fmt.Println(result)

	// Output:
	// aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d
}

func ExampleSha1WithBase64() {
	result := cryptor.Sha1WithBase64("hello")
	fmt.Println(result)

	// Output:
	// qvTGHdzF6KLavt4PO0gs2a6pQ00=
}

func ExampleSha256() {
	result := cryptor.Sha256("hello")
	fmt.Println(result)

	// Output:
	// 2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
}

func ExampleSha256WithBase64() {
	result := cryptor.Sha256WithBase64("hello")
	fmt.Println(result)

	// Output:
	// LPJNul+wow4m6DsqxbninhsWHlwfp0JecwQzYpOLmCQ=
}

func ExampleSha512() {
	result := cryptor.Sha512("hello")
	fmt.Println(result)

	// Output:
	// 9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043
}

func ExampleSha512WithBase64() {
	result := cryptor.Sha512WithBase64("hello")
	fmt.Println(result)

	// Output:
	// m3HSJL1i83hdltRq0+o9czGb+8KJDKra4t/3JRlnPKcjI8PZm6XBHXx6zG4UuMXaDEZjR1wuXDre9G9zvN7AQw==
}

func ExampleRsaEncryptOAEP() {
	pri, pub := cryptor.GenerateRsaKeyPair(1024)

	data := []byte("hello world")
	label := []byte("123456")

	encrypted, err := cryptor.RsaEncryptOAEP(data, label, *pub)
	if err != nil {
		return
	}

	decrypted, err := cryptor.RsaDecryptOAEP([]byte(encrypted), label, *pri)
	if err != nil {
		return
	}

	fmt.Println(string(decrypted))

	// Output:
	// hello world
}
