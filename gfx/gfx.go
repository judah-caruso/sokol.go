package gfx

/*
#include "../sokol/sokol_gfx.h"

#cgo nocallback sg_setup
#cgo nocallback sg_shutdown
#cgo nocallback sg_isvalid
*/
import "C"

import "unsafe"

type Desc struct {
	_ uint32

	BufferPoolSize                            int
	ImagePoolSize                             int
	SamplerPoolSize                           int
	ShaderPoolSize                            int
	PipelinePoolSize                          int
	AttachmentsPoolSize                       int
	UniformBufferSize                         int
	MaxDispatchCallsPerPass                   int // max expected number of dispatch calls per pass (default: 1024)
	MaxCommitListeners                        int
	DisableValidation                         bool // disable validation layer even in debug mode, useful for tests
	D3d11ShaderDebugging                      bool // if true, HLSL shaders are compiled with D3DCOMPILE_DEBUG | D3DCOMPILE_SKIP_OPTIMIZATION
	MtlForceManagedStorageMode                bool // for debugging: use Metal managed storage mode for resources even with UMA
	MtlUseCommandBufferWithRetainedReferences bool // Metal: use a managed MTLCommandBuffer which ref-counts used resources
	WgpuDisableBindgroupsCache                bool // set to true to disable the WebGPU backend BindGroup cache
	WgpuBindgroupsCacheSize                   int  // number of slots in the WebGPU bindgroup cache (must be 2^N)
	Allocator                                 Allocator
	Logger                                    Logger // optional log function override
	Environment                               Environment

	_ uint32
}

type (
	Environment struct {
		Defaults EnvironmentDefaults
		Metal    MetalEnvironment
		D3d11    D3D11Environment
		Wgpu     WgpuEnviornment
	}

	EnvironmentDefaults struct {
		ColorFormat PixelFormat
		DepthFormat PixelFormat
		SampleCount int
	}
	MetalEnvironment struct {
		Device unsafe.Pointer
	}
	D3D11Environment struct {
		Device        unsafe.Pointer
		DeviceContext unsafe.Pointer
	}
	WgpuEnviornment struct {
		Device unsafe.Pointer
	}
)

type PixelFormat uint32

const (
	PixelFormatDefault PixelFormat = iota // value 0 reserved for default-init
	PixelFormatNone

	PixelFormatR8
	PixelFormatR8SN
	PixelFormatR8UI
	PixelFormatR8SI

	PixelFormatR16
	PixelFormatR16SN
	PixelFormatR16UI
	PixelFormatR16SI
	PixelFormatR16F
	PixelFormatRG8
	PixelFormatRG8SN
	PixelFormatRG8UI
	PixelFormatRG8SI

	PixelFormatR32UI
	PixelFormatR32SI
	PixelFormatR32F
	PixelFormatRG16
	PixelFormatRG16SN
	PixelFormatRG16UI
	PixelFormatRG16SI
	PixelFormatRG16F
	PixelFormatRGBA8
	PixelFormatSRGB8A8
	PixelFormatRGBA8SN
	PixelFormatRGBA8UI
	PixelFormatRGBA8SI
	PixelFormatBGRA8
	PixelFormatRGB10A2
	PixelFormatRG11B10F
	PixelFormatRGB9E5

	PixelFormatRG32UI
	PixelFormatRG32SI
	PixelFormatRG32F
	PixelFormatRGBA16
	PixelFormatRGBA16SN
	PixelFormatRGBA16UI
	PixelFormatRGBA16SI
	PixelFormatRGBA16F

	PixelFormatRGBA32UI
	PixelFormatRGBA32SI
	PixelFormatRGBA32F

	PixelFormatDepth
	PixelFormatDepthStencil

	PixelFormatBC1_RGBA
	PixelFormatBC2_RGBA
	PixelFormatBC3_RGBA
	PixelFormatBC3_SRGBA
	PixelFormatBC4_R
	PixelFormatBC4_RSN
	PixelFormatBC5_RG
	PixelFormatBC5_RGSN
	PixelFormatBC6H_RGBF
	PixelFormatBC6H_RGBUF
	PixelFormatBC7_RGBA
	PixelFormatBC7_SRGBA
	PixelFormatETC2_RGB8
	PixelFormatETC2_SRGB8
	PixelFormatETC2_RGB8A1
	PixelFormatETC2_RGBA8
	PixelFormatETC2_SRGB8A8
	PixelFormatEAC_R11
	PixelFormatEAC_R11SN
	PixelFormatEAC_RG11
	PixelFormatEAC_RG11SN

	PixelFormatASTC_4x4_RGBA
	PixelFormatASTC_4x4_SRGBA

	PixelFormatNum
)

type Logger struct {
	Func     func(*C.char, uint32, uint32, *C.char, uint32, *C.char, unsafe.Pointer)
	UserData unsafe.Pointer
}

type Allocator struct {
	AllocFunc func(C.size_t, unsafe.Pointer)
	FreeFunc  func(C.size_t, unsafe.Pointer)
	UserData  unsafe.Pointer
}

func Setup(desc *Desc) {
	C.sg_setup((*C.sg_desc)(unsafe.Pointer(desc)))
}

func Shutdown() {
	C.sg_shutdown()
}

func IsValid() bool {
	return bool(C.sg_isvalid())
}
