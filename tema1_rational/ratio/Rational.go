package ratio

type Rational struct {
	Numa, Numi int
}

func (r Rational) GetNuma() int {
	return r.Numa
}

func (r Rational) GetNumi() int {
	return r.Numi
}

func (r Rational) NewInstance(x int, y int) Rational {
	t := Rational{x, y}
	return t
}

func (r Rational) Add(t Rational) Rational {
	r.Numa = r.Numa*t.Numi + r.Numi*t.Numa
	r.Numi = r.Numi * t.Numi
	return r
}

func (r Rational) Substract(t Rational) Rational {
	r.Numa = r.Numa*t.Numi - r.Numi*t.Numa
	r.Numi = r.Numi * t.Numi
	return r
}

func (r Rational) Multiply(t Rational) Rational {
	r.Numa = r.Numa * t.Numa
	r.Numa = r.Numi * t.Numi
	return r
}

func (r Rational) MultiplyInt(t int) Rational {
	r.Numa = r.Numa * t
	return r
}

func (r Rational) DivideInt(t int) Rational {
	r.Numa = r.Numa / t
	return r
}

func (r Rational) SubstractInt(t int) Rational {
	r.Numi = r.Numi - t
	return r
}

func (r Rational) AddInt(t int) Rational {
	r.Numa = r.Numa + r.Numi*t
	return r
}

func (r Rational) AddNuma(t int) Rational {
	r.Numa = r.Numa + t
	return r
}

func (r Rational) AddNumi(t int) Rational {
	r.Numi = r.Numi + t
	return r
}

func (r Rational) AddNumaAndNumi(t int) Rational {
	r.Numa = r.Numa + t
	r.Numi = r.Numi + t
	return r
}

func (r Rational) SubstractNuma(t int) Rational {
	r.Numa = r.Numa - t
	return r
}

func (r Rational) SubstractNumi(t int) Rational {
	r.Numi = r.Numi - t
	return r
}

func (r Rational) SubstractNumaAndNumi(t int) Rational {
	r.Numa = r.Numa - t
	r.Numi = r.Numi - t
	return r
}

func (r Rational) IsNull() bool {
	return r.Numa == 0
}

func (r Rational) GetRealValue() float32 {
	var t float32 = float32(r.Numa / r.Numi)
	return t
}

func (r Rational) GetAbsValue() Rational {

	if r.Numi < 0 {
		r.Numi = r.Numi * -1
	}
	if r.Numa < 0 {
		r.Numa = r.Numa * -1
	}

	return r
}

func (r Rational) Divide(t Rational) Rational {
	r.Numa = r.Numa / t.Numi
	r.Numi = r.Numi / t.Numa
	return r
}

func (r Rational) Simplify(t Rational) Rational {
	c := Cmmdc(r.Numa, r.Numi)
	r.Numa /= c
	r.Numi /= c
	return r
}

func (r Rational) Pow(n int) Rational {
	for i := 0; i < n; i++ {
		r.Numi = r.Numi * r.Numi
		r.Numa = r.Numa * r.Numa
	}
	return r
}

func (r Rational) BigerThan(t Rational) bool {

	var val1 int = r.Numa * t.Numi
	var val2 int = r.Numi * t.Numa
	if val1 > val2 {
		return true
	}
	return false
}

func (r Rational) SmallerThan(t Rational) bool {
	var val1 int = r.Numa * t.Numi
	var val2 int = r.Numi * t.Numa
	if val1 < val2 {
		return true
	}
	return false
}

func (r Rational) Equals(t Rational) bool {

	var val1 int = r.Numa * t.Numi
	var val2 int = r.Numi * t.Numa
	if val1 == val2 {
		return true
	}
	return false
}

func (r Rational) Inverse() Rational {
	var i int = r.Numa
	var j int = r.Numi
	r.Numa = j
	r.Numi = i
	return r
}

func (r Rational) IsNatural() bool {
	if r.Numa/r.Numi == 0 {
		return true
	}
	return r.Numi == 1
}
func isIntegral(val float32) bool {
	return val == float32(int(val))
}
func num_after_point(x float32) int {
	var count int = 0

	var y int = int(x)
	var residue float32 = x - float32(y)
	if residue != 0 {
		var multiplier float32 = 1.0
		for !isIntegral((x * multiplier)) {
			count += 1
			multiplier = 10 * multiplier
		}

	}

	return count
}

// 0.4324 => 4324 / 1000 => simplificare
func GetFromFloat32(x float32) Rational {
	var zecimale int = num_after_point(x)
	var nu =1
	for i := 0; i < zecimale; i++ {
		 nu = nu * 10
	}

	var n int = int( float32(nu)  * x)

	return Rational{n, nu}
}

func Cmmdc(a, b int) int {
	for a%b != 0 {
		r := a % b
		a = b
		b = r
	}
	return b
}

func (r Rational) ShowNr() {

	print(r.Numa)
	print("\\")
	println(r.Numi)
}
