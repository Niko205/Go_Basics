package listen_schleifen

// Hier siehts du wie die Terminal Ausgabe der Übung aussehen soll
func ExampleUebung() {
	Uebung()

	// Output:
	// 7210810832108116
	// <<>>><<>>>
	// Hallo Welt
}

// Mögliche Lösung, erst danach schauen oder wenn du es nicht hinbekommst
/*
	1.
	for i := 0; i < len(intlist); i++ {
		if intlist[i]%2 == 0 {
			fmt.Print(intlist[i])
		}
	}
	2.

	fmt.Println()
	for i := 0; i < len(intlist); i++ {
		if intlist[i] == 100 {
			fmt.Print("=")
		} else if intlist[i] < 100 {
			fmt.Print("<")
		} else {
			fmt.Print(">")
		}
	}

	3.
	bytelist := make([]byte, len(intlist))

	4.
	for i := 0; i < len(bytelist); i++ {
		bytelist[i] = byte(intlist[i])
	}

	5.
	fmt.Println()
	for i := 0; i < len(bytelist); i++ {
		fmt.Print(string(bytelist[i]))
	}
*/
