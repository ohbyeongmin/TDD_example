package calculate

var opMap map[string]func(int, int)int

func InitOpMap() {
	opMap = make(map[string]func(int, int)int)
	opMap["+"] = add
	opMap["-"] = sub
	opMap["*"] = mul
	opMap["/"] = div
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func Calculate(op string, a, b int) int {
	if v, ok := opMap[op]; ok {
		return v(a,b)
	}
	return 0
}

