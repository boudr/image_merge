package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color/palette"
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
		Path: "/Users/cee/go/src/github.com/boudr/test/sylas.png",
		Perm: 0644,
	}

	lux := &img{
		Path: "/Users/cee/go/src/github.com/boudr/test/lux.png",
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

	for i := 0; i < 28900; i++ {
		newPix = append(newPix, sylas.Imge.(*image.Paletted).Pix[i], lux.Imge.(*image.Paletted).Pix[i])
	}

	fmt.Println(len(newPix))
	fmt.Println(sylas.Imge.(*image.Paletted).Rect)
	fmt.Println(lux.Imge.(*image.Paletted).Rect)

	merged := &img{
		Path: "/Users/cee/go/src/github.com/boudr/test/merged-mini.jpeg",
		Perm: 0644,
		Imge: image.NewPaletted(image.Rect(0, 0, 170*2, 170), palette.WebSafe),
	}

	fmt.Println(merged.Imge.(*image.Paletted).Stride)

	merged.Imge.(*image.Paletted).Pix = newPix

	enc := png.Encoder{
		CompressionLevel: png.NoCompression,
	}

	f, err := os.Create(merged.Path)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	err = enc.Encode(f, merged.Imge)
	if err != nil {
		log.Fatal(err)
	}
}
