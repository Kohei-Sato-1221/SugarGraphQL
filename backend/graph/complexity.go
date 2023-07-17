package graph

import (
	"fmt"

	"github.com/Kohei-Sato-1221/SugarGraphQL/backend/generated"
)

func ComplexityConfig() generated.ComplexityRoot {
	var c generated.ComplexityRoot

	c.Repository.Issues = func(childComplexity int, after *string, before *string, first *int, last *int) int {
		var cnt int
		switch {
		case first != nil && last != nil:
			if *first < *last {
				cnt = *last
			} else {
				cnt = *first
			}
		case first != nil && last == nil:
			cnt = *first
		case first == nil && last != nil:
			cnt = *last
		default:
			cnt = 1
		}
		fmt.Printf("ComplexityConfig:Issues complexity is %v", cnt*childComplexity)
		return cnt * childComplexity
	}
	return c
}
