package util

import (
	"crypto/rand"
	"errors"
	"math/big"

	"github.com/oklog/ulid/v2"
)

// Một số charset phổ biến để sử dụng
const (
	AlphaNumeric   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	AlphaLowerCase = "abcdefghijklmnopqrstuvwxyz"
	AlphaUpperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numeric        = "0123456789"
	AlphaOnly      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	HexLowerCase   = "0123456789abcdef"
	HexUpperCase   = "0123456789ABCDEF"
	Base64Safe     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"
	SpecialChars   = "!@#$%^&*()_+-=[]{}|;:,.<>?"
)

// GenerateRandomString tạo chuỗi ngẫu nhiên từ charset với độ dài length
// Sử dụng crypto/rand để đảm bảo tính ngẫu nhiên an toàn
func GenerateRandomString(charset string, length int) (string, error) {
	if length <= 0 {
		return "", errors.New("length must be greater than 0")
	}

	if len(charset) == 0 {
		return "", errors.New("charset cannot be empty")
	}

	result := make([]byte, length)
	charsetLen := big.NewInt(int64(len(charset)))

	for i := range length {
		// Sử dụng crypto/rand để tạo số ngẫu nhiên thực sự
		randomIndex, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			return "", err
		}
		result[i] = charset[randomIndex.Int64()]
	}

	return string(result), nil
}

// Các hàm helper cho các trường hợp sử dụng phổ biến
func GenerateAlphaNumeric(length int) (string, error) {
	return GenerateRandomString(AlphaNumeric, length)
}

func GenerateNumeric(length int) (string, error) {
	return GenerateRandomString(Numeric, length)
}

func GenerateAlphaOnly(length int) (string, error) {
	return GenerateRandomString(AlphaOnly, length)
}

func GenerateHex(length int) (string, error) {
	return GenerateRandomString(HexLowerCase, length)
}

// GeneratePassword tạo mật khẩu mạnh với tất cả các loại ký tự
func GeneratePassword(length int) (string, error) {
	charset := AlphaNumeric + SpecialChars
	return GenerateRandomString(charset, length)
}

// GenerateULID tạo ulid
func GenerateULID() string {
	return ulid.Make().String()
}
