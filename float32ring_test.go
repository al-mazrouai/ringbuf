package ringbuf

import (
	"fmt"
	"testing"
)

func TestCreateRing(t *testing.T) {
	_, err := NewRingFloat32(5)
	if err != nil {
		t.Fatal("err: %v", err)
	}
}

func TestInsert(t *testing.T) {
	r, err := NewRingFloat32(5)
	if err != nil {
		t.Fatal("err: %v", err)
	}
	for i := 0; i < 7; i++ {
		r.Insert(float32(i))
	}
	fmt.Println(r)
}

func TestAvgMax(t *testing.T) {
	r, err := NewRingFloat32(10)
	if err != nil {
		t.Fatal("err: %v", err)
	}
	for i := 0; i < 101; i++ {
		r.Insert(float32(i % 7))
	}
	fmt.Println(r)
	fmt.Printf("avg: %.2f, max: %.2f \n", r.Avg(), r.Max())
}
