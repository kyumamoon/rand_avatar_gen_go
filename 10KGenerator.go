package main

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func createavatar(set string, filename int) {
	// Filename Setup
	exportname := strconv.Itoa(filename) + ".jpg"

	// Opens Base Image
	baseimg, err := os.Open("Base/Base.jpg")
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}

	// Decode Base Img
	basedecode, err := jpeg.Decode(baseimg)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer baseimg.Close()

	// Open First Set
	color := string(set[0])
	color = "1/" + color + ".png"

	colorimg, err := os.Open(color)
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}
	colordecode, err := png.Decode(colorimg)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer colorimg.Close()

	// Open Second Set
	eyes := string(set[1])
	eyes = "2/" + eyes + ".png"

	eyesimg, err := os.Open(eyes)
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}
	eyesdecode, err := png.Decode(eyesimg)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer eyesimg.Close()

	// Open Mouth Set
	mouth := string(set[2])
	mouth = "3/" + mouth + ".png"

	mouthimg, err := os.Open(mouth)
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}
	mouthdecode, err := png.Decode(mouthimg)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer mouthimg.Close()

	// Open Head Set
	head := string(set[3])
	head = "4/" + head + ".png"

	headimg, err := os.Open(head)
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}
	headdecode, err := png.Decode(headimg)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer headimg.Close()

	// Merge All Sets
	offset := image.Pt(0, 0)   // offset
	b := basedecode.Bounds()   // base bound
	image3 := image.NewRGBA(b) // new image based on base's bounds

	draw.Draw(image3, b, basedecode, image.ZP, draw.Over)                                 // create a new image with base on top.
	draw.Draw(image3, colordecode.Bounds().Add(offset), colordecode, image.ZP, draw.Over) // put color ontop of base
	draw.Draw(image3, eyesdecode.Bounds().Add(offset), eyesdecode, image.ZP, draw.Over)   // put eyes ontop of color
	draw.Draw(image3, mouthdecode.Bounds().Add(offset), mouthdecode, image.ZP, draw.Over) // put mouth ontop of eyes
	draw.Draw(image3, headdecode.Bounds().Add(offset), headdecode, image.ZP, draw.Over)   // put hat ontop of mouth

	third, err := os.Create(exportname)
	if err != nil {
		log.Fatalf("failed to create: %s", err)
	}
	jpeg.Encode(third, image3, &jpeg.Options{jpeg.DefaultQuality})
	defer third.Close()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	combolist := [10000]string{}

	// Generate 10000 Numbers
	for i := 0; i < 10000; i++ {
		combo := strconv.Itoa(i)
		if len(combo) == 1 {
			combo = "000" + combo
		} else if len(combo) == 2 {
			combo = "00" + combo
		} else if len(combo) == 3 {
			combo = "0" + combo
		} else if len(combo) == 4 {
			// nothing
		}
		combolist[i] = combo
	}

	// Shuffle the Combo Set
	for i := 0; i < 10000; i++ {
		randompos := rand.Intn(10000)
		tempA := combolist[randompos]
		combolist[randompos] = combolist[i]
		combolist[i] = tempA
	}

	// Generate Avatar
	for i := 0; i < 10000; i++ {
		createavatar(combolist[i], i)
	}

	//fmt.Println(combolist)
	//fmt.Println((combolist[9999]))
}
