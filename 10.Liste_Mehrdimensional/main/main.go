package main

import "fmt"

func main() {
	// Vor ein paar Kapiteln haben wir Listen kennengelernt
	// Listen können auch mehrdimensional sein, indem man eine Liste von Listen erstellt
	// Bsp:
	var list = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	// Dabei ist jede Zeile eine Liste vom Typ int, welche in der Liste list gespeichert ist
	// Mann kann nun auf zwei Weisen auf die Werte zugreifen
	// 1. mit list[0] erhält man die erste Zeile, somit eine eindiemensionale Liste vom Typ int
	var list1 = list[0] // Hoverd man über list1 sieht man den Datenyp der Variable, hier []int
	fmt.Println(list1)  // Gibt die erste Zeile aus: [1 2 3]
	fmt.Println()

	// 2. mit list[0][0] erhält man den ersten Wert der ersten Zeile
	var value = list[0][0] // Hoverd man über value sieht man den Datenyp der Variable, hier int

	// Also gibt einen die erste eckige Klammer die Zeile an und die zweite den Wert in der Zeile
	// Nach diesem Prinziep kann man wie bei eindimensionalen Listen durch iterieren
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list[i]); j++ {
			fmt.Print(list[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()
	// Diese doppelte-for-Schleife gibt alle Werte der Liste aus
	// Dabei geht die äußere Schleife die Zeilen durch und die innere die Werte in den Zeilen

	// Nach dem selben Prinziep kann man auch Werte suchen oder verändern
	// Bsp: wir wollen den in value gespeicherten Wert in der Liste durch 10 ersetzen
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list[i]); j++ {
			if list[i][j] == value {
				list[i][j] = 10
				break
			}
		}
	}
	// Diese Schleife geht wieder alle Werte der Liste durch und vergleicht sie mit value
	// Wenn der Wert gefunden wurde, wird er durch 10 ersetzt und die Schleife wird abgebrochen
	// In diesem Fall wird die Stelle list[0][0] durch 10 ersetzt

	// Noch ein Beispiel: wir wollen eine Zeile mit der nächsten addieren
	for i := 0; i < len(list)-1; i++ {
		for j := 0; j < len(list[i]); j++ {
			list[i][j] += list[i+1][j]
		}
	}
	// Diese Schleife geht alle Zeilen durch, außer der letzten
	// Dabei wird jeder Wert der Zeile mit dem Wert der nächsten Zeile addiert und gespeichert

	for _, l := range list {
		fmt.Println(l)
	}
	fmt.Println()

	// Noch ein Beispiel: wir wollen die Liste transponieren (Zeilen und Spalten vertauschen)
	var newList = make([][]int, len(list))
	for i := 0; i < len(list); i++ {
		newList[i] = make([]int, len(list[i]))
	}
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list[i]); j++ {
			newList[j][i] = list[i][j]
		}
	}
	for _, l := range newList {
		fmt.Println(l)
	}
	// Diese Schleife erstellt eine neue Liste mit den Dimensionen vertauscht

}
