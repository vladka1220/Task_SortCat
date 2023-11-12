package main

import "sort"

func SortByLength(countryDict map[string][]string) map[string][]string {
	sortedDict := make(map[string][]string)
	for country, breeds := range countryDict {
		sort.Slice(breeds, func(i, j int) bool {
			return len(breeds[i]) < len(breeds[j])
		})
		sortedDict[country] = breeds
	}
	return sortedDict
}
