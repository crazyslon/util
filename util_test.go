package util

import (
	"encoding/base64"
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMd5(t *testing.T) {
	result := Md5("123456")
	assert.Equal(t, "e10adc3949ba59abbe56e057f20f883e", result)
}

func TestDistinctUint(t *testing.T) {
	result := DistinctUint([]uint{1, 2, 2, 3, 4, 4, 5, 5, 5})
	assert.Equal(t, []uint{1, 2, 3, 4, 5}, result)
}

func TestDistinctInt64(t *testing.T) {
	result := DistinctInt64([]int64{1, 2, 2, 3, 4, 4, 5, 5, 5})
	assert.Equal(t, []int64{1, 2, 3, 4, 5}, result)
}

func TestDistincString(t *testing.T) {
	result := DistinctString([]string{"1", "2", "2", "3", "4", "4", "5", "5", "5"})
	assert.Equal(t, []string{"1", "2", "3", "4", "5"}, result)
}

func TestDistinctLowerCase(t *testing.T) {
	result := DistinctLowerCase([]string{"TEST1", "TEST2", "test2", "TeSt3", "tEsT4", "tEST4", "TEST5", "test5", "Test5"})
	assert.Equal(t, []string{"test1", "test2", "test3", "test4", "test5"}, result)
}

func TestContainsStringContains(t *testing.T) {
	result := ContainsString([]string{"1", "2", "3", "4", "5"}, "4")
	assert.True(t, result)
}

func TestContainsStringNotContains(t *testing.T) {
	result := ContainsString([]string{"1", "2", "3", "4", "5"}, "6")
	assert.False(t, result)
}

func TestEncrypt(t *testing.T) {
	encryptionKey := []byte("1234567891234568")

	text := []byte("some text")
	ciphertext, err := Encrypt(text, encryptionKey)
	assert.Nil(t, err)
	assert.NotEmpty(t, ciphertext)

	originalText, err := Decrypt(ciphertext, encryptionKey)
	assert.Nil(t, err)
	assert.Equal(t, text, originalText)
}

func TestEncryptWhenKeyInvalid(t *testing.T) {
	encryptionKey := []byte("1234567891234568")
	text := []byte("some text")
	ciphertext, err := Encrypt(text, encryptionKey)
	assert.Nil(t, err)
	assert.NotEmpty(t, ciphertext)

	anotherencryptionKey := []byte("1234567891234569")
	_, err = Decrypt(ciphertext, anotherencryptionKey)
	assert.NotNil(t, err)
}

func TestEncryptBase64(t *testing.T) {
	encryptionKey := []byte("1234567891234568")

	text := []byte("some text")
	ciphertext, err := EncryptBase64(text, encryptionKey)
	assert.Nil(t, err)
	assert.NotEmpty(t, ciphertext)

	originalText, err := DecryptBase64(ciphertext, encryptionKey)
	assert.Nil(t, err)
	assert.Equal(t, text, originalText)
}

func TestEncryptBase64WhenHashInvalid(t *testing.T) {
	encryptionKey := []byte("1234567891234568")
	text := []byte("some text")
	ciphertext, err := EncryptBase64(text, encryptionKey)
	assert.Nil(t, err)
	assert.NotEmpty(t, ciphertext)

	invalidCiphertext := base64.URLEncoding.EncodeToString(text)
	_, err = DecryptBase64(invalidCiphertext, encryptionKey)
	assert.NotNil(t, err)
}

func TestEncryptBase64JSON(t *testing.T) {

	type TestType struct {
		Something1 uint64
		Something2 uint64
		Something3 uint64
		Something4 uint64
		Something5 string
		Something6 string
		Something7 int64
	}

	testType := TestType{
		Something1: 1,
		Something2: 1,
		Something3: 1,
		Something4: 1,
		Something5: "test.com",
		Something6: "https://test.com/search?newwindow=1&source=hp&ei=ANetWvW7Jca2sQGbirDYDQ&q=test&oq=test&gs_l=psy-ab.12..0l10.3653.4088.0.5146.4.4.0.0.0.0.113.351.3j1.4.0....0...1c.1.64.psy-ab..0.4.349...0i131k1.0.Ho3HCjOhEjw",
		Something7: time.Now().Add(time.Second * 60).Unix(),
	}
	testTypeJSON, _ := json.Marshal(testType)

	encryptionKey := []byte("1234567891234568")
	ciphertext, err := EncryptBase64(testTypeJSON, encryptionKey)
	assert.Nil(t, err)
	assert.NotEmpty(t, ciphertext)

	originalJSON, err := DecryptBase64(ciphertext, encryptionKey)
	assert.Nil(t, err)
	assert.Equal(t, testTypeJSON, originalJSON)
}

func BenchmarkEncryptBase64(b *testing.B) {

	type TestType struct {
		Something1 uint64
		Something2 uint64
		Something3 uint64
		Something4 uint64
		Something5 string
		Something6 string
		Something7 int64
	}

	testType := TestType{
		Something1: 1,
		Something2: 1,
		Something3: 1,
		Something4: 1,
		Something5: "test.com",
		Something6: "https://test.com/search?newwindow=1&source=hp&ei=ANetWvW7Jca2sQGbirDYDQ&q=test&oq=test&gs_l=psy-ab.12..0l10.3653.4088.0.5146.4.4.0.0.0.0.113.351.3j1.4.0....0...1c.1.64.psy-ab..0.4.349...0i131k1.0.Ho3HCjOhEjw",
		Something7: time.Now().Add(time.Second * 60).Unix(),
	}
	testTypeJSON, _ := json.Marshal(testType)
	encryptionKey := []byte("1234567891234568")
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		ciphertext, _ := EncryptBase64(testTypeJSON, encryptionKey)
		DecryptBase64(ciphertext, encryptionKey)
	}
}
