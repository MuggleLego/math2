package math2

import (
	"fmt"
)

type GF28 byte

type GF interface {
	AddSub(GF) GF
	Multiply(GF) GF
	Divide(GF) GF
	Inverse() GF
}

func (p GF28) AddSub(q GF28) GF28 {
	return p ^ q
}

//Times function receives a polynomial P in GF(2^8)
//Outputs {0x2}*P
func Times(p GF28) GF28 {
	if p&0x80 != 0 { //{0x2}*p will carry
		return (p << 1) ^ 0x1B //mod x^4+x^3+x+1
	} else {
		return p << 1
	}
}

func (p GF28) Multiply(q GF28) GF28 {
	tmp := []GF28{p}
	for i := 1; i < 8; i++ {
		tmp = append(tmp, Times(tmp[i-1])) //{P,P*0x2,P*0x4,P*0x8,...,P*0x80}
	}
	res := GF28(0x0)
	for i := 0; i < 8; i++ {
		if q&0x1 == 1 {
			res = res.AddSub(tmp[i])
		}
		q >>= 1
	}
	return res
}

func (p GF28) Order() int {
	k := 0
	for p != 0x0 {
		p >>= 1
		k += 1
	}
	if k == 0 {
		return k
	} else {
		return k - 1
	}
}

//PolyDivide function receive polynomial P(x),Q(x)
//Let P(x) = Q(x)B(x) + R(x)
//Output B(x)
//Attention:just a simple polynomial division (mod 2)
//Not involve any modulus polynomial
func (p GF28) PolyDivide(q GF28) GF28 {
	if q == GF28(0x0) {
		panic("Can not divided by zero!")
	}
	if q == GF28(0x1) {
		return p
	}
	if p.Order() < q.Order() {
		return GF28(0)
	}
	quotient := GF28(0)
	for p.Order()-q.Order() >= 0 {
		deg := p.Order() - q.Order()
		tmp := GF28(1 << deg)
		quotient = quotient.AddSub(tmp)
		p = p.AddSub(q.Multiply(tmp))
	}
	return quotient
}

//Inverse function receive an element P(x) in GF28 and output its inverse by egcd method
//The modolus polynomial M(x)=x^8+x^4+x^3+x+1=0x11B
//In order to calculate M(x)/P(x) which M(x) not in GF28
//We use the simple polynomial divison method
//Let M(x)=P(x)*Q(x)+R(x)
//[1,P(x),0,M(x)] -> [1,P(x),-Q(x),R(x)] -> [1,P(x),Q(x),R(x)]

func (p GF28) Inverse() GF28 {
	if p == 0x0 || p == 0x1 {
		return p
	}
	_m := 0x11B
	q := GF28(1 << (8 - p.Order())) //q(x) is the highest term of Q(x)
	_p := int(p) << (8 - p.Order()) //p(x)q(x)
	_r := _m ^ _p
	r := GF28(_r)

	x1 := GF28(0x1)
	x2 := p
	y1 := q
	y2 := r
	for y2 != GF28(0x0) {
		quotient := x2.PolyDivide(y2)
		tmp1, tmp2 := y1, y2
		y1 = x1.AddSub(quotient.Multiply(y1))
		y2 = x2.AddSub(quotient.Multiply(y2))
		x1, x2 = tmp1, tmp2
	}
	return x1
}

func GF28InverseTable() [16][16]GF28 {
	table := [16][16]GF28{}
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			table[i][j] = GF28(((i << 4) & 0xf0) ^ (j & 0xf))
			table[i][j] = table[i][j].Inverse()
		}
	}
	/*
		for _, i := range table {
			for _, v := range i {
				fmt.Printf("%x ", v)
			}
			fmt.Println()
		}*/
	return table
}

func (p GF28) transform() GF28 {
	res := GF28(0)
	c := GF28(63)
	for i := 0; i < 8; i++ {
		tmp := (p & (1 << i)) ^ (p & (1 << ((i + 4) % 8))) ^ (p & (1 << ((i + 5) % 8))) ^ (p & (1 << ((i + 6) % 8))) ^ (p & (1 << ((i + 7) % 8))) ^ (c & (1 << i))
		res ^= (tmp << i)
	}
	return res
}

func GetSbox() [16][16]GF28 {
	table := GF28InverseTable()
	for _, i := range table {
		for _, v := range i {
			v = v.transform()
			fmt.Printf("%x ", v)
		}
		fmt.Println()
	}
	return table
}
