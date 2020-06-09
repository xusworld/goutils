package test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

type SupportedMiniProgramType struct {
	SupportedMiniProgramType []uint32 `json:"support_mini_program_type,omitempty"`
}

// 获取终端支持的小程序类型
func getSupportMPType(str string) []uint32 {

	// 注意，这里经常会带个\n结尾导致不合法,[3,1]\n,要做兼容
	sTmp := strings.Replace(str, `\n`, "", -1)

	fmt.Println("stmp ", sTmp)

	beg := strings.IndexAny(sTmp, "[")
	end := strings.IndexAny(sTmp, "]")

	sMini := string([]byte(sTmp)[beg : end+1])
	fmt.Println("sMini ", sMini)


	smpt := SupportedMiniProgramType{}

	if len(sMini) > 2 {
		strMini := "{\"support_mini_program_type\":" + sMini + "}"
		strMini = `{"support_mini_program_type":[11]}`
		fmt.Printf("sMini %s \n", strMini)
		err := json.Unmarshal([]byte(strMini), &smpt)
		if err != nil {
			fmt.Printf(" Error %s", err)
			return []uint32{}
		}
		fmt.Println("stSxxxxmpt", smpt)
		fmt.Println("stSxxxxmpt", smpt.SupportedMiniProgramType)
	}
	return smpt.SupportedMiniProgramType
}

func TestDemo(t *testing.T) {
	ret := getSupportMPType("[11]\n")
	fmt.Println(ret)
}