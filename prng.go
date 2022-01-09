package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	prng "github.com/ericlagerg/go-prng/xorshift"
)

func main() {
	// Read prng-service.txt
	cmd, err := ioutil.ReadFile("prng-service.txt")
	

	if strings.ToLower(string(cmd)) == "run" {
		fmt.Println(string(cmd))
	}

	// err is non-nil
	if err != nil {
		log.Fatal(err)
	}

	// PRNG Service (Returns uint64)
	s := new(prng.Shift128Plus)
	s.Seed()
	rand_num := s.Next()

	fmt.Println("RANDOM: ", rand_num)

	
	
	// Convert uint64 --> int
	num_int := int(rand_num)
	// Convert int --> str
	num_str := strconv.Itoa(num_int)

	// Open file
	//OpenFile(filename, read/write/create, file permissions)
    f2, err2:= os.OpenFile("prng-service.txt", os.O_RDWR|os.O_CREATE, 0755)
    if err2 != nil {
        panic(err)
    }


	// Write str to file
	// Getting error here
    _, err3 := f2.WriteString(num_str)

    if err3 != nil {
        log.Fatal(err3)
    }

	// Close immediately after opening file
	defer f2.Close()


    fmt.Println("done")

}
