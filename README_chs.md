```
          ▄▄                            ▄▄
          ██         ██                 ███
          ██         ██                 ██
 ▄██▀███  ██▄████▄ ██████   ▄██▀███   ███████▄
▄█▀   ██  ██    ▀██  ██     ██        ██   ██
██    ██  ██     ██  ██     ▀██████▄  ██   ██
██▄   ██  ██▄   ▄██  ██     █▄   ███  ██   ██
 ▀██████  █▀█████▀   ▀████  ██████▀   ██  ████▄
      ██
   ▄████▄
```

- 使 `qBittorrent-nox` 在 `Linux` 内核系统上更容易安装、卸载及重载。
- 应用安装在 `/usr/local/bin/qbittorrent-nox`，应用数据存储在 `/usr/local/etc/qBittorrent/`，避免将来系统变动带来的问题。
- 静态编译的 `qBittorrent-nox`
  使用 [userdocs/qbittorrent-nox-static](https://github.com/userdocs/qbittorrent-nox-static/releases)。

## 使用说明

```
NAME:
   qbtsh.exe - qBittorrent quick install tool

USAGE:
   qbtsh.exe [global options] command [command options] 

VERSION:
   v2.00

COMMANDS:
   install, i  Install qBittorrent
   uninstall   Remove config,cache and uninstall qBittorrent
   update      Update qBittorrent
   reload      Reload service
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## 安装说明

```sh
# 例如宿主机是 amd64 架构的基于 linux 内核的系统
# 下载应用
curl -Lo /usr/local/bin/qbtsh https://github.com/unix755/qbtsh/releases/latest/download/qbtsh-linux-amd64
# 给应用执行权限
chmod +x /usr/local/bin/qbtsh
# 查看帮助
/usr/local/bin/qbtsh -h
```

## 编译说明

```sh
# 下载应用源代码
git clone https://github.com/unix755/qbtsh.git
# 编译源代码
cd qbtsh
export CGO_ENABLED=0
go build -v -trimpath -ldflags "-s -w"
```

## 常见问答

### 本工具是否支持 FreeBSD?

- 不支持，[userdocs/qbittorrent-nox-static](https://github.com/userdocs/qbittorrent-nox-static/releases) 仅支持 `linux`
  内核的系统, 并且在 `FreeBSD` 上编译一个完全静态的 `qBittorrent-nox` 是一件相当困难的事情。

### 有什么办法能在FreeBSD上使用吗?

- 你可以使用我为FreeBSD系统编写的 `qBittorrent-nox`
  编译脚本来尝试自行编译动态链接的 `qBittorrent-nox` [freebsd build](https://github.com/unix755/qbittorrent-nox)
  ，我已经为最新的FreeBSD系统编译好了 `amd64` 架构的二进制文件，你也可以直接使用。

### 运行以后该如何访问?

- 服务成功运行后访问 `SERVER_IP:8080` 即可，默认用户名是 `admin`
  ，默认密码可以在服务日志中找到，虽然此时你已经可以正常使用，但是我还是建议你及时更改密码，使用 `apache` 或者 `nginx`
  提供的反向代理和开启 `https`，并将 `qBittorrent-nox` 的监听地址从 `0.0.0.0` 更变为 `127.0.0.1` 以提高安全性。

## 许可证

- **GPL-3.0 License**
- 查看 `LICENSE` 获取详细内容

## 致谢

- [userdocs/qbittorrent-nox-static](https://github.com/userdocs/qbittorrent-nox-static/releases)
- [goland](https://www.jetbrains.com/go/)
- [vscode](https://code.visualstudio.com/)
