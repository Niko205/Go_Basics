package main

import "fmt"

func main() {
	// Wie bereits gesagt gibt es eine einfachere Variante Variablen anzulegen als im ersten Kapitell gezeigt wurde
	// Hierfür benötigen wir nun das ":=" und benutzt es wie folgt
	// "Name der Variable" := Wert, der Datentyp der Vriable legt sich hierbei von allein fest (in dem Beispiel ein String, drüberhovern)
	stringvar := "Hallo"

	// Nach diesem Prinziep kann man Variablen jedes Datentyps anlegen, eine Liste der Datentypen
	// string <- Speichert Wörter und Sätze, ist eine Liste von bytes dazu aber später mehr
	// byte <- Speichert ein einzelnes Zeichen (in '' geschrieben) aus dem Unicode (Tabbelle von zeichen mit zugehörigen Zahlenwert)
	// int <- Speichert ganze Zahlen im positiven und negativen Bereich
	// float <- Speichert Gleitkommazahlen, diese gibt es in unterschiedlicher Genauigkeit (32 oder 64) bei dieser Variante wird aber automatisch 64 genommen
	// bool <- Speichert einen Warheitswert der die zwei Werte "true" oder "fals" annehmen kann
	// Hier Beipiele vom Anlegen mit der obigen Variante:
	bytevar := 'A'
	intvar := 1
	floatvar := 1.87
	boolvar := true

	// Möchte man einem Attribut nun einen anderen Wert zuteilen schreibt man wieder den Namen der Variable
	// Danach ein = und im Anschluss den neuen Wert. MERKEN ":=" wenn man die Variable anlegt und einen Wert gibt und "=" wenn man einer Variable einen neuen Wert gibt
	stringvar = "Tschau"
	bytevar = 'Z'
	intvar = -1
	floatvar = -1.87
	boolvar = false

	// Nun müssen wir wieder die erstellten Variablen benutzen, da es sonst einen Compilerfehler gibt
	// Hierführ geben wir die Variablen einfach wieder mittel "fmt.Println" im Terminal aus
	fmt.Println(stringvar)
	fmt.Println(bytevar)
	fmt.Println(intvar)
	fmt.Println(floatvar)
	fmt.Println(boolvar)

	// Führt man nun die Main-Funktion aus (Im Integrated Termin: "go run .")
	// So erhalten wir folgenden Output:
	// Tschau
	// 90
	// -1
	// -1.87
	// false

	// Beim ausgegebenen Byte fällt auf das wir kein Zeichen sondern eine Zahl erhalten
	// Dies liegt daran das ein Byte gleichzeitig einem Zeichen, aber auch einer Zahl entspricht
	// Diese zugehörigen Zeichen und Zahlen lassen sich in der ASCII-Tabelle nachschauen, in diesem Fall entspricht 'Z' der Zahl 90
	// Im nächsten Kapitel gibt es wieder einen Test um die Theorie zu üben
}
