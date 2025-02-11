package pub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

// 所有时间、时间戳均为UNIX时间，单位为毫秒。
// HTTP 返回代码
// HTTP 4XX 错误码用于指示错误的请求内容、行为、格式。问题在于请求者。
// HTTP 403 错误码表示违反WAF限制(Web应用程序防火墙)。
// HTTP 429 错误码表示警告访问频次超限，即将被封IP。 访问限制是基于IP的，而不是API Key
// HTTP 418 表示收到429后继续访问，于是被封了。
// HTTP 5XX 错误码用于指示Binance服务侧的问题。
// 建议您尽可能多地使用websocket消息获取相应数据，以减少请求带来的访问限制压力。

// 接口的基本信息
// GET 方法的接口, 参数必须在 query string中发送。
// POST, PUT, 和 DELETE 方法的接口,参数可以在内容形式为application/x-www-form-urlencoded的 query string 中发送，也可以在 request body 中发送。 如果你喜欢，也可以混合这两种方式发送参数。
// 对参数的顺序不做要求。但如果同一个参数名在query string和request body中都有，query string中的会被优先采用。

// 如果需要 API-key，应当在HTTP头中以X-MBX-APIKEY字段传递

// request data
type ParamData = map[string]interface{}

var serverTimeAhead = 0

// func init() {
// 	AdjustTime()
// }

func AdjustTime() {
	// totalSa := int64(0)
	// for i := 0; i < 5; i++ {
	// 	_, sa := ServerTime()
	// 	totalSa += sa
	// 	time.Sleep(200 * time.Millisecond)
	// }
	// avgSa := totalSa / 5
	// serverTimeAhead = int(avgSa)
	// log.Printf("serverTimeAhead %v", serverTimeAhead)
}

// 发送需要签名的GET请求
func GetWithSign(key *Key, path string, data ParamData) (resBody []byte, err error) {
	return getWithSign(futureBaseUrl, key, path, data)
}

