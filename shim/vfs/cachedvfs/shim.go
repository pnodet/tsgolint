
// Code generated by tools/gen_shims. DO NOT EDIT.

package cachedvfs

import "github.com/microsoft/typescript-go/internal/vfs"
import "github.com/microsoft/typescript-go/internal/vfs/cachedvfs"
import _ "unsafe"

type FS = cachedvfs.FS
//go:linkname From github.com/microsoft/typescript-go/internal/vfs/cachedvfs.From
func From(fs vfs.FS) *cachedvfs.FS
