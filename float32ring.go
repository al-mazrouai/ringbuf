package ringbuf

import (
	"fmt"
)

type RingFloat32 struct {
	buff []float32
	tail int
	max  float32
}

func NewRingFloat32(size int) (*RingFloat32, error) {
	if size <= 0 {
		return nil, fmt.Errorf("Size must be positive")
	}

	r := &RingFloat32{
		buff: make([]float32, 0, size),
		tail: 0,
		max:  0,
	}

	return r, nil
}

func (r *RingFloat32) Insert(n float32) {

	if len(r.buff) < cap(r.buff) {
		r.buff = append(r.buff, n)

	} else {
		r.buff[r.tail] = n
		r.tail = (r.tail + 1) % cap(r.buff)
	}

	if n > r.max {
		r.max = n
	}
}

func (r *RingFloat32) Avg() float32 {
	var total float32

	for _, v := range r.buff {
		total += v
	}

	return total / float32(len(r.buff))
}

func (r *RingFloat32) Max() float32 {
	return r.max
}
