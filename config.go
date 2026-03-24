package main

import (
	"embed"
	"errors"
	"net/url"
	"os"
	"os/exec"
	"runtime"
)

//go:embed service/*
var container embed.FS

func GetDownloadURL(tagName string) (downloadURL string, err error) {
	var baseURL = ""
	var assetName = ""

	if tagName != "" {
		baseURL = "https://github.com/userdocs/qbittorrent-nox-static/releases/download/" + tagName + "/"
	} else {
		baseURL = "https://github.com/userdocs/qbittorrent-nox-static/releases/latest/download/"
	}

	switch runtime.GOARCH {
	case "amd64":
		assetName = "x86_64-qbittorrent-nox"
	case "386":
		assetName = "x86-qbittorrent-nox"
	case "arm":
		assetName = "armv7-qbittorrent-nox"
	case "arm64":
		assetName = "aarch64-qbittorrent-nox"
	}
	return url.JoinPath(baseURL, assetName)
}

func GetService() (initSystem string, serviceContent []byte, err error) {
	serviceFile := ""

	// 通过查找 /proc/1/comm 文件
	b, err := os.ReadFile("/proc/1/comm")
	if err == nil {
		switch string(b) {
		case "systemd":
			initSystem = "systemd"
			serviceFile = "service/qbittorrent.service"
		case "procd":
			initSystem = "procd"
			serviceFile = "service/qbittorrent.procd"
		}
	}

	// 通过查找有特异性的二进制文件
	_, err = exec.LookPath("openrc")
	if err == nil {
		initSystem = "openrc"
		serviceFile = "service/qbittorrent.openrc"
	}

	// 通过查找系统变量
	if runtime.GOOS == "freebsd" {
		initSystem = "rc.d"
		serviceFile = "service/qbittorrent.rcd"
	}

	// 找不到初始化系统返回错误
	if initSystem == "" {
		return "", nil, errors.New("init system not found")
	}

	// 读取文件并返回文件内容
	bytes, err := container.ReadFile(serviceFile)
	if err != nil {
		return initSystem, nil, err
	}
	return initSystem, bytes, nil
}
