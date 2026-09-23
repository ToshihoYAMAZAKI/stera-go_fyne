package internal

import (
	"log"
	"math"
	"math/rand"
	"time"
)

// RandStory は建物階数をランダムに設定する
func RandStory(anum int) (story int) {
	seed := time.Now().UnixNano()
	log.Println("seed=", seed)
	r := rand.New(rand.NewSource(seed))
	log.Println("r=", r)
	val := r.Float64()
	log.Println("val=", val)
	// s = int(val)
	// log.Println("s=", s)

	var trh float64
	if anum == 1 || anum == 2 || anum == 12 {
		trh = 1.0
	} else if anum == 11 {
		trh = 0.95
	} else if anum == 3 || anum == 4 || anum == 5 || anum == 6 {
		trh = 0.90
	} else if anum == 7 || anum == 8 || anum == 10 {
		trh = 0.85
	} else if anum == 9 {
		trh = 0.80
	}
	if val <= trh {
		spr := 0.700458600675584 * math.Exp(2.23524816406936*val)
		log.Println("spr=", spr)
		story = int(spr + 0.5)
		log.Println("story1=", story)
	} else {
		spr1 := 0.700458600675584 * math.Exp(2.23524816406936*trh)
		spr2 := 3.35035038917538 * math.Exp(1.36764285404573*(val-trh)*10/((1-trh)*10))
		spr3 := 3.35035038917538 * math.Exp(1.36764285404573*0.0)
		sprt := spr1 + spr2 - spr3
		log.Println("sprt=", sprt)
		story = int(sprt + 0.5)
		log.Println("story2=", story)
	}

	return story
}
