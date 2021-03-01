package main
import(
	"fmt"
)
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for _, s := range strs[1:]{
		if len(s) == 0 {
			return ""
		}
		for i, _ := range(prefix){
			if i >= len(s) || prefix[i] != s[i]{
				if i == 0 {
					return ""
				}
				prefix = s[:i]				
				break
			}
		}
	}
	return prefix
}

func main() {
	strs := [] string {"flower", "flowertest"}
	answer := longestCommonPrefix(strs)
	fmt.Println(answer)
}