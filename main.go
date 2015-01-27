package main

import "image"
import "image/color"
import "os"
import "golang.org/x/image/bmp"
import "fmt"
import "github.com/kloster/ulam/prime"
import "flag"

func main() {
    var dim int
    flag.IntVar(&dim, "dim", 100, "The dimension of the resulting image (it will be square)")
    flag.Parse()

    outputFile := "./UlamSpiral.bmp"
    if len(flag.Args()) > 0 {
        outputFile = flag.Arg(0)
    }

    // Force the dimension to be odd to avoid black lines in the image
    if dim % 2 == 0 {
        dim++
    }

    // Create the image
    img := image.NewRGBA(image.Rect(0,0,dim,dim))

    // Fill it with the Ulam spiral
    fillWithUlamSpiral(img,color.RGBA{255,255,255,255}, color.RGBA{0,150,0,0})

    // Create the ouput file
    imgFile, err := os.Create(outputFile)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer imgFile.Close()

    // Encode the image as a bmp file
    if err = bmp.Encode(imgFile, img); err != nil {
        fmt.Println(err)
    }
}

// A straightforward implementation
func fillWithUlamSpiral(img *image.RGBA, compositeColor, primeColor color.RGBA) {
    bounds := img.Bounds()
    dim := bounds.Dx()                      // One of the dimensions of the image (it is square)
    x, y := dim/2, dim/2                    // The points to be painted. We start from the center
    dx, dy := 1, -1                         // The increment direction
    alreadyMovedX, alreadyMovedY := 0,0     // The points we already have moved in the current row/column
    rowSize, columSize := 1,1               // The current row/column size

    totalPoints := dim*dim
    for i := 0; i < totalPoints; i++ {
        numToCheck := i + 1
        if prime.Is(numToCheck) {
            img.Set(x, y, primeColor)
        } else {
            img.Set(x, y, compositeColor)
        }

        if alreadyMovedX < rowSize {
            x += dx
            alreadyMovedX++
        } else {
            if alreadyMovedY < columSize {
                y += dy;
                alreadyMovedY++
            }
            if alreadyMovedY >= columSize {
                dx *= -1
                dy *= -1
                alreadyMovedX, alreadyMovedY = 0,0
                rowSize++
                columSize++
            }
        }
    }
}

