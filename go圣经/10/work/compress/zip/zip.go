package zip

import (
	"archive/zip"
	"compress/compress"
	"errors"
	"io"
	"os"
	"path/filepath"
)

const fext = ".zip"
var (
	ErrSrcNotExists     = errors.New("source file not exists")
	ErrDstExists        = errors.New("dest file is exists")
	ErrFileFormat       = errors.New("error file format")
	ErrSrcNotSupportDir = errors.New("not support dir yet")
)

// zip的压缩函数
func Compress(dst,src string) error {
	// 检查目标文件是否存在
	_ , err := os.Stat(dst)
	if os.IsExist(err) {
			return ErrDstExists
		}

	// 检查源文件是否存在
	srcinfo , err := os.Stat(src)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrSrcNotExists
		}
		return err
	}

	// 检查源文件
	if srcinfo.IsDir() {
		return ErrSrcNotSupportDir
	}

	// 仅支持*.zip
	ff := filepath.Ext(dst)
	if  ff != fext {
		return ErrFileFormat
	}

	// 创建压缩文件
	fdst , err := os.Create(dst)
	if err != nil {
		return err
	}
	defer fdst.Close()

	fdstw := zip.NewWriter(fdst)
	defer fdstw.Close()

	// 读取源文件内容
	srcFile , err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	// write zip file
	fileZip, err := fdstw.Create(srcinfo.Name())
	if err != nil { return err}
	if srcinfo.IsDir() {
		return ErrSrcNotSupportDir
	}

	_,err = io.Copy(fileZip,srcFile)
	if err != nil { return err }
	
	return nil

}

func Extract(dst, src string) error {
	return nil
}
func  init() {
	compress.RegisterFormat(fext,Compress,Extract)
}
