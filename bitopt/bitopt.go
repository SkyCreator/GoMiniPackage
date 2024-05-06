//bitopt.go add by skycreator 2024/5/6
//位操作，最多64位
package bitopt

type BitData struct {
	value uint64
}
func (bd* BitData)SetBit(index int, value byte) {
	if value == 1 {
		bd.value |= 1 << index
	} else {
		bd.value &= ^(1 << index)
	}
}

func (bd *BitData)GetBit(index int) byte {
	return byte((bd.value >> index) & 1)
}

func (bd *BitData)GetValue() uint64 {
	return bd.value
}

func (bd *BitData)BitCount() int {
	var count int
	for i := 0; i < 64; i++ {
		if bd.GetBit(i) == 1 {
			count++
		}
	}
	return count
}