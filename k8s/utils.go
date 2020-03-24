package k8s

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
)

// 加密
func Jiami(code string) string {
	w := md5.New()
	io.WriteString(w, code)
	md5str2 := fmt.Sprintf("%x", w.Sum(nil))
	return md5str2
}

// 加密base64
func EncodeBase64(in string) string {
	return base64.StdEncoding.EncodeToString([]byte(in))
}
