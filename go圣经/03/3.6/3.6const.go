package main

import (
	"fmt"
)


type Weekday int
const(
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Flags uint

const (
	FlagUp Flags = 1 << iota // is up
	FlagBroadcast            // supports broadcast access capability
	FlagLoopback             // is a loopback interface
	FlagPointToPoint         // belongs to a point-to-point link
	FlagMulticast            // supports multicast access capability
)

func IsUp(v Flags) bool     { return v&FlagUp == FlagUp }
func TurnDown(v *Flags)     { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }

const (
	_ = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776             (exceeds 1 << 32)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424    (exceeds 1 << 64)
	YiB // 1208925819614629174706176
)

const (
	KB =  1000
	MB = KB * KB

)

func main () {
	fmt.Println(Sunday,Friday)
	fmt.Println(FlagBroadcast,FlagMulticast)
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 true"
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10000 false"
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))   // "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 true"

	fmt.Println(YiB/ZiB) // "1024"
}
