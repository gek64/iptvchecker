package main

import (
	"fmt"
	"net/url"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

type validResult struct {
	id  int
	url string
	err error
}

// 执行并行检查URL任务
func checkAllURL(urlBegin string, urlEnd string, channelCodeBegin int, channelCodeEnd int, channelCodeInterval int, bufferSize int) (validResults []validResult) {
	// 建立缓存结果信道
	resultChan := make(chan validResult, bufferSize)
	// 流量控制信道
	limit := make(chan int, bufferSize)

	// 并行检查所有URL
	for i := channelCodeBegin; i <= channelCodeEnd; i += channelCodeInterval {
		go checkURLinParallel(urlBegin, urlEnd, i, resultChan, limit)
	}

	// 从结果信道收取所有结果
	for i := channelCodeBegin; i <= channelCodeEnd; i += channelCodeInterval {
		result := <-resultChan
		if result.url != "" {
			validResults = append(validResults, result)
		}
	}

	// 按任务id从小到大排序 https://pkg.go.dev/sort#Slice
	sort.Slice(validResults, func(i, j int) bool {
		return validResults[i].id < validResults[j].id
	})

	return validResults
}

// 并行检查URL
func checkURLinParallel(urlBegin string, urlEnd string, id int, resultChan chan validResult, limit chan int) {
	// 流量控制
	// 函数开始的时候占用一个信道缓冲,如无缓冲可用则等待
	limit <- 1
	// 函数结束时消除一个信道缓冲,等待的其他携程函数可运行
	defer func() {
		<-limit
	}()

	// 组合完整 URL
	channelUrl := urlBegin + strconv.Itoa(id) + urlEnd

	// 解析URL
	parsedURL, err := url.Parse(channelUrl)
	if err != nil {
		// 不是有效的URL,通过结果信道传递消息给主程序
		resultChan <- validResult{
			id:  id,
			url: "",
			err: fmt.Errorf("%s is not a valid url", channelUrl),
		}
		return
	}

	// 执行curl指令获取http访问返回数据
	cmd := exec.Command("curl", "-Ls", parsedURL.String())
	output, err := cmd.CombinedOutput()
	if err != nil {
		// curl执行出错,通过结果信道传递消息给主程序
		resultChan <- validResult{
			id:  id,
			url: "",
			err: err,
		}
		return
	}

	// curl获取到的返回数据包含m3u头
	if strings.Contains(string(output), "#EXTM3U") {
		// 找到了,通过结果信道传递消息给主程序
		resultChan <- validResult{
			id:  id,
			url: parsedURL.String(),
			err: nil,
		}
		return
	} else {
		// 没有找到,通过结果信道传递消息给主程序
		resultChan <- validResult{
			id:  id,
			url: "",
			err: fmt.Errorf("can not find a valid mu3 header"),
		}
	}
}

// 参考资料 https://chai2010.cn/advanced-go-programming-book/ch1-basic/ch1-06-goroutine.html
