package endian

import (
    "encoding/binary"
    "unsafe"
)

var NativeEndian binary.ByteOrder

// In lack of binary.NativeEndian ...
func init() {
    var i uint32 = 0x01020304
    u := unsafe.Pointer(&i)
    pb := (*byte)(u)
    b := *pb
    if b == 0x04 {
        NativeEndian = binary.LittleEndian
    } else {
        NativeEndian = binary.BigEndian
    }
}

func To16(v uint16) (r uint16) {
    b := make([]byte, unsafe.Sizeof(v))
    binary.BigEndian.PutUint16(b, v)
    r = NativeEndian.Uint16(b)
    return
}

func To32(v uint32) (r uint32) {
    b := make([]byte, unsafe.Sizeof(v))
    binary.BigEndian.PutUint32(b, v)
    r = NativeEndian.Uint32(b)
    return
}

func To64(v uint64) (r uint64) {
    b := make([]byte, unsafe.Sizeof(v))
    binary.BigEndian.PutUint64(b, v)
    r = NativeEndian.Uint64(b)
    return
}
