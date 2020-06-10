package WordGuess

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var (
	Start int
	GuessValue int
	count int
	RandValue = RandInt(0,100)
	Flag = false
	End = 100
)
func initAll(){
	Start = GuessValue
	End = GuessValue
	Flag = true
	Start =0
	End = 100
	count=0
	RandValue = RandInt(0,100)
}
func WordGuess(w http.ResponseWriter,guessValue int) {
	if guessValue < Start || guessValue > End {
		fmt.Fprintf(w,"请输入 %d ~ %d的值\n",Start,End)
		return
	}
	GuessValue =  guessValue
	diff(RandValue,w)
	if Flag == false {
		fmt.Fprintf(w, "现在的区间是:\t[%d~%d]\n", Start, End)
		fmt.Printf("现在的区间是:\t[%d~%d]\n", Start, End)
	}
}

func WordGuessByAuto(w http.ResponseWriter) {

	RandValue = RandInt(Start,End)
	var count int
	for Start < End {
		diff(RandValue,w)
		fmt.Fprintf(w,"现在的区间是:\t[%d~%d]\n",Start,End)
		count++
	}
	fmt.Fprintf(w,"你一共猜了%d 次, 恭喜你!",count)
	fmt.Printf("你一共猜了%d 次, 恭喜你!",count)
}

func diff(myrand int, w http.ResponseWriter) {
	count++
	fmt.Printf("答案是:%d,\t你猜的是:%d\t\n",myrand, GuessValue)
	fmt.Fprintf(w,"你猜的是:%d\t",GuessValue)
	if GuessValue > myrand {
		fmt.Fprintln(w,"你猜大了")
		End = GuessValue-1
	}else if GuessValue < myrand {
		fmt.Fprintln(w,"你猜小了")
		Start = GuessValue+1
	}else if GuessValue == myrand {
		fmt.Fprintln( w,"你猜对了")
		fmt.Fprintf(w,"你一共猜了%d 次, 恭喜你!\n",count)
		initAll()
		fmt.Println("本次结束")
	}

}




func RandInt(min, max int) int {
	if min >= max || max ==0 || min < 0 {
		return max
	}
	//方法一
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
	////方法二
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//return r.Int63n(max-min) + min
}


