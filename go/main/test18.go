package main

import (
    "fmt"
    "image/gif"
    "image/png"
    "os"
    "strings"
)
func main(){
    width, height, err := ImageGetInfo("/smartrtb/smartrtb-server/files/IMAGE/5c8e52dc39c5859b2a7ddab0c9268c00_1369785.png")
    if err != nil {
       fmt.Println(err)
       return
    }
    fmt.Println(width, height)
}
func ImageGetInfo(srcFullFile string) (width int, height int, error error) {
    suffix := srcFullFile[strings.LastIndex(srcFullFile, ".")+1:] //文件后缀
    file, _ := os.Open(srcFullFile)
    defer file.Close()
    if suffix == "gif" {
        img, err := gif.DecodeConfig(file)
        if err != nil {
            return 0, 0, err
        }
        return img.Width, img.Height, nil
    } else { //png
        img, err := png.DecodeConfig(file)
        if err != nil {
            return 0, 0, err
        }
        return img.Width, img.Height, nil
    }
}
