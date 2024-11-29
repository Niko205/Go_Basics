package listemehrdimensional

import "fmt"

// Uebung1 bekommt einen int-Wert und erstellt eine zwei-dimensionale Liste mit der Länge des Wertes, Bsp 3 -> 3x3 Liste
func ExampleUebung1() {
	list := Uebung1(3)
	for _, l := range list {
		fmt.Println(l)
	}
	// Output:
	// [1 2 3]
	// [4 5 6]
	// [7 8 9]
}

// Uebung2 bekommt eine zwei-dimensionale Liste und soll die Werte der Hauptdiagonale zurückgeben
func ExampleUebung2() {
	list := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(Uebung2(list))
	// Output:
	// [1 5 9]
}

// Uebung3 bekommt eine zwei-dimensionale Liste und eine Zahl n
func ExampleUebung3() {
	list := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(Uebung3(list, 5))
	// Output:
	// Reihe: 2, Spalte: 2
}

// Uebung4 bekommt eine zwei-dimensionale Liste und soll die Liste um 90 Grad drehen und zurückgeben
func ExampleUebung4() {
	list := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	newList := Uebung4(list)
	for _, l := range newList {
		fmt.Println(l)
	}
	// Output:
	// [7 4 1]
	// [8 5 2]
	// [9 6 3]
}

/*
Lösung für Uebung1:
func Uebung1(n int) [][]int {
	list := make([][]int, n)
	for i := 0; i < n; i++ {
		list[i] = make([]int, n)
		for j := 0; j < n; j++ {
			list[i][j] = i*n + j + 1
		}
	}
	return list
}

Lösung für Uebung2:
func Uebung2(list [][]int) []int {
	var diagonal []int
	for i := 0; i < len(list); i++ {
		diagonal = append(diagonal, list[i][i])
	}
	return diagonal
}

Lösung für Uebung3:
func Uebung3(list [][]int, n int) string {
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list[i]); j++ {
			if list[i][j] == n {
				return fmt.Sprintf("Reihe: %d, Spalte: %d", i+1, j+1)
			}
		}
	}
	return "Nicht gefunden"
}

Lösung für Uebung4:
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
*/
