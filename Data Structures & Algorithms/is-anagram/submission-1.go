func isAnagram(s string, t string) bool {

freqS := make(map[rune]int) 
freqT := make(map[rune]int) 

for _, value := range s {
	freqS[value] += 1
}
for _, value := range t {
	freqT[value] += 1
}

// printFrequency(freqS, freqT)

if len(freqS) != len(freqT) {
	return false 
}

for keyS, countS := range freqS{
		if freqT[keyS] != countS{
		return false
	} 
}

return true
}


func printFrequency(freqS, freqT map[rune]int) {
	fmt.Printf("string S: LEN: %d | Frequency \n ", len(freqS))
	for char, count := range freqS {
		fmt.Printf("%c: %d ", char, count)
	}

	fmt.Printf("\n\nstring T: LEN: %d | Frequency \n ", len(freqT))
	for char, count := range freqT {
		fmt.Printf("%c: %d ", char, count)
	}
}