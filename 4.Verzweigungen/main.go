package main

import "fmt"

func main() {
	// In diesem Kapitel geht es nun um Vergleichsoperatoren und deren Einsatz in If-Else Verzweigungen
	// Die folgenden Vergleichsoperatoren gibt es:
	// == <- istgleich, Prüft ob ein Wert/Variable einer anderen entspricht
	// != <- istungleich, Prüft ob eine Wert/Variable nicht einer anderen entspricht
	// <  <- kleiner, Prüft ob ein Wert/Variable kleiner als eine andere ist
	// >  <- größer, Prüft ob ein Wert/Variable größer als eine andere ist
	// <= <- kleinergleich, Prüft ob ein Wert/Variable kleiner oder gleich als eine andere ist
	// >= <- größergleich, kleinergleich, Prüft ob ein Wert/Variable größer oder gleich als eine andere ist
	// All diese Operatoren benutzt man um eine Condition in eine If-Else verzweigung zu machen
	// Diese verzweigung wird nur ausgeführt, wenn die Condition erfüllt wird/true ist
	// Die if verzwigung startet man mit "if", dann die Condition, dann eine geschweifte Klammer
	intvar := 10
	if intvar == 10 { //Diese Condition entspricht true
		fmt.Println("intvar = 10") // Code in den geschweiften Klammern wird nur ausgeführt wenn die Condition true ergibt
		intvar++
	}

	// Dies war die If-Verzweigung, diese führt nur etwas aus, wenn die Condition zutrifft
	// Zudem gibt es noch die If-Else-Verzweigung, heißt so viel wie entwerde oder
	// Sie beginnt identisch zur If-Verzwigung, jedoch hängt man nach der geshweiften Klammer noch die Else-Verzweigung an

	if intvar == 10 {
		fmt.Println("intvar =", intvar) // Dieser Code wird wieder nur ausgeführt wenn die Condition stimmt
	} else {
		fmt.Println("intvar != 10") // Dieser Code wird ausgeführt wenn die Condition nicht stimmt
		intvar++
	}

	// Man kann auch mehrere Conditions mit logischen AND oder OR kombinieren
	// && <- verbindet zwei Conditions mit einem logischen AND
	// || <- verbindet zwei Conditions mit einem logischen OR
	if intvar >= 10 && intvar <= 15 { // Prfüt ob intvar im Zahlenbereich von 10 bis 15 liegt
		fmt.Println("intvar liegt im Bereich von 10 bis 15")
		intvar++
	}

	// Zuletzt kann man noch mehrere If-Verzweigungen aneinanderhängen
	// Dafür benutzt man nach dem ersten "if cond {}" ein else if cond {}
	if intvar <= 12 {
		fmt.Println("intvar ist gleich 12 oder kleiner")
	} else if intvar == 13 {
		fmt.Println("Intvar ist gleich 13")
		intvar++
	} else { // Jetzt könnte man noch ein else if mit intvar >= 12 machen, jedoch ist dies mit den anderen zwei schon festgelegt
		fmt.Println("intvar ist gleich 14 oder größer") // Darum reicht das einfache else
	}
	// Hat man eine ganze Reihe solcher Else-If-Verzwigungen lohnt sich möglicherweise ein Switch-Case um Zeit zu sparen
	// Dieser wird wie folgt angelegt:
	switch intvar { // Sagt das sich für den Switch-Case die Variable intvar angeschaut werden soll
	case 13: // Der Case 13 wird ausgeführt, wenn intvar den Wert 13 hat, genau so funktionieren auch die anderen Cases
		fmt.Println("intvar gleich 113")
	case 14:
		fmt.Println("intvar gleich 14")
	case 15:
		fmt.Println("intvar gleich 15")
	default:
		fmt.Println("intvar entspricht keinem Case ") // Der Default wird ausgegeben wenn die untersuchte Variable keinem Case entspricht
	}
	// Dieses Kapitel endet wieder ohne Test, jedoch finden sich Inhalte im Test des nächsten Kapitels wieder
}
