
// Code generated by tools/gen_shims. DO NOT EDIT.

package osvfs

import "github.com/microsoft/typescript-go/internal/vfs"
import _ "github.com/microsoft/typescript-go/internal/vfs/osvfs"
import _ "unsafe"

//go:linkname FS github.com/microsoft/typescript-go/internal/vfs/osvfs.FS
func FS() vfs.FS
