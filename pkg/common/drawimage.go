package common

import (
    "os"
    "image"
    "time"
    "strings"
    "math/rand"
    "image/draw"
    "image/color"
    "image/png"
    "io/ioutil"

    "github.com/golang/freetype"
    "github.com/golang/freetype/truetype"
    "golang.org/x/image/font"
)


type DrawTextInfo struct {
    Text string 
    X   int 
    Y   int 
}


type TextBrush struct {
    FontType      *truetype.Font  
    FontSize      float64 
    FontColor     *image.Uniform 
    TextWidth     int 
}




func DrawLogo(corpname,srcdir,imagePath string)error{
    data := strings.Split(corpname,"")
    texts := []*DrawTextInfo{&DrawTextInfo{Text: data[0]+" "+data[1],X:8,Y:60},&DrawTextInfo{Text:  data[2]+" "+data[3],X:8,Y:130}}
    if err := DrawStringOnImageAndSave(imagePath,srcdir,texts);err != nil{
       return  err 
    }
    return  nil 
}

func NewTextBrush(FontFilePath string,FontSize float64,FontColor *image.Uniform,textWidth int)(*TextBrush,error){
    fontFile,err := ioutil.ReadFile(FontFilePath)
    if err != nil{
        return nil,err
    }
    fontType,err := truetype.Parse(fontFile)
    if err != nil {
        return nil, err
    }
    if textWidth <= 0 {
        textWidth = 42
    }
    return &TextBrush{FontType:fontType,FontSize:FontSize,FontColor:FontColor,TextWidth:textWidth},nil
}



func DrawStringOnImageAndSave(imagePath,srcdir string,infos []*DrawTextInfo)(err error){
    const width, height = 150, 150
    rand.Seed(time.Now().UnixNano())
    tmpcolor := rand.Intn(255)
    for {
        if tmpcolor < 245{
            break
        }
        tmpcolor = rand.Intn(255)
    }

    background := image.NewRGBA(image.Rect(0, 0, width, height))
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            background.Set(x, y, color.NRGBA{
                R: uint8((tmpcolor) & 255),
                G: uint8((tmpcolor) << 1 & 255),
                B: uint8((tmpcolor) << 2 & 255),
                A: 255,
            })
        }
    }

    des := image.NewRGBA(background.Bounds()) 
    textBrush,_ := NewTextBrush(srcdir+"/PingFang.ttf",60,image.White,60)
    c := freetype.NewContext()
    c.SetDPI(72)
    c.SetFont(textBrush.FontType)
    c.SetHinting(font.HintingFull)
    c.SetFontSize(textBrush.FontSize)
    c.SetClip(des.Bounds()) 
    c.SetDst(des)
    c.SetSrc(textBrush.FontColor)

    for _, info := range infos{
        c.DrawString(info.Text,freetype.Pt(info.X,info.Y))
    }
    
    draw.Draw(background,background.Bounds(),des,image.Pt(0,0),draw.Over)
    fSave, err := os.Create(imagePath)
    if err != nil{
        return err
    }
    defer fSave.Close()
    err = png.Encode(fSave,background)
    if err != nil{
        return err
    }
    return nil 

}

