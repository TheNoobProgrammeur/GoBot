package commands

import (
	"myTest/src/models"
	"myTest/src/services"
)

func Fibonacci(number uint64 ) uint64  {


	var fib models.Fibonacci

	dbservice := services.New()

	dbservice.GetDB().Where("input = ?",number).First(&fib)

	if fib.Output == 0 {
		fib.Input = number
		if number <= 2 {
			fib.Output = 1
			dbservice.GetDB().Save(&fib)
			return 1
		}
		fib.Output = Fibonacci(number-1) + Fibonacci(number-2)
		dbservice.GetDB().Save(&fib)
		return fib.Output
	}
	return fib.Output
}
