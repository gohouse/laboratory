package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := `ay:inline-block}.head_cont .logo_cont{width:154px;height:24px}.head_cont .scope_cont{margin-left:72px}.head_cont .logo{position:absolute;width:154px;height:24px;transform:translate(0,12px) scale(1.16);filter:drop-shadow(0 0 24px rgba(0,0,0,.45))}.head_cont .logo .squares .top_l{fill:#f26522}.head_cont .logo .squares .top_r{fill:#8dc63f}.head_cont .logo .squares .bom_l{fill:#00aeef}.head_cont .logo .squares .bom_r{fill:#ffc20e}.head_cont .logo .ms_text,.head_cont .logo .b_text{fill:white}.scope_cont,.logo_cont{z-index:3}.scopes,#idCont{display:none;z-index:3;margin:0;padding:0;vertical-align:text-bottom}.hpn_top_container{overflow:hidden}.bottom_row,.below_sbox{display:none}@media screen and (-ms-high-contrast:black-on-white){.logo,#hp_trivia_icon,.mappin,.musCardCont .share,.edit_interests{background-color:#000}#leftNavCaro,#rightNavCaro{background-color:#fff}}@media screen and (-ms-high-contrast:active){#leftNavCaro,#rightNavCaro{background-color:#000}}</style></head><body data-priority="2"><div class="hpapp"><div class="hp_body  "><div class="hpl"><div class="img_cont" style="background-image: url(/th?id=OHR.BentsGeneral_ROW0566713395_1920x1080.jpg&amp;rf=LaDigue_1920x1080.jpg); opacity: 1;"><div class="shaders"><div class="t`
	//matched, err := regexp.MatchString(`url\(/th\?id=(OHR\.BentsGeneral_ROW\d+)_\d+x\d+\.jpg`, str)
	compile := regexp.MustCompile(`url\(/th\?id=(OHR\.BentsGeneral_ROW\d+)_\d+x\d+\.jpg`)
	submatch := compile.FindAllStringSubmatch(str, -1)
	stringSubmatch := compile.FindStringSubmatch(str)
	fmt.Println(submatch)
	fmt.Println(stringSubmatch)
}