package ipnet

import (
    "testing"
)

func TestIPV4MaskFromString(t *testing.T) {
    {        
		addresses := map[string]int{
			"0.0.0.0": 0,
            "255.255.255.0" : 24,
            "255.255.255.255" : 32,
		}
		for address, maskSize := range addresses {
			mask, err := IPv4MaskFromString(address) 
			if err != nil {
				t.Errorf("raise error: address=%s, err=%v", address, err)
			}
            a, b := mask.Size()
            if a != maskSize || b != 32 {
				t.Errorf("mask.Size() is invalid, expect=%d/32, got=%d/%d", maskSize, a, b)
            }
		}
    }    
    
    {        
		invalidAddresses := []string{
			"foo",
            "a.1.2.3",
            "1",
            "1.1",
            "1.1.1",
            "1.1.1.1.1",
            "-1.1.2.3",
            "1.2.3.256",
		}
		for _, address := range invalidAddresses {
			_, err := IPv4MaskFromString(address) 
			if err == nil {
				t.Errorf("not raise error: address=%s", address)
			}
		}
    }    
}

