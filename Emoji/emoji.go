package emoji

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"regexp"
	"unicode"
)

//需要转换的字符
const (
	ReplacePadding = " "
)

func CodeMap() map[string]string {
	return emojiCodeMap
}

//匹配规则
var flagRegexp = regexp.MustCompile(":flag-(a-z{2}):")

//字符与unicode字符匹配
func emojize(x string) string {
	str, ok := emojiCodeMap[x]
	if ok {
		return str + ReplacePadding
	}
	if match := flagRegexp.FindStringSubmatch(x); len(match) == 2 {
		return regionlIndicator(match[1][0]) + regionlIndicator(match[1][1])
	}
	return x
}

//将字符转换成Unicode字符
func regionlIndicator(i byte) string {
	return string('\U0001F1E6' + rune(i) - 'a')
}

//replaseEmoji 传入的字符与对应的unicode更换判断函数
func replaseEmoji(input *bytes.Buffer) string {
	//建立一个内容是":"的缓冲器(初始化缓冲器)
	emoji := bytes.NewBufferString(":")
	for {
		//读取传入参数的第一个字符(.ReadRune()读取缓冲器头部第一个rune,返回值:rune，size，error)
		i, _, err := input.ReadRune()
		if err != nil {
			//如果出错则直接返回emoji不做任何更换
			return emoji.String()
		}
		//如果第一个rune字符是 : ,并且此时emoji缓冲器长度是1
		//这个循环遍历了传入的字符，emoji缓冲器长度是1且rune的头部是 : 的时候才刚刚开始遍历这个字符
		if i == ':' && emoji.Len() == 1 {
			return emoji.String() + replaseEmoji(input)
		}
		//上面情况都不满足的的话，则将 i 写入emoji
		emoji.WriteRune(i)
		switch {
		//参考http://www.cnblogs.com/golove/p/3273585.html
		// IsSpace 判断 r 是否为一个空白字符
		// 在 Latin-1 字符集中，空白字符为：\t, \n, \v, \f, \r,
		// 空格, U+0085 (NEL), U+00A0 (NBSP)
		// 其它空白字符的定义有“类别 Z”和“Pattern_White_Space 属性”
		case unicode.IsSpace(i):
			//如果是空白字符则返回转换后的emoji
			return emoji.String()
		case i == ':':
			//当此时 i == ':'的时候，整个input中的缓冲器的字符已经读取完成，返回匹配完成的unicode字符
			return emojize(emoji.String())
		}
	}
}

func compile(x string) string {
	if x == "" {
		return ""
	}
	input := bytes.NewBufferString(x)
	output := bytes.NewBufferString("")

	for {
		i, _, err := input.ReadRune()
		if err != nil {
			break
		}
		switch i {
		default:
			output.WriteRune(i)
		case ':':
			output.WriteString(replaseEmoji(input))
		}
	}
	return output.String()
}

//*[]interface{} 空接口  接受任何类型
func compileValues(a *[]interface{}) {
	for i, x := range *a {
		if str, ok := x.(string); ok {
			fmt.Println(str)
			(*a)[i] = compile(str)
		}
	}
}

/*=================对 fmt 的 封 装==================*/

// Print is fmt.Print which supports emoji
func Print(a ...interface{}) (int, error) {
	compileValues(&a)
	return fmt.Print(a...)
}

// Println is fmt.Println which supports emoji
func Println(a ...interface{}) (int, error) {
	compileValues(&a)
	return fmt.Println(a...)
}

// Printf is fmt.Printf which supports emoji
func Printf(format string, a ...interface{}) (int, error) {
	format = compile(format)
	compileValues(&a)
	return fmt.Printf(format, a...)
}

// Fprint is fmt.Fprint which supports emoji
func Fprint(w io.Writer, a ...interface{}) (int, error) {
	compileValues(&a)
	return fmt.Fprint(w, a...)
}

// Fprintln is fmt.Fprintln which supports emoji
func Fprintln(w io.Writer, a ...interface{}) (int, error) {
	compileValues(&a)
	return fmt.Fprintln(w, a...)
}

// Fprintf is fmt.Fprintf which supports emoji
func Fprintf(w io.Writer, format string, a ...interface{}) (int, error) {
	format = compile(format)
	compileValues(&a)
	return fmt.Fprintf(w, format, a...)
}

// Sprint is fmt.Sprint which supports emoji
func Sprint(a ...interface{}) string {
	compileValues(&a)
	return fmt.Sprint(a...)
}

// Sprintf is fmt.Sprintf which supports emoji
func Sprintf(format string, a ...interface{}) string {
	format = compile(format)
	compileValues(&a)
	return fmt.Sprintf(format, a...)
}

// Errorf is fmt.Errorf which supports emoji
func Errorf(format string, a ...interface{}) error {
	compileValues(&a)
	return errors.New(Sprintf(format, a...))
}
