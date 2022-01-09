package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)


func check(err error){
	if err != nil {
        fmt.Println(err.Error())
        return
    }

}

func main(){


	var first string
	var second string

	fmt.Println("Would you like to run PRNG microservice?")
	fmt.Scanln(&first) // Need to convert to all lowercase

	if strings.ToLower(first) == "yes"{
		fmt.Println("Running PRNG Microservice?")
		// Call PRNG microservice here
		
		// Create prng-service.txt if doesn't exist
		f, err := os.Create("prng-service.txt")
		check(err)
		// Write "run" to prng-service.txt
		_, err2 := f.WriteString("run")
		check(err2)

		
		// Run command in terminal
		cmd := exec.Command("go", "run", "prng.go")
		_, err3 := cmd.Output()
		check(err3)

		// Image Microservice
		fmt.Println("Would you like to run Image microservice?") // Runs before if statemetn below
		fmt.Scanln(&second) // Need to convert to all lowercase
	
		if strings.ToLower(second) == "yes"{
			fmt.Println("Running Image Microservice?")

			cmd := exec.Command("go", "run", "image.go")
			_, err := cmd.Output()
			check(err)
		}

	}else {
		fmt.Println("Exiting...See you later!")
	}


}