func getWithSign(baseUrl string, key *Key, path string, data ParamData) (resBody []byte, err error) {
	// 添加recvWindow
	if data == nil {
		data = ParamData{"recvWindow": recvWindow}
	} else {
		data["recvWindow"] = recvWindow
	}

	// 参与计算签名的参数
	signData := make(map[string]interface{})
	// 请求所有参数都参与签名计算
	for k, v := range data {
		signData[k] = v
	}

	// 取币安的服务器时间
	timestamp := time.Now().UnixMilli() + int64(serverTimeAhead) // ServerTime()
	signData["timestamp"] = fmt.Sprintf("%v", timestamp)

	sign := NewSign(key.ApiKey, key.SecretKey)
	str, s := sign.BinanceGetSign(signData)

	var req *http.Request
	path += "?" + str + "&" + url.QueryEscape("signature") + "=" + url.QueryEscape(s)
	req, err = http.NewRequest("GET", baseUrl+path, nil)

	if err != nil {
		return
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36")
	req.Header.Add("Accept-Language", "en-US,en,zh-cn")
	req.Header.Add("X-MBX-APIKEY", sign.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	resBody, err = io.ReadAll(res.Body)
	if err != nil {
		return
	}

	var errMsg ErrMsg
	json.Unmarshal(resBody, &errMsg)
	if errMsg.Code != 0 {
		err = fmt.Errorf("httpreq.GetWithSign error %v", errMsg)
		switch errMsg.Code {
		case -1021: // -1021 Timestamp for this request was 1000ms ahead of the server's time.
			AdjustTime()
		}
	}
	return
}

// / 发送需要签名的POST请求
func PostWithSign(key *Key, path string, data ParamData) (resBody []byte, errMsg ErrMsg, err error) {
	// 添加recvWindow
	if data == nil {
		data = ParamData{"recvWindow": recvWindow}
	} else {
		data["recvWindow"] = recvWindow
	}

	signData := make(map[string]interface{})
	// GET 请求所有参数都参与签名计算，POST 请求业务参数不参与签名计算
	for k, v := range data {
		signData[k] = v
	}

	timestamp := time.Now().UnixMilli() + int64(serverTimeAhead) // ServerTime()
	signData["timestamp"] = fmt.Sprintf("%v", timestamp)
	sign := NewSign(key.ApiKey, key.SecretKey)
	str, s := sign.BinanceGetSign(signData)

	path += "?" + str + "&" + url.QueryEscape("signature") + "=" + url.QueryEscape(s)

	var req *http.Request
	req, err = http.NewRequest("POST", futureBaseUrl+path, nil)
	if err != nil {
		return
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36")
	req.Header.Add("Accept-Language", "zh-cn")
	req.Header.Add("X-MBX-APIKEY", sign.ApiKey)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	resBody, err = io.ReadAll(res.Body)
	if err != nil {
		return
	}

	json.Unmarshal(resBody, &errMsg)
	if errMsg.Code != 0 {
		err = fmt.Errorf("httpreq.PostWithSign resp err %v", errMsg)
		switch errMsg.Code {
		case -1021: // -1021 Timestamp for this request was 1000ms ahead of the server's time.
			AdjustTime()
		}
	}
	return
}

// 发送需要签名的Put请求
func PutWithSign(key *Key, path string, data ParamData) (resBody []byte, err error) {
	// 添加recvWindow
	if data == nil {
		data = ParamData{"recvWindow": recvWindow}
	} else {
		data["recvWindow"] = recvWindow
	}

	// 参与计算签名的参数
	signData := make(map[string]interface{})

	// 请求所有参数都参与签名计算
	for k, v := range data {
		signData[k] = v
	}

	// 取币安的服务器时间
	timestamp := time.Now().UnixMilli() + int64(serverTimeAhead) // ServerTime()
	signData["timestamp"] = fmt.Sprintf("%v", timestamp)

	sign := NewSign(key.ApiKey, key.SecretKey)
	str, s := sign.BinanceGetSign(signData)

	var req *http.Request
	path += "?" + str + "&" + url.QueryEscape("signature") + "=" + url.QueryEscape(s)
	req, err = http.NewRequest("PUT", futureBaseUrl+path, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36")
	req.Header.Add("Accept-Language", "zh-cn")
	req.Header.Add("X-MBX-APIKEY", sign.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	resBody, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var errMsg ErrMsg
	json.Unmarshal(resBody, &errMsg)
	if errMsg.Code != 0 {
		err = fmt.Errorf("httpreq.GetWithSign resp err %v", errMsg)
		switch errMsg.Code {
		case -1021: // -1021 Timestamp for this request was 1000ms ahead of the server's time.
			AdjustTime()
		}
	}

	return
}

// 发送需要签名的Delete请求
func DeleteWithSign(key *Key, path string, data ParamData) (resBody []byte, err error) {
	// 添加recvWindow
	if data == nil {
		data = ParamData{"recvWindow": recvWindow}
	} else {
		data["recvWindow"] = recvWindow
	}

	// 参与计算签名的参数
	signData := make(map[string]interface{})

	// 请求所有参数都参与签名计算
	for k, v := range data {
		signData[k] = v
	}

	// 取币安的服务器时间
	timestamp := time.Now().UnixMilli() + int64(serverTimeAhead) // ServerTime()
	signData["timestamp"] = fmt.Sprintf("%v", timestamp)

	sign := NewSign(key.ApiKey, key.SecretKey)
	str, s := sign.BinanceGetSign(signData)

	var req *http.Request
	path += "?" + str + "&" + url.QueryEscape("signature") + "=" + url.QueryEscape(s)
	req, err = http.NewRequest(http.MethodDelete, futureBaseUrl+path, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36")
	req.Header.Add("Accept-Language", "zh-cn")
	req.Header.Add("X-MBX-APIKEY", sign.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	resBody, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var errMsg ErrMsg
	json.Unmarshal(resBody, &errMsg)
	if errMsg.Code < 0 {
		err = fmt.Errorf("httpreq.GetWithSign resp err %v", errMsg)
		switch errMsg.Code {
		case -1021: // -1021 Timestamp for this request was 1000ms ahead of the server's time.
			AdjustTime()
		}
	}

	return
}

// / 发送原始请求，不需要签名
func GetNoSign(path string, params ParamData) (resBody []byte, err error) {
	url := futureBaseUrl + path
	if params != nil {
		url = url + "?" + EncodeQueryString(params, false)
	}

	var req *http.Request
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Mobile Safari/537.36")
	req.Header.Add("accept-language", "en,zh-CN;q=0.9,zh;q=0.8")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("upgrade-insecure-requests", "1")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// log.Println("httpreq.GetNoSign res.StatusCode", res.StatusCode)

	resBody, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var errMsg ErrMsg
	json.Unmarshal(resBody, &errMsg)
	if errMsg.Code != 0 {
		return nil, fmt.Errorf("httpreq.GetNoSign: %+v", errMsg)
	}
	return
}

// / 发送原始请求
func PostNoSign(path string, data ParamData) (resBody []byte, err error) {
	var body *bytes.Buffer
	if data == nil {
		data = ParamData{}
	}
	path += "?" + EncodeQueryString(data, false)

	// POST 请求 JSON
	if b, err := json.Marshal(data); err != nil {
		return nil, err
	} else {
		body = bytes.NewBuffer(b)
	}

	var req *http.Request
	if body != nil {
		req, err = http.NewRequest("POST", futureBaseUrl+path, body)
	}
	if err != nil {
		return
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36")
	req.Header.Add("Accept-Language", "zh-cn")
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	resBody, err = io.ReadAll(res.Body)
	if err != nil {
		return
	}

	var errMsg ErrMsg
	err = json.Unmarshal(resBody, &errMsg)
	if err == nil {
		return nil, fmt.Errorf("httpreq.GetWithSign: %+v", errMsg)
	}

	return
}

func SpotGetWithSign(key *Key, path string, data ParamData) (resBody []byte, err error) {
	return getWithSign(spotBaseUrl, key, path, data)
}
