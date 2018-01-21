package sparse

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func Example() {
	// Construct a new DOK (Dictionary Of Keys) matrix
	dokMatrix := NewDOK(3, 2)

	// Populate it with some non-zero values
	dokMatrix.Set(0, 0, 5)
	dokMatrix.Set(2, 1, 7)

	// Demonstrate accessing values (could use mat.Formatted() to
	// pretty print but this demonstrates element access)
	m, n := dokMatrix.Dims()
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%.0f", dokMatrix.At(i, j))
			if j < n-1 {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}

	// Convert DOK matrix to mat.Dense just for fun
	// (not required for upcoming multiplication operation)
	denseMatrix := dokMatrix.ToDense()

	// Confirm the two matrices in different formats are equal
	// Using the mat.Equal function
	if !mat.Equal(dokMatrix, denseMatrix) {
		fmt.Println("DOK and converted Dense are not equal")
	}

	// Create a random 2x3 CSR (Compressed Sparse Row) matrix with
	// density of 0.5 (half the elements will be non-zero)
	csrMatrix := Random(CSRFormat, 2, 3, 0.5)

	// Create a new CSR (Compressed Sparse Row) matrix
	csrProduct := &CSR{}

	// Multiply the 2 matrices together and store the result in the
	// receiver
	csrProduct.Mul(csrMatrix, denseMatrix)

	// Output: 5 0
	//0 0
	//0 7
}
