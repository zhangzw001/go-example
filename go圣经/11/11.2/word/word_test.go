package word

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)
//
//func TestIsPalindrome(t *testing.T) {
//	if !IsPalindrome("datartrated") {
//		t.Error(`IsPalindrome("detartrated") = false`)
//	}
//	if !IsPalindrome("kayak") {
//		t.Error(`IsPalindrome("kayak") = false`)
//	}
//}
//func TestNonPalindrome(t *testing.T) {
//	if IsPalindrome("palindrome") {
//		t.Error(`IsPalindrome("palindrome") = true`)
//	}
//}
//
//func TestFrenchPalindrome(t *testing.T) {
//	if !IsPalindrome("été") {
//		t.Error(`IsPalindrome("été") = false`)
//	}
//}
//
//func TestCanalPalindrome(t *testing.T) {
//	input := "A man, a plan, a canal: Panama"
//	if !IsPalindrome(input) {
//		t.Errorf(`IsPalindrome(%q) = false`, input)
//	}
//}


//
//func TestIsPalindrome2(t *testing.T) {
//	if !IsPalindrome2("datartrated") {
//		t.Error(`IsPalindrome2("detartrated") = false`)
//	}
//	if !IsPalindrome2("kayak") {
//		t.Error(`IsPalindrome2("kayak") = false`)
//	}
//}
func TestNonPalindrome2(t *testing.T) {
	if IsPalindrome2("palindrome") {
		t.Error(`IsPalindrome2("palindrome") = true`)
	}
}

func TestFrenchPalindrome2(t *testing.T) {
	if !IsPalindrome2("été") {
		t.Error(`IsPalindrome2("été") = false`)
	}
}

func TestCanalPalindrome2(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome2(input) {
		t.Errorf(`IsPalindrome2(%q) = false`, input)
	}
}


func TestIsPalindrome2(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // non-palindrome
		{"desserts", false},   // semi-palindrome
	}
	for _, test := range tests {
		//if got := IsPalindrome(test.input); got != test.want {
		if got := IsPalindrome2(test.input); got != test.want {
			t.Errorf("IsPalindrome2(%q) = %v", test.input, got)
			//t.Fatalf("IsPalindrome2(%q) = %v", test.input, got)
		}
	}
}




func TestRandomPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d",seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0 ; i < 10 ; i ++ {
		p := RandomPalindrome(rng)
		fmt.Printf("%s \t %v\n",p,IsPalindrome2(p))
		if !IsPalindrome2(p) {
			t.Errorf("IsPalindrome(%q) = false",p)
		}
	}
}

//练习 11.3: TestRandomPalindromes测试函数只测试了回文字符串。编写新的随机测试生成器，用于测试随机生成的非回文字符串。
func TestRandomNonPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d",seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0 ; i < 10 ; i ++ {
		p := RandomNonPalindrome(rng)
		fmt.Printf("%s \t %v\n",p,IsPalindrome2(p))
		if IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = true",p)
		}
	}
}
