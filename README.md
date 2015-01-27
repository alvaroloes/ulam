# Ulam Spiral
This is a very simple program to create images with an Ulam Spiral on it. 
The output of the program with a image dimension of 500 will be the following:

![Ulam Spiral](UlamSpiral.bmp)

In this case the prime numbers are highlighted in green color. As you see, the mysterious (or maybe not so) diagonals are shown,
indicating that there could be a pattern in the prime number distribution.

If you don't know who is Stanislaw Ulam and what he discovered, or you want to know more about the 
Ulam Spiral, here are some links:

* [Ulam Spiral](http://en.wikipedia.org/wiki/Ulam_spiral)
* [Ulamspiral.com](http://ulamspiral.com/)
* [Prime Spiral](http://mathworld.wolfram.com/PrimeSpiral.html)

# How to use
You first need to have [Go properly installed](http://golang.org/doc/install#download). 
Ensure that your `GOPATH/bin` is added to your PATH environment variable.

After that you can "go get" the project easily with:

    go get github.com/kloster/ulam
    
To generate an Ulam spiral image with default values you can just call

    ulam
    
This will generate a 100 x 100 image in the current working directory with the name `UlamSpiral.bmp`

You can specify the dimensions and/or the name of the image in this way:

    ulam -dim=2000 ~/Images/SuperUlamSpiral.bmp

The algorithm to check if a number is prime or not is under the `prime` package.
It uses the best simple algorithm as shown in [Wikipedia - Primality test](http://en.wikipedia.org/wiki/Primality_test)



# Why?
Well, I like the kind of mystery around the prime numbers and I am liking a lot the Go programing language a lot, 
so I decided to start learning the image processing in Go with this tiny program.
 
If I find the time, I will try to play a little more with other kind of prime numbers graphical representations...

Who knows, maybe the formula to get Prime numbers is about to be discovered! (Well, in that case we should not 
be that happy. The Internet security would explode!)

# TODO

*Improve the UlamSpiral algorithm
*Add concurrency
