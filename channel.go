package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

// 检查频道号是否返回对应的m3u文件,urlBegin前部分链接,urlEnd后部分链接,channelCodeBegin频道号起始,channelCodeEnd频道号终止,channelCodeInterval频道号间隔
func checkChannelCode(urlBegin string, urlEnd string, channelCodeBegin int, channelCodeEnd int, channelCodeInterval int) (validChannel []string) {
	for i := channelCodeBegin; i <= channelCodeEnd; i += channelCodeInterval {
		channelUrl := urlBegin + strconv.Itoa(i) + urlEnd

		fmt.Println("channel id:", i, "channel url:", channelUrl)

		cmd := exec.Command("curl", "-Ls", channelUrl)
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Println(err)
		}
		if strings.Contains(string(output), "#EXTM3U") {
			fmt.Println("Found")
			validChannel = append(validChannel, channelUrl)
		} else {
			fmt.Println("Not Found")
		}
	}
	return validChannel
}
