package popcount

//// pc[i] is the population count of i.
//var pc [256]byte
//
//func init() {
//	for i := range pc {
//		pc[i] = pc[i/2] + byte(i&1)
//	}
//}

var pc [256]byte = func() ( pc [256]byte) {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	return
}()

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	//fmt.Println(pc)
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountByLoop(x uint64) int {
	var n int
	for i:=0 ;i < 8 ; i++ {
		n += int(pc[byte(x>>(i*8))])
	}
	return n
}

func PopCountByLoopShift(x uint64) int {
	var n int
	for i:=0 ; i< 64; i++ {
		if x&1 != 0 {
			n++
		}
		x = x >> 1
	}
	return n
}

func PopCountByAnd(x uint64) int {
	//7 		1001	n=1
	//7-1		1000	n=1
	//&			1000	n=1,x=8
	//6-1		0111	n=2
	//&			0000	n=2,x=0
	var n int
	for x!=0 {
		x = x&(x-1)
		n++
	}
	return n
}