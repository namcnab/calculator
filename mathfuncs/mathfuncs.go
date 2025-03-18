package mathfuncs

func AddTwo(a, b int) int {
	return a + b
}

func SubtractTwo(a, b int) int {
	return a - b
}

func MultiplyTwo(a, b int) int {
	return a * b
}

func DivideTwo(a, b int) int {
	if b == 0 {
		return 0 // Handle division by zero
	}
	return a / b
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
		return "Unsupported type"
	}

	return sum
}

func SubtractList(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	result := numbers[0]
	for _, num := range numbers[1:] {
		result -= num
	}
	return result
}
