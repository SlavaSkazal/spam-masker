package masking

func FindAndMaskLinks(sourceStr string) string {
	const mask = "http://"

	if len(sourceStr) < len(mask) {
		return sourceStr
	}

	maskingStr, maskedStr := make([]byte, len(sourceStr)), make([]byte, len(sourceStr))
	copy(maskingStr, sourceStr)
	copy(maskedStr, sourceStr)

	spaceByte := []byte(" ")[0]
	starByte := []byte("*")[0]
	var itsLink bool
	var indBeginLink int

	for i := 0; i < len(maskingStr); i++ {
		if !itsLink && i > len(maskingStr)-len(mask) {
			break
		}

		if !itsLink && string(maskingStr[i:i+len(mask)]) == mask && (i == 0 || maskingStr[i-1] == spaceByte) {
			itsLink = true
			i += len(mask)
			indBeginLink = i
		}

		if itsLink && (maskingStr[i] == spaceByte || i+1 >= len(maskingStr)) {
			lenStars := i - indBeginLink
			if i+1 >= len(maskingStr) {
				lenStars++
			}

			for k := 0; k < lenStars; k++ {
				maskedStr[indBeginLink] = starByte
				indBeginLink++
			}
			itsLink = false
		}
	}

	return string(maskedStr)
}
