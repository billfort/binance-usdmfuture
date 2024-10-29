package pub

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

type Sign struct {
	ApiKey           string
	SecretKey        string
	SignatureMethod  string
	SignatureVersion string
}

func NewSign(apiKey, secretKey string) *Sign {
	return &Sign{
		ApiKey:           apiKey,
		SecretKey:        secretKey,
		SignatureMethod:  "HmacSHA256",
		SignatureVersion: "2",
	}
}

// 返回：签名串，签名
func (s *Sign) BinanceGetSign(params map[string]interface{}) (str string, sign string) {
	str = EncodeQueryString(params, false)
	return str, BinanceHmac256(str, s.SecretKey)
}

func BinanceHmac256(data string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func (s *Sign) HuobiGetSign(method, host, path, timestamp string, params map[string]interface{}) (string, error) {
	var str = method + "\n" + host + "\n" + path + "\n"
	params["ApiKey"] = s.ApiKey
	params["SignatureMethod"] = s.SignatureMethod
	params["SignatureVersion"] = s.SignatureVersion
	params["Timestamp"] = timestamp
	str += EncodeQueryString(params, true)
	return HuobiHmac256(str, s.SecretKey), nil
}

func HuobiHmac256(data string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func GetMapKeys(m map[string]interface{}) (keys []string) {
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func SortKeys(keys []string) []string {
	sort.Strings(keys)
	return keys
}

// / 拼接query字符串
func EncodeQueryString(query map[string]interface{}, bSort bool) string {
	var keys []string
	if bSort {
		keys = SortKeys(GetMapKeys(query))
	} else {
		keys = make([]string, len(query))
		var i int = 0
		for k := range query {
			keys[i] = k
			i = i + 1
		}
	}
	var len = len(keys)
	var lines = make([]string, len)
	for i := 0; i < len; i++ {
		var k = keys[i]
		lines[i] = url.QueryEscape(k) + "=" + url.QueryEscape(fmt.Sprintf("%v", query[k]))
	}
	return strings.Join(lines, "&")
}
