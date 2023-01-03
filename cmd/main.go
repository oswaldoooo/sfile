package sfile_command

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)
func checkerror(er error) bool{
	if er!=nil{
		fmt.Println(er)
		return false
	}
	return true
}

// add file to filesystem
func AddFile(filename string) {
	_,err:=os.Open(filename)
	if !checkerror(err){
		os.Exit(1)
	}
	if !Exist_Dir(FILESAVEPATH){
		os.Mkdir(FILESAVEPATH,os.ModePerm)
	}
	dirpath,_:=os.Getwd()
	filepath:=fmt.Sprintf("%v\\%v",dirpath,filename)
	linkmap:=ReadDir()
	// fmt.Println(linkmap)
	// if path not exist in filesystem,put it into filesystem
	if _,ok:=linkmap[filename];!ok{
		linkmap[filename]=filepath
		WriteToDir(&linkmap)
	}else{
		// if file exist in filesystem,upload to filesystem
		f,_:=ioutil.ReadFile(filename)
		wf,err:=os.OpenFile(linkmap[filename],os.O_TRUNC|os.O_CREATE,os.ModePerm)
		checkerror(err)
		_,err=wf.Write(f)
		checkerror(err)
		fmt.Println("work done")
	}
}

// get file from filesystem
func GetFIle(filename string){
	linkmap:=ReadDir()
	if _,ok:=linkmap[filename];!ok{
		fmt.Println("file dont exist")
		os.Exit(1)
	}
	filepath:=linkmap[filename]
	fe,err:=ioutil.ReadFile(filepath)
	checkerror(err)
	f,err:=os.OpenFile(filename,os.O_CREATE|os.O_TRUNC,os.ModePerm)
	if !checkerror(err){
		os.Exit(1)
	}
	_,err=f.Write(fe)
	checkerror(err)
}

// Check the file exist
func CheckPath(path string) bool{
	f,err:=os.OpenFile(FINALFILE,os.O_RDONLY,os.ModePerm)
	if !checkerror(err){
		os.Exit(1)
	}
	var buff []byte
	n,err:=f.Read(buff)
	checkerror(err)
	ms:=string(buff[:n])
	normalarr:=strings.Split(ms, "\n")
	if len(normalarr)<1{
		return false
	}
	for i := 0; i < len(normalarr); i++ {
		newarr:=strings.Split(normalarr[i], "::")
		if newarr[0]==path{
			return true
		}
	}
	return false
}

// upgrade file origin path in filesystem
func Upgrade(filename string){
	linkmap:=ReadDir()
	if _,ok:=linkmap[filename];!ok{
		fmt.Println("not this file in filesystem")
		os.Exit(0)
	}
	paths,err:=os.Getwd()
	checkerror(err)
	filepath:=fmt.Sprintf("%v\\%v",paths,filename)
	linkmap[filename]=filepath
	if ok:=WriteToDir(&linkmap);ok{
		fmt.Println("upgrade success")
	}else{
		fmt.Println("upgrade failed")
	}
}

// get List of filesystem
func List(){
	linkmap:=ReadDir()
	for k,v:=range linkmap{
		fmt.Printf("%-20s %v\n",k,v)
	}
}

// clear the file that dont exist
func Reload(){
	linkmap:=ReadDir()
	for k,v:=range linkmap{
		_,err:=os.Stat(v)
		if err!=nil{
			delete(linkmap,k)
		}
	}
	ok:=WriteToDir(&linkmap)
	if !ok{
		fmt.Println("reload failed")
	}
}