// Copyright 2020 ConsenSys Software Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by consensys/gnark-crypto DO NOT EDIT

package fr

// /!\ WARNING /!\
// this code has not been audited and is provided as-is. In particular,
// there is no security guarantees such as constant time implementation
// or side-channel attack resistance
// /!\ WARNING /!\

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"io"
	"math/big"
	"math/bits"
	"reflect"
	"strconv"
	"sync"
)

// Element represents a field element stored on 4 words (uint64)
// Element are assumed to be in Montgomery form in all methods
// field modulus q =
//
// 11502027791375260645628074404575422495959608200132055716665986169834464870401
type Element [4]uint64

// Limbs number of 64 bits words needed to represent Element
const Limbs = 4

// Bits number bits needed to represent Element
const Bits = 253

// Bytes number bytes needed to represent Element
const Bytes = Limbs * 8

// field modulus stored as big.Int
var _modulus big.Int

// Modulus returns q as a big.Int
// q =
//
// 11502027791375260645628074404575422495959608200132055716665986169834464870401
func Modulus() *big.Int {
	return new(big.Int).Set(&_modulus)
}

// q (modulus)
var qElement = Element{
	1860204336533995521,
	14466829657984787300,
	2737202078770428568,
	1832378743606059307,
}

// rSquare
var rSquare = Element{
	6242551132904523857,
	16951295617263545407,
	10923821274252739203,
	584663452775307866,
}

var bigIntPool = sync.Pool{
	New: func() interface{} {
		return new(big.Int)
	},
}

func init() {
	_modulus.SetString("11502027791375260645628074404575422495959608200132055716665986169834464870401", 10)
}

// NewElement returns a new Element from a uint64 value
//
// it is equivalent to
// 		var v NewElement
// 		v.SetUint64(...)
func NewElement(v uint64) Element {
	z := Element{v}
	z.Mul(&z, &rSquare)
	return z
}

// SetUint64 z = v, sets z LSB to v (non-Montgomery form) and convert z to Montgomery form
func (z *Element) SetUint64(v uint64) *Element {
	*z = Element{v}
	return z.Mul(z, &rSquare) // z.ToMont()
}

// Set z = x
func (z *Element) Set(x *Element) *Element {
	z[0] = x[0]
	z[1] = x[1]
	z[2] = x[2]
	z[3] = x[3]
	return z
}

// SetInterface converts provided interface into Element
// returns an error if provided type is not supported
// supported types: Element, *Element, uint64, int, string (interpreted as base10 integer),
// *big.Int, big.Int, []byte
func (z *Element) SetInterface(i1 interface{}) (*Element, error) {
	switch c1 := i1.(type) {
	case Element:
		return z.Set(&c1), nil
	case *Element:
		return z.Set(c1), nil
	case uint64:
		return z.SetUint64(c1), nil
	case int:
		return z.SetString(strconv.Itoa(c1)), nil
	case string:
		return z.SetString(c1), nil
	case *big.Int:
		return z.SetBigInt(c1), nil
	case big.Int:
		return z.SetBigInt(&c1), nil
	case []byte:
		return z.SetBytes(c1), nil
	default:
		return nil, errors.New("can't set fr.Element from type " + reflect.TypeOf(i1).String())
	}
}

// SetZero z = 0
func (z *Element) SetZero() *Element {
	z[0] = 0
	z[1] = 0
	z[2] = 0
	z[3] = 0
	return z
}

// SetOne z = 1 (in Montgomery form)
func (z *Element) SetOne() *Element {
	z[0] = 18291444782079148022
	z[1] = 2905656009828539926
	z[2] = 9521467359714817544
	z[3] = 122956637648958544
	return z
}

// Div z = x*y^-1 mod q
func (z *Element) Div(x, y *Element) *Element {
	var yInv Element
	yInv.Inverse(y)
	z.Mul(x, &yInv)
	return z
}

// Bit returns the i'th bit, with lsb == bit 0.
// It is the responsability of the caller to convert from Montgomery to Regular form if needed
func (z *Element) Bit(i uint64) uint64 {
	j := i / 64
	if j >= 4 {
		return 0
	}
	return uint64(z[j] >> (i % 64) & 1)
}

