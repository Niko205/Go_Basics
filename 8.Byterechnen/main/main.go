package main

import "fmt"

func main() {
	// In diesem Kapitel geht es um das Rechnnen mit dem Datentyp byte
	// Wieschonmal erwähnt, hat der Datentyp Byte eine Besonderheit
	// Er speicher nicht nur Zeichen, sondern eine zu dem Zeichen zugehörige Zahl
	// Dies lässt sich in der ASCII-Tabelle nachschauen, welche ich mit in Kapitel 8 gepackt habe
	// Schaut man in die Tabelle, sieht man, dass das zeichen 'A' die Zahl 65 hat
	// Addiert man nun z.B. 'A' mit 1, erhält man das Zeichen 'B'
	if 'A'+1 == 'B' {
		fmt.Println("A + 1 = B")
	}
	// Die Zeichen sind zudem alle geordnet, spricht A-Z, a-z und 0-9 folgen einander
	// Mit diesem Prinziep kann man ganz unterschiedliche Dinge machen
	// Hier ein Beispiele:
	// Kleinbuchstaben in Großbuchstaben umwandeln
	// Der Großbuchstabe 'A' hat den Wert 65, der Kleinbuchstabe 'a' den Wert 97
	// Die Differenz zwischen Groß- und Kleinbuchstaben beträgt somit 32
	// Addiert man also einen Großbuchstaben mit 32, erhält man den zugehörigen Kleinbuchstaben, und umgedreht
	fmt.Println(string('A' + 32)) // a muss zum String gecastet werden, da sonst eine Zahl ausgegeben wird
	fmt.Println(string('a' - 32)) // A -||-

	// Zudem kann man das Zeichen auch als Zahl ausgeben
	// Hier ein Beispiel:
	// Die Zeichen '0' bis '9' haben die Werte 48 bis 57
	// Möchte man also diese Zeichen als int ausgeben, kann man den Byte-Wert einfach um 48 subtrahieren
	fmt.Println('0' - 48) // 0
	fmt.Println('1' - 48) // 1

	// Wenn man weiß wie man mit Bytes rechnen kann, kann man dies für ganz verschiedene Dinge nutzen
	// Dies könnt ihr nun in den Übungen ausprobieren, da wir am Ende des Kapitels sind
}
