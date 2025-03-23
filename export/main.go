package main

import (
	"bufio"
	"fmt"
	"github.com/dmlvhh/mobile_sdk"
	"log"
	"net/url"
	"os"
	"sync"
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
	outputFile, err := os.Create("./output2.txt")
	if err != nil {
		log.Fatalf("无法创建输出文件: %v", err)
	}
	defer outputFile.Close()

	// 创建错误记录文件
	errFile, err := os.Create("./err.txt")
	if err != nil {
		log.Fatalf("无法创建错误文件: %v", err)
	}
	defer errFile.Close()

	// 创建一个 WaitGroup 来等待所有 goroutine 完成
	var wg sync.WaitGroup

	// 创建一个 channel 用于接收处理结果
	resultChan := make(chan string)
	errChan := make(chan string)

	// 启动一个 goroutine 来处理结果写入文件
	go func() {
		for result := range resultChan {
			_, err := outputFile.WriteString(result)
			if err != nil {
				fmt.Printf("写入文件失败: %v\n", err)
			}
		}
	}()

	// 启动一个 goroutine 来处理错误写入文件
	go func() {
		for errLine := range errChan {
			_, err := errFile.WriteString(errLine + "\n")
			if err != nil {
				fmt.Printf("写入错误文件失败: %v\n", err)
			}
		}
	}()

	// 逐行读取文件内容
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text() // 获取当前行的内容

		// 启动一个 goroutine 处理每一行
		wg.Add(1)
		go func(line string, lineNumber int) {
			defer wg.Done()

			// 调用 API 获取 SIM 卡信息
			res, err2 := cf.SimImei(&mobile_sdk.SimBasicInfoReq{
				Transid: cf.TransID,
				Token:   token,
				Msisdn:  line, // 使用当前行的内容作为 MSISDN
			})
			if err2 != nil {
				// 将失败的 Msisdn 发送到 errChan
				errChan <- line
				return
			}

			// 如果 API 返回结果不为空，将结果发送到 resultChan
			if len(res.Result) != 0 {
				resultChan <- fmt.Sprintf("%s,%s\n", line, res.Result[0].Imei)
			}
		}(line, lineNumber)

		lineNumber++
	}

	// 等待所有 goroutine 完成
	wg.Wait()

	// 关闭 channel，通知结果处理 goroutine 退出
	close(resultChan)
	close(errChan)

	// 检查是否有读取错误
	if err := scanner.Err(); err != nil {
		log.Fatalf("读取文件时出错: %v", err)
	}

	fmt.Println("处理完成，结果已写入 output2.txt，失败记录已写入 err.txt")
}
