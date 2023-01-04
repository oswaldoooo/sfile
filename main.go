package main

import (
	"fmt"
	"os"
	"sfile/cmd"
)
func main(){
	if len(os.Args)<2{
		os.Exit(0)
	}else if len(os.Args)==3{
		switch os.Args[1]{
		case "add":
			sfile_command.AddFile(os.Args[2])
		case "get":
			sfile_command.GetFIle(os.Args[2])
		case "upgrade":
			sfile_command.Upgrade(os.Args[2])
		case "upload":
			if os.Args[2]=="list"{
				sfile_command.UploadToServer()
			}else{
				sfile_command.UploadFile(os.Args[2])
			}
		default:
			fmt.Println("no this command")
		}
	}else if len(os.Args)==2{
		switch os.Args[1]{
		case "list":
			sfile_command.List()
		case "reload":
			sfile_command.Reload()
		case "init":
			Init()
		default:
			fmt.Println("no this command")
		}
	}
}

// init all programmer
func Init(){
	if !sfile_command.Exist_Dir(sfile_command.FILESAVEPATH){
		err:=os.Mkdir(sfile_command.FILESAVEPATH,os.ModePerm)
		fmt.Println(err)
	}

}