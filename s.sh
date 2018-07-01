#!/bin/sh
s() {
	cd "$(scd "$@")" || return 1
}
