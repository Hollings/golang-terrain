package main

import (
    "image"
    "image/color"
    "image/png"
    //"reflect"
    "math/rand"
    "os"
    "fmt"
    "github.com/skratchdot/open-golang/open"
    "github.com/ojrac/opensimplex-go"
    // "image/gif"

)
var (
	colorGrey = color.RGBA{150,150,150,255}
    colorWhite = color.RGBA{255,255,255,255}
    colorBrown = color.RGBA{139,69,19,255}
    colorDarkGrey = color.RGBA{100,100,100,255}
    colorDarkBrown = color.RGBA{100,60,10,255}
    colorBlack = color.RGBA{0,0,0,255}

)

func main() {
	fmt.Println("starting")
    img := image.NewRGBA(image.Rect(0, 0, 1000, 1000))
    for i := 0; i < 1000; i++ {
    	for q := 0; q < 1000; q++ {
			img.Set(i,q,colorBlack)        		
    	}
    }
    for o := 50; o < 1000; o+=50 {
	   
	    baseHeight := rand.Intn(50)+10
	   	candleWidth := rand.Intn(10)*2
	    candleHeight := rand.Intn(100)
	    //wickHeight := rand.Intn(10)
	    //flameHeight := rand.Intn(10)

	    //candleBoundsX = candleWidth
	    //candleBoundsY = candleHeight + baseHeight + wickHeight + flameHeight
	    
	   

	    // Create the waxy part of the candle
	    for z := 0; z < candleHeight; z++ {
	    	for i := candleWidth; i > 0; i-- {
	    		if i<2{
				img.Set(o+i+2,51-z,colorDarkGrey)    	
	    		}else{
				img.Set(o+i+2,51-z,colorGrey)    	

	    		}
	    	}
	    }

	    // Create candle "holder"
	    for i := candleWidth+4; i > 0; i-- {
			img.Set(o+i,50,colorDarkBrown)    	
	    }
	   

	    // Randomly shape the candle holder thing
	    currentSides := [2]int{3+o,o+2+candleWidth}
	    lastDirection := 0.0;
	    lastColor := colorBrown;
	  	simp := opensimplex.NewWithSeed(int64(o+1241512));

	    for i := 0; i < baseHeight; i++ {
	   
	   		direction := simp.Eval2(float64(i),0.0)
	    	if (((lastDirection < 0 && direction > 0 ) ||  (lastDirection > 0 && direction < 0))&& lastColor!=colorDarkBrown) {
	    		img.Set(currentSides[0],51+i,colorDarkBrown)    	
	    		img.Set(currentSides[1],51+i,colorDarkBrown)
	    		for z := currentSides[0]; z < currentSides[1]; z++ {
	    			img.Set(z,51+i,colorDarkBrown)    	
	    		}
	    		lastColor = colorDarkBrown
	    	}else{
	    		img.Set(currentSides[0],51+i,colorBrown)    	
	    		img.Set(currentSides[1],51+i,colorBrown)
	    		for z := currentSides[0]; z < currentSides[1]; z++ {
	    			img.Set(z,51+i,colorBrown)    	
	    		}
	    		lastColor = colorBrown
	    	}

	    	lastDirection = direction


	   		if direction > 0.3 {
	   			currentSides[0],currentSides[1] = currentSides[0]-1,currentSides[1]+1
	   		}else if direction < -0.3{
	   			currentSides[0],currentSides[1] = currentSides[0]+1,currentSides[1]-1
	   		}


	    }
	}
    // Save to out.png
    f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
    defer f.Close()
    png.Encode(f, img)
    open.Run("out.png")

}