package main

import(
        "fmt"
        b64 "encoding/base64"
        "crypto/md5"
        _ "io"
        "flag"
        "reflect"
)

func main() {
        url := flag.String("url", "yeah", "url to shorten")
        flag.Parse()
        fmt.Println("url is ", *url)
        
        data := []byte(*url)
        md5ed := md5.Sum(data) //returns md5 checksum of the data
        fmt.Printf("%x", md5ed)

        fmt.Println(reflect.TypeOf(md5ed)) 
        fmt.Println(reflect.TypeOf(md5ed[:]))

        sEnc := b64.StdEncoding.EncodeToString(md5ed[:])

        fmt.Println(sEnc)
        fmt.Println("first 6: ", sEnc[0:6])
}
