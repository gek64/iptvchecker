# iptvchecker
- A tool to help you get iptv urls and then convert all found urls to m3u files

## Usage
```
Usage:
  iptvchecker {Parameters} [Commands]

Parameters:
  -ub       : URL begin
  -ue       : URL end
  -cb       : Channel Code Begin
  -ce       : Channel Code End
  -i        : Channel Code Interval
  -o        : Output File
  -x        : Number of parallel processing, default 3

Commands:
  -h        : show help
  -v        : show version

Example:
  1) iptvchecker -ub "http://127.0.0.1/PLTV/88888888/999/" -ue "/index.m3u8" -cb 0 -ce 1000 -i 1 -o "iptv.m3u"
  2) iptvchecker -h
  3) iptvchecker -v

```

## Install
```sh
# system is linux(debian,redhat linux,ubuntu,fedora...) and arch is amd64
curl -Lo /usr/local/bin/iptvchecker https://github.com/gek64/iptvchecker/releases/latest/download/iptvchecker-linux-amd64
chmod +x /usr/local/bin/iptvchecker

# system is freebsd and arch is amd64
curl -Lo /usr/local/bin/iptvchecker https://github.com/gek64/iptvchecker/releases/latest/download/iptvchecker-freebsd-amd64
chmod +x /usr/local/bin/iptvchecker
```

## Compile
### How to compile if prebuilt binaries are not found
```sh
git clone https://github.com/gek64/iptvchecker.git
cd iptvchecker
go build -v -trimpath -ldflags "-s -w"
```

## QA
### Q: Windows Security detect `.exe` as `Trojan:Win32/Wacatac.B!ml`
A: This application does not contain any malware, backdoors, and advertisements, all released files are build by github actions. For more information, see https://go.dev/doc/faq#virus

## License
- **GPL-3.0 License**
- See `LICENSE` for details
