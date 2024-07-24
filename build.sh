#!/bin/sh
out_file="dscServer"
[ $# -lt 3 ] && {
	echo "Usage: $0 1.0.0-SNAPSHOT linux amd64"
	exit 1
}
build() {
  local version="$1"
  local os="$2"
  local arch="$3"
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
  build $1 $2 $3
}

main $1 $2 $3