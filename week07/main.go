package main

import (
	"fmt"
	"strings"
)

func main() {
	var f string = `	복무신조 
	우리의 결의 
	우리는 국가와 국민의 충성을 다하는 대한민국 육군이다.
	하나. 우리는 자유민주주의를 수호하며 조국통일의 역군이 된다.
	하나. 우리는 실전과 같은 훈련으로 지상전의승리자가 된다.
	하나. 우리는 법규를 준수하고 상관의 명령에 복종한다.
	하나. 우리는 명예와 신의를 지키며 전우애로 굳게 단결한다.`
	r := strings.NewReplacer("우", "Woo")

	// fmt.Println(f)
	fmt.Println(r.Replace(f))
}
