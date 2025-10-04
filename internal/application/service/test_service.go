package service

import testdto "github.com/Hona-Tahlil/Backend/internal/application/dto/test"

func Test(data testdto.TestRequest) testdto.TestResponse {
	avg := ((float64)(data.Num1) + (float64)(data.Num2)) / 2.0
	mul := data.Num1 * data.Num2

	return testdto.TestResponse{
		Average:        avg,
		Multiplication: mul,
	}
}
