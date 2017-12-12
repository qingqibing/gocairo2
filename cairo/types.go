// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Tue, 12 Dec 2017 17:29:45 MSK.
// By https://git.io/c-for-go. DO NOT EDIT.

package cairo

/*
#cgo pkg-config: cairo
#include "cairo.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// Bool type as declared in cairo/cairo.h:107
type Bool int32

// Cairo as declared in cairo/cairo.h:124
type Cairo C.cairo_t

// Surface as declared in cairo/cairo.h:153
type Surface C.cairo_surface_t

// Device as declared in cairo/cairo.h:171
type Device C.cairo_device_t

// Matrix as declared in cairo/cairo.h:196
type Matrix struct {
	Xx             float64
	Yx             float64
	Xy             float64
	Yy             float64
	X0             float64
	Y0             float64
	ref6fd9c783    *C.cairo_matrix_t
	allocs6fd9c783 interface{}
}

// Pattern as declared in cairo/cairo.h:220
type Pattern C.cairo_pattern_t

// DestroyFunc type as declared in cairo/cairo.h:232
type DestroyFunc func(data unsafe.Pointer)

// UserDataKey as declared in cairo/cairo.h:248
type UserDataKey struct {
	Unused         int32
	refb2afcf30    *C.cairo_user_data_key_t
	allocsb2afcf30 interface{}
}

// WriteFunc type as declared in cairo/cairo.h:445
type WriteFunc func(closure unsafe.Pointer, data string, length uint32) Status

// ReadFunc type as declared in cairo/cairo.h:467
type ReadFunc func(closure unsafe.Pointer, data []byte, length uint32) Status

// RectangleInt as declared in cairo/cairo.h:486
type RectangleInt struct {
	X              int32
	Y              int32
	Width          int32
	Height         int32
	ref182a0d2b    *C.cairo_rectangle_int_t
	allocs182a0d2b interface{}
}

// Rectangle as declared in cairo/cairo.h:1004
type Rectangle struct {
	X              float64
	Y              float64
	Width          float64
	Height         float64
	ref9b663fb6    *C.cairo_rectangle_t
	allocs9b663fb6 interface{}
}

// RectangleList as declared in cairo/cairo.h:1021
type RectangleList struct {
	Status         Status
	Rectangles     []Rectangle
	NumRectangles  int32
	ref149ca897    *C.cairo_rectangle_list_t
	allocs149ca897 interface{}
}

// ScaledFont as declared in cairo/cairo.h:1059
type ScaledFont C.cairo_scaled_font_t

// FontFace as declared in cairo/cairo.h:1080
type FontFace C.cairo_font_face_t

// Glyph as declared in cairo/cairo.h:1112
type Glyph struct {
	Index          uint
	X              float64
	Y              float64
	ref5b16733c    *C.cairo_glyph_t
	allocs5b16733c interface{}
}

// TextCluster as declared in cairo/cairo.h:1143
type TextCluster struct {
	NumBytes       int32
	NumGlyphs      int32
	refd989a040    *C.cairo_text_cluster_t
	allocsd989a040 interface{}
}

// TextExtents as declared in cairo/cairo.h:1200
type TextExtents struct {
	XBearing       float64
	YBearing       float64
	Width          float64
	Height         float64
	XAdvance       float64
	YAdvance       float64
	refe2add9c7    *C.cairo_text_extents_t
	allocse2add9c7 interface{}
}

// FontExtents as declared in cairo/cairo.h:1251
type FontExtents struct {
	Ascent         float64
	Descent        float64
	Height         float64
	MaxXAdvance    float64
	MaxYAdvance    float64
	refbc0b0f3c    *C.cairo_font_extents_t
	allocsbc0b0f3c interface{}
}

// FontOptions as declared in cairo/cairo.h:1385
type FontOptions C.cairo_font_options_t

// UserScaledFontInitFunc type as declared in cairo/cairo.h:1722
type UserScaledFontInitFunc func(scaledFont []ScaledFont, cr []Cairo, extents []FontExtents) Status

// UserScaledFontRenderGlyphFunc type as declared in cairo/cairo.h:1769
type UserScaledFontRenderGlyphFunc func(scaledFont []ScaledFont, glyph uint, cr []Cairo, extents []TextExtents) Status

// UserScaledFontTextToGlyphsFunc type as declared in cairo/cairo.h:1839
type UserScaledFontTextToGlyphsFunc func(scaledFont []ScaledFont, utf8 string, utf8Len int32, glyphs [][]Glyph, numGlyphs []int32, clusters [][]TextCluster, numClusters []int32, clusterFlags []TextClusterFlags) Status

// UserScaledFontUnicodeToGlyphFunc type as declared in cairo/cairo.h:1886
type UserScaledFontUnicodeToGlyphFunc func(scaledFont []ScaledFont, unicode uint, glyphIndex []uint) Status

// PathData as declared in cairo/cairo.h:2061
const sizeofPathData = unsafe.Sizeof(C.cairo_path_data_t{})

type PathData [sizeofPathData]byte

// Path as declared in cairo/cairo.h:2097
type Path struct {
	Status         Status
	Data           []PathData
	NumData        int32
	reffd10d41e    *C.cairo_path_t
	allocsfd10d41e interface{}
}

// SurfaceObserverCallback type as declared in cairo/cairo.h:2251
type SurfaceObserverCallback func(observer []Surface, target []Surface, data unsafe.Pointer)

// RasterSourceAcquireFunc type as declared in cairo/cairo.h:2622
type RasterSourceAcquireFunc func(pattern []Pattern, callbackData unsafe.Pointer, target []Surface, extents []RectangleInt) *Surface

// RasterSourceReleaseFunc type as declared in cairo/cairo.h:2642
type RasterSourceReleaseFunc func(pattern []Pattern, callbackData unsafe.Pointer, surface []Surface)

// RasterSourceSnapshotFunc type as declared in cairo/cairo.h:2663
type RasterSourceSnapshotFunc func(pattern []Pattern, callbackData unsafe.Pointer) Status

// RasterSourceCopyFunc type as declared in cairo/cairo.h:2681
type RasterSourceCopyFunc func(pattern []Pattern, callbackData unsafe.Pointer, other []Pattern) Status

// RasterSourceFinishFunc type as declared in cairo/cairo.h:2696
type RasterSourceFinishFunc func(pattern []Pattern, callbackData unsafe.Pointer)

// Region as declared in cairo/cairo.h:3072
type Region C.cairo_region_t
