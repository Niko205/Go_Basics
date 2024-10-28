package main

import "fmt"

// In diesem kapitel geht es um den Aufbau und Nutzen von Funktionen
// In den letzten Kapitel haben wir beraits Funktionen erstellt wie die main Funktion, oder die Test Funktionen
// Und zudem auch Funktionen benutzt wie fmt.Println() oder len()
// Nun wird erleutert wie genau man Funktionen erstellt und nutzt
// Schauen wir uns dazu erstmal den Funktionskopf der main Funktion an
// Er start mit dem Schlüsselwort func gefolgt von dem Namen der Funktion, in diesem Fall main
// Danach kommt eine Klammer, in welchen festgelgt ob und welche Variablen der Funktoin übergeben werden sollen
//Zum Schluss öffnet man eine geschweifte Klammer in welcher der Funktionskörper/-ablauf steht
func main() {
	// Wie berteits erklärt ist die main Funktion die Funktion die beim Start des Programms aufgerufen wird
	// In dieser Funktion wird der Ablauf des Programms festgelegt
	// Wird dieser unübersichtlich oder Passagen des Codes wiederholt, lohnt es sich Code in Funktionen auszulagern

	// Zudem kann man Funktionen in zwei Arten unterscheiden, welche mit und welche ohne Rückgabewert
	// Beide Typen haben wir schon mehrmals benutzt

	// Beispiel fmt.Println() gibt keinen Wert zurück, erzeugt jediglich eine Ausgabe im Terminal
	fmt.Println("Hallo Welt")
	// Die Funktion direkt aufgerufen und der String "Hallo Welt" übergeben
	// Solche Funktionen führen meistens eine Aktion aus, wie eine Ausgabe und geben keinen Wert zurück

	// Beispiel len() gibt einen Wert zurück, die Länge einer Liste, Strings, etc.
	intvar1 := len("Hallo Welt")
	// Diese Funktion kann nicht direkt aufgerufen werden, da sie einen Wert zurückgibt
	// Darum speichern wir den Wert in einer Variable, können ihn ausgeben oder auch in einer Condition nutzen (Schleifen, If-Abfragen)
	// Solche Funktionen können benutzt werden um Werte zu berechnen, oder zu überprüfen und geben das Ergebnis zurück

	// Das Aufrufen einer Funktion erfolgt immer nach den selben Schema
	// Befindet sich die Funktion in einem anderen Package, muss dieses importiert werden und vor dem Funktionsaufruf angegeben werden
	// Darauf schreibt man den Funktionsnamen, öffnet eine Klammer. Dies ist der einzige Verwendungszweck der nromalen Klammer in Go
	// In dieser Klammer können Variablen übergeben werden, wenn die Funktion welche erwartet

	// Nun legen wir unsere Erste Funktion an
	// Angenommen man müsst in einem Programm öfters zwei Variablen addieren, lohnt sich dies in eine Funktion auszulagern
	// Dammit man nicht jedes man var1 + var2 schreiben muss, sondern nur die Funktion aufrufen muss
	intvar1 = add(5, 3)  // Die Funktion add wird aufgerufen und die Werte 5 und 3 übergeben und das Ergebnis in intvar1 gespeichert
	fmt.Println(intvar1) // Nun geben wir das Ergebnis noch aus, damit das Programm auch läuft
	// Nun existiert die Funkion add nicht automatisch, sondern muss erstellt werden, dies macht man außerhalb der main Funktion

}

// Hier wird die Funktion add erstellt
// Der Funktionskopf besteht aus dem Schlüsselwort func, gefolgt von dem Namen der Funktion, in diesem Fall add
// Danach öfnnen wir eine Klammer in welcher die Variablen übergeben werden, in diesem Fall zwei Integer
// Nach der Klammer öffnen wir eine geschweifte Klammer in welcher der Funktionskörper steht
func add(var1 int, var2 int) int {
	// In dieser Funktion wird die Summe der beiden Variablen berechnet und zurückgegeben
	return var1 + var2
	// Das Schlüsselwort return beendet die Funktion und gibt den Wert dahinter zurück
	// Welchen wir in Zeile 40 in intvar1 speichern
}

// Damit kommen wir zum Ende des Kapitels Funktionen, in der uebung.go Datei gibt es noch eine Übung zu diesem Thema
