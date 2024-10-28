package funktionen

// Hier kannst du dir anschauen wie die Ausgabe aussehen soll
func ExampleUebung() {
	Uebung([]int{2, 9, 7, 4, 6, 3, 8, 1})

	//Output:
	// truefalsefalsetruetruefalsetruefalse
	// 418148126162

}

// Eine mögliche Lösung für die Übung:
/*
1.
func gerade(zahl int) bool {
	return zahl%2 == 0
}

2.
for i := 0; i < len(intlsit); i++ {
	fmt.Print(gerade(intlsit[i]))
}

3.
func multiplizieren(zahl int) int {
	return zahl * 2
}

4.
for i := 0; i < len(intlsit); i++ {
	intlsit[i] = multiplizieren(intlsit[i])
}

5.
func ausgabe(intlist []int) {
	fmt.Println()
	for i := 0; i < len(intlist); i++ {
		fmt.Print(intlist[i])
	}
}

6.
ausgabe(intlsit)
*/
