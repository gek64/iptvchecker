package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	cliURLBegin            string
	cliURLEnd              string
	cliChannelCodeBegin    int
	cliChannelCodeEnd      int
	cliChannelCodeInterval int
	cliOutputFile          string
	cliHelp                bool
	cliVersion             bool
)

func init() {
	flag.StringVar(&cliURLBegin, "ub", "", "-ub http://127.0.0.1/PLTV/88888888/224/")
	flag.StringVar(&cliURLEnd, "ue", "", "-ue /index.m3u8")
	flag.IntVar(&cliChannelCodeBegin, "cb", 0, "-cb 0")
	flag.IntVar(&cliChannelCodeEnd, "ce", 1000, "-ce 1000")
	flag.IntVar(&cliChannelCodeInterval, "i", 1, "-i 1")
	flag.StringVar(&cliOutputFile, "o", "iptv.m3u", "-o iptv.m3u")
	flag.BoolVar(&cliHelp, "h", false, "show help")
	flag.BoolVar(&cliVersion, "v", false, "show version")
	flag.Parse()

	// 重写显示用法函数
	flag.Usage = func() {
		var helpInfo = `Usage:
  iptvchecker {Parameters} [Commands]

Parameters:
  -ub       : URL begin
  -ue       : URL end
  -cb       : Channel Code Begin
  -ce       : Channel Code End
  -i        : Channel Code Interval
  -o        : Output File

Commands:
  -h        : show help
  -v        : show version

Example:
  1) iptvchecker -ub "http://127.0.0.1/PLTV/88888888/999/" -ue "/index.m3u8" -cb 0 -ce 1000 -i 1 -o "iptv.m3u"
  2) iptvchecker -h
  3) iptvchecker -v`

		fmt.Println(helpInfo)
	}

	// 打印帮助信息
	if cliHelp || len(os.Args) == 1 {
		flag.Usage()
		os.Exit(0)
	}

	// 打印版本信息
	if cliVersion {
		fmt.Println("v1.01")
		os.Exit(0)
	}
}

func main() {
	if cliURLBegin != "" && cliURLEnd != "" {
		// 单线程
		//validChannel := checkChannelCode(cliURLBegin, cliURLEnd, cliChannelCodeBegin, cliChannelCodeEnd, cliChannelCodeInterval)
		//err := makeM3u(validChannel, cliOutputFile)
		//if err != nil {
		//	log.Panicln(err)
		//}

		//	多线程
		validResults := checkAllURL(cliURLBegin, cliURLEnd, cliChannelCodeBegin, cliChannelCodeEnd, cliChannelCodeInterval, 3)
		err := makeM3u(validResultsToChannels(validResults), cliOutputFile)
		if err != nil {
			log.Panicln(err)
		}

	}
}
