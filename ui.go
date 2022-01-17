package main

import (
	"fmt"
	"io/ioutil"
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

	fmt.Println("Would you like to run PRNG microservice? (yes/no)")
	fmt.Scanln(&first)

	if strings.ToLower(first) == "yes"{
		fmt.Println("Running PRNG Microservice?")
	
		// Create prng-service.txt if doesn't exist
		f, err := os.Create("prng-service.txt")
		check(err)
		// Write "run" to prng-service.txt
		_, err2 := f.WriteString("run")
		check(err2)


		// ------------------ Call PRNG Microservice ------------------
		cmd1 := exec.Command("go", "run", "prng.go")
		_, err3 := cmd1.Output()
		check(err3)

		fmt.Println("Would you like to run Image microservice? (yes/no)") 
		fmt.Scanln(&second) 
	
		if strings.ToLower(second) == "yes"{
			// Read prng-service.txt to get pseudo-random #
			f2, err4 := ioutil.ReadFile("prng-service.txt")
			check(err4)

			// Convert file contents (pseudo-random #) to str
			str := string(f2)
			f7, err7:= os.OpenFile("image-service.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
			check(err7)
		

			// Write str (pseduo-random #) to file
			_, err8 := f7.WriteString(str)
			check(err8)
			



			fmt.Println("Running Image Microservice?")

			// ------------------ Call Image Microservice ------------------
			cmd2 := exec.Command("go", "run", "image.go")
			_, err6 := cmd2.CombinedOutput()
			check(err6)
			
			// Read image-service.txt to get image path
			image, err9 := ioutil.ReadFile("image-service.txt")
			check(err9)
			
			// Display image path to user in terminal
			fmt.Println(string(image))

		}

	}else {
		fmt.Println("Exiting...See you later!")
	}


}