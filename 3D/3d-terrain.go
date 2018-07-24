package main

import (
    "image"
    "image/color"
    // "image/png"
    // "reflect"
    // "math/rand"
    "os"
    // "fmt"
    "github.com/skratchdot/open-golang/open"
    "github.com/ojrac/opensimplex-go"
    "image/gif"

)
var (

	colorBlue = color.RGBA{70,130,180,255}
    colorGreen = color.RGBA{34,139,34,255}
    colorYellow = color.RGBA{255,250,205,255}
    colorWhite = color.RGBA{255,255,255,255}
    colorBrown = color.RGBA{139,69,19,255}
    zoom = 0.02 //smaller is more zoomed in
    height = 500
    width = 500
    z = 0.00
    images []*image.Paletted
    delays []int

    palette = []color.Color{
        color.RGBA{70,130,180,255},
        color.RGBA{34,139,34,255},
        color.RGBA{255,250,205,255},
        color.RGBA{255,255,255,255},
        color.RGBA{139,69,19,255},
    }

)

func main() {
	simp := opensimplex.New();

    for step := 0; step < 200; step++ {
        img := image.NewPaletted(image.Rect(0, 0, width, height), palette)
        images = append(images, img)
        delays = append(delays, 0)
        z+=0.01

        // drawing code (too long)
  
        for i := 0; i < height; i++ {
        	for q := 0; q < width; q++ {
        		height := simp.Eval3(zoom*float64(q),zoom*float64(i+750),z)
        		if(height > -.5){
    		        img.Set(q, i, colorWhite)
        		}else{
    				img.Set(q, i, colorBlue)
        		}
        	}
        }
    }

    f, _ := os.OpenFile("out.gif", os.O_WRONLY|os.O_CREATE, 0600)
    defer f.Close()
    gif.EncodeAll(f, &gif.GIF{
        Image: images,
        Delay: delays,
    })

    // Save to out.png
    open.Run("out.gif")

}