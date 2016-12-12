package terminal

import (
	"fmt"
	"math/rand"
	"time"
)

type output struct {
	d time.Duration
}

func (o *output) Write(b []byte) (n int, err error) {
	for _, c := range b {
		fmt.Print(string(c))
		time.Sleep(o.d)
		n++
	}
	fmt.Println()
	return
}

func newOutout() *output {
	l := []int{0, 300, 600, 1200, 1600}
	n := l[rand.Intn(len(l))]
	return &output{
		d: time.Duration(int(time.Microsecond) * n),
	}
}
