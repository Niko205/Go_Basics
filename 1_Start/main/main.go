package main

import "fmt"

// Das Main-Package und die Main-Funktion
// Nur diese Funktion aus diesem Package lässt sich im Terminal starten
// In der Main-Funktion lässt sich der Ablauf eines Programms realisieren
// Dabei schreibt man den Code der ausgeführt werden soll in die Funktion und kann auch andere Funktionen aufrufen
// Hier machen wir ein Klassisches Hello World zum Anfang
func main() {
	// Hier rufen wir die Funktion "Print" vom Package "fmt" auf, welche Dinge auf der Konsole ausgeben kann, um sie benutzten zu kommen muss oben "fmt" importiert werden
	// Hoverd man über das "Print" wird die Klassenbeschreibung angezeigt
	// Diese zeigt das irgendein Datentyp der Klasse übergeben werden kann, d.h. in die Klammern der Funktion geschrieben wird
	// In diesem Beispiel übergeben wir den String "Hello World", wenn man Strings schtreibt müssen diese immer in Anführungszeichen sein
	fmt.Print("Hello World")

	// Das \n erzeugt einen Zeilenumbruch auf der Konsole, sprich spätere Ausgaben starten in der Zeile darunter
	fmt.Print("\n")

	// Diesen Effekt hat auch die Funktion "Println" der Klasse "fmt", diese erzeugt nach der Ausgabe einen Zeilenumruch
	// Hier habe ich mal eine Zahl (int) der Funktion übergeben, um zu zeigen dass man ihr verschiedene Datentypen übergeben kann
	fmt.Println(1)

	//Man kann ebenfalls Variablen einer Funktion übergeben, diese erstellt man mit:
	// var "Name der Variable" Datentyp:
	// Hier die Variable "helloworld" vom Typ string, wird gelb markiert weil es eine einfachere Möglichkeit gibt, welche im nächsten Abschnitt erklärt wird
	var helloworld string

	//Nun haben wir eine Variable, jetzt müssen wir ihr nurnoch einen Wert zuteilen
	// Dies geht indem wir den Namen der Variable wiederverwenden und EIN "=" dahinter schreiben (Es gibt auch noch ":=" und "==" welche jedoch für komplett andere Zwecke verwendet werden)
	// Daraufhin kann man den Wert der Variable festlegen, da es sich hierbei wieder um einen String handelt muss dieser in Anführungszeichen geschrieben werden
	helloworld = "Hello World"

	// Unbenutzte Variablen gelten in Go als Compilerfehler, sprich eine Variable muss im späteren Verlauf des Codes verwendet werden
	// Mann kann eine Variable benutzen, indem man ihren Namen wiederverwendet, wie eben beim Wertzuteilen
	// In diesem Beispiel übergebe ich sie der Funktion "Println"
	// man Beachte: das helloworld ist nicht in Anführungszeichen, daran erkent man das es sich um eine Methode und nicht um einen String-Text handelt
	fmt.Println(helloworld)

	// Führt man nun die Main-Funktion mittels: "go run ."" im Integrated Terminal der main-Datei aus (rechtklick auf Datei, dann auswählen)
	// So erhalten wir folgende Terminelausgabe:
	// Hello World
	// 1
	// Hello World

	// In den nächsten Zwei Dateien kannst du Üben, Dinge auf der Konsole auszugeben
	// Öffne dafür die Dateien uebung.go und
}
