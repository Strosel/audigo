package audigo

import "math"

//Wave Describeas a soundwave as a set of points
type Wave []int16

func GenerateSine(env Envelope, sampleRate uint32, frequency ...float64) Wave {
	bufferSize := uint32(float64(sampleRate) * env.Duration().Seconds())

	timePeriod := make([]float64, len(frequency))
	for i, freq := range frequency {
		timePeriod[i] = (math.Pi * 2 * freq) / float64(sampleRate)
	}

	sinesum := make([]float64, bufferSize)
	sinemax := 0.0
	for i := range sinesum {
		sum := 0.0
		for _, tp := range timePeriod {
			sum += math.Sin(tp * float64(i))
		}

		if sum > sinemax {
			sinemax = sum
		}

		sinesum[i] = sum
	}

	data := make(Wave, bufferSize)
	for i, sine := range sinesum {
		sec := (float64(i) / float64(bufferSize)) * float64(env.Duration())
		amp := env.Amplitude(sec)

		data[i] = int16(amp * sine * (1 / sinemax))
	}

	return data
}
