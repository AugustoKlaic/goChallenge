package exception

import "strconv"

// CheckMatrix - check if the matrix have all elements as int type
func CheckMatrix(records [][]string) bool {
	for _, row := range records {
		for _, num := range row {
			_, err := strconv.Atoi(num)
			if err != nil {
				return false
			}
		}
	}
	return true
}
