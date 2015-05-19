package main
import (
    "fmt"
    "image"
    "image/jpeg"
    "os"
    "image/color"
    "image/draw"
    "math"
)

func init() {
    image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

func imageToSlice(img image.Image) [][][]uint8{
    width := img.Bounds().Max.X
    height := img.Bounds().Max.Y
    test := make([][][]uint8, height)
    for y := 0; y < height; y += 1 {
        test[y] = make([][]uint8, width)
        for x := 0; x < width; x += 1 {
            test[y][x] = make([]uint8, 3)
            a := img.At(x, y)
            rIn, gIn, bIn, _ := a.RGBA()
            rIn, gIn, bIn = rIn / 257, gIn / 257, bIn / 257
            test[y][x][0], test[y][x][1], test[y][x][2] = uint8(rIn), uint8(gIn), uint8(bIn)
        }
    }
    return test
}

func rgbToGray(arr [][][]uint8, width int, height int) [][]uint8{
    test := make([][]uint8, height)
    for y := 0; y < height; y += 1 {
        test[y] = make([]uint8, width)
        for x := 0; x < width; x += 1 {
            var gray uint32
            rIn, gIn, bIn := uint32(arr[y][x][0]), uint32(arr[y][x][1]), uint32(arr[y][x][2])
            gray = (rIn * 30 + gIn * 59 + bIn * 11 + 50) / 100
            test[y][x] = uint8(gray)
        }
    }
    return test

}

func sobel(arr [][]uint8, result [][]uint8, width int, height int){
    Sx := [][]int {{-1, 0, 1},{-2, 0, 2}, {-1, 0, 1}}
    Sy := [][]int {{-1, -2, -1},{0, 0, 0}, {1, 2, 1}}
    for y := 0; y < height; y += 1 {
        result[y] = make([]uint8, width)
        for x := 0; x < width; x += 1 {
            if y == 0 || y == height - 1 || x == 0 || x == width - 1 {
                result[y][x] = 0
            }else {
                Gx, Gy := 0, 0
                for i := 0; i < 3; i += 1 {
                    for j := 0; j < 3; j += 1 {
                        tmp := int(arr[y - 1 + j][x - 1 + j])
                        Gx += tmp * Sx[j][i]
                        Gy += tmp * Sy[j][i]
                    }
                }
                G := math.Sqrt(float64(Gx * Gx) + float64(Gy * Gy))
                if G > 255 {
                    result[y][x] = 255
                }else {
                    result[y][x] = uint8(G)
                }
            }

        }
    }
    fmt.Println(result[1][1])
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
    width := imgIn.Bounds().Max.X
    height := imgIn.Bounds().Max.Y


    x := imageToSlice(imgIn)
    arr := rgbToGray(x, width, height)

    result := make([][]uint8, height)
    sobel(arr, result, width, height)
    fmt.Println(result[1][1])


    imgOut, err := os.Create("output/output.jpg")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }





    imgRect := image.Rect(0, 0, width, height)
    img := image.NewRGBA(imgRect)
    draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)
    for y := 0; y < height; y += 1 {
        for x := 0; x < width; x += 1 {
            draw.Draw(
                      img,
                      image.Rect(x, y, x+1, y+1),
                      &image.Uniform{color.RGBA{
                                                result[y][x],
                                                result[y][x],
                                                result[y][x],
                                                0}},
                      image.ZP, draw.Src)
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
