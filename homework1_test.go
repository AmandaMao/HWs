package selector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNameSpaceSelector(t *testing.T) {
	Source = []Score{
		Score{1, "Tom", "One", 60},
		Score{2, "Jerry", "One", 70},
		Score{3, "Woof", "Two", 90},
		Score{4, "Mary", "Two", 85},
		Score{5, "Billy", "Three", 35},
		Score{6, "Leo", "Two", 66},
		Score{7, "Louis Vuitton", "Two", 25},
		Score{8, "Lucy", "One", 75},
		Score{9, "Alex", "Two", 15},
		Score{10, "Christian Dior", "One", 40}}
	groupSelectByNameSpace := SelectNameSpace(Source)
	assert.NotNil(t, groupSelectByNameSpace)
	for _, v := range groupSelectByNameSpace {
		nameSpaceLargerThan10 := len(v)
		assert.Equal(t, 2, nameSpaceLargerThan10)
	}
}

func TestClassSelector(t *testing.T) {

	groupSelectByClass := SelectClass(Source)
	assert.NotNil(t, groupSelectByClass)
	for _, v := range groupSelectByClass {
		classOne := len(v)
		assert.Equal(t, 4, classOne)
	}
}

func TestScoreSelector(t *testing.T) {

	groupSelectByScore := SelectScore(Source)
	assert.NotNil(t, groupSelectByScore)
	for _, v := range groupSelectByScore {
		scoreLargerThan59 := len(v)
		assert.Equal(t, 6, scoreLargerThan59)
	}
}
