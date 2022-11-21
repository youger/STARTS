package dynamic_programming

import (
	"fmt"
	"testing"
)

func TestUniquePath(t *testing.T) {
	path := uniquePaths(23, 12)
	fmt.Println("path", path)
}

func TestUniquePath2(t *testing.T) {
	path := uniquePaths2(23, 12)
	fmt.Println("path", path)
}
