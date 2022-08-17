package main

import (
	"log"
	"os"
	"strconv"
)

// 提取validResults中的URL拼接为channels
func validResultsToChannels(validResults []validResult) (channels []string) {
	for _, result := range validResults {
		channels = append(channels, result.url)
	}
	return channels
}

// 制作m3u文件,channels频道切片,file输出文件
func makeM3u(channels []string, file string) (err error) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	_, err = f.WriteString("#EXTM3U" + "\n")
	if err != nil {
		return err
	}

	for i, channel := range channels {
		_, err = f.WriteString("#EXTINF:-1,Channel" + strconv.Itoa(i+1) + "\n")
		if err != nil {
			return err
		}
		_, err = f.WriteString(channel + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