// Equal returns z == x
func (z *Element) Equal(x *Element) bool {
	return (z[3] == x[3]) && (z[2] == x[2]) && (z[1] == x[1]) && (z[0] == x[0])
}

// IsZero returns z == 0
func (z *Element) IsZero() bool {
	return (z[3] | z[2] | z[1] | z[0]) == 0
}

// IsUint64 returns true if z[0] >= 0 and all other words are 0
func (z *Element) IsUint64() bool {
	return (z[3] | z[2] | z[1]) == 0
}

// Cmp compares (lexicographic order) z and x and returns:
//
//   -1 if z <  x
//    0 if z == x
//   +1 if z >  x
//
func (z *Element) Cmp(x *Element) int {
	_z := *z
	_x := *x
	_z.FromMont()
	_x.FromMont()
	if _z[3] > _x[3] {
		return 1
	} else if _z[3] < _x[3] {
		return -1
	}
	if _z[2] > _x[2] {
		return 1
	} else if _z[2] < _x[2] {
		return -1
	}
	if _z[1] > _x[1] {
		return 1
	} else if _z[1] < _x[1] {
		return -1
	}
	if _z[0] > _x[0] {
		return 1
	} else if _z[0] < _x[0] {
		return -1
	}
	return 0
}

// LexicographicallyLargest returns true if this element is strictly lexicographically
// larger than its negation, false otherwise
func (z *Element) LexicographicallyLargest() bool {
	// adapted from github.com/zkcrypto/bls12_381
	// we check if the element is larger than (q-1) / 2
	// if z - (((q -1) / 2) + 1) have no underflow, then z > (q-1) / 2

	_z := *z
	_z.FromMont()

	var b uint64
	_, b = bits.Sub64(_z[0], 930102168266997761, 0)
	_, b = bits.Sub64(_z[1], 7233414828992393650, b)
	_, b = bits.Sub64(_z[2], 10591973076239990092, b)
	_, b = bits.Sub64(_z[3], 916189371803029653, b)

	return b == 0
}

