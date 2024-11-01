package main

import (
	funktionen "basics/6.Funktionen"
	"fmt"
)

func main() {
	// In diesem Kapitel geht es um Packages, ihren Nutzen und wie man sie verwendet
	// Packages sind eine Sammlung von Funktionen, die in einer Datei oder in mehreren Dateien zusammengefasst sind
	// Diese Funktionen ergänzen die Standardbibliothek von Go
	// Benötigt man also eine Funktion, die nicht in der Standardbibliothek vorhanden, kann man sich ein passendes Package suchen
	// Dafür sucht man am besten auf https://pkg.go.dev/ und gibt ein passendes Stichwort ein

	// Ein Package wird in Go mit dem Schlüsselwort "import" eingebunden, dies geschieht am Anfang der Datei
	// Beispiel: import "fmt" -> ein Package, welches wir in den vorherigen Kapiteln schon verwendet haben
	// Dieses ermöglicht uns die Ausgabe von Text auf der Konsole, was wir in den meisten unserer Programme benötigen
	// Kennt man den Namen des Packages, kann dieses im Code direkt verwenden und das Package wird automatisch eingebunden

	fmt.Println("Automatische eingebunden")
	// Falls dies mal nicht funktioniert muss es manuell einbinden, wie man es in Zeile 3 sieht

	// Weitere Packages die wir gebrauchen können:
	// - math: erweiterte Mathematische Funktionen für komplexe Berechnungen
	// - strings: Funktionen zum bearbeiten und zerlegen von Strings

	// Um Funktionen aus einem Package zu verwenden, schreibt man den Namen des Packages gefolgt von einem Punkt
	// Daraudhin werden die Funktionen des Packages vorgeschlagen. Man kann dann eine Auswählen oder den Namen der Funktion eingeben

	// Man kann auch eigene Packages erstellen, wie man daran erkennt das jede unserer Dateien mit "package" beginnt
	// Möchte mann also eigene Funktionen erstellen und in anderen Programmen verwenden, kann man diese in einem Package zusammenfassen
	// Wenn man diese Funktionen in einem anderen Programm verwenden möchte, muss man den Ordner des Packages importieren
	funktionen.Uebung([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	// Hier importieren wir das Package "funktionen" aus Kapitel 6 und rufen die Funktion "Uebung" auf

	// Dieses Kapitel endet wieder ohne Übung
}
