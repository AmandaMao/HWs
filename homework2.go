package song

import (
	"fmt"
)

const rhyme1 = "There was an old lady who swallowed a "
const rhyme2 = "I don't know why she swallowed a "
const rhyme3 = " - perhaps she'll die!"
const rhyme4 = "Amazing that swallow a "
const rhyme5 = "She swallowed the "
const rhyme6 = " to catch the "
const rhyme7 = "...She's dead, of course!"

func main() {
	animals := []string{"fly", "spider", "bird", "cat", "dog", "cow", "horse"}
	rhyme := song(animals)
	fmt.Println(rhyme)
}

func song(objects []string) string {
	count := 0
	rhyme := ""
	for key, _ := range objects {
		rhyme = rhyme + rhyme1 + objects[key] + ",\n"
		if objectsInTheMiddle(count, objects) {
			rhyme = rhyme + rhyme4 + objects[key] + ",\n"
			i := count
			for i > 0 {
				rhyme = rhyme + rhyme5 + objects[i] + rhyme6 + objects[i-1] + ",\n"
				i = i - 1
			}
		} else if objectsInTheEnd(count, objects) {
			rhyme = rhyme + rhyme7
			return rhyme
		}
		rhyme = rhyme + rhyme2 + objects[0] + rhyme3 + "\n"
		count++
	}
	return rhyme
}

func objectsInTheMiddle(count int, objects []string) bool {
	if count > 0 && count < len(objects)-1 {
		return true
	} else {
		return false
	}
}

func objectsInTheEnd(count int, objects []string) bool {
	if count == len(objects)-1 {
		return true
	} else {
		return false
	}
}