// SetRandom sets z to a random element < q
func (z *Element) SetRandom() (*Element, error) {
	var bytes [32]byte
	if _, err := io.ReadFull(rand.Reader, bytes[:]); err != nil {
		return nil, err
	}
	z[0] = binary.BigEndian.Uint64(bytes[0:8])
	z[1] = binary.BigEndian.Uint64(bytes[8:16])
	z[2] = binary.BigEndian.Uint64(bytes[16:24])
	z[3] = binary.BigEndian.Uint64(bytes[24:32])
	z[3] %= 1832378743606059307

	// if z > q --> z -= q
	// note: this is NOT constant time
	if !(z[3] < 1832378743606059307 || (z[3] == 1832378743606059307 && (z[2] < 2737202078770428568 || (z[2] == 2737202078770428568 && (z[1] < 14466829657984787300 || (z[1] == 14466829657984787300 && (z[0] < 1860204336533995521))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 1860204336533995521, 0)
		z[1], b = bits.Sub64(z[1], 14466829657984787300, b)
		z[2], b = bits.Sub64(z[2], 2737202078770428568, b)
		z[3], _ = bits.Sub64(z[3], 1832378743606059307, b)
	}

	return z, nil
}

// One returns 1 (in montgommery form)
func One() Element {
	var one Element
	one.SetOne()
	return one
}

// Halve sets z to z / 2 (mod p)
func (z *Element) Halve() {
	if z[0]&1 == 1 {
		var carry uint64

		// z = z + q
		z[0], carry = bits.Add64(z[0], 1860204336533995521, 0)
		z[1], carry = bits.Add64(z[1], 14466829657984787300, carry)
		z[2], carry = bits.Add64(z[2], 2737202078770428568, carry)
		z[3], _ = bits.Add64(z[3], 1832378743606059307, carry)

	}

	// z = z >> 1

	z[0] = z[0]>>1 | z[1]<<63
	z[1] = z[1]>>1 | z[2]<<63
	z[2] = z[2]>>1 | z[3]<<63
	z[3] >>= 1

}

// API with assembly impl

// Mul z = x * y mod q
// see https://hackmd.io/@zkteam/modular_multiplication
func (z *Element) Mul(x, y *Element) *Element {
	mul(z, x, y)
	return z
}

// Square z = x * x mod q
// see https://hackmd.io/@zkteam/modular_multiplication
func (z *Element) Square(x *Element) *Element {
	mul(z, x, x)
	return z
}

// FromMont converts z in place (i.e. mutates) from Montgomery to regular representation
// sets and returns z = z * 1
func (z *Element) FromMont() *Element {
	fromMont(z)
	return z
}

// Add z = x + y mod q
func (z *Element) Add(x, y *Element) *Element {
	add(z, x, y)
	return z
}

// Double z = x + x mod q, aka Lsh 1
func (z *Element) Double(x *Element) *Element {
	double(z, x)
	return z
}

// Sub  z = x - y mod q
func (z *Element) Sub(x, y *Element) *Element {
	sub(z, x, y)
	return z
}

// Neg z = q - x
func (z *Element) Neg(x *Element) *Element {
	neg(z, x)
	return z
}

// Generic (no ADX instructions, no AMD64) versions of multiplication and squaring algorithms

func _mulGeneric(z, x, y *Element) {

	var t [4]uint64
	var c [3]uint64
	{
		// round 0
		v := x[0]
		c[1], c[0] = bits.Mul64(v, y[0])
		m := c[0] * 2184305180030271487
		c[2] = madd0(m, 1860204336533995521, c[0])
		c[1], c[0] = madd1(v, y[1], c[1])
		c[2], t[0] = madd2(m, 14466829657984787300, c[2], c[0])
		c[1], c[0] = madd1(v, y[2], c[1])
		c[2], t[1] = madd2(m, 2737202078770428568, c[2], c[0])
		c[1], c[0] = madd1(v, y[3], c[1])
		t[3], t[2] = madd3(m, 1832378743606059307, c[0], c[2], c[1])
	}
	{
		// round 1
		v := x[1]
		c[1], c[0] = madd1(v, y[0], t[0])
		m := c[0] * 2184305180030271487
		c[2] = madd0(m, 1860204336533995521, c[0])
		c[1], c[0] = madd2(v, y[1], c[1], t[1])
		c[2], t[0] = madd2(m, 14466829657984787300, c[2], c[0])
		c[1], c[0] = madd2(v, y[2], c[1], t[2])
		c[2], t[1] = madd2(m, 2737202078770428568, c[2], c[0])
		c[1], c[0] = madd2(v, y[3], c[1], t[3])
		t[3], t[2] = madd3(m, 1832378743606059307, c[0], c[2], c[1])
	}
	{
		// round 2
		v := x[2]
		c[1], c[0] = madd1(v, y[0], t[0])
		m := c[0] * 2184305180030271487
		c[2] = madd0(m, 1860204336533995521, c[0])
		c[1], c[0] = madd2(v, y[1], c[1], t[1])
		c[2], t[0] = madd2(m, 14466829657984787300, c[2], c[0])
		c[1], c[0] = madd2(v, y[2], c[1], t[2])
		c[2], t[1] = madd2(m, 2737202078770428568, c[2], c[0])
		c[1], c[0] = madd2(v, y[3], c[1], t[3])
		t[3], t[2] = madd3(m, 1832378743606059307, c[0], c[2], c[1])
	}
	{
		// round 3
		v := x[3]
		c[1], c[0] = madd1(v, y[0], t[0])
		m := c[0] * 2184305180030271487
		c[2] = madd0(m, 1860204336533995521, c[0])
		c[1], c[0] = madd2(v, y[1], c[1], t[1])
		c[2], z[0] = madd2(m, 14466829657984787300, c[2], c[0])
		c[1], c[0] = madd2(v, y[2], c[1], t[2])
		c[2], z[1] = madd2(m, 2737202078770428568, c[2], c[0])
		c[1], c[0] = madd2(v, y[3], c[1], t[3])
		z[3], z[2] = madd3(m, 1832378743606059307, c[0], c[2], c[1])
	}

	// if z > q --> z -= q
	// note: this is NOT constant time
	if !(z[3] < 1832378743606059307 || (z[3] == 1832378743606059307 && (z[2] < 2737202078770428568 || (z[2] == 2737202078770428568 && (z[1] < 14466829657984787300 || (z[1] == 14466829657984787300 && (z[0] < 1860204336533995521))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 1860204336533995521, 0)
		z[1], b = bits.Sub64(z[1], 14466829657984787300, b)
		z[2], b = bits.Sub64(z[2], 2737202078770428568, b)
		z[3], _ = bits.Sub64(z[3], 1832378743606059307, b)
	}
}

func _fromMontGeneric(z *Element) {
	// the following lines implement z = z * 1
	// with a modified CIOS montgomery multiplication
	{
		// m = z[0]n'[0] mod W
		m := z[0] * 2184305180030271487
		C := madd0(m, 1860204336533995521, z[0])
		C, z[0] = madd2(m, 14466829657984787300, z[1], C)
		C, z[1] = madd2(m, 2737202078770428568, z[2], C)
		C, z[2] = madd2(m, 1832378743606059307, z[3], C)
		z[3] = C
	}
	{
		// m = z[0]n'[0] mod W
		m := z[0] * 2184305180030271487
		C := madd0(m, 1860204336533995521, z[0])
		C, z[0] = madd2(m, 14466829657984787300, z[1], C)
		C, z[1] = madd2(m, 2737202078770428568, z[2], C)
		C, z[2] = madd2(m, 1832378743606059307, z[3], C)
		z[3] = C
	}
	{
		// m = z[0]n'[0] mod W
		m := z[0] * 2184305180030271487
		C := madd0(m, 1860204336533995521, z[0])
		C, z[0] = madd2(m, 14466829657984787300, z[1], C)
		C, z[1] = madd2(m, 2737202078770428568, z[2], C)
		C, z[2] = madd2(m, 1832378743606059307, z[3], C)
		z[3] = C
	}
	{
		// m = z[0]n'[0] mod W
		m := z[0] * 2184305180030271487
		C := madd0(m, 1860204336533995521, z[0])
		C, z[0] = madd2(m, 14466829657984787300, z[1], C)
		C, z[1] = madd2(m, 2737202078770428568, z[2], C)
		C, z[2] = madd2(m, 1832378743606059307, z[3], C)
		z[3] = C
	}

	// if z > q --> z -= q
	// note: this is NOT constant time
	if !(z[3] < 1832378743606059307 || (z[3] == 1832378743606059307 && (z[2] < 2737202078770428568 || (z[2] == 2737202078770428568 && (z[1] < 14466829657984787300 || (z[1] == 14466829657984787300 && (z[0] < 1860204336533995521))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 1860204336533995521, 0)
		z[1], b = bits.Sub64(z[1], 14466829657984787300, b)
		z[2], b = bits.Sub64(z[2], 2737202078770428568, b)
		z[3], _ = bits.Sub64(z[3], 1832378743606059307, b)
	}
}

func _addGeneric(z, x, y *Element) {
	var carry uint64

	z[0], carry = bits.Add64(x[0], y[0], 0)
	z[1], carry = bits.Add64(x[1], y[1], carry)
	z[2], carry = bits.Add64(x[2], y[2], carry)
	z[3], _ = bits.Add64(x[3], y[3], carry)

	// if z > q --> z -= q
	// note: this is NOT constant time
	if !(z[3] < 1832378743606059307 || (z[3] == 1832378743606059307 && (z[2] < 2737202078770428568 || (z[2] == 2737202078770428568 && (z[1] < 14466829657984787300 || (z[1] == 14466829657984787300 && (z[0] < 1860204336533995521))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 1860204336533995521, 0)
		z[1], b = bits.Sub64(z[1], 14466829657984787300, b)
		z[2], b = bits.Sub64(z[2], 2737202078770428568, b)
		z[3], _ = bits.Sub64(z[3], 1832378743606059307, b)
	}
}

func _doubleGeneric(z, x *Element) {
	var carry uint64

	z[0], carry = bits.Add64(x[0], x[0], 0)
	z[1], carry = bits.Add64(x[1], x[1], carry)
	z[2], carry = bits.Add64(x[2], x[2], carry)
	z[3], _ = bits.Add64(x[3], x[3], carry)

	// if z > q --> z -= q
	// note: this is NOT constant time
	if !(z[3] < 1832378743606059307 || (z[3] == 1832378743606059307 && (z[2] < 2737202078770428568 || (z[2] == 2737202078770428568 && (z[1] < 14466829657984787300 || (z[1] == 14466829657984787300 && (z[0] < 1860204336533995521))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 1860204336533995521, 0)
		z[1], b = bits.Sub64(z[1], 14466829657984787300, b)
		z[2], b = bits.Sub64(z[2], 2737202078770428568, b)
		z[3], _ = bits.Sub64(z[3], 1832378743606059307, b)
	}
}

func _subGeneric(z, x, y *Element) {
	var b uint64
	z[0], b = bits.Sub64(x[0], y[0], 0)
	z[1], b = bits.Sub64(x[1], y[1], b)
	z[2], b = bits.Sub64(x[2], y[2], b)
	z[3], b = bits.Sub64(x[3], y[3], b)
	if b != 0 {
		var c uint64
		z[0], c = bits.Add64(z[0], 1860204336533995521, 0)
		z[1], c = bits.Add64(z[1], 14466829657984787300, c)
		z[2], c = bits.Add64(z[2], 2737202078770428568, c)
		z[3], _ = bits.Add64(z[3], 1832378743606059307, c)
	}
}

func _negGeneric(z, x *Element) {
	if x.IsZero() {
		z.SetZero()
		return
	}
	var borrow uint64
	z[0], borrow = bits.Sub64(1860204336533995521, x[0], 0)
	z[1], borrow = bits.Sub64(14466829657984787300, x[1], borrow)
	z[2], borrow = bits.Sub64(2737202078770428568, x[2], borrow)
	z[3], _ = bits.Sub64(1832378743606059307, x[3], borrow)
}

func _reduceGeneric(z *Element) {

	// if z > q --> z -= q
	// note: this is NOT constant time
	if !(z[3] < 1832378743606059307 || (z[3] == 1832378743606059307 && (z[2] < 2737202078770428568 || (z[2] == 2737202078770428568 && (z[1] < 14466829657984787300 || (z[1] == 14466829657984787300 && (z[0] < 1860204336533995521))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 1860204336533995521, 0)
		z[1], b = bits.Sub64(z[1], 14466829657984787300, b)
		z[2], b = bits.Sub64(z[2], 2737202078770428568, b)
		z[3], _ = bits.Sub64(z[3], 1832378743606059307, b)
	}
}

func mulByConstant(z *Element, c uint8) {
	switch c {
	case 0:
		z.SetZero()
		return
	case 1:
		return
	case 2:
		z.Double(z)
		return
	case 3:
		_z := *z
		z.Double(z).Add(z, &_z)
	case 5:
		_z := *z
		z.Double(z).Double(z).Add(z, &_z)
	default:
		var y Element
		y.SetUint64(uint64(c))
		z.Mul(z, &y)
	}
}

// BatchInvert returns a new slice with every element inverted.
// Uses Montgomery batch inversion trick
func BatchInvert(a []Element) []Element {
	res := make([]Element, len(a))
	if len(a) == 0 {
		return res
	}

	zeroes := make([]bool, len(a))
	accumulator := One()

	for i := 0; i < len(a); i++ {
		if a[i].IsZero() {
			zeroes[i] = true
			continue
		}
		res[i] = accumulator
		accumulator.Mul(&accumulator, &a[i])
	}

	accumulator.Inverse(&accumulator)

	for i := len(a) - 1; i >= 0; i-- {
		if zeroes[i] {
			continue
		}
		res[i].Mul(&res[i], &accumulator)
		accumulator.Mul(&accumulator, &a[i])
	}

	return res
}

func _butterflyGeneric(a, b *Element) {
	t := *a
	a.Add(a, b)
	b.Sub(&t, b)
}

// BitLen returns the minimum number of bits needed to represent z
// returns 0 if z == 0
func (z *Element) BitLen() int {
	if z[3] != 0 {
		return 192 + bits.Len64(z[3])
	}
	if z[2] != 0 {
		return 128 + bits.Len64(z[2])
	}
	if z[1] != 0 {
		return 64 + bits.Len64(z[1])
	}
	return bits.Len64(z[0])
}

// Exp z = x^exponent mod q
func (z *Element) Exp(x Element, exponent *big.Int) *Element {
	var bZero big.Int
	if exponent.Cmp(&bZero) == 0 {
		return z.SetOne()
	}

	z.Set(&x)

	for i := exponent.BitLen() - 2; i >= 0; i-- {
		z.Square(z)
		if exponent.Bit(i) == 1 {
			z.Mul(z, &x)
		}
	}

	return z
}

// ToMont converts z to Montgomery form
// sets and returns z = z * r^2
func (z *Element) ToMont() *Element {
	return z.Mul(z, &rSquare)
}

// ToRegular returns z in regular form (doesn't mutate z)
func (z Element) ToRegular() Element {
	return *z.FromMont()
}

// String returns the string form of an Element in Montgomery form
func (z *Element) String() string {
	zz := *z
	zz.FromMont()
	if zz.IsUint64() {
		return strconv.FormatUint(zz[0], 10)
	} else {
		var zzNeg Element
		zzNeg.Neg(z)
		zzNeg.FromMont()
		if zzNeg.IsUint64() {
			return "-" + strconv.FormatUint(zzNeg[0], 10)
		}
	}
	vv := bigIntPool.Get().(*big.Int)
	defer bigIntPool.Put(vv)
	return zz.ToBigInt(vv).String()
}

// ToBigInt returns z as a big.Int in Montgomery form
func (z *Element) ToBigInt(res *big.Int) *big.Int {
	var b [Limbs * 8]byte
	binary.BigEndian.PutUint64(b[24:32], z[0])
	binary.BigEndian.PutUint64(b[16:24], z[1])
	binary.BigEndian.PutUint64(b[8:16], z[2])
	binary.BigEndian.PutUint64(b[0:8], z[3])

	return res.SetBytes(b[:])
}

// ToBigIntRegular returns z as a big.Int in regular form
func (z Element) ToBigIntRegular(res *big.Int) *big.Int {
	z.FromMont()
	return z.ToBigInt(res)
}

// Bytes returns the regular (non montgomery) value
// of z as a big-endian byte array.
func (z *Element) Bytes() (res [Limbs * 8]byte) {
	_z := z.ToRegular()
	binary.BigEndian.PutUint64(res[24:32], _z[0])
	binary.BigEndian.PutUint64(res[16:24], _z[1])
	binary.BigEndian.PutUint64(res[8:16], _z[2])
	binary.BigEndian.PutUint64(res[0:8], _z[3])

	return
}

// Marshal returns the regular (non montgomery) value
// of z as a big-endian byte slice.
func (z *Element) Marshal() []byte {
	b := z.Bytes()
	return b[:]
}

// SetBytes interprets e as the bytes of a big-endian unsigned integer,
// sets z to that value (in Montgomery form), and returns z.
func (z *Element) SetBytes(e []byte) *Element {
	// get a big int from our pool
	vv := bigIntPool.Get().(*big.Int)
	vv.SetBytes(e)

	// set big int
	z.SetBigInt(vv)

	// put temporary object back in pool
	bigIntPool.Put(vv)

	return z
}

// SetBigInt sets z to v (regular form) and returns z in Montgomery form
func (z *Element) SetBigInt(v *big.Int) *Element {
	z.SetZero()

	var zero big.Int

	// fast path
	c := v.Cmp(&_modulus)
	if c == 0 {
		// v == 0
		return z
	} else if c != 1 && v.Cmp(&zero) != -1 {
		// 0 < v < q
		return z.setBigInt(v)
	}

	// get temporary big int from the pool
	vv := bigIntPool.Get().(*big.Int)

	// copy input + modular reduction
	vv.Set(v)
	vv.Mod(v, &_modulus)

	// set big int byte value
	z.setBigInt(vv)

	// release object into pool
	bigIntPool.Put(vv)
	return z
}

// setBigInt assumes 0 <= v < q
func (z *Element) setBigInt(v *big.Int) *Element {
	vBits := v.Bits()

	if bits.UintSize == 64 {
		for i := 0; i < len(vBits); i++ {
			z[i] = uint64(vBits[i])
		}
	} else {
		for i := 0; i < len(vBits); i++ {
			if i%2 == 0 {
				z[i/2] = uint64(vBits[i])
			} else {
				z[i/2] |= uint64(vBits[i]) << 32
			}
		}
	}

	return z.ToMont()
}

// SetString creates a big.Int with number and calls SetBigInt on z
//
// The number prefix determines the actual base: A prefix of
// ''0b'' or ''0B'' selects base 2, ''0'', ''0o'' or ''0O'' selects base 8,
// and ''0x'' or ''0X'' selects base 16. Otherwise, the selected base is 10
// and no prefix is accepted.
//
// For base 16, lower and upper case letters are considered the same:
// The letters 'a' to 'f' and 'A' to 'F' represent digit values 10 to 15.
//
// An underscore character ''_'' may appear between a base
// prefix and an adjacent digit, and between successive digits; such
// underscores do not change the value of the number.
// Incorrect placement of underscores is reported as a panic if there
// are no other errors.
//
func (z *Element) SetString(number string) *Element {
	// get temporary big int from the pool
	vv := bigIntPool.Get().(*big.Int)

	if _, ok := vv.SetString(number, 0); !ok {
		panic("Element.SetString failed -> can't parse number into a big.Int " + number)
	}

	z.SetBigInt(vv)

	// release object into pool
	bigIntPool.Put(vv)

	return z
}

// Legendre returns the Legendre symbol of z (either +1, -1, or 0.)
func (z *Element) Legendre() int {
	var l Element
	// z^((q-1)/2)
	l.expByLegendreExp(*z)

	if l.IsZero() {
		return 0
	}

	// if l == 1
	if (l[3] == 122956637648958544) && (l[2] == 9521467359714817544) && (l[1] == 2905656009828539926) && (l[0] == 18291444782079148022) {
		return 1
	}
	return -1
}

// Sqrt z = √x mod q
// if the square root doesn't exist (x is not a square mod q)
// Sqrt leaves z unchanged and returns nil
func (z *Element) Sqrt(x *Element) *Element {
	// q ≡ 1 (mod 4)
	// see modSqrtTonelliShanks in math/big/int.go
	// using https://www.maa.org/sites/default/files/pdf/upload_library/22/Polya/07468342.di020786.02p0470a.pdf

	var y, b, t, w Element
	// w = x^((s-1)/2))
	w.expBySqrtExp(*x)

	// y = x^((s+1)/2)) = w * x
	y.Mul(x, &w)

	// b = x^s = w * w * x = y * x
	b.Mul(&w, &y)

	// g = nonResidue ^ s
	var g = Element{
		2675275753227370406,
		18180984726441494600,
		9289909143059162211,
		12979261504110204,
	}
	r := uint64(22)

	// compute legendre symbol
	// t = x^((q-1)/2) = r-1 squaring of x^s
	t = b
	for i := uint64(0); i < r-1; i++ {
		t.Square(&t)
	}
	if t.IsZero() {
		return z.SetZero()
	}
	if !((t[3] == 122956637648958544) && (t[2] == 9521467359714817544) && (t[1] == 2905656009828539926) && (t[0] == 18291444782079148022)) {
		// t != 1, we don't have a square root
		return nil
	}
	for {
		var m uint64
		t = b

		// for t != 1
		for !((t[3] == 122956637648958544) && (t[2] == 9521467359714817544) && (t[1] == 2905656009828539926) && (t[0] == 18291444782079148022)) {
			t.Square(&t)
			m++
		}

		if m == 0 {
			return z.Set(&y)
		}
		// t = g^(2^(r-m-1)) mod q
		ge := int(r - m - 1)
		t = g
		for ge > 0 {
			t.Square(&t)
			ge--
		}

		g.Square(&t)
		y.Mul(&y, &t)
		b.Mul(&b, &g)
		r = m
	}
}

// Inverse z = x^-1 mod q
// Algorithm 16 in "Efficient Software-Implementation of Finite Fields with Applications to Cryptography"
// if x == 0, sets and returns z = x
func (z *Element) Inverse(x *Element) *Element {
	if x.IsZero() {
		z.SetZero()
		return z
	}

	// initialize u = q
	var u = Element{
		1860204336533995521,
		14466829657984787300,
		2737202078770428568,
		1832378743606059307,
	}

	// initialize s = r^2
	var s = Element{
		6242551132904523857,
		16951295617263545407,
		10923821274252739203,
		584663452775307866,
	}

	// r = 0
	r := Element{}

	v := *x

	var carry, borrow uint64
	var bigger bool

	for {
		for v[0]&1 == 0 {

			// v = v >> 1

			v[0] = v[0]>>1 | v[1]<<63
			v[1] = v[1]>>1 | v[2]<<63
			v[2] = v[2]>>1 | v[3]<<63
			v[3] >>= 1

			if s[0]&1 == 1 {

				// s = s + q
				s[0], carry = bits.Add64(s[0], 1860204336533995521, 0)
				s[1], carry = bits.Add64(s[1], 14466829657984787300, carry)
				s[2], carry = bits.Add64(s[2], 2737202078770428568, carry)
				s[3], _ = bits.Add64(s[3], 1832378743606059307, carry)

			}

			// s = s >> 1

			s[0] = s[0]>>1 | s[1]<<63
			s[1] = s[1]>>1 | s[2]<<63
			s[2] = s[2]>>1 | s[3]<<63
			s[3] >>= 1

		}
		for u[0]&1 == 0 {

			// u = u >> 1

			u[0] = u[0]>>1 | u[1]<<63
			u[1] = u[1]>>1 | u[2]<<63
			u[2] = u[2]>>1 | u[3]<<63
			u[3] >>= 1

			if r[0]&1 == 1 {

				// r = r + q
				r[0], carry = bits.Add64(r[0], 1860204336533995521, 0)
				r[1], carry = bits.Add64(r[1], 14466829657984787300, carry)
				r[2], carry = bits.Add64(r[2], 2737202078770428568, carry)
				r[3], _ = bits.Add64(r[3], 1832378743606059307, carry)

			}

			// r = r >> 1

			r[0] = r[0]>>1 | r[1]<<63
			r[1] = r[1]>>1 | r[2]<<63
			r[2] = r[2]>>1 | r[3]<<63
			r[3] >>= 1

		}

		// v >= u
		bigger = !(v[3] < u[3] || (v[3] == u[3] && (v[2] < u[2] || (v[2] == u[2] && (v[1] < u[1] || (v[1] == u[1] && (v[0] < u[0])))))))

		if bigger {

			// v = v - u
			v[0], borrow = bits.Sub64(v[0], u[0], 0)
			v[1], borrow = bits.Sub64(v[1], u[1], borrow)
			v[2], borrow = bits.Sub64(v[2], u[2], borrow)
			v[3], _ = bits.Sub64(v[3], u[3], borrow)

			// s = s - r
			s[0], borrow = bits.Sub64(s[0], r[0], 0)
			s[1], borrow = bits.Sub64(s[1], r[1], borrow)
			s[2], borrow = bits.Sub64(s[2], r[2], borrow)
			s[3], borrow = bits.Sub64(s[3], r[3], borrow)

			if borrow == 1 {

				// s = s + q
				s[0], carry = bits.Add64(s[0], 1860204336533995521, 0)
				s[1], carry = bits.Add64(s[1], 14466829657984787300, carry)
				s[2], carry = bits.Add64(s[2], 2737202078770428568, carry)
				s[3], _ = bits.Add64(s[3], 1832378743606059307, carry)

			}
		} else {

			// u = u - v
			u[0], borrow = bits.Sub64(u[0], v[0], 0)
			u[1], borrow = bits.Sub64(u[1], v[1], borrow)
			u[2], borrow = bits.Sub64(u[2], v[2], borrow)
			u[3], _ = bits.Sub64(u[3], v[3], borrow)

			// r = r - s
			r[0], borrow = bits.Sub64(r[0], s[0], 0)
			r[1], borrow = bits.Sub64(r[1], s[1], borrow)
			r[2], borrow = bits.Sub64(r[2], s[2], borrow)
			r[3], borrow = bits.Sub64(r[3], s[3], borrow)

			if borrow == 1 {

				// r = r + q
				r[0], carry = bits.Add64(r[0], 1860204336533995521, 0)
				r[1], carry = bits.Add64(r[1], 14466829657984787300, carry)
				r[2], carry = bits.Add64(r[2], 2737202078770428568, carry)
				r[3], _ = bits.Add64(r[3], 1832378743606059307, carry)

			}
		}
		if (u[0] == 1) && (u[3]|u[2]|u[1]) == 0 {
			z.Set(&r)
			return z
		}
		if (v[0] == 1) && (v[3]|v[2]|v[1]) == 0 {
			z.Set(&s)
			return z
		}
	}

}
