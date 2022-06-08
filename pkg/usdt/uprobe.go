package usdt

import (
	"bytes"
	"debug/elf"
	"encoding/binary"
	"unsafe"
)

type Elf_Nhdr struct {
	SzName uint32
	SzDesc uint32
	Type uint32
}

type SdtDesc struct {
	Addr uint64
	Base uint64
	Sem uint64
	//Provider string
	//ProbeName string
	//Args string
}

func Bytes2String(b []byte) string {
	return string(b[:bytes.IndexByte(b, 0x0)])
}

func Align(l, a int) int {
	return ((l + a - 1)/a)*a
}

func Uprobes(path string) (map[string]uint64, error) {

	var nhdr Elf_Nhdr
	var desc SdtDesc
	var load uint64

	um := make(map[string]uint64)

	f, err := elf.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	for _, prog := range f.Progs {
		if prog.Type == elf.PT_LOAD {
			load = prog.Vaddr
			break
		}
	}

	for _, sec := range f.Sections {

		if sec.Type != elf.SHT_NOTE {
			continue
		}

		data, err := sec.Data()
		if err != nil {
			return nil, err
		}

		err = binary.Read(bytes.NewReader(data), NativeEndian, &nhdr)
		if err != nil {
			return nil, err
		}
		data = data[unsafe.Sizeof(nhdr):]

		if nhdr.Type != 3 {
			continue
		}

		vendor := Bytes2String(data)
		data = data[Align(int(nhdr.SzName), 4):]
		if vendor != "stapsdt" {
			continue
		}

		err = binary.Read(bytes.NewReader(data), NativeEndian, &desc)
		if err != nil {
			return nil, err
		}
		data = data[unsafe.Sizeof(desc):]

		prov := Bytes2String(data)
		data = data[len(prov)+1:]

		prob := Bytes2String(data)
		if len(prob) > 0 {
			um[prob] = desc.Addr - load
		}
	}

	return um, nil
}
