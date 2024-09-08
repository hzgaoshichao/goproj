package musicinterpreter

import (
	"fmt"
	"strconv"
	"strings"
)

type PlayContext struct {
	playText string
}

func (p *PlayContext) GetPlayText() string {
	return p.playText
}

func (p *PlayContext) SetPlayText(pt string) {
	p.playText = pt
}

type Express interface {
	Interpret(ctx *PlayContext, ex Express)
	execute(key string, value string)
}

type ExpressCommon struct {
}

func (ec *ExpressCommon) Interpret(ctx *PlayContext, ex Express) {
	if len(ctx.GetPlayText()) == 0 {
		return
	}

	playKey := string(ctx.GetPlayText()[:1])
	ctx.SetPlayText(ctx.GetPlayText()[2:])

	posBlank := strings.Index(ctx.GetPlayText(), " ")
	playValue := string(ctx.GetPlayText()[:posBlank])
	ctx.SetPlayText(ctx.GetPlayText()[posBlank+1:])

	// fmt.Printf("Key: %v Value: %v \n", playKey, playValue)

	ex.execute(playKey, playValue)

}

type Note struct {
	ExpressCommon
}

func (n *Note) execute(key string, value string) {
	note := ""
	switch key {
	case "C":
		note = "1"
	case "D":
		note = "2"
	case "E":
		note = "3"
	case "F":
		note = "4"
	case "G":
		note = "5"
	case "A":
		note = "6"
	case "B":
		note = "7"
	}

	fmt.Printf("%v ", note)
}

type Scale struct {
	ExpressCommon
}

func (s *Scale) execute(key string, value string) {
	scale := ""
	switch value {
	case "1":
		scale = "低音"
	case "2":
		scale = "中音"
	case "3":
		scale = "高音"
	}
	fmt.Printf("%v ", scale)
}

type Speed struct {
	ExpressCommon
}

func (s *Speed) execute(key string, value string) {
	speed := ""
	valuei, _ := strconv.Atoi(value)
	if valuei < 500 {
		speed = "快速"
	} else if valuei >= 1000 {
		speed = "慢速"
	} else {
		speed = "中速"
	}

	fmt.Printf("%v ", speed)
}
