// Your friend is typing his name into a keyboard. Sometimes, when typing a character c, the key might get long pressed, and the character will be typed 1 or more times.
//
// You examine the typed characters of the keyboard. Return True if it is possible that it was your friends name, with some characters (possibly none) being long pressed.
//
// Input: name = "alex", typed = "aaleex"
// Output: true
// Input: name = "saeed", typed = "ssaaedd"
// Output: false

package longpressedname

// O(n+m) time, O(1) space
func isLongPressedName(name string, typed string) bool {
	np, tp := 0, 0
	for tp < len(typed) {
		if np < len(name) && name[np] == typed[tp] {
			np++
			tp++
		} else if tp > 0 && typed[tp-1] == typed[tp] {
			tp++
		} else {
			return false
		}
	}
	return np == len(name)
}
