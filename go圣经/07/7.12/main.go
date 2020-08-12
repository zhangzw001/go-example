package main

import (
	"io"
)

//因为Write方法需要传入一个byte切片而我们希望写入的值是一个字符串，所以我们需要使用[]byte(...)进行转换。这个转换分配内存并且做一个拷贝，但是这个拷贝在转换后几乎立马就被丢弃掉。让我们假装这是一个web服务器的核心部分并且我们的性能分析表示这个内存分配使服务器的速度变慢。这里我们可以避免掉内存分配么？
func writeHeader(w io.Writer, contentType string) error {
	if _, err := w.Write([]byte("Content-Type: ")); err != nil {
		return err
	}
	if _, err := w.Write([]byte(contentType)); err != nil {
		return err
	}
	return nil
}
//我们不能对任意io.Writer类型的变量w，假设它也拥有WriteString方法。但是我们可以定义一个只有这个方法的新接口并且使用类型断言来检测是否w的动态类型满足这个新接口。
//为了避免重复定义，我们将这个检查移入到一个实用工具函数writeString中，但是它太有用了以致标准库将它作为io.WriteString函数提供。这是向一个io.Writer接口写入字符串的推荐方法。
func writeString(w io.Writer, s string ) (n int , err error ) {
	type stringWriter interface {
		WriteString(string) (n int , err error )
	}
	if sw , ok := w.(stringWriter); ok {
		return sw.WriteString(s)	//avoid a copy
	}
	return w.Write([]byte(s))	// allocate temporary copy
}
//上面的writeString函数使用一个类型断言来知道一个普遍接口类型的值是否满足一个更加具体的接口类型；并且如果满足，它会使用这个更具体接口的行为。这个技术可以被很好的使用不论这个被询问的接口是一个标准的如io.ReadWriter或者用户定义的如stringWriter。



//这也是fmt.Fprintf函数怎么从其它所有值中区分满足error或者fmt.Stringer接口的值。在fmt.Fprintf内部，有一个将单个操作对象转换成一个字符串的步骤，像下面这样：
//func formatOneValue(x interface{}) string {
//	if err, ok := x.(error); ok {
//		return err.Error()
//	}
//	if str, ok := x.(Stringer); ok {
//		return str.String()
//	}
//	// ...all other types...
//}

func main() {
}