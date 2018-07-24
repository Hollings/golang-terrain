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
    colorBlack = color.RGBA{0,0,0,255}
    frames = 2000 // How many frames to render
    fineAmplitude = 5.0 // I should think of different names for these. it just controls how "intense" the small and large 
    largeAmplitude = 100.0 // features are
    zoom = 0.04 //smaller is more zoomed in
    imageHeight = 500
    imageWidth = 500
    z = 0.00
    images []*image.Paletted
    delays []int

    palette = []color.Color{
        color.RGBA{70,130,180,255},
        color.RGBA{34,139,34,255},
        color.RGBA{255,250,205,255},
        color.RGBA{255,255,255,255},
        color.RGBA{139,69,19,255},
        color.RGBA{0,0,0,255},
    }

)

func main() {
	simp := opensimplex.New();
    simp2 := opensimplex.New();
    for step := 0; step < frames; step++ {
        img := image.NewPaletted(image.Rect(0, 0, imageWidth, imageHeight), palette)
        images = append(images, img)
        delays = append(delays, 0)
        z+=0.01

        // drawing code (too long)
  
        for i := 0; i < imageHeight; i++ {
        	for q := 0; q < imageWidth; q++ {

                // "surface"
                height := simp2.Eval3(zoom/2*float64(q),zoom/2*float64(i+750),z)
                height2 := simp2.Eval3(zoom*2*float64(q),zoom*2*float64(i+750),z+9999)

    
                if (int( height2 * fineAmplitude) + int( height * largeAmplitude  ) > i-int(largeAmplitude)) {
                    img.Set(q,i,colorBlue);
                }else{
                     // "caves"
                    height := simp.Eval3(zoom*float64(q),zoom*float64(i+750),z)
                    if(height > -.5){
                        img.Set(q, i, colorWhite)
                    }else{
                        img.Set(q, i, colorBlack)
                    }
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