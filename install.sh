#!/usr/bin/env bash
set -euf

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ]; do # resolve $SOURCE until the file is no longer a symlink
  DIR="$(cd -P "$(dirname "$SOURCE")" && pwd)"
  SOURCE="$(readlink "$SOURCE")"
  [[ $SOURCE != /* ]] && SOURCE="$DIR/$SOURCE" # if $SOURCE was a relative symlink, we need to resolve it relative to the path where the symlink file was located
done
DIR="$(cd -P "$(dirname "$SOURCE")" && pwd)"

go get -v github.com/Nekroze/scd/...
go install github.com/Nekroze/scd

if ! grep -q "^source .*/s\\.sh$" ~/.bashrc; then
	echo "source $DIR/s.sh" >> ~/.bashrc
fi
