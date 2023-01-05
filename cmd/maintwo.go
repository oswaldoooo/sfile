package sfile_command

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
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
}

// download from server filesystem
func Download(arg []string){
	args:=[]string{"address"}
	res:=remindmetools.ReadConfPlus("Server",args,"site-conf.ini")
	con,err:=net.Dial("tcp",res["address"])
	checkerror(err)
	if len(args)==1{
		allcmd:="get-->"+arg[0]
		// fmt.Println(allcmd)
		con.Write([]byte(allcmd))
	}else if len(args)==2{
		allcmd:="get-->"+arg[0]+"-->"+arg[1]
		con.Write([]byte(allcmd))
	}
	var buff [100*MB]byte
	n,err:=con.Read(buff[:])
	checkerror(err)
	body:=string(buff[:n])
	switch body{
	case "502":
		fmt.Println(body)
	case "404":
		fmt.Println(body)
	case "403":
		fmt.Println(body)
	default:
		// dont exist local filesystem
		if !Exist_File(arg[0]){
			f,err:=os.OpenFile(arg[0],os.O_CREATE|os.O_WRONLY|os.O_TRUNC,0755)
			checkerror(err)
			_,err=f.Write(buff[:n])
			if checkerror(err){
				AddFile(arg[0])
			}
		}else{
			// exist local filesystem
			linkmap:=ReadDir()
			fp:=linkmap[arg[0]]
			f,err:=os.OpenFile(fp,os.O_WRONLY|os.O_TRUNC|os.O_CREATE,0755)
			if checkerror(err){
				f.Write(buff[:n])
			}

		}
	}
}

// get list filesystem from host
func ListAll(){
	args:=[]string{"address"}
	res:=remindmetools.ReadConfPlus("Server",args,"site-conf.ini")
	con,err:=net.Dial("tcp",res["address"])
	checkerror(err)
	allcmd:="list-->--all"
	_,err=con.Write([]byte(allcmd))
	CheckError(err)
	var buff [KB]byte
	n,err:=con.Read(buff[:])
	CheckError(err)
	fmt.Print(string(buff[:n]))
}