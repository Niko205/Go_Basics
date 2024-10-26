package main

import "fmt"

func main() {
	// In diesem Kapitel geht es um Operatoren
	// Es gibt Rechen- und Vergleichsoperatoren, wir schauen uns erstmal ersteres an
	// Wie es der Name schon sagt kann man Rechenoperatoren zum rechnen benutzen
	// Folgende Rechenoperatoren stehen uns zur verfügung:
	// + <- zum Addieren von Werten
	// - <- zum Subtrahiern von Werten
	// * <- zum Multiplizieren von Werten
	// / <- zum Dividieren
	// % <- der Modulo berechnet den Rest einer Divison
	// Diese Operatoren lassen sich in ganz verschiedenen Situationen einsetzten
	// Hier werden mal ein paar gezeigt:

	intvar1 := 2 + 3             // Beim erstellen der Variable zwei Werte addieren
	intvar2 := intvar1 - 2       // Auch auf bereits existierende Variablen anwendbar
	intvar3 := intvar1 * intvar2 // Oder zwischen zwei Variablen anwendbar

	// Bei der Diviosion vom Datentyp int muss man beachten, dass es keine Kommastellen gibt, sprich immer eine gerade Zahl rauskommt
	// Um dies zu zeigen benutzte ich nun den Divisionsoperator in der fmt.Println() Funktion:
	fmt.Println(intvar3 / 3) // = 5, stimmt weil intvar3 (=15) geteilt durch 3 eine gerade Zahl ergibt
	fmt.Println(intvar3 / 4) // = 3, stimmt nicht, es wird auch nicht gerundet, sondern nur geguckt wie oft die eine Zahl in die andere passt
	fmt.Println(intvar3 % 4) // = 3, hier können wir uns auch mal den Modulo anschauen, bei der Division 15 / 4 bleibt ein Rest von 3 (15-3*4=3)

	// Wenn man bei der Division Kommastellen will, muss man den Datentyp float benutzten
	floatvar1 := 4.0 / 2.5 // Damit auch ein richtiger Kommawert rauskommt, müssen beide Zahlen von diesem Datentyp sein. Beachte: das Komma ist ein "."
	fmt.Println(floatvar1) // Nun wird auch im Terminal eine Kommazahl ausgegeben

	// Möchte man auf eine Variable etwas draufaddiern, kann man nach dem "=" einfach nochmal die Variable aufrufen:
	intvar1 = intvar1 + 1
	intvar1 += 1 // Hat den selben effekt wie darüber
	intvar1++    // Erhöt die Variable ebenfalls um +1
	// Alle drei auch durch - austauschbar

	// Zudem lassen sich Strings mit den Rechenoperator "+" zusammenfügen
	stringvar1 := "Hallo"
	stringvar2 := "Welt"
	stringvar3 := stringvar1 + stringvar2 // Hängt die Strings direkt aneinander, ohne Leerzeichen
	fmt.Println(stringvar3)
	fmt.Println(stringvar1, stringvar2) // Andere Option, fügt zudem ein Leerzeichen ein

	// Man kann Rechenoperatoren ebenfalls auf den Datentyp Byte anwenden, dieses Thema ist jedoch etwas größer und bekommt ein eigenes Kapitel

	// Zuletzt wird noch die Funktion "Printf" vom package "fmt" vorgestellt
	// Diese erzeugt eine Formatierte Ausgabe, sprich fügt eine Variable and einer bestimmten Stelle in der Ausgabe ein
	// Man ruft sie ganz normal mittel "fmt.Printf()" auf, jedoch muss man ihr nun zwei Werte in der Klammer übergeben
	// Inhlat der Klammer: ("Text der ausgegebn werden soll" , Name der Variable)
	// Nun kann man an einer Beliebigen stelle im Text den platzhalter "%" eingestzt werden
	// Dieser legt fest an welcher Stelle der Wert der Variable eingefügt werden soll
	// Hinter das "%" muss der Anfangsbuchstabe des gewollten Datentyps geschrieben werden:
	// %d <- dezimal für Zahlen, %s <- string, %b <- byte
	fmt.Printf("Den Wert: %d hat die Variable intvar3", intvar3)

	// In diesem Kapitel gibt es wieder einen Test, öffne nun also die Dateien uebung.go und uebung_test.go
}
