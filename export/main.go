package main

import (
	"bufio"
	"fmt"
	"github.com/dmlvhh/mobile_sdk"
	"log"
	"net/url"
	"os"
)

func main() {
	// 打开文件
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatalf("无法打开文件: %v", err)
	}
	defer file.Close() // 确保文件在函数结束时关闭

	// 创建一个 Scanner 来逐行读取文件
	scanner := bufio.NewScanner(file)

	// 初始化 API 配置
	Tp := mobile_sdk.NewTP2Config("https://ac.buleideiot.com/api/get-chinamobile-onelink-token", "NNIYDQcitiDJngWA", "17")
	params := url.Values{}
	params.Add("channel_id", Tp.ChannelID)
	params.Add("sign", Tp.Sign)
	request, err := Tp.ApiGetRequest(params)
	if err != nil {
		log.Fatalf("获取 Token 失败: %v", err)
	}
	token := request.Data.Token
	cf := mobile_sdk.NewConfig("https://api.iot.10086.cn", "PMMxeJGqamJw7Jrh13361", "mUzJa4znlE7f9ZtdIReVFMFubUzx2vs6")

	// 创建输出文件
	outputFile, err := os.Create("./output.txt")
	if err != nil {
		log.Fatalf("无法创建输出文件: %v", err)
	}
	defer outputFile.Close()

	// 逐行读取文件内容
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text() // 获取当前行的内容
		//fmt.Printf("第 %d 行: %s\n", lineNumber, line)

		// 调用 API 获取 SIM 卡信息
		res, err2 := cf.SimImei(&mobile_sdk.SimBasicInfoReq{
			Transid: cf.TransID,
			Token:   token,
			Msisdn:  line, // 使用当前行的内容作为 MSISDN
		})
		if err2 != nil {
			fmt.Printf("API 调用失败: %v\n", err2)
			continue // 跳过当前行，继续处理下一行
		}

		// 如果 API 返回结果不为空，将结果写入输出文件
		if len(res.Result) != 0 {
			//outputLine := fmt.Sprintf("MSISDN: %s, IMEI: %s\n", line, res.Result)
			_, err := outputFile.WriteString(line + "," + res.Result[0].Imei + "\n")
			if err != nil {
				fmt.Printf("写入文件失败: %v\n", err)
			}
		}

		lineNumber++
	}

	// 检查是否有读取错误
	if err := scanner.Err(); err != nil {
		log.Fatalf("读取文件时出错: %v", err)
	}

	fmt.Println("处理完成，结果已写入 output.txt")
}
