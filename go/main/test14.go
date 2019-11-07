package main

import (
	"archive/zip"
	"io"
	"os"
	"time"
)
func zipCompress(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = zipCompress(f, prefix, zw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
func ZipCompress(files []*os.File, dest string) error {
	d, _ := os.Create(dest)
	defer d.Close()
	w := zip.NewWriter(d)
	defer w.Close()
	for _, file := range files {
		err := zipCompress(file, "", w)
		if err != nil {
			return err
		}
	}
	return nil
}
func main(){
	date := time.Now()
	day,_ := time.ParseDuration("-24h")
	var f *os.File
	var files = []*os.File{}
	for {
		str := date.String()[:10]
		f, _ = os.Open("./log/smartrtb.log."+str)
		defer f.Close()
		if f == nil {
			date = date.Add(day)
		}else{
			files = append(files, f)
			ZipCompress(files, "./log/a.zip")
			break
		}
	}


}
