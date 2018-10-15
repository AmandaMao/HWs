package selector

import (
	"fmt"
)

var Source []Score

type Score struct {
	ID    int
	Name  string
	Class string
	Score int
}

func Select(source []Score, selector func() string, condition func(s Score) bool) map[string][]Score {
	resultMap := make(map[string][]Score)
	for _, v := range source {
		key := selector()
		if condition(v) {
			resultMap[key] = append(resultMap[key], v)
		}
	}
	fmt.Println(resultMap)
	return resultMap
}

func SelectNameSpace(source []Score) map[string][]Score {
	return Select(source, func() string { return "NameSpace > 10" }, func(s Score) bool {
		if len(s.Name) > 10 {
			return true
		}
		return false
	})
}

func SelectClass(source []Score) map[string][]Score {
	return Select(source, func() string { return "Class One" }, func(s Score) bool {
		if s.Class == "One" {
			return true
		}
		return false
	})
}

func SelectScore(source []Score) map[string][]Score {
	return Select(source, func() string { return "Score > 59" }, func(s Score) bool {
		if s.Score > 59 {
			return true
		}
		return false
	})
}

// func main() {
// 	Source = []Score{
// 		Score{1, "Tom", "One", 60},
// 		Score{2, "Jerry", "One", 70},
// 		Score{3, "Woof", "Two", 90},
// 		Score{4, "Mary", "Two", 85},
// 		Score{5, "Billy", "Three", 35},
// 		Score{6, "Leo", "Two", 66},
// 		Score{7, "Louis Vuitton", "Two", 25},
// 		Score{8, "Lucy", "One", 75},
// 		Score{9, "Alex", "Two", 15},
// 		Score{10, "Christian Dior", "One", 40}}

// 	SelectNameSpace(Source)
// 	SelectClass(Source)
// 	SelectScore(Source)
// }
