// +build linux

package raylib

// #cgo CFLAGS: -g -std=c99 -O0 -Wno-missing-braces -D_DEFAULT_SOURCE -fsanitize=undefined -I${SRCDIR}/../c/include
// #cgo LDFLAGS: -L${SRCDIR}/../c/lib -lubsan -lraylib -lraylib -lGL -lm -lpthread -ldl -lrt -fsanitize=address
import "C"
