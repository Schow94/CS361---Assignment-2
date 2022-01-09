package main

import (
	"io/ioutil"
	"os"
	"strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func main(){
	// Read prng-service.txt
	b, err := ioutil.ReadFile("image-service.txt")
	check(err)

	// Convert file contents to str & then to int64
	str := string(b)
	n, err2 := strconv.ParseInt(str, 10, 64)
	rand := int(n)
	check(err2)

	// Check # of images in ./images directory
	files, err := ioutil.ReadDir("./images")
	var num_files = len(files)

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

	// fmt.Println(reflect.TypeOf(files))

	// Write pseudo-random # to image-service.txt
	// OpenFile(filename, read/write/create, file permissions)
	f7, err7:= os.OpenFile("image-service.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	check(err7)


	// Write image path to image-service.txt
	_, err8 := f7.WriteString(image)
	check(err8)


}