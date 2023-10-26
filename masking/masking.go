package masking

func FindAndMaskLinks(sourceStr string) string {
	maskingStr := make([]byte, len(sourceStr))
	copy(maskingStr, sourceStr)

	const maskHttp = "http://"
	spaceByte := []byte(" ")[0]
	starByte := []byte("*")[0]
	maskedStr := []byte(maskingStr)
	var itsLink bool
	var indBeginLink, indFinalLink int
	var y int

	for i := 0; i < len(maskingStr); i++ {
		if !itsLink {
			if y+1 >= len(maskHttp) {
				itsLink = true
				indBeginLink = i + 1
			} else if maskingStr[i] == maskHttp[y] {
				if y != 0 || i == 0 || maskingStr[i-1] == spaceByte {
					y++
				}
			} else {
				y = 0
			}
		} else {
			if maskingStr[i] == spaceByte {
				indFinalLink = i - 1
			} else if i+1 == len(maskingStr) {
				indFinalLink = i
			}
		}

		if indBeginLink > 0 && indFinalLink > 0 {
			endMaskedStr := maskedStr[indBeginLink:]
			lenStars := indFinalLink - indBeginLink + 1
			for k := 0; k < lenStars; k++ {
				endMaskedStr[k] = starByte
			}

			maskedStr = append(maskedStr[:indBeginLink], endMaskedStr...)
			indBeginLink, indFinalLink, y = 0, 0, 0
			itsLink = false
		}
	}
	return string(maskedStr)
}
