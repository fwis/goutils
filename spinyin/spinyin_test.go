package spinyin

import (
	"fmt"
	"testing"
)

func TestSPinyin(t *testing.T) {
	fmt.Printf("中国人民=%s\n", SPinyin([]rune("中国人民")))
	fmt.Printf("abcd12 34=%s\n", SPinyin([]rune("abcd1234")))
	fmt.Printf("，。　＆×a=%s\n", SPinyin([]rune("，。　＆×")))
}
