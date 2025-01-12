# lxd_exporter

[![Release CI](https://github.com/eumel8/lxd_exporter/actions/workflows/build.yaml/badge.svg)](https://github.com/eumel8/lxd_exporter/actions/workflows/build.yaml)

[![Go Report Card](https://goreportcard.com/badge/github.com/eumel8/lxd_exporter)](https://goreportcard.com/report/github.com/eumel8/lxd_exporter)

LXD metrics exporter for Prometheus. Improved version of [viveksing/lxdexporter_golang](https://github.com/viveksing/lxdexporter_golang) and
[nieltg/lxd_exporter](https://github.com/nieltg/lxd_exporter) and


## Usage

Download latest precompiled binary of this exporter from [the release page](https://github.com/eumel8/lxd_exporter/releases).

Extract archive, then run the exporter:
```
./lxd_exporter
```

The exporter must have access to LXD socket which can be guided by:
- Specifying `LXD_SOCKET` environment variable to LXD socket path, or
- Specifying `LXD_DIR` environment variable to LXD socket's parent directory.

Incus Support (LXD 4):

In Incus only secured connections are allowed by default. This is disabled here while connection mostly local. To prevent exposed data set the listen address to localhost, the lxd bridge ip address, or what fits your requirements:

```
./lxd_exporter --help
usage: lxd_exporter [<flags>]


Flags:
  --[no-]help         Show context-sensitive help (also try --help-long and --help-man).
  --port=9472         Provide the port to listen on
  --listen="0.0.0.0"  Provide the interface to listen on
  --socket="/var/snap/lxd/common/lxd/unix.socket"
                      Provide the socket to listen on
  --[no-]version      Show application version.
  ```

For more information, you can see the documentation from [Go LXD client library](https://godoc.org/github.com/lxc/incus/client#ConnectLXDUnix).

## Hacking

Install [Go](https://golang.org/dl) before hacking this library.

To run all tests:
```
go test ./...
```

To build exporter binary:
```
mkdir build
go build -o build ./...
```

Binary will be available on `build/` directory.

## License

[MIT](LICENSE).
