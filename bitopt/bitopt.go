// bitopt.go add by skycreator 2024/5/6
// 位操作，最多64位
package bitopt

import "fmt"

type BitData struct {
	value uint64
}

func (bd *BitData) SetBit(index int, value byte) error {
	if index < 0 || index > 63 {
		return fmt.Errorf("bitopt:SetBit index out of range.index is 0~63")
	}
	if value == 1 {
		bd.value |= 1 << index
	} else {
		bd.value &= ^(1 << index)
	}
	return nil
}

func (bd *BitData) GetBit(index int) (byte, error) {
	if index < 0 || index > 63 {
		return 0, fmt.Errorf("bitopt:GetBit index out of range.index is 0~63")
	}
	return byte((bd.value >> index) & 1), nil
}
func (bd *BitData) SetValue(value uint64) {
	bd.value = value
}

func (bd *BitData) GetValue() uint64 {
	return bd.value
}

func (bd *BitData) BitCount() (int, error) {
	var count int
	var value byte
	var err error
	for i := 0; i < 64; i++ {
		value, err = bd.GetBit(i)
		if err != nil {
			return 0, fmt.Errorf("bitopt:BitCount error is %v", err)
		}
		if value == 1 {
			count++
		}
	}
	return count, nil
}
func (bd *BitData) BitCountByRange(startidx int, endidx int) (int, error) {
	if startidx < 0 || startidx > 63 || endidx < 0 || endidx > 63 || startidx > endidx {
		return 0, fmt.Errorf("bitopt:BitCountByRange index out of range.startidx and endidx is 0~63,and startidx <= endidx")
	}
	var count int
	var value byte
	var err error
	for i := startidx; i <= endidx; i++ {
		value, err = bd.GetBit(i)
		if err != nil {
			return 0, fmt.Errorf("bitopt:BitCountByRange %v", err)
		}
		if value == 1 {
			count++
		}
	}
	return count, nil
}
