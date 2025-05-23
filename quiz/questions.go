package quiz

type Questions struct {
	Id			string   `json:"id"`
	Question	string   `json:"question"`
	Options 	[]string `json:"options"`
	Answer 		int      `json:"answer"`
}

var GeographyQuiz = []Questions{
	{
		Id: 			"1",
		Question:		"In what city in Switzerland where the Eurovison 2025 finals held?",	
		Options:		[]string{"Bern", "Zürich", "Basel", "Geneva"},
		Answer:			2,	// Basel
	},
	{
		Id: 			"2",
		Question:         	"What was the capital of the former country German Democratic Republic?",
		Options:		[]string{"Berlin", "Leipzig", "East Berlin", "Bonn"},		
		Answer:			2,	// East Germany
	},
	{
		Id: 			"3",
		Question:		"What is the smallest country in Europe?",
		Options:		[]string{"Vatican City", "San Marino", "Liechtenstein", "Monaco"},		
		Answer:			0,	// Vatican City
	},
	{
		Id: 			"4",
		Question:		"Which European country has the most islands?",
		Options:		[]string{"Czech Republic", "Norway", "Finland", "Sweden"},		
		Answer:			3,	// Sweden
	},
	{
		Id: 			"5",
		Question: 		"Which European city has a famous peeing statue called 'Manneken Pis'?",
		Options: 	[]string{"Amsterdam", "Brussels", "Vienna", "Prague"},
		Answer: 	1, 	// Brussels
	},
	{
		Id: 			"6",
		Question:		"Which country of 'Big Five' for the second year in a row.\n Did not get a single point in the 2025 Eurovison Song Contest finals?",		// raw string input 
		Options:		[]string{"UK", "Spain", "Germany", "France"},		
		Answer:			0,	// UK
	},
	{
		Id: 			"7",
		Question:		"What is the capital of Malta?",
		Options:		[]string{"Senglea", "Valletta", "Cospicua", "Sliema"},		
		Answer: 		1,	// Valletta
	},		
	{
		Id: 			"8",
		Question:		"Which country was Not part of Yugoslavia?",
		Options:		[]string{"Croatia", "Greece", "Slovenia", "Montenegro"},	
		Answer:			1,	// Greecce
	},
	{
		Id: 			"9",
		Question:		"What country borders both the Atlatic & the Mediterranean Sea?",
		Options:		[]string{"Bulgaria", "Hungary", "Turkey", "France"}	,	
		Answer:			3,	// This is a hard one
	},
	{
		Id: 			"10",
		Question: 		"What name did the newly elected pope get/take?",
		Options: 		[]string{"Leo XIV", "Johannes III", "Benedict XVII", "John Paul III"},
		Answer: 		0, // Leo XIV	
	},
}