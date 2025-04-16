//go:build darwin

package app

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Metal -framework Foundation

#define SOKOL_APP_IMPL
#define SOKOL_METAL
#include "../sokol/sokol_app.h"
*/
import "C"
