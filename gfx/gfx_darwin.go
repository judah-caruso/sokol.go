//go:build darwin

package gfx

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Metal -framework Foundation

#define SOKOL_GFX_IMPL
#define SOKOL_METAL
#include "../sokol/sokol_gfx.h"
*/
import "C"
