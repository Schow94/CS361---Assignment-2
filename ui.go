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

	fmt.Println("Would you like to run PRNG microservice?")
	fmt.Scanln(&first) // Need to convert to all lowercase

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

		fmt.Println("Would you like to run Image microservice?") // Runs before if statement below
		fmt.Scanln(&second) 
	
		if strings.ToLower(second) == "yes"{
			// Read prng-service.txt to get pseudo-random #
			f2, err4 := ioutil.ReadFile("prng-service.txt")
			check(err4)

			// Convert file contents (pseudo-random #) to str
			str := string(f2)
			// fmt.Println("STRING TO WRITE: ", str)
			// Write pseudo-random # to image-service.txt
			// OpenFile(filename, read/write/create, file permissions)
			f7, err7:= os.OpenFile("image-service.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
			check(err7)
		

			// Write str (pseduo-random #) to file
			_, err8 := f7.WriteString(str)
			check(err8)
			
			// // Check that writing 3 to image-service.txt was successful
			// f10, err := ioutil.ReadFile("image-service.txt")
			// check(err)
			// f10_str := string(f10)
			// fmt.Println("image service content: ", string(f10_str))



			fmt.Println("Running Image Microservice?")

			// ------------------ Call Image Microservice ------------------
			// Unable to run ui.go twice in a row
			// Must be a bug somewhere in my code
			// Somehow # & image are being combined in image-service.txt
			// Believe issue is due to not checking if an image or a # is already in image-service.txt
			cmd2 := exec.Command("go", "run", "image.go")
			_, err6 := cmd2.CombinedOutput()
			check(err6)
			
			// Read image-service.txt to get image path
			image, err9 := ioutil.ReadFile("image-service.txt")
			check(err9)
			
			// Display image path to user in terminal
			// Change this to a path. Currenlty just displaying the name of the image
			fmt.Println(string(image))

		}

	}else {
		fmt.Println("Exiting...See you later!")
	}


}