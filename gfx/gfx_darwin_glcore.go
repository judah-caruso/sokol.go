//go:build darwin && !SOKOL_METAL

package gfx

/*
#cgo LDFLAGS: -framework OpenGL

#define SOKOL_GFX_IMPL
#define SOKOL_GLCORE
#include "../sokol/sokol_gfx.h"
*/
import "C"
