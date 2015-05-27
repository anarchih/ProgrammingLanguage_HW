package main
import (
    "fmt"
    "image"
    "image/jpeg"
    "os"
    "image/color"
    "image/draw"
    "math"
    "time"
    "runtime"
)

func init() {
    image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

func imageToSlice(img image.Image) [][][]uint8 {
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
            test[y][x][0], test[y][x][1], test[y][x][2] = uint8(rIn),
                                                          uint8(gIn),
                                                          uint8(bIn)
        }
    }
    return test
}

func rgbToGray(arr [][][]uint8, width int, height int) [][]uint8 {
    test := make([][]uint8, height)
    for y := 0; y < height; y += 1 {
        test[y] = make([]uint8, width)
        for x := 0; x < width; x += 1 {
            var gray uint32
            rIn, gIn, bIn := uint32(arr[y][x][0]),
                             uint32(arr[y][x][1]),
                             uint32(arr[y][x][2])
            gray = (rIn * 30 + gIn * 59 + bIn * 11 + 50) / 100
            test[y][x] = uint8(gray)
        }
    }
    return test
}

func sobel(
        arr [][]uint8,
        result [][]uint8,
        width int,
        height int,
        start int,
        c chan int,
        flag int) {

    Sx := [][]int {{-1, 0, 1},{-2, 0, 2}, {-1, 0, 1}}
    Sy := [][]int {{-1, -2, -1},{0, 0, 0}, {1, 2, 1}}
    for y := 1; y < height - 1; y += 1 {
        result[y] = make([]uint8, width)
        for x := 0; x < width; x += 1 {
            if x == 0 || x == width - 1 {
                continue
            }else {
                Gx, Gy := 0, 0
                for i := 0; i < 3; i += 1 {
                    for j := 0; j < 3; j += 1 {
                        tmp := int(arr[y - 1 + j + start][x - 1 + j])
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
    if flag == 1 {
        c <- 1
    }
}

func concurrencySobel(arr [][]uint8, result [][]uint8, width int, height int) {
    c := make(chan int)
    result[0] = make([]uint8, width)
    tmp := height / 2
    go sobel(arr, result[0:tmp], width, tmp, 0, c, 1)
    go sobel(arr, result[tmp - 2:], width, height - (tmp - 2), tmp - 2, c, 1)
    result[height - 1] = make([]uint8, width)
    _ = <-c
    _ = <-c
}

func singleSobel(arr [][]uint8, result [][]uint8, width int, height int) {
    result[0] = make([]uint8, width)
    c := make(chan int)
    sobel(arr, result, width, height, 0, c, 0)
    result[height - 1] = make([]uint8, width)
}

func main() {
    runtime.GOMAXPROCS(4)
    // read file
    imgfile, err := os.Open("data/test.jpg")

    if err != nil {
        fmt.Println("img.jpg file not found!")
        os.Exit(1)
    }

    defer imgfile.Close()

    imgIn, _, err := image.Decode(imgfile)
    width := imgIn.Bounds().Max.X
    height := imgIn.Bounds().Max.Y


    x := imageToSlice(imgIn)
    arr := rgbToGray(x, width, height)
    result := make([][]uint8, height)

    start := time.Now()

    concurrencySobel(arr, result, width, height)

    elapsed := time.Since(start)
    fmt.Println(elapsed)
    imgOut, err := os.Create("output/output.jpg")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }


    imgRect := image.Rect(0, 0, width, height)
    img := image.NewRGBA(imgRect)
    draw.Draw(img, img.Bounds(),
              &image.Uniform{color.White}, image.ZP, draw.Src)
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
                      image.ZP,
                      draw.Src)
        }
    }
    var opt jpeg.Options

    opt.Quality = 100

    err = jpeg.Encode(imgOut, img, &opt)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
