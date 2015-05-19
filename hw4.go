package main
import (
    "fmt"
    "image"
    "image/jpeg"
    "os"
    "image/color"
    "image/draw"
)

func init() {
    image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

func imageToArray(img image.Image) [][][]uint8{
    test := make([][][]uint8, 538)
    for y := 0; y < 538; y += 1 {
        test[y] = make([][]uint8, 718)
        for x := 0; x < 718; x += 1 {
            test[y][x] = make([]uint8, 3)
            a := img.At(x, y)
            rIn, gIn, bIn, _ := a.RGBA()
            rIn, gIn, bIn = rIn / 257, gIn / 257, bIn / 257
            test[y][x][0], test[y][x][1], test[y][x][2] = uint8(rIn), uint8(gIn), uint8(bIn)
        }
    }
    return test
}

func rgbToGray(arr [][][]uint8) [][]uint8{
    test := make([][]uint8, 538)
    for y := 0; y < 538; y += 1 {
        test[y] = make([]uint8, 718)
        for x := 0; x < 718; x += 1 {
            var gray uint32
            rIn, gIn, bIn := uint32(arr[y][x][0]), uint32(arr[y][x][1]), uint32(arr[y][x][2])
            gray = (rIn * 30 + gIn * 59 + bIn * 11 + 50) / 100
            test[y][x] = uint8(gray)
        }
    }
    return test

}

func main() {

    // read file
    imgfile, err := os.Open("data/test.jpg")

    if err != nil {
        fmt.Println("img.jpg file not found!")
        os.Exit(1)
    }

    defer imgfile.Close()

    imgIn, _, err := image.Decode(imgfile)
    a := imgIn.At(0, 0)
    rIn, gIn, bIn, _ := a.RGBA()
    fmt.Println(rIn, gIn, bIn)
    bounds := imgIn.Bounds()
    fmt.Println(bounds)


    x := imageToArray(imgIn)
    arr := rgbToGray(x)

    fmt.Println(arr[1][1])
    imgOut, err := os.Create("output/output.jpg")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }





    imgRect := image.Rect(0, 0, 718, 538)
    img := image.NewRGBA(imgRect)
    draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)
    for y := 0; y < 538; y += 1 {
        for x := 0; x < 718; x += 1 {
            draw.Draw(img, image.Rect(x, y, x+1, y+1), &image.Uniform{color.RGBA{arr[y][x], arr[y][x], arr[y][x], 0}}, image.ZP, draw.Src)
        }
    }
    var opt jpeg.Options

    opt.Quality = 100

    err = jpeg.Encode(imgOut, img, &opt)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Println("Generated image to output.jpg \n")
}