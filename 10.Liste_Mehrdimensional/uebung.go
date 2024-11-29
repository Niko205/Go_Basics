package listemehrdimensional

// Uebung1 bekommt einen int-Wert und erstellt eine zwei-dimensionale Liste mit der Länge des Wertes, Bsp 3 -> 3x3 Liste
// Diese Liste soll aufzählend gefüllt werden beginnend mit 1, 2 usw.
// Die Liste soll zurückgegeben werden
func Uebung1(n int) [][]int {
	// Hier die Lösung für die erste Übung
	return nil
}

// Uebung2 bekommt eine zwei-dimensionale Liste und soll die Werte der Hauptdiagonale zurückgeben
// Die Hauptdiagonale ist die Diagonale von links oben nach rechts unten
func Uebung2(list [][]int) []int {
	// Hier die Lösung für die zweite Übung
	return nil
}

// Uebung3 bekommt eine zwei-dimensionale Liste und eine Zahl n
// Sie soll die Liste nach n absuchen und die Koordinaten des ersten Vorkommens zurückgeben
// Die Koordinaten sollen als String zurückgegeben werden, Bsp: "Reihe: 1, Spalte: 2"
// Wenn die Zahl nicht gefunden wird, soll "Nicht gefunden" zurückgegeben werden

func Uebung3(list [][]int, n int) string {
	// Hier die Lösung für die dritte Übung
	return ""
}

// Uebung4 bekommt eine zwei-dimensionale Liste und soll die Liste um 90 Grad nach rechts drehen und zurückgeben
func Uebung4(list [][]int) [][]int {
	newList := make([][]int, len(list))
	for i := 0; i < len(list); i++ {
		newList[i] = make([]int, len(list))
	}
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list[i]); j++ {
			newList[j][len(list)-i-1] = list[i][j]
		}
	}
	return newList
}
