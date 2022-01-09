package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func findImage(x string) string {
	fmt.Println(x)
	return x
}



func main(){
	// Read prng-service.txt
	b, err := ioutil.ReadFile("prng-service.txt")
	check(err)

	// Convert file contents to str & then to int64
	str := string(b)
	n, err := strconv.ParseInt(str, 10, 64)
	rand := int(n)
	check(err)

	// Check # of images in ./images directory
	files, err := ioutil.ReadDir("./images")
	var num_files = len(files)
	// fmt.Println(num_files)

	// Create a slice (more flexible than array)
	var images []string

	// Create arr w/ names of files in ./images
	for _, file := range files {
		images = append(images, file.Name())
    }

	//Check rand_num against length of arr & use modulo
	var mod int
	if int(rand) > num_files{
		mod = rand % num_files
	}

	image := images[mod]

	fmt.Println(image)
	// fmt.Println(reflect.TypeOf(files))

	// Want image.go to return image for UI to display
	findImage(image)
}