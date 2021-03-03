package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	names := make([]string, 0)
	lines := make([]int, 0)
	temp := make([]string, 0)
	for scanner.Scan() {
		text := scanner.Text()
		tmp := strings.Split(text, "\\")
		str := strings.Split(tmp[len(tmp)-1], " ")
		if len(str) != 2 {
			continue
		}
		name := str[0]
		line, err := strconv.Atoi(str[1])
		if err != nil {
			continue
		}
		if len(name) > 16 {
			temp = append(temp, name[len(name)-16:]+" "+str[1])
		} else {
			temp = append(temp, tmp[len(tmp)-1])
		}
		names = append(names, name)
		lines = append(lines, line)
	}
	type NameNum struct {
		name string
		num  int
		line int
	}
	rue := make(map[string]*NameNum, 0)
	order := make([]string, 0)
	for i := len(lines) - 1; i >= 0; i-- {
		if _, ok := rue[temp[i]]; ok {
			rue[temp[i]].num++
		} else {
			order = append(order, temp[i])
			rue[temp[i]] = &NameNum{
				name: names[i],
				num:  1,
				line: lines[i],
			}
		}

	}
	for i := len(order) - 1; i >= 0; i-- {
		nm, ok := rue[order[i]]
		if !ok {
			return
		}
		if len(nm.name) > 16 {
			nm.name = nm.name[len(nm.name)-16:]
		}
		fmt.Printf("%s %d %d\n", nm.name, nm.line, nm.num)
	}
}

/*
题目描述
开发一个简单错误记录功能小模块，能够记录出错的代码所在的文件名称和行号。

处理：

1、 记录最多8条错误记录，循环记录，最后只用输出最后出现的八条错误记录。对相同的错误记录只记录一条，但是错误计数增加。
    最后一个斜杠后面的带后缀名的部分（保留最后16位）和行号完全匹配的记录才做算是”相同“的错误记录。
2、 超过16个字符的文件名称，只记录文件的最后有效16个字符；
3、 输入的文件可能带路径，记录文件名称不能带路径。
4、循环记录时，只以第一次出现的顺序为准，后面重复的不会更新它的出现时间，仍以第一次为准

输入描述:
每组只包含一个测试用例。一个测试用例包含一行或多行字符串。每行包括带路径文件名称，行号，以空格隔开。

输出描述:
将所有的记录统计并将结果输出，格式：文件名 代码行数 数目，一个空格隔开，如：

示例1
输入
复制
D:\zwtymj\xccb\ljj\cqzlyaszjvlsjmkwoqijggmybr 645
E:\je\rzuwnjvnuz 633
C:\km\tgjwpb\gy\atl 637
F:\weioj\hadd\connsh\rwyfvzsopsuiqjnr 647
E:\ns\mfwj\wqkoki\eez 648
D:\cfmwafhhgeyawnool 649
E:\czt\opwip\osnll\c 637
G:\nt\f 633
F:\fop\ywzqaop 631
F:\yay\jc\ywzqaop 631
输出
复制
rzuwnjvnuz 633 1
atl 637 1
rwyfvzsopsuiqjnr 647 1
eez 648 1
fmwafhhgeyawnool 649 1
c 637 1
f 633 1
ywzqaop 631 2




G:\rp\onajqj\maahmq 631
E:\njfgjkcrh 641
C:\co\zk\ao\bxgxjfgrwckfxekeqro 629
D:\mf\si\jmfdahkeffyjjsf 646
E:\wn\arefkiz 633
C:\gpjleb\cinhhx\zjydgr\njfgjkcrh 640
E:\nwrrhx\qyw\bxgxjfgrwckfxekeqro 636
G:\usgsl\ywr\tve\cqekvaxypemktyurn 647
C:\jftbig\arefkiz 650
F:\rgk\cai\arefkiz 640
D:\tvse\vs\dhzrmy\njfgjkcrh 634
E:\coba\qbs\xagq\njfgjkcrh 628
F:\wnfsmf\oxrvbv\njfgjkcrh 632
C:\khqx\nv\jmfdahkeffyjjsf 637
F:\hm\ra\uaxckn\bxgxjfgrwckfxekeqro 647
D:\soq\jmfdahkeffyjjsf 642
F:\moxnw\szxcdhlaytgj 639
E:\avcop\jd\vwtrt\njfgjkcrh 650
E:\hou\vv\szxcdhlaytgj 631
C:\uozkwd\bxgxjfgrwckfxekeqro 650
F:\jmfdahkeffyjjsf 650
E:\hgoxms\nwax\szxcdhlaytgj 633
F:\vylww\zhh\cqekvaxypemktyurn 643
C:\njfgjkcrh 637
F:\bfn\dxwjje\jmfdahkeffyjjsf 632
E:\bxgxjfgrwckfxekeqro 634
G:\gwuusj\ized\qq\szxcdhlaytgj 646
F:\arefkiz 644
G:\zsw\uewu\arefkiz 634
E:\ja\zg\njfgjkcrh 644
D:\gfute\ju\wuy\szxcdhlaytgj 636
C:\mpgcx\kcgi\arefkiz 645
C:\zayn\jmfdahkeffyjjsf 648
F:\kkplu\avvw\hbzmwj\jmfdahkeffyjjsf 648
E:\maahmq 631
E:\hs\xnto\jmfdahkeffyjjsf 645
G:\cqekvaxypemktyurn 633
D:\maahmq 646
E:\jmfdahkeffyjjsf 636
G:\hbvm\szxcdhlaytgj 642


szxcdhlaytgj 636 1
arefkiz 645 1
jmfdahkeffyjjsf 648 2
jmfdahkeffyjjsf 645 1
qekvaxypemktyurn 633 1
maahmq 646 1
jmfdahkeffyjjsf 636 1
szxcdhlaytgj 642 1


arefkiz 645 1
jmfdahkeffyjjsf 648 2
maahmq 631 1
jmfdahkeffyjjsf 645 1
qekvaxypemktyurn 633 1
maahmq 646 1
jmfdahkeffyjjsf 636 1
szxcdhlaytgj 642 1
*/
