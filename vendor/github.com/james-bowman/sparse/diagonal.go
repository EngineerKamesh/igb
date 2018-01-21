package sparse

import (
	"gonum.org/v1/gonum/mat"
)

var (
	diagonal *DIA

	_ Sparser = diagonal

	_ mat.Matrix = diagonal

	_ mat.ColViewer    = diagonal
	_ mat.RowViewer    = diagonal
	_ mat.RawColViewer = diagonal
	_ mat.RawRowViewer = diagonal
)

// DIA matrix type is a specialised matrix designed to store DIAgonal values of square symmetrical
// matrices (all zero values except along the diagonal running top left to bottom right).  The DIA matrix
// type is specifically designed to take advantage of the sparsity pattern of square symmetrical matrices.
type DIA struct {
	m    int
	data []float64
}

// NewDIA creates a new DIAgonal format sparse matrix.
// The matrix is initialised to the size of the specified m * m dimensions (rows * columns)
// (creating a square) with the specified slice containing it's diagonal values.  The diagonal slice
// will be used as the backing slice to the matrix so changes to values of the slice will be reflected
// in the matrix.
func NewDIA(m int, diagonal []float64) *DIA {
	if uint(m) < 0 || m != len(diagonal) {
		panic(mat.ErrRowAccess)
	}

	return &DIA{m: m, data: diagonal}
}

// Dims returns the size of the matrix as the number of rows and columns
func (d *DIA) Dims() (int, int) {
	return d.m, d.m
}

// At returns the element of the matrix located at row i and column j.  At will panic if specified values
// for i or j fall outside the dimensions of the matrix.
func (d *DIA) At(i, j int) float64 {
	if uint(i) < 0 || uint(i) >= uint(d.m) {
		panic(mat.ErrRowAccess)
	}
	if uint(j) < 0 || uint(j) >= uint(d.m) {
		panic(mat.ErrColAccess)
	}

	if i == j {
		return d.data[i]
	}
	return 0
}

// T returns the matrix transposed.  In the case of a DIA (DIAgonal) sparse matrix this method
// simply returns the receiver as the matrix is perfectly symmetrical and transposing has no effect.
func (d *DIA) T() mat.Matrix {
	return d
}

// NNZ returns the Number of Non Zero elements in the sparse matrix.
func (d *DIA) NNZ() int {
	return d.m
}

// Diagonal returns the diagonal values of the matrix from top left to bottom right.
// The values are returned as a slice backed by the same array as backing the receiver
// so changes to values in the returned slice will be reflected in the receiver.
func (d *DIA) Diagonal() []float64 {
	return d.data
}

// RowView slices the matrix and returns a Vector containing a copy of elements
// of row i.
func (d *DIA) RowView(i int) mat.Vector {
	return mat.NewVecDense(d.m, d.slice(i))
}

// ColView slices the matrix and returns a Vector containing a copy of elements
// of column j.
func (d *DIA) ColView(j int) mat.Vector {
	return mat.NewVecDense(d.m, d.slice(j))
}

// RawRowView returns a slice representing row i of the matrix.  This is a copy
// of the data within the matrix and does not share underlying storage.
func (d *DIA) RawRowView(i int) []float64 {
	return d.slice(i)
}

// RawColView returns a slice representing col i of the matrix.  This is a copy
// of the data within the matrix and does not share underlying storage.
func (d *DIA) RawColView(j int) []float64 {
	return d.slice(j)
}

// nativeSlice slices the DIAgonal matrix.
func (d *DIA) slice(i int) []float64 {
	if i >= d.m || i < 0 {
		panic(mat.ErrRowAccess)
	}

	slice := make([]float64, d.m)

	slice[i] = d.data[i]

	return slice
}
