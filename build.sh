#!/usr/bin/env bash
version=${1-"1.0.0-SNAPSHOT"}
out_file="dscServer"

build() {
  local os="$1"
  local arch="$2"
  local dir="build/dscServer-$os-$arch"
  out_file="${out_file}-${version}"
  go env -w GOPROXY=https://goproxy.cn,direct
  [ "$os" = "windows" ] && {
  		out_file="${out_file}.exe"
  	}
  rm -rf $dir
  mkdir -p $dir
  cp bootstrap*.yml $dir
  GOOS=$os GOARCH=$arch CGO_ENABLED=0 go build -o "${dir}/${out_file}" .

}

main() {
  echo "mod download"
  go get -t .
  go mod download
  build
}

main