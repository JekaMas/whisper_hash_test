package bloom_filter

import (
	"testing"
	"fmt"
)

func TestBloom(t *testing.T) {
	top := [4]byte{0,0,0,0}
	hashes := make(map[string]struct{})
	var hash string
	var collisionCount int
	var count uint64
	total := uint64(256*256*256*256)

	for i0:=0; i0<256; i0++ {
		top[0] = byte(i0)

		for i1:=0; i1<256; i1++ {
			top[1] = byte(i1)

			for i2:=0; i2<256; i2++ {
				top[2] = byte(i2)

				for i3:=0; i3<256; i3++ {
					top[3] = byte(i3)

					hash = string(TopicToBloom(TopicType(top)))

					_, ok := hashes[hash]
					if ok {
						collisionCount++
					} else {
						hashes[hash] = struct{}{}
					}

					count++
					if count%100000000 == 0 {
						fmt.Printf("Total count %v\tCollisions %v\tProbability of collision %.2f\tProgress %.2f\n",
							count, collisionCount, float64(collisionCount)/float64(count)*100, float64(count)/float64(total)*100)
					}
				}
			}
		}
	}

	fmt.Println("\n\nFinal result:")
	fmt.Printf("Total count %v\tCollisions %v\tProbability of collision %.2f\tProgress %.2f\n",
		count, collisionCount, float64(collisionCount)/float64(count)*100, float64(count)/float64(total)*100)
	fmt.Println("Unique values:", len(hashes))
}
