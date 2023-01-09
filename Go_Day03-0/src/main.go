package main

import (
	"htpp/api"
	"log"
	"net/http"
)

type Gun struct {
    On bool
    Ammo, Power int
}

func (g *Gun) Shoot() bool {
    if g.On && g.Ammo != 0 {
        g.Ammo--
        return true
    }
    return false
}
func (g *Gun) RideBike() bool {
    if g.On && g.Power != 0 {
        g.Power--
        return true
    }
    return false
}


func main() {
	testStruct := &Gun{}
	srv := api.NewServer()
	if err := http.ListenAndServe(":8081", srv); err != nil {
		log.Fatalln(err)
	}
}
