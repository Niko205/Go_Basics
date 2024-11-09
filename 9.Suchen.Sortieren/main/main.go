package main

import "fmt"

func main() {
	// In diesem Kapielt geht es um das Suchen und Sortieren innerhalb von Listen
	// Dies braucht man, wenn man wissen will ob ein Element in einer Liste vorhand oder wo es sich befindet
	// Oder der Inhalt einer Liste sortiert werden soll
	// Möchte mann also wissen ob ein Element in einer Liste vorhanden ist,
	// so kann man mit einer Schleife durch die Liste iterieren
	// und bei jedem Element prüfen ob es dem gesuchten Element entspricht
	// Bsp Simples Suchen:
	var list = []int{2, 4, 6, 3, 7, 1}
	var search = 3
	for i := 0; i < len(list); i++ {
		if list[i] == search {
			fmt.Println("Element gefunden")
			fmt.Println("Index:", i)
			break // Wenn das Element gefunden wurde, kann die Schleife abgebrochen werden
		}
	}

	// Binäre Suche:
	// Die Binäre Suche ist schneller als die einfache Suche, jedoch muss die Liste sortiert sein
	// Bei der Binären such wird das mittlere Element der Liste mit dem gesuchten Element verglichen
	// Ist das mittlere Element kleiner als das gesuchte Element, so wird die Suche im rechten Teil der Liste fortgesetzt
	// Ist das mittlere Element größer als das gesuchte Element, so wird die Suche im linken Teil der Liste fortgesetzt
	// Ist das mittlere Element gleich dem gesuchten Element, so wurde das Element gefunden
	// Dies wird solange wiederholt, bis das Element gefunden wurde
	// Bsp Binäre Suche: (Nicht Rekursiv. Rekursiv ist auch möglich Bsp. Dafür in prog1-material/code-vorlesung/2024-10-18/find-list)
	var list2 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var search2 = 3
	var left = 0
	var right = len(list2) - 1
	for left <= right {
		var middle = (left + right) / 2 // Berechnung des mittleren Index zwischen dem linken und rechten Index
		if list2[middle] == search2 {
			fmt.Println("Element gefunden")
			fmt.Println("Index:", middle)
			break
		} else if list2[middle] < search2 { // Element ist im rechten Teil der Liste
			left = middle + 1 // Somit wird der linke Index auf das mittlere Element +1 gesetzt
		} else { // Element ist im linken Teil der Liste
			right = middle - 1 // Somit wird der rechte Index auf das mittlere Element -1 gesetzt
		}
	}
	// Simple Sortierung:
	// Ein Beispiel für eine einfache Sortierung ist der Selection Sort
	// Bei diesem Sortierverfahren wird jedes Element mit den anderen Elementen der Liste verglichen und falls dieses kleiner ist die Stellen getauscht
	// Dies macht man mittels zwei ineinander verschachtelten Schleifen
	// Bsp Selection Sort: (es gibt mehere simple Sortierverfahren, wie Bubblesort, Insertionsort, etc. Dieser ist jedoch mein Favorit)
	var list3 = []int{2, 4, 6, 3, 7, 1}
	for i := 0; i < len(list3); i++ { // Äußere Schleife: Geht jedes Element der Liste durch
		for j := i + 1; j < len(list3); j++ { // Innere Schleife: Vergleicht das aktuelle Element mit den nächsten Elementen
			if list3[i] > list3[j] { // Wenn das aktuelle Element größer ist als das nächste Element, werden die Stellen getauscht
				var temp = list3[i]
				list3[i] = list3[j]
				list3[j] = temp
			}
		}
	}
	fmt.Println(list3)

	// Es gibt auch komplexere Sortierverfahren, wie den Quicksort oder Mergesort
	// Diese sind jedoch rekursiv und komplexer zu verstehen und bekommen vllt ein eigenes Kapitel
	// Das war es für das Kapitel Suchen und Sortieren, es gibt wieder Übungen in den folgenden Dateien
	// Noch zu aller letzt, es gibt schon fertige Sortier oder Scuhfunktionen in bestimmten Packages
	// Trotzdem ist es gut zu wissen wie man dies selbst implementiert, falls es doch mal nötig ist
}
