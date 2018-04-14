package bitwiseversusstandard

func isEvenStandard(number int) bool {
	return number%2 == 0
}

func isEvenBitwise(number int) bool {
	return byte(number&1) == 0
}
