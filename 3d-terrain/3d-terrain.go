package main

import (
    "image"
    "image/color"
    "image/png"
    "math"
    // "reflect"
    // "math/rand"
    "os"
    "fmt"
    "github.com/skratchdot/open-golang/open"
    "github.com/ojrac/opensimplex-go"
    // "image/g?if"

)
var (

	colorBlue = color.RGBA{70,130,180,255}
    colorGreen = color.RGBA{34,139,34,255}
    colorYellow = color.RGBA{255,250,205,255}
    colorWhite = color.RGBA{255,255,255,255}
    colorBrown = color.RGBA{210,180,140,255}
    colorBlack = color.RGBA{100,100,100,1}
    colorClear = color.RGBA{0,0,0,0}
    frames = 200 // How many frames to render
    fineAmplitude = 20.0 // I should think of different names for these. it just controls how "intense" the small and large 
    largeAmplitude = 20.0 // features are
    zoom = 0.02 //smaller is more zoomed in
    caveAmount = -0.30 // higher = bigger and more caves. Must be between -1 and 1
    imageHeight = 500
    imageWidth = 500
    z = 0.00

)

func combine(c1, c2 color.Color) color.Color {
    // Used for transparency
    r, g, b, a := c1.RGBA()
    r2, g2, b2, a2 := c2.RGBA()

    return color.RGBA{
        uint8((r + r2) >> 9), // div by 2 followed by ">> 8"  is ">> 9"
        uint8((g + g2) >> 9),
        uint8((b + b2) >> 9),
        uint8((a + a2) >> 9),
    }
}

func main() {
    fmt.Println("starting")
	simp := opensimplex.New();
    simp2 := opensimplex.New();
    img := image.NewRGBA(image.Rect(0, 0, imageHeight*2, imageWidth*2))
    n := 1
    for step := 1; step < frames; step++ {
        z+=0.1
        n++
        // drawing code (too long)
        
        for i := 0; i < imageHeight; i++ {
        	for q := 0; q < imageWidth; q++ {

                // "surface"
                height := simp2.Eval3(zoom/2*float64(q),zoom/2*float64(i+750),z)
                height2 := math.Abs(simp2.Eval3(zoom*2*float64(q),zoom*2*float64(i+750),z+9999))
                caveSimp := (simp.Eval3(zoom*float64(q),zoom*float64(i+750),z) - height2) / 2

        
                if (int( height2 * fineAmplitude) + int( height * largeAmplitude  ) > i-int(largeAmplitude)) {

                    // SKY (right now outputs nothing)
                    //img.Set(q+step,i+step,colorWhite)
                    //img.Set(q+step, i+step, combine(img.At(q+step, i+step), color.RGBA{0, 0, 0, 50}))
                }else if(int( height2 * fineAmplitude) + int( height * largeAmplitude  ) <= i-int(largeAmplitude) &&
                         int( height2 * fineAmplitude) + int( height * largeAmplitude  ) >  i-int(largeAmplitude) - 5){

                    // SURFACE
                    img.Set(q+step, i+step, color.RGBA{0,uint8(200*float32(n)/float32(frames)),0,255})
                }else{
                    if( caveSimp <= caveAmount &&
                        caveSimp > caveAmount - 0.15){

                        // CAVE WALLS
                        img.Set(q+step,i+step,color.RGBA{uint8(150.0*(150-(float32(n)/float32(frames)))),
                                                         uint8(150.0*(150-(float32(n)/float32(frames)))),
                                                         uint8(150.0*(150-(float32(n)/float32(frames)))),255});
                    }
                    if(caveSimp > caveAmount){

                        // SOLID GROUND
                        if q < 3 {
                        	img.Set(q+step,i+step,color.RGBA{uint8(210.0*(float32(n)/float32(frames))),
                                                         uint8(180.0*(float32(n)/float32(frames))),
                                                         uint8(140.0*(float32(n)/float32(frames))),255});
                        }else{
                        	 img.Set(q+step,i+step,color.RGBA{uint8(210.0*(float32(n-50)/float32(frames))),
                                                         uint8(180.0*(float32(n-50)/float32(frames))),
                                                         uint8(140.0*(float32(n-50)/float32(frames))),255});
                        }
                       
                    }
                }

               

        	}
        }
    }

   f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
    defer f.Close()
    png.Encode(f, img)
    open.Run("out.png")
}