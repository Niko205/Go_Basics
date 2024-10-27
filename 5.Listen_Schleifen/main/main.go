package main

import "fmt"

func main() {
	// In diesem Kapitel geht es um Listen und Schleifen
	// In Listen kann man eine Kette von Daten speichern, jede davon ist an einer bestimmten Stelle (Index)
	// Dabei hat eine List ähnlich wie Variablen einen Namen und Datentyp. Dazu kommt jedoch noch eine feste Länge
	// Man legt sie auch ähnlich wie eine Variable an, wobei man hier ebenfalls zwei Varianten hat
	// Die erste Variante ist: var "Name" []Datentyp
	var stringlist1 []string                // ist gelb markiert weil es besser geht
	stringlist1 = []string{"Hallo", "Welt"} // So kann man der Liste dann die in der geschweiften Klammer enthaltenen Werte zuteilen

	// Die schnellere Variante, später kommt noch eine dritte Variant für Sonderfälle
	stringlist2 := []string{"H", "a", "l", "l", "o"} // Diese Liste hat die Länge 5
	fmt.Println(stringlist1)                         // Mann kann die List zwar auf diese Art ausgeben, jedoch wird der Inhalt dann in eckigen Klammern ausgegben

	// Wenn man die einzelnen Werte der Liste ausgeben will, kann man auch bestimmte Stellen der Liste auswählen
	// Dafür schauen wir uns erstmal am Beispiel der stringlist2 an, wie Werte in einer Liste gespeichert werden
	// Werte: "H" "a" "l" "l" "o"
	// Index:  0   1   2   3   4
	// Um nun auf einen bestimmten Wert zuzugreifen, schreibt man hinter den namen der List eine eckige Klammer und den gewollten Index
	// Möchten wir nun von stringlsit2 den dritten Buchstaben (l) ausgeben, müssen wir den Index 2 in die klammer schreiben
	fmt.Println(stringlist2[2])  // Bei der ersten Ausgabe haben wir "Println" noch eine List des Typs string übergeben, nun übergeben wir einen String
	stringvar1 := stringlist2[2] // Hoverd man über stringvar sieht man, dass die Variable vom Typ string ist. Somit entspricht stringlist2[3] einem string

	// Möchte man einen bestimmten Wert in einer List ändern, geht das auf die selbe Weise, indem man in der eckigen Klammer den gewollten Index angibt
	stringlist2[4] = stringvar1 // Fügen das l aus stringvar am Index 4 in der Liste ein, überschreibt das "o"
	fmt.Println(stringlist2)    // "Halll" wird im Terminal ausgegeben

	// In Gewissen fällen, wenn man zum Beispiel eine List anlegt, die die Länge einer anderen haben soll, muss man make() verwenden
	intlist1 := make([]int, len(stringlist2)) // len(list) gibt einem die Länge einer Liste
	// Hier haben wir jetzt eine Liste vom Datentyp int erstellt mit der Länge der stringlist2, jedoch hat sie keine Werte bzw. den Standartwert für int (0)
	fmt.Println(intlist1) // Gibt fünf 0er aus, legt man keinen Wert für einen int fest, hat er den Wert 0
	// Man kann Listen von jedem Datentyp erstellen

	// Jetzt geht es noch um Schleifen, mit deren Hilfe man gut durch Listen gehen kann um Daten auszulesen oder zu speichern
	// Wir wollen in der intlist1 die zahlen 0-4 speichern, dafür eignet sich eine Schleife ideal
	// Eine Standart for-Schleife legt man wie folgt an
	// i := 0 <- ist eine Zählvariable, hier startet sie bei 0
	// i < len(intlist1) <- Schleife läuft solang i kleiner als Länge der Liste ist, damit wir jeder Stelle einen Wert geben können
	// i++ <- zählt i nach jedem Schleifenfurchlauf hoch, damit wir quasi die Liste entlang laufen
	for i := 0; i < len(intlist1); i++ {
		intlist1[i] = i // Speichert am Index i den Wert i in der Liste (hier aufsteigend 0, 1, 2, 3, 4)
	}
	fmt.Println(intlist1) // Gibt die Liste aus: [0 1 2 3 4]

	// Nach dem selben Prinziep kann man auch die einzelnen Werte einer Liste ausgeben
	// Dies verbinden wir noch mit der Besonderheit des Datentyp Strings
	// Auch wenn wir einen String nicht selber als Liste anlegen, ist er im enteffekt eine Liste vom Datentyp byte (Zeichen)
	// Ich zeige nun die Ausgabe am Beispier eines strings, lässt sich jedoch genau so mit einer Liste machen (sind ja prinzipiell auch beides eine Liste)
	stringvar2 := "Hallo"
	//Wieder haben wir den Zähler i; wieder läuft die Schleife solange i kleiner Länge von stringvar2 ist; wieder zählt i mit jedem durchlauf hoch
	for i := 0; i < len(stringvar2); i++ { // len() lässt sich auf strings genauso wie auf Listen anwenden
		//fmt.Print(stringvar2[i]) würde den byte den stringvar2[i] liefert im Terminal ausgeben, jedoch währe dies der Zahlenwert des bytes
		fmt.Print(string(stringvar2[i])) // Darum muss string() benutzt werden um den byte wieder in einen string umzuwandeln und somit lesbar zu machen
	} // Die Ausgabe währe ansonsten: 7297108108111 gewesem, wenn man in der ASCII-Tabelle nachscheut entspricht das Hallo
	fmt.Println() // Erzeugt einen Zeilenumbruch weil wir in der Schleife keinen erzeugen konnten
	// INFO: man kann die bytes eines String auf diese Weise nur auslesen, nicht verändern!!

	// Es gibt noch komplexere Weisen Schleifen zu benutzten, welche in einem eigenen Kapitel drankommen werden

	// Jetzt kannst du zum Test übergehen

}
