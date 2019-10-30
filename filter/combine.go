package filter

import (
	"acrcloud-scan-tool-golang/models"
	"fmt"
)

func CombineFilter(results []models.Result) []models.Result {

	for _, v := range results {
		fmt.Println(v.Title)
	}

	return results
}
