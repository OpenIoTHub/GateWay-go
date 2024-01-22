# gateway-go
[![Build Status](https://travis-ci.com/OpenIoTHub/gateway-go.svg?branch=master)](https://travis-ci.com/OpenIoTHub/gateway-go)

[![Get it from the Snap Store](https://snapcraft.io/static/images/badges/en/snap-store-white.svg)](https://snapcraft.io/gateway-go)

You can install the pre-compiled binary (in several different ways),
use Docker.

Here are the steps for each of them:

## Install the pre-compiled binary

**openwrt/entware/optware (Usually on the router)**:
#### use snapshot branch：https://downloads.openwrt.org/snapshots/
```sh
opkg update
opkg install gateway-go
```

**homebrew tap** :

```sh
$ brew install OpenIoTHub/tap/gateway-go
```

**homebrew** (may not be the latest version):

```sh
$ brew install gateway-go
```
homebrew pr [gateway-go](https://github.com/Homebrew/homebrew-core/blob/master/Formula/gateway-go.rb)
```text
*** config file : 
/usr/local/etc/gateway-go/gateway-go.yaml
```


**snapcraft**:

```sh
$ sudo snap install gateway-go
```
```text
*** config file :
 /root/snap/gateway-go/current/gateway-go.yaml
```


**scoop**:

```sh
$ scoop bucket add OpenIoTHub https://github.com/OpenIoTHub/scoop-bucket.git
$ scoop install gateway-go
```

**deb/rpm**:

Download the `.deb` or `.rpm` from the [releases page][releases] and
install with `dpkg -i` and `rpm -i` respectively.
```text
*** config file :
 /etc/gateway-go/gateway-go.yaml
```

**manually**:

Download the pre-compiled binaries from the [releases page][releases] and
copy to the desired location.

## Running with Docker

You can also use it within a Docker container. To do that, you'll need to
execute something more-or-less like the following:

```sh
$ docker run -it --net=host openiothub/gateway-go:latest -t <your Token>
```

Note that the image will almost always have the last stable Go version.

[releases]: https://github.com/OpenIoTHub/gateway-go/releases
## Build dynamic/static Library
```shell
#build and push mobile lib
#install gomobile(at system cli)
go install golang.org/x/mobile/cmd/gobind@latest
go install golang.org/x/mobile/cmd/gomobile@latest
gomobile init
gomobile version
go get -u golang.org/x/mobile/...
#
export GO111MODULE="off"
gomobile bind -target=android  -o gateway.aar
gomobile bind -ldflags '-w -s -extldflags "-lresolve"' --target=ios,macos,iossimulator -o OpenIoTHubGateway.xcframework ./client
#
#https://gitee.com/OpenIoThub/mobile-lib-podspec
#git tag -a 0.0.1 -m '0.0.1'
#git pus --tags
#pod trunk push ./OpenIoTHubGateway.podspec --skip-import-validation --allow-warnings
#
mvn gpg:sign-and-deploy-file -DrepositoryId=ossrh -Dfile=gateway.aar -DpomFile=gateway.pom -Durl=https://s01.oss.sonatype.org/service/local/staging/deploy/maven2/
mvn deploy:deploy-file -Dfile=client.aar -DgroupId=cloud.iothub -DartifactId=gateway -Dversion=0.0.1 -Dpackaging=aar -DrepositoryId=github -Durl=https://maven.pkg.github.com/OpenIoTHub/gateway-go
```
```shell
#for build windows dll
echo "building windows dll"
#brew install mingw-w64
#sudo apt-get install binutils-mingw-w64
# shellcheck disable=SC2034
export CGO_ENABLED=1
export CC=x86_64-w64-mingw32-gcc
export CXX=x86_64-w64-mingw32-g++
export GOOS=windows GOARCH=amd64
go build -tags windows -ldflags=-w -trimpath -o ./build/windows/gateway_amd64.dll -buildmode=c-shared main.go
```
```shell
#for build linux/android so file
echo "building linux/android so file"
#linux和Android共用动态链接库
export CGO_ENABLED=1
export GOARCH=amd64
export GOOS=linux
go build -tags linux -ldflags=-w -trimpath -o build/linux/libgateway_amd64.so -buildmode c-shared main.go
# shellcheck disable=SC2034
export CGO_ENABLED=1
export GOARCH=arm64
export GOOS=linux
#sudo apt install gcc-aarch64-linux-gnu
export CC=aarch64-linux-gnu-gcc
##sudo apt install g++-aarch64-linux-gnu
#export CXX=aarch64-linux-gnu-g++
##sudo apt-get install binutils-aarch64-linux-gnu
#export AR=aarch64-linux-gnu-ar
go build -tags linux -ldflags=-w -trimpath -o build/linux/libgateway_arm64.so -buildmode c-shared main.go
```