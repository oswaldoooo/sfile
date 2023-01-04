package sfile_command

import (
	"fmt"
	"io/ioutil"
	"net"
	"strings"

	remindmetools "github.com/oswaldoooo/octools/reminde"
)

// upload all filesystem to server
func UploadToServer(){
	basiccmd:="upload-->list"
	linkmap:=ReadDir()
	nodearr:=[]string{}
	for k,v:=range linkmap{
		if !Exist_File_M(v){
			continue
		}else{
			body,err:=ioutil.ReadFile(v)
			checkerror(err)
			node:=fmt.Sprintf("%v::-::%v",k,string(body))
			nodearr = append(nodearr, node)
		}
	}
	seccmd:=strings.Join(nodearr, "|-|")
	allcmd:=basiccmd+"::--::"+seccmd
	args:=[]string{"address"}
	res:=remindmetools.ReadConfPlus("Server",args,"site-conf.ini")
	con,err:=net.Dial("tcp",res["address"])
	checkerror(err)
	_,err=con.Write([]byte(allcmd))
	checkerror(err)
	con.Close()
	// fmt.Println(allcmd)
	// checkerror(err)
	// var buff []byte
	// n,err:=con.Read(buff)
	// checkerror(err)
	// fmt.Println(string(buff[:n]))
}
// upload single file to server
func UploadFile(filename string){
	if !Exist_File(filename){
		AddFile(filename)
	}
	basiccmd:=fmt.Sprintf("upload-->%v",filename)
	linkmap:=ReadDir()
	path:=linkmap[filename]
	f,err:=ioutil.ReadFile(path)
	checkerror(err)
	seccmd:=fmt.Sprintf("body::-::%v",string(f))
	allcmd:=basiccmd+"::--::"+seccmd
	args:=[]string{"address"}
	res:=remindmetools.ReadConfPlus("Server",args,"site-conf.ini")
	con,err:=net.Dial("tcp",res["address"])
	checkerror(err)
	_,err=con.Write([]byte(allcmd))
	checkerror(err)
	con.Close()
	// var buff []byte
	// n,err:=con.Read(buff)
	// checkerror(err)
	// fmt.Println(string(buff[:n]))
}