package base

import (
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"sort"
	"strings"
)

// 获取签名
func GetSign(params url.Values, apisecret string) string {

	params.Add("apisecret", apisecret)

	keys := make([]string, 0)
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	tmp := make([]string, 0)
	for _, v := range keys {
		tmp = append(tmp, v+"="+params.Get(v))
	}

	sign := Md5V(strings.Join(tmp, "&"))
	params.Del("apisecret")

	return sign
}

func Md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
