package main

import (

	"gowave/utils"
)


func main() {
	// if len(os.Args) >= 2 && os.Args[1] == "generate" {
	// 	generate()
	// 	return
	// }

	utils.Boot();
}


// func generate() {
// 	if len(os.Args) < 4 {
// 		log.Fatalf("GoWave : Invalid arguments.")
// 	}

// 	folderName := os.Args[2]
// 	fileName := os.Args[3]

// 	utils.Generate(folderName, fileName)
// }