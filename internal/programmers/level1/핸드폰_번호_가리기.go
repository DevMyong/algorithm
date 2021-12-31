func solution(phoneNumber string) string {
	blurredPhoneNumber := ""
	for i, ch := range phoneNumber {
		if i+4 < len(phoneNumber){
			blurredPhoneNumber += "*"
		}else {
			blurredPhoneNumber += string(ch)
		}
	}
	return blurredPhoneNumber
}