// Package time provides functions for simple cross-platform time measurement
package time

/*
#define SOKOL_TIME_IMPL
#include "../sokol/sokol_time.h"

#cgo nocallback stm_setup
#cgo nocallback stm_now
#cgo nocallback stm_diff
#cgo nocallback stm_since
#cgo nocallback stm_laptime
#cgo nocallback stm_round_to_common_refresh_rate
#cgo nocallback stm_sec
#cgo nocallback stm_ms
#cgo nocallback stm_us
#cgo nocallback stm_ns

#cgo noescape stm_setup
#cgo noescape stm_now
#cgo noescape stm_diff
#cgo noescape stm_since
#cgo noescape stm_laptime
#cgo noescape stm_round_to_common_refresh_rate
#cgo noescape stm_sec
#cgo noescape stm_ms
#cgo noescape stm_us
#cgo noescape stm_ns
*/
import "C"

// Call once before any other functions to initialize sokol_time
// (this calls for instance QueryPerformanceFrequency on Windows)
func Setup() {
	C.stm_setup()
}

// Get current point in time in unspecified 'ticks'. The value that
// is returned has no relation to the 'wall-clock' time and is
// not in a specific time unit, it is only useful to compute
// time differences.
func Now() uint64 {
	return uint64(C.stm_now())
}

// Computes the time difference between new and old. This will always
// return a positive, non-zero value.
func Diff(new, old uint64) uint64 {
	return uint64(C.stm_diff(C.uint64_t(new), C.uint64_t(old)))
}

// Takes the current time, and returns the elapsed time since start
// (this is a shortcut for "stm_diff(stm_now(), start)")
func Since(start uint64) uint64 {
	return uint64(C.stm_since(C.uint64_t(start)))
}

// This is useful for measuring frame time and other recurring
// events. It takes the current time, returns the time difference
// to the value in last_time, and stores the current time in
// last_time for the next call. If the value in last_time is 0,
// the return value will be zero (this usually happens on the
// very first call).
func Laptime(last *uint64) uint64 {
	return uint64(C.stm_laptime((*C.uint64_t)(last)))
}

// This oddly named function takes a measured frame time and
// returns the closest "nearby" common display refresh rate frame duration
// in ticks. If the input duration isn't close to any common display
// refresh rate, the input duration will be returned unchanged as a fallback.
// The main purpose of this function is to remove jitter/inaccuracies from
// measured frame times, and instead use the display refresh rate as
// frame duration.
// NOTE: for more robust frame timing, consider using the
// sokol_app.h function sapp_frame_duration()
func RoundToCommonRefreshRate(duration uint64) uint64 {
	return uint64(C.stm_round_to_common_refresh_rate(C.uint64_t(duration)))
}

// Converts a tick value into seconds.
func Sec(ticks uint64) float64 {
	return float64(C.stm_sec(C.uint64_t(ticks)))
}

// Converts a tick value into milliseconds.
func Ms(ticks uint64) float64 {
	return float64(C.stm_ms(C.uint64_t(ticks)))
}

// Converts a tick value into microseconds.
// Note that not all platforms will have microsecond precision.
func Us(ticks uint64) float64 {
	return float64(C.stm_us(C.uint64_t(ticks)))
}

// Converts a tick value into nanoseconds.
// Note that not all platforms will have nanosecond precision.
func Ns(ticks uint64) float64 {
	return float64(C.stm_ns(C.uint64_t(ticks)))
}
