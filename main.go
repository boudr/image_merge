package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

type img struct {
	Path string
	Data []byte
	Perm os.FileMode
	Imge image.Image
}

func main() {
	sylas := &img{
		Path: "/Users/cee/go/src/github.com/boudr/test/t1.png",
		Perm: 0644,
	}

	lux := &img{
		Path: "/Users/cee/go/src/github.com/boudr/test/t2.png",
		Perm: 0644,
	}

	var err error

	sylas.Data, err = ioutil.ReadFile(sylas.Path)
	if err != nil {
		log.Fatal(err)
	}

	lux.Data, err = ioutil.ReadFile(lux.Path)
	if err != nil {
		//Fatal will print the error to stderr and os.Exit(1)
		log.Fatal(err)
	}

	sylas.Imge, err = png.Decode(bytes.NewReader(sylas.Data))
	if err != nil {
		log.Fatal(err)
	}

	lux.Imge, err = png.Decode(bytes.NewReader(lux.Data))
	if err != nil {
		log.Fatal(err)
	}

	var newPix []uint8

	for i := 0; i < (len(sylas.Imge.(*image.RGBA).Pix)+len(lux.Imge.(*image.RGBA).Pix))/2; i++ {
		newPix = append(newPix, sylas.Imge.(*image.RGBA).Pix[i], lux.Imge.(*image.RGBA).Pix[i])
	}

	fmt.Println(len(newPix))

	merged := &img{
		Path: "/Users/cee/go/src/github.com/boudr/test/merged.jpeg",
		Perm: 0644,
		Imge: image.NewRGBA(image.Rect(0, 0, 1920*2, 1080)), //sylas.Imge.(*image.Paletted).Palette),
	}

	merged.Imge.(*image.RGBA).Pix = newPix
	fmt.Println(len(merged.Imge.(*image.RGBA).Pix))
	fmt.Println(len(sylas.Imge.(*image.RGBA).Pix))

	fmt.Println(merged.Imge.(*image.RGBA).Bounds())

	// enc := png.Encoder{
	// 	CompressionLevel: png.NoCompression,
	// }

	f, err := os.Create(merged.Path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(f)
	defer f.Close()

	err = jpeg.Encode(f, merged.Imge, &jpeg.Options{
		Quality: 100,
	})
	if err != nil {
		log.Fatal(err)
	}

}

//Read and write an image
// func main() {
// 	fmt.Println("Reading image from path")

// 	b, err := ioutil.ReadFile("/Users/cee/go/src/github.com/boudr/test/sylas.png")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println(b)

// 	fmt.Println("Outputting to a new file")

// 	err = ioutil.WriteFile("/Users/cee/go/src/github.com/boudr/test/sylas-new.png", b, 0644)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println("File written")

// }
