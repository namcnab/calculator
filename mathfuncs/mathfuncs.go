package mathfuncs

type TwoNums struct {
	A float64 `json:"a" validate:"required"`
	B float64 `json:"b" validate:"required"`
}

type ListOfNums struct {
	List []float64
}

func AddTwo[T int | float32 | float64](a, b T) T {
	return a + b
}

func AddList[T []int | []float32 | []float64](numbers T) any {
	var sum any

	if len(numbers) == 0 {
		return 0
	}

	switch nums := any(numbers).(type) {
	case []int:
		sum := 0
		for _, num := range nums {
			sum += num
		}
	case []float32:
		sum := float32(0)
		for _, num := range nums {
			sum += num
		}
	case []float64:
		sum := float64(0)
		for _, num := range nums {
			sum += num
		}

	default:
		sum = 0
	}

	return sum
}

func DivideTwo[T int | float32 | float64](a, b T) T {
	if b == 0 {
		return 0 // Handle division by zero
	}
	return a / b
}

func MultiplyTwo[T int | float32 | float64](a, b T) T {
	return a * b
}

func SubtractTwo[T int | float32 | float64](a, b T) T {
	return a - b
}
