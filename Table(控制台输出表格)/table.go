package main

/*
	原作者：https://github.com/modood/table/blob/master/table.go
*/

import (
	"fmt"
	"reflect"
)

//bd 表格结构体
type bd struct {
	H  rune // 制表符水平 '─'
	V  rune // 制表符垂直 '│'
	VH rune // 制表符垂直和水平 '┼'
	HU rune // 制表符水平和向上 '┴'
	HD rune // 制表符水平和向下 '┬'
	VL rune // 制表符垂直和向左 '┤'
	VR rune // 制表符垂直和向右 '├'
	DL rune // 制表符向下和向左 '┐'
	DR rune // 制表符下和右 '┌'
	UL rune // 制表符上和左 '┘'
	UR rune // 制表符上和右 '└'

}

// 定义制表符
var m = map[string]bd{
	"box-drawing": bd{'─', '│', '┼', '┴', '┬', '┤', '├', '┐', '┌', '┘', '└'},
}

//Output 输出表格的主函数
func Output(slice interface{}) {
	fmt.Println(Table(slice))
}

//Table 格式化数据
func Table(slice interface{}) string {
	coln, colw, rows := parse(slice)
	table := table(coln, colw, rows, m["box-drawing"])
	return table
}

//parse(解析传过来的切片数组)
func parse(slice interface{}) (
	coln []string, //表格第一行的title
	colw []int, //表格的行宽
	rows [][]string, //表格每行的内容
) {
	for i, u := range sliceconv(slice) {
		v := reflect.ValueOf(u)
		t := reflect.TypeOf(u)
		if v.Kind() != reflect.Struct {
			panic("Table: items of slice should be on struct value")
		}
		var row []string

		m := 0
		for n := 0; n < v.NumField(); n++ {
			if t.Field(n).PkgPath != "" {
				m++
				continue
			}
			cn := t.Field(n).Name
			cv := fmt.Sprintf("%+v", v.FieldByName(cn).Interface())
			//填充表格第一行标题及计算行宽
			if i == 0 {
				coln = append(coln, cn)
				colw = append(colw, len(cn))
			}
			//统计每格的宽度
			if colw[n-m] < len(cv) {
				colw[n-m] = len(cv)
			}
			//填充单行数据
			row = append(row, cv)
		}
		//填充每行数据
		rows = append(rows, row)
	}
	//返回第一行title数组，所有行宽的数组，内容
	return coln, colw, rows
}

//table 根据传来的所需参数绘制表格
func table(coln []string, colw []int, rows [][]string, b bd) (table string) {
	//设定表格的顶部 '┌' '│' '├'
	head := [][]rune{[]rune{b.DR}, []rune{b.V}, []rune{b.VR}}
	//'└'
	bttm := []rune{b.UR}
	//遍历所有行宽
	for i, v := range colw {
		//绘制左上角格子
		head[0] = append(head[0], []rune(repeat(v+2, b.H)+string(b.HD))...)

		head[1] = append(head[1], []rune(" "+coln[i]+repeat(v-len(coln[i])+1, ' ')+string(b.V))...)
		head[2] = append(head[2], []rune(repeat(v+2, b.H)+string(b.VH))...)
		bttm = append(bttm, []rune(repeat(v+2, b.H)+string(b.HU))...)
	}
	head[0][len(head[0])-1] = b.DL
	head[2][len(head[2])-2] = b.VL
	bttm[len(bttm)-1] = b.UL

	var body [][]rune
	for _, r := range rows {
		row := []rune{b.V}
		for i, v := range colw {
			l := length([]rune(r[i]))
			row = append(row, []rune(" "+r[i]+repeat(v-l+1, ' ')+string(b.V))...)
		}
		body = append(body, row)
	}

	for _, v := range head {
		table += string(v) + "\n"
	}
	for _, v := range body {
		table += string(v) + "\n"
	}
	table += string(bttm)

	return table
}

/*-==========================辅助函数=========================-*/

/*
通过Value实现泛型
为了解决method接受不同类型的slice为入参，可以用反射来完成。对于可记长度和可随机访问的类型，可以通过v.Len()和v.Index(i)获取他们的第几个元素。
v.Index(i).Interface()将reflect.Value反射回了interface类型
来源:https://ninokop.github.io/2017/10/30/Go-%E5%8F%8D%E5%B0%84%E4%B8%8Einterface%E6%8B%BE%E9%81%97/
代码参考来源:https://segmentfault.com/q/1010000000198391
*/
func sliceconv(slice interface{}) []interface{} {
	//通过反射检测传入的slice的值的类型，不符合类型的报错退出程序
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		panic("sliceconv: param \"slice\" should be on slice value")
	}
	//如果符合要求,通过反射获取数组中的元素个数
	l := v.Len()
	//初始化一个长度为l(上一步得到的元素个数)，类型为interface{}的数组
	r := make([]interface{}, l)
	for i := 0; i < l; i++ {
		//通过Interface()方法恢复reflect对象为一个interface
		r[i] = v.Index(i).Interface()
	}
	return r
}

//repeat 重复制表符
func repeat(time int, char rune) string {
	var s = make([]rune, time)
	for i := range s {
		s[i] = char
	}
	return string(s)
}

//length 处理非ascii字符
func length(r []rune) int {
	//CJK(Chinese, Japanese, Korean)
	type cjk struct {
		from rune
		to   rune
	}

	// References:
	// -   [Unicode Table](http://www.tamasoft.co.jp/en/general-info/unicode.html)
	// -   [汉字 Unicode 编码范围](http://www.qqxiuzi.cn/zh/hanzi-unicode-bianma.php)
	var a = []cjk{
		{0x2E80, 0x9FD0},   // Chinese, Hiragana, Katakana, ...
		{0xAC00, 0xD7A3},   // Hangul
		{0xF900, 0xFACE},   // Kanji
		{0xFE00, 0xFE6C},   // Fullwidth
		{0xFF00, 0xFF60},   // Fullwidth again
		{0x20000, 0x2FA1D}, // Extension
		// More? PRs are aways welcome here.
	}
	length := len(r)
l:
	for _, v := range r {
		for _, c := range a {
			if v >= c.from && v <= c.to {
				length++
				continue l
			}
		}
	}
	return length
}

/*===========================测试部分===============================*/
type House struct {
	Name  string
	Sigil string
	Motto string
}

func main() {
	hs := []House{
		{"Stark", "direwolf", "Winter is coming"},
		{"Targaryen", "dragon", "Fire and Blood"},
		{"Lannister", "lion", "Hear Me Roar"},
	}

	// Output to stdout
	Output(hs)

	// Or just return table string and then do something
	s := Table(hs)
	fmt.Println(s)
}
