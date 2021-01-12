package main

import (
	"fmt"
	"time"
	"./myBcrypt"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)


const salt = "OOv2wLxbNjUxVcc1sjysau"

var n = 0

func hashPassword(password []byte) (string, error) {
	hashedPassword, err := myBcrypt.GenerateFromPassword(password, myBcrypt.DefaultCost, salt)
    if err != nil {
        return "", err
		
    }
    return string(hashedPassword), nil
}


func hashHandler(ctx *fasthttp.RequestCtx) {
	query := ctx.QueryArgs()
	password := query.Peek("password")

	enc, err := hashPassword(password)
	if err != nil {
		fmt.Print(err)
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}
	fmt.Fprintf(ctx, "{\"hashedPassword\":\"%s\"}", enc)
	n++
}


func main() {
	r := router.New()
	r.GET("/hash", hashHandler)
	fmt.Println("Server started...")
	go func() {
		for true {
			fmt.Printf("\rRequest #%d", n)
			time.Sleep(time.Second * 1)
		}
	}()
	fasthttp.ListenAndServe(":14567", r.Handler)
}