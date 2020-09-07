package compress

import (
	"errors"
	"fmt"
	"path/filepath"
)

var (
	ErrNotSupport = errors.New("Not suitable driver.")
)

var  (
	// drivers slice
	drivers []driver
)

type driver struct {
	format 	 string
	compress compressFunc
	extract  extractFunc
}

type compressFunc func(dst , src string ) error
type extractFunc func(dst, src string ) error

// 提供zip,tar注册
func RegisterFormat(format string, compress compressFunc, extract extractFunc) {
	fmt.Println("RegisterFormat: ", format)
	drivers = append(drivers, driver{format,compress,extract})
}

// 获取压缩文件后缀, 不支持非*.ext的ext格式文件(例如不支持文件名为 azip 的zip文件,仅支持*.zip)
func getFileFormat(filename string) string {
	return filepath.Ext(filename)
}

// 查看是否支持的格式
func matchDriver(format string ) (driver , error) {
	fmt.Println("all drivers: ", drivers)
	for _, d := range drivers {
		if d.format == format {
			return d ,nil
		}
	}
	return driver{}, ErrNotSupport
}


// 导出的压缩方法

func Compress(dst,src string ) error {
	ff := getFileFormat(dst)
	dri,err := matchDriver(ff)
	if err != nil {
		return err
	}
	return dri.compress(dst,src)
}