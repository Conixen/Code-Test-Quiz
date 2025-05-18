package questions

type Questions struct {
	Question		string
	Alternatvs		[]string
	RightAnswear	int	
}

var GeographyQuiz = []Questions{
	{
		Question:		"In what city in Switzerland where the Eurovison 2025 finals held?"	
		Alternatvs:		[]string{"Bern", "Zürich", "Basel", "Geneva"}
		RightAnswear:	2	// Basel
	},
	{
		Text:         	"What was the capital of the former country German Democratic Republic?",
		Alternatvs:		[]string{"Berlin", "Leipzig", "East Berlin", "Bonn"}		
		RightAnswear:	2	// East Germany
	},
	{
		Question:		"What is the smallest country in Europe?"
		Alternatvs:		[]string{"Vatican City", "San Marino", "Liechtenstein", "Monaco"}		
		RightAnswear:	0	// Vatican City
	},
	{
		Question:		"Which European country has the most islands?"
		Alternatvs:		[]string{"Czech Republic", "Norway", "Finland", "Sweden"}		
		RightAnswear:	3	// Sweden
	},
	{
		Question: 		"Which European city has a famous peeing statue called 'Manneken Pis'?",
		Alternatvs: 	[]string{"Amsterdam", "Brussels", "Vienna", "Prague"},
		RightAnswear: 	1, 	// Brussels
	},
	{
		Question:		"Which country of 'Big Five' for the second year in a row. 
						\nDid not get a single point in the 2025 Eurovison Song Contest finals?"
		Alternatvs:		[]string{"UK", "Spain", "Germany", "France"}		
		RightAnswear:	0	// UK
	},
	{
		Question:		"What is the capital of Malta?"
		Alternatvs:		[]string{"Senglea", "Valletta", "Cospicua", "Sliema"}		
		RightAnswear: 	1	// Valletta
	},		
	{
		Question:		"Which country was Not part of Yugoslavia?"
		Alternatvs:		[]string{"Croatia", "Greece", "Slovenia", "Montenegro"}		
		RightAnswear:	1	// Greecce
	},
	{
		Question:		"What country borders both the Atlatic & the Mediterranean Sea?"
		Alternatvs:		[]string{"Bulgaria", "Hungary", "Turkey", "France"}		
		RightAnswear:	3	// This is a hard one
	},
}