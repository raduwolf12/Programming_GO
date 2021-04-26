package main

import r "./ratio"

func main() {
	val3 := r.Rational{}
	val3 = r.Rational.NewInstance(val3, 2, 2)
	val1 := r.Rational{
		Numa: 2,
		Numi: 3,
	}
	val2 := r.Rational{
		Numa: 3,
		Numi: 3,
	}
	println("Numar rational 1:")
	val1.ShowNr()
	println("Numar rational 2:")
	val2.ShowNr()
	println("Numar rational 3:")
	val3.ShowNr()

	println("Numar rational 3 + Numar rational 2:")
	val3=val3.Add(val2)

	val3.ShowNr()


	println("Comparare val1 cu val2:")
	println(val1.BigerThan(val2))
	println(val1.SmallerThan(val2))
	println(val1.Equals(val2))
	println(val1.IsNatural())

	val1.Add(val2).ShowNr()
	val1.AddInt(2).ShowNr()

	val1.Multiply(val2).ShowNr()
	val1.MultiplyInt(2).ShowNr()

	val1.Divide(val2).ShowNr()
	val1.DivideInt(2).ShowNr()

	val1.Substract(val2)
	val1.SubstractInt(2).ShowNr()
	val1.Divide(val2).ShowNr()

	val1.AddNuma(2).ShowNr()
	val1.AddNumi(3).ShowNr()
	val1.AddNumaAndNumi(4).ShowNr()

	val1.SubstractNuma(2)
	val1.SubstractNumi(1).ShowNr()
	val1.SubstractNumaAndNumi(4).ShowNr()



	println(val1.IsNull())
	println(val1.GetRealValue())
	val1.GetAbsValue().ShowNr()

	val1.Simplify(val2).ShowNr()
	val1.Pow(3).ShowNr()
	val1.Inverse().ShowNr()

	r.GetFromFloat32(3.124).ShowNr()

}
