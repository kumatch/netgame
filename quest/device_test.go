package quest

import (
	"testing"
)

func TestPickupInterfaceNumber(t *testing.T) {
	{
		for i := 0; i < 10000; i++ {
			if1, if2 := pickupInterfaceNumbers(2)
			switch {
			case if1 == 0, if2 == 1:
				continue
			case if1 == 1, if2 == 0:
				continue
			default:
				t.Errorf("invalid interface number 2 pickup: if1 = %d, if2 = %d", if1, if2)
			}
		}
	}

	{
		for i := 0; i < 10000; i++ {
			if1, if2 := pickupInterfaceNumbers(24)
			if if1 < 0 || if1 > 23 || if2 < 0 || if2 > 23 || if1 == if2 {
				t.Errorf("invalid interface number 24 pickup: if1 = %d, if2 = %d", if1, if2)
			}
		}
	}
}
