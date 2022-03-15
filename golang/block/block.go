package block

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func Block(dataStr string, decrypted string) {

	//Güvenliği arttırmak için anahtarı rastgele döngü halinde değiştirmek.

}
func Encrypt(stringToEncrypt string, keyString string) (encryptedString string) {

	//Anahtar hex. olduğundan, kodu bayt'a dönüştürmemiz gerekiyor.
	key, _ := hex.DecodeString(keyString)
	plaintext := []byte(stringToEncrypt)

	//Anahtarı kullanarak yeni şifre örneği alın.
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//GCM oluştur. GCM simetrik anahtar şifreleme blok şifrelemeleri için kullanılır. örnekleri https://golang.hotexamples.com/examples/crypto.cipher/-/NewGCM/golang-newgcm-function-examples.html

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Bir nonce oluşturun. Nonce, GCM'den olmalıdır. anahtar tekrarlama riskine karşılık.
	nonce := make([]byte, aesGCM.NonceSize()) // nonce size yerine byte eklentiside yapılabilir.
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	//aesGCM.Seal kullanarak verileri şifreleyin
	//Bu durumda nonce'yi başka bir yere kaydetmek istemediğimiz için onu şifrelenmiş verilere örnek olarak ekliyoruz. Seal'deki ilk nonce argümanı örnektir.
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil) //şifrele. örnekler https://golang.hotexamples.com/examples/crypto.cipher/AEAD/Open/golang-aead-open-method-examples.html
	return fmt.Sprintf("%x", ciphertext)
}

func Decrypt(encryptedString string, keyString string) (decryptedString string) {

	key, _ := hex.DecodeString(keyString)
	enc, _ := hex.DecodeString(encryptedString)

	//Anahtardan yeni bir Şifre Bloğu oluşturun.
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//Yeni bir GCM oluşturun.
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//nonce boyutunu al.
	nonceSize := aesGCM.NonceSize()

	//Şifrelenmiş verilerden nonce'yi çıkarın.
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Verilerin şifresini çöz
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", plaintext)
}
