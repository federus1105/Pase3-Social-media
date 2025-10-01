package pkg

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type HashConfig struct {
	Memory  uint32
	Time    uint32
	Thread  uint8
	KeyLen  uint32
	SaltLen uint32
}

func NewHashConfig() *HashConfig {
	hc := &HashConfig{}
	hc.UseRecommended() // langsung set config standar
	return hc
}

func (h *HashConfig) SetConfig(memory, time, keylen, saltlen uint32, thread uint8) {
	h.KeyLen = keylen
	h.SaltLen = saltlen
	h.Memory = memory
	h.Time = time
	h.Thread = thread
}

func (h *HashConfig) UseRecommended() {
	h.KeyLen = 32
	h.SaltLen = 16
	h.Memory = 64 * 1024
	h.Time = 2
	h.Thread = 4
}

func (hc *HashConfig) HashPassword(password string) (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, hc.Time, hc.Memory, hc.Thread, hc.KeyLen)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("%d.%d.%d.%s.%s", hc.Time, hc.Memory, hc.Thread, b64Salt, b64Hash)

	return encodedHash, nil
}

func (h *HashConfig) GenHash(password string) (string, error) {
	salt, err := h.genSalt()
	if err != nil {
		return "", err
	}
	hash := argon2.IDKey([]byte(password), salt, h.Time, h.Memory, h.Thread, h.KeyLen)
	// dalam penulisan hash ada format
	// $jenisKey$versiKey$konfigurasi(memory, time, thread)$salt$hash
	version := argon2.Version
	saltStr := base64.RawStdEncoding.EncodeToString(salt)
	hashStr := base64.RawStdEncoding.EncodeToString(hash)
	hashedPwd := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", version, h.Memory, h.Time, h.Thread, saltStr, hashStr)
	return hashedPwd, nil
}
func (h *HashConfig) genSalt() ([]byte, error) {
	salt := make([]byte, h.SaltLen)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	return salt, nil
}

func (h *HashConfig) CompareHashAndPassword(password, hashedPassword string) (bool, error) {
	result := strings.Split(hashedPassword, "$")
	if len(result) != 6 {
		return false, errors.New("invalid hash format")
	}
	if result[1] != "argon2id" {
		return false, errors.New("invalid crypto method")
	}
	var version int
	fmt.Sscanf(result[2], "v=%d", &version)
	if version != argon2.Version {
		return false, errors.New("invalid argon2id version")
	}

	var memory, time, threads uint32
	_, err := fmt.Sscanf(result[3], "m=%d,t=%d,p=%d", &memory, &time, &threads)
	if err != nil {
		return false, errors.New("invalid format")
	}
	h.Memory = memory
	h.Time = time
	h.Thread = uint8(threads)

	salt, err := base64.RawStdEncoding.DecodeString(result[4])
	if err != nil {
		return false, err
	}
	h.SaltLen = uint32(len(salt))

	hash, err := base64.RawStdEncoding.DecodeString(result[5])
	if err != nil {
		return false, err
	}
	h.KeyLen = uint32(len(hash))

	hashPwd := argon2.IDKey([]byte(password), salt, h.Time, h.Memory, h.Thread, h.KeyLen)
	if subtle.ConstantTimeCompare(hash, hashPwd) != 1 {
		return false, nil
	}
	return true, nil
}
