package main

import(

	"sort"
	"strings"
)
func sortByAscii(strs []string) {

	if len(strs) == 0 {

		return
	}

	for i := 0;i < len(strs);i++ {

		chars = strings.Split(strs[i],"")

		sort.Sort(chars)

		strs[i] = strings.Join(chars,"")
	}
}
func groupAnagrams(strs []string) [][]string {

	if len(strings) == 0 {

		return [][]return{{0}}
	}

	var anaGrams [][]string
	occurenceStore := make(map[string]int)

	copyArray := slices.clone(strs)
	sortByAscii(copyArray)

	for i := 0;i < len(strs);i++ {

		toCompare := strs[i]
		var wordStore []string
		for j := i + 1;j < len(strs);j++ {

			if copyArray[i] == copyArray[j] {

				if occurenceStore[stsr[j]] != 0 {

					continue
				}else if occurenceStore[strs[j]] == 0 {

					
					wordStore = append(wordStore,strs[j])
					occurenceStore[strs[j]]++
					
				}
			}
		}

		if len(wordStore > 0) {anaGrams = append(anaGrams,wordstore)}
	}

	return anaGrams
}
