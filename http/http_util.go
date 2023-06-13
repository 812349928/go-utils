package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Request(url, method, data string, contentType string, header map[string]string) ([]byte, error) {
	payload := strings.NewReader(data)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if len(contentType) > 0 {
		req.Header.Set("Content-Type", contentType)
	} else {
		req.Header.Set("Content-Type", "application/json")
	}

	for k, v := range header {
		req.Header.Set(k, v)
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return body, nil
}

func Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func GetIp() (string, error) {
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		//fmt.Println("获取公网 IP 失败:", err)
		return "", err
	}
	defer resp.Body.Close()

	ipBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//fmt.Println("读取响应失败:", err)
		return "", err
	}

	ip := string(ipBytes)
	//fmt.Println("公网 IP 地址:", ip)
	return ip, nil
}

func GetIpPosition(ip string) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	// 构建 API 请求 URL
	url := fmt.Sprintf("http://ip-api.com/json/%s", ip)

	// 发送 HTTP GET 请求
	resp, err := http.Get(url)
	if err != nil {
		//fmt.Println("发送 HTTP 请求失败:", err)
		return result, err
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//fmt.Println("读取响应失败:", err)
		return result, err
	}

	// 解析 JSON 响应
	err = json.Unmarshal(body, &result)
	if err != nil {
		//fmt.Println("解析 JSON 响应失败:", err)
		return nil, err
	}

	//// 判断 IP 地址所属国家是否为中国
	//country := result["country"].(string)
	//if country == "China" {
	//	fmt.Println("IP 地址属于中国")
	//} else {
	//	fmt.Println("IP 地址不属于中国")
	//}

	return result, nil
}
