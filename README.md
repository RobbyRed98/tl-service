# tl-service
A port of the traffic-lights-service to golang.

## Packaging
Requires: `goreleaser` https://github.com/goreleaser/goreleaser
```bash
$ goreleaser release --skip-publish --snapshot --rm-dist
```

Creates the following packages:
* `deb`: arm6 arm7 arm8 i386 amd64
* `linux.tar.gz`: arm6 arm7 arm8 i386 amd64
* `windows.tar.gz`: arm6 arm7 arm8 i386 amd64
* `windows.exe`: arm6 arm7 arm8 i386 amd64

## Usage

### Linux

#### As app
```bash
$ tl-service
```

#### As service

##### Start
```bash
$ service tl-service start
```

##### Stop
```bash
$ service tl-service start
```

### Windows

```powershell
tl-service_<x.y.z>.exe
```
