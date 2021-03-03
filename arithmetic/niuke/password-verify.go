package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) <= 8 {
			fmt.Println("NG")
			continue
		}
		if verify(text) {
			fmt.Println("OK")
			continue
		}
		fmt.Println("NG")
	}

}

func verify(password string) bool {
	tmp := make(map[int]struct{})
	for _, t := range password {
		if len(tmp) >= 3 {
			break
		}
		if t >= '0' && t <= '9' {
			tmp[0] = struct{}{}
		} else if t >= 'a' && t <= 'z' {
			tmp[1] = struct{}{}
		} else if t >= 'A' && t <= 'Z' {
			tmp[2] = struct{}{}
		} else {
			tmp[3] = struct{}{}
		}
	}
	if len(tmp) < 3 {
		return false
	}
	for i := 1; i < len(password)-6; i++ {
		if strings.Count(password, password[i:i+3]) > 1 {
			return false
		}
	}
	return true
}

/*
题目描述
密码要求:

1.长度超过8位
2.包括大小写字母.数字.其它符号,以上四种至少三种
3.不能有相同长度大于2的子串重复

输入描述:
一组或多组长度超过2的子符串。每组占一行

输出描述:
如果符合要求输出：OK，否则输出NG

示例1
输入
021Abc9000
021Abc9Abc1
021ABC9000
021$bc9000
输出
OK
NG
NG
OK


^Awsl7w222*hV
0p^6#HJ5v4@3zxQv#~1q85
$5^$q6
s*z8F1O~7~3%83~y46~6&5w15)F46-7483
z(9Zh+$ZT*8(9c9&P202*CFM
pYr71#@nG5@)
8#I682T090!32gth&Wl)O*!H-&-%$En^2*3S$-
w3214(08$*o*7j$I-6!5$1b90974b25)p550#D*)@^SU%
W(4+1Q2#*&*aX3+02ig8m
W!8nDUP0l$m6!k1en8)j4-JQ+
5o04ac-l#$-eM021$6*66n9b#o7d06~)^f4
6~@
Yw-09-Q78G**81O&J210
))20Uq%0
L83M!g!$~!7&u2B-#ovw~7LKRz5+!
!3^@^%45*892nv&50~0S6e2y9U!&i~8iWu2K8z
~32@M#
3&!8-~6c1-0l#!ommNW%2dJ6U)0JmdR@586!7#+4+-e^)2$
++^(348~8N^d$19@2GU3&(3$07i@tt1#)S-*6!(+!34u
$y2*+@^7w8~4^0+EN7%1#n94Fpdb4&
55
a
R!g9FpT78$F8WAq~!#(491#F67bf*23X6^48Q
00W^20&+L$+B&tH!12^%))4@3F^*OVK(6l6#%27ja
(-%&5Et6-2n&%#Tk38(iz&0A7qa8u9((
&-4d3$)9&9nMA3@%J0@#~X8wi17kv1V*88%9X2$h+R3@#5v
b225H04-!~33#
31L$(!i2^8m$2^p
p3L^04Pg5923!2$$9pE&009Kjm&1d5(
r-^$t6@9O)ReC(8^7b9!T-g0+y+5!
n6$++Us66h20!E61$%E%3gMw#-0!
^-$ezL%-$30k2plAN9+i8544$B-($aS611#H*0
*+5B459-u17&s9$6JRJ@%MD07-5^03^806SN44I@$(*y5-~$2(
05ag
-CUuJ5Pg39~Wby@^$093Y)kdQ8&Zg51*
)-q4(@~l~)ATN
)37^%%
jC0$4T003#v6*@65&w87a6r3^!x(r%OA28540G33Nl!3
&Lh3bvG@!2mGk%+3k57N8s
$##467F61vI5JN204oB@hK!0
35HI^k89&x4)+107~cJ-5d258Y*)4%mPHs^8
07A4D8$X9C8pZ*4v690y#W5#@0$0-+B99Jfd
ZH^7%0k3(n5460217!-41H%
dSwg*6r$%6tQr(
z5!hn%&1$67s*c
-&~OBtk#$UCI1-q!Dc(nx43!W8Up!*4-6&12&ba$3~4-
80-K!+@#y67854*7%GbB1(*~b2I^u9jCE5~999N0$%~
#~5^R7G!&k(7P-%48@m
-59V*R(j+~*
7!(@0f#+C0$(7$00-A#A4+-KxH8M&^l-8-%9+MB+g4X+b%-*
#Z@#!xu6x9
A!(630sKb06$#bI12~!4!3%%Sa1H+2%(~8G-3A4#)
@+##+*#g534+$+05$D@4~%Z5370
F858@TG%&Y0@2$Ba0C&g
6V1W49o~ow%S@W8@@
z87
47$Z^
&+36grQ7%9+6g0D*925oBp
+19%55@+2oW^i79R)&JGL8&64@836~%%H318l@5G1I)&&@54&1
I@4-(85$-6!B@4+j)@
CS6(5%x50v%#3D686&7@C!1!-kd-pf7r16
*/
