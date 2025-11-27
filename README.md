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

[中文说明](https://github.com/unix755/qbtsh/blob/main/README_chs.md)

- Make `qBittorrent-nox` easier to install, uninstall and reload on `Linux` kernel system.
- The application is installed in `/usr/local/bin/qbittorrent-nox`, and the application data is stored
  in `/usr/local/etc/qBittorrent/` to avoid problems caused by future system changes.
- Statically compiled `qBittorrent-nox`
  using [userdocs/qbittorrent-nox-static](https://github.com/userdocs/qbittorrent-nox-static/releases)

## Usage

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

## Install

```sh
# For example, the host is a linux kernel based system with amd64 architecture
# Download the app
curl -Lo /usr/local/bin/qbtsh https://github.com/unix755/qbtsh/releases/latest/download/qbtsh-linux-amd64
# Give the app execute permission
chmod +x /usr/local/bin/qbtsh
# Show help
/usr/local/bin/qbtsh -h
```

## Compile

```sh
# Download application source code
git clone https://github.com/unix755/qbtsh.git
# Compile the source code
cd qbtsh
export CGO_ENABLED=0
go build -v -trimpath -ldflags "-s -w"
```

## FAQ

### Does this tool support FreeBSD?

- Not supported, [userdocs/qbittorrent-nox-static](https://github.com/userdocs/qbittorrent-nox-static/releases) only
  supports `linux` kernel systems, and builds a fully static `qBittorrent-nox` on `FreeBSD` is quite a difficult thing.

### Is there any way to get this working on FreeBSD?

- You can try compiling a dynamically linked `qBittorrent-nox` yourself using the `qBittorrent-nox` build script I wrote
  for `FreeBSD` [freebsd build](https://github.com/unix755/qbittorrent-nox), also I have already compiled binaries for
  the
  latest `FreeBSD`/`amd64`, you can also use them directly.

### How to access after running?

- After the service runs successfully, you can access it at `SERVER_IP:8080`. The default username is `admin`, and the
  default password can be found in service log. Although you can use it normally at this time, I still recommend you to
  change the
  password in time, use reverse proxy that provided by `apache` or `nginx` and enable `https`, then
  change the listening address of `qBittorrent-nox` from `0.0.0.0` to `127.0.0.1` to improve security.

## License

- **GPL-3.0 License**
- See `LICENSE` for details

## Credits

- [userdocs/qbittorrent-nox-static](https://github.com/userdocs/qbittorrent-nox-static/releases)
- [goland](https://www.jetbrains.com/go/)
- [vscode](https://code.visualstudio.com/)
