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

//匹配规则 对国旗字符表情单独处理
var flagRegexp = regexp.MustCompile(":flag-(a-z{2}):")

//字符与unicode字符匹配
func emojize(x string) string {
	//将传入的表情字符做键查询map对应的值
	str, ok := emojiCodeMap[x]
	if ok {
		//如果找到对应的unicode字符则返回(这里在字符后面加个空格不影响，只是为了输出好看一点)
		return str + ReplacePadding
	}
	//如果没有找到对应的则进行自定义规则匹配
	//.FindStringSubmatch(x)返回一个 [][]string 数组，返回的数组如果是二维的说明匹配成功
	if match := flagRegexp.FindStringSubmatch(x); len(match) == 2 {
		//将两个字母分别转换成对应的unicode字符然后拼接起来
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
		/*这里是为了排除传入的表情符号前面是接着冒号的这种情况(eg:I like::apple:)*,出现这种情况则继续读取*/
		if i == ':' && emoji.Len() == 1 {
			return emoji.String() + replaseEmoji(input)
		}
		//上面情况都不满足的的话，则将 i 写入emoji
		emoji.WriteRune(i)

		//switch语句:都满足随机执行case，都不满足不执行case
		switch {
		//参考http://www.cnblogs.com/golove/p/3273585.html
		// IsSpace 判断 r 是否为一个空白字符
		// 在 Latin-1 字符集中，空白字符为：\t, \n, \v, \f, \r,
		// 空格, U+0085 (NEL), U+00A0 (NBSP)
		// 其它空白字符的定义有“类别 Z”和“Pattern_White_Space 属性”
		case unicode.IsSpace(i):
			//如果是空白字符则说明不满足表情字符的要求则直接返回退出该函数
			return emoji.String()
		case i == ':':
			//当此时 i == ':'的时候，一个完整的表情字符已经读取完成，返回匹配完成的unicode字符
			return emojize(emoji.String())
		}
	}
}

//对传入的字符串做分割判断(分隔符判断标识符为 :: )
func compile(x string) string {
	//如果没有传入的字符串则直接返回""
	if x == "" {
		return ""
	}
	//将传入的字符放的缓冲器中并初始化
	input := bytes.NewBufferString(x)
	//初始化一个空字符串的缓冲器
	output := bytes.NewBufferString("")

	//遍历缓冲器中的每个rune字符
	for {
		i, _, err := input.ReadRune()
		if err != nil {
			break
		}
		//switch语句，不满足case则进行default操作
		switch i {
		default:
			//默认操作:将i写入output缓冲器中
			output.WriteRune(i)
		case ':':
			//如果当前字符为 : (也就是传入的表情字符分隔符)，则将input缓冲器做为参数传给replaseEmoji函数
			output.WriteString(replaseEmoji(input))
		}
	}
	return output.String()
}

//*[]interface{} 空接口  接受任何类型
func compileValues(a *[]interface{}) {
	for i, x := range *a {
		//str为传入的字符
		if str, ok := x.(string); ok {
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
