package sfile_command

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)
const(
	FILESAVEPATH="E:\\WorkStudio\\Golang\\SFile\\conf"
	FINALFILE=FILESAVEPATH+"\\filemap"
)
func Exist_Dir(path string) bool {
	_, err := os.Stat(path)
	return checkerror(err)
}

// read from database to map
func ReadDir() map[string]string{
	f,err:=ioutil.ReadFile(FINALFILE)
	if !checkerror(err){
		os.Exit(1)
	}
	// var buff []byte
	ms:=string(f)
	normalarr:=strings.Split(ms, "\n")
	filelinkmap:=make(map[string]string)
	for i := 0; i < len(normalarr); i++ {
		newarr:=strings.Split(normalarr[i], "::")
		if len(newarr)!=2{
			break
		}
		filelinkmap[newarr[0]]=newarr[1]
	}
	return filelinkmap
}

// write from map to database
func WriteToDir(v *map[string]string) bool{
	emptykey:=""
	// fmt.Println(v)
	for k,ve:=range(*v){
		basicletter:=fmt.Sprintf("%v::%v\n",k,ve)
		emptykey+=basicletter
		// fmt.Println(emptykey)
	}
	// fmt.Println(emptykey)
	f,err:=os.OpenFile(FINALFILE,os.O_CREATE|os.O_WRONLY|os.O_TRUNC,os.ModePerm)
	checkerror(err)
	_,err=f.Write([]byte(emptykey))
	checkerror(err)
	return true
}