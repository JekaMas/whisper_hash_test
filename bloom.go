package bloom_filter

type TopicType [4]byte
const bloomFilterSize = 64

// from `whisper/whisperv6/envelope.go:243`
func TopicToBloom(topic TopicType) []byte {
	b := make([]byte, bloomFilterSize)
	var index [3]int
	for j := 0; j < 3; j++ {
		index[j] = int(topic[j])
		if (topic[3] & (1 << uint(j))) != 0 {
			index[j] += 256
		}
	}

	for j := 0; j < 3; j++ {
		byteIndex := index[j] / 8
		bitIndex := index[j] % 8
		b[byteIndex] = (1 << uint(bitIndex))
	}
	return b
}
