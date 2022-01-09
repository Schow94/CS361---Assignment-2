# CS361 - Assignment 2

## Description

Demonstrate microservices architecture, writing software composed of 3 separate programs

### Microservice 1

A program that generates pseudo-random numbers (PRNG Service)

### Microservice 2

A program that, given a non-negative integer i, returns the ith image in a set (order
doesn’t matter) (Image Service)

- If i is >= the number of images, modulo i by the size of the image set

### Microservice 3

A user interface (UI) that either has a button or can receive a user command. When the button is pushed or the command is entered...

- (a) UI calls the PRNG Service
- (b) UI calls the Image Service using the pseudo-random number from the PRNG Service
- (c) UI displays the image (or a path to it)
