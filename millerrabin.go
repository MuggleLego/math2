package math2

import (
	"math/rand"
	"time"
)

//assuming a,b positive
func GCD(a int64, b int64) int64 {
	if b == 0 {
		panic("divided by zero!")
	}
	if a%b == 0 {
		return b
	}
	return GCD(b, a%b)
}

//return a^b mod n
func power(a, b, n int64) int64 {
	if b < 0 {
		panic("The power should not be negative")
	}
	var res int64 = 1
	a = a % n
	for b > 0 {
		if b&1 == 1 {
			res = (res * a) % n
		}
		a = (a * a) % n
		b >>= 1
	}
	return res

}

func MillerTest(n int64) bool {
	if n <= 1 {
		return false
	}
	rand.Seed(time.Now().UnixNano())
	a := rand.Int63n(n)
	for GCD(a, n) != 1 {
		a = rand.Int63n(n)
	} //generate random number a that gcd(a,n) ==  1
	q := n - 1
	for q&1 == 0 {
		q >>= 1
	}
	tmp := power(a, q, n)
	if tmp == 1 || tmp == n-1 { //case 1: a^q \equiv 1 \pmod n
		return true
	}
	for q != n-1 {
		tmp = (tmp * tmp) % n
		q <<= 1
		if tmp == n-1 {
			return true
		}
	}
	return false
}

//Testing primality of n with k times of Miller-Rabin
func MillerRabin(n int64, k int64) bool {
	for i := 0; i < int(k); i++ {
		if !MillerTest(n) {
			return false
		}
	}
	return true
}
