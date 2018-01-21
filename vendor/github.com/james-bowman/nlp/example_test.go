package nlp

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func Example() {
	testCorpus := []string{
		"The quick brown fox jumped over the lazy dog",
		"hey diddle diddle, the cat and the fiddle",
		"the cow jumped over the moon",
		"the little dog laughed to see such fun",
		"and the dish ran away with the spoon",
	}

	query := "the brown fox ran around the dog"

	vectoriser := NewCountVectoriser(true)
	transformer := NewTfidfTransformer()

	// set k (the number of dimensions following truncation) to 4
	reducer := NewTruncatedSVD(4)

	// Transform the corpus into an LSI fitting the model to the documents in the process
	matrix, _ := vectoriser.FitTransform(testCorpus...)
	matrix, _ = transformer.FitTransform(matrix)
	lsi, _ := reducer.FitTransform(matrix)

	// run the query through the same pipeline that was fitted to the corpus and
	// to project it into the same dimensional space
	matrix, _ = vectoriser.Transform(query)
	matrix, _ = transformer.Transform(matrix)
	queryVector, _ := reducer.Transform(matrix)

	// iterate over document feature vectors (columns) in the LSI and compare with the
	// query vector for similarity.  Similarity is determined by the difference between
	// the angles of the vectors known as the cosine similarity
	highestSimilarity := -1.0
	var matched int
	_, docs := lsi.Dims()
	for i := 0; i < docs; i++ {
		similarity := CosineSimilarity(queryVector.(mat.ColViewer).ColView(0), lsi.(mat.ColViewer).ColView(i))
		if similarity > highestSimilarity {
			matched = i
			highestSimilarity = similarity
		}
	}

	fmt.Printf("Matched '%s'", testCorpus[matched])
	// Output: Matched 'The quick brown fox jumped over the lazy dog'
}
