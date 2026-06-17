package findcommoncharacters

func commonChars(words []string) []string {
	minFreq := charFreq(words[0])
	for _, w := range words[1:] {
		wFreq := charFreq(w)
		for i := range minFreq {
			if wFreq[i] < minFreq[i] {
				minFreq[i] = wFreq[i]
			}
		}
	}

	result := []string{}
	for i, count := range minFreq {
		for range count {
			result = append(result, string(rune('a'+i)))
		}
	}
	return result
}

func charFreq(w string) [26]int {
	var f [26]int
	for _, r := range w {
		f[r-'a']++
	}
	return f
}
