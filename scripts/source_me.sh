# shellcheck shell=bash
# go stuff
GOPATH="$(path_add "${GOPATH}" "${HOME}/install/go")"
export GOPATH
PATH="$(path_add "${PATH}" "${HOME}/install/go/bin")"
export PATH
