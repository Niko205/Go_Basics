package byterechnen

import "fmt"

// Hier kannst du deinen Funktionen testen
func ExampleUebungBuchstabe() {
	fmt.Println(UebungBuchstabe('A'))
	fmt.Println(UebungBuchstabe('a'))
	fmt.Println(UebungBuchstabe('1'))

	//Output:
	// a
	// A
	//

}

func ExampleUebungZahl() {
	fmt.Println(UebungZahl('0'))
	fmt.Println(UebungZahl('1'))
	fmt.Println(UebungZahl('a'))

	//Output:
	// 0
	// 1
	// -1
}

// Mögliche Lösung, erst danach schauen oder wenn du es nicht hinbekommst
/*
	1.
	if x >= 'A' && x <= 'Z' {
		return string(x + 32)
	} else if x >= 'a' && x <= 'z' {
		return string(x - 32)

	}
	return ""

	2.
	if x >= '0' && x <= '9' {
		return int(x - 48)
	}
	return -1
*/
