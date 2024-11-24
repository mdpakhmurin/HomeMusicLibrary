package mathutils

// Ограничивает число, приводя его к значению границ, при их превышении.
func BoundNumber(num int, lowerBound int, upperBound int) int {
	if num < lowerBound {
		return lowerBound
	} else if num > upperBound {
		return upperBound
	}
	return num
}
