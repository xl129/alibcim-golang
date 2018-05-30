package im

import (
	"crypto/md5"
	"encoding/hex"
	"crypto/hmac"
	"crypto/sha1"
	"strings"
	"sort"
)

func Md5(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	md5Data := h.Sum([]byte(""))
	return strings.ToUpper(hex.EncodeToString(md5Data))
}

func Hmac(key, data string) string {
	h := hmac.New(md5.New, []byte(key))
	h.Write([]byte(data))
	return strings.ToUpper(hex.EncodeToString(h.Sum([]byte(""))))
}

func Sha1(data string) string {
	h := sha1.New()
	h.Write([]byte(data))
	return strings.ToUpper(hex.EncodeToString(h.Sum([]byte(""))))
}

func Sign(data map[string]string, secret_key string) string {
	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var sign_str string
	for _, k := range keys {
		sign_str += k + data[k]
	}
	return Hmac(secret_key, sign_str)
}
