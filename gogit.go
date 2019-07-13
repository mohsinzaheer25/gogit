package main

import (
	"bytes"
	"fmt"
	"github.com/manifoldco/promptui"
	"log"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strings"
)
var (
	files string
	comment string
	branch  string
	gitInput string
	branchName string
	uniqueCommit string
	branchout bytes.Buffer
	modout bytes.Buffer
	uncommitout bytes.Buffer
	untrackout bytes.Buffer
	undoOut bytes.Buffer
)

const version  = "v1.0.0"

func init(){
	addCmd.PersistentFlags().StringVarP(&files,"files","f","", "Single File or Multiple Files With Spaces Within Quotes")
	addCmd.PersistentFlags().StringVarP(&comment,"comment", "c", "", "Comment")
	addCmd.PersistentFlags().StringVarP(&branch,"branch", "b", "", "Git Branch")
}

var addCmd = &cobra.Command{
	Use:   "gogit",
	Short: "gogit is simple commandline tool for git commands",
	Long: `This tool do all git commands in simple way
	Usage: gogit add -f "[Filenames]" -c "[Comment]" -b "[Branch Name]"`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			os.Exit(1)
		}else if len(files) < 1 || comment == "" || branch == "" {
			cmd.Help()
			os.Exit(1)
		}
		addFunc()
	},
}


func main(){
	if len(os.Args) < 2 || os.Args[1] == "help" || os.Args[1] == "--help" || os.Args[1] == "-h" {
		fmt.Println("GOGIT is command line tool to use git command in simple way.")
		fmt.Printf("\nUsage: gogit [COMMAND]\n")
		fmt.Printf("\nCommands:\n")
		fmt.Println("ls          List All The Change Files")
		fmt.Println("add         Add Files To The Repository")
		fmt.Println("get         Get a Repository or Branch Updated Files")
		fmt.Println("undo        Reset your commit to particular commit or Reset last commit")
		fmt.Println("newbranch   Creates A New Branch")
		fmt.Println("version     Display GOGIT Version")
		os.Exit(0)
	}
	inputCMD(os.Args[1])
}

func inputCMD(cmd string){
	switch cmd {
	case "version":
		versionFunc()
	case "get":
		getFunc()
	case "newbranch":
		newbranchFunc()
	case "ls":
		listFunc()
	case "add":
		if err := addCmd.Execute(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "undo":
		undoFunc()

	}
}

func cmdFunc(getRepocmd *exec.Cmd){
		getRepocmd.Stdout = os.Stdout
		getRepocmd.Stderr = os.Stderr
		getRepocmd.Run()
}

func versionFunc(){
		fmt.Println("GOGIT",version)
}

func getFunc(){
	if len(os.Args) != 3 || os.Args[2] == "help" || os.Args[2] == "--help" || os.Args[2] == "-h"{
		fmt.Println("GOGIT is command line tool to use git command in simple way.")
		fmt.Printf("\nUSAGE: gogit get [URL] or [BranchName]\n\n")
		os.Exit(0)
	}
	gitInput = os.Args[2]
	if strings.HasPrefix(gitInput, "https://github.com/") == true {
		getRepocmd:= exec.Command("git","clone",gitInput)
		cmdFunc(getRepocmd)
	}else {
		getRepocmd:= exec.Command("git","pull", "origin",gitInput)
		cmdFunc(getRepocmd)
	}

}

func newbranchFunc(){
	if len(os.Args) != 3 || os.Args[2] == "help" || os.Args[2] == "--help" || os.Args[2] == "-h" {
		fmt.Println("GOGIT is command line tool to use git command in simple way.")
		fmt.Printf("\nUSAGE: gogit newbranch [Branch Name]\n\n")
		os.Exit(0)
	}
	statusCMD := exec.Command("git","status")
	statusCMD.Stderr = os.Stderr
	error := statusCMD.Run()
	if error != nil {
		os.Exit(1)
	}
	branchName = os.Args[2]
	newBranchcmd := exec.Command("git","checkout","-b",branchName)
	cmdFunc(newBranchcmd)
}

func listFunc(){
	statusCMD := exec.Command("git","status")
	statusCMD.Stderr = os.Stderr
	error := statusCMD.Run()
	if error != nil {
		os.Exit(1)
	}
	branCMD := exec.Command("git","rev-parse", "--abbrev-ref", "HEAD")
	branCMD.Stdout = &branchout
	branCMD.Run()
	fmt.Println("On branch",branchout.String())



	modCMD := exec.Command("git","ls-files", "-m")
	modCMD.Stdout = &modout
	modCMD.Run()
	if modout.String() != ""{
		fmt.Println("Files Modified But Not Added & Commited")
		fmt.Printf("  (use 'gogit add ..' to commit and push files to repo)\n\n")
		fmt.Println(modout.String())
		fmt.Println()
	}

	// This will be remove in final stage

	uncommitCmd := exec.Command("git","diff", "--cached", "--name-only", "--diff-filter=A")
	uncommitCmd.Stdout = &uncommitout
	uncommitCmd.Run()
	if uncommitout.String() != ""{
		fmt.Println("Files Added But Need To Be Committed")
		fmt.Printf("  (use 'gogit add ..' to commit and push files to repo )\n\n")
		fmt.Println(uncommitout.String())
		fmt.Println()
	}


	untrackCMD := exec.Command("git","ls-files", "-o")
	untrackCMD.Stdout = &untrackout
	untrackCMD.Run()
	if untrackout.String() != ""{
		fmt.Println("Untracked Files")
		fmt.Printf("  (use 'gogit add ..' to commit and push files to repo)\n\n")
		fmt.Println(untrackout.String())
	}
}

func addFunc(){
	statusCMD := exec.Command("git","status")
	statusCMD.Stderr = os.Stderr
	error := statusCMD.Run()
	if error != nil {
		os.Exit(1)
	}
	if files == "." {
		addsCMD := exec.Command("git","add",".")
		addsCMD.Run()
	} else {
		files := strings.Split(files," ")
		for _,f := range files{
			addsCMD := exec.Command("git","add",f)
			addsCMD.Run()
		}
	}
	commitCMD := exec.Command("git","commit","-m",comment)
	commitCMD.Run()
	pushCMD := exec.Command("git","push","origin",branch)
	cmdFunc(pushCMD)
}

func undoFunc(){
	if len(os.Args) == 3 && ( os.Args[2] == "help" || os.Args[2] == "--help" || os.Args[2] == "-h") {
		fmt.Println("GOGIT is command line tool to use git command in simple way.")
		fmt.Printf("\nUSAGE: gogit [COMMAND] \n")
		fmt.Printf("\nCommands:\n")
		fmt.Println("undo,        			  Reset Commit To Last Commit")
		fmt.Printf("undo -h [Unqiue Hash Commit],     Reset Commit To Unqiue Hash Commit\n\n")
		os.Exit(0)
	}else if len(os.Args) == 3 {
		fmt.Println("Did you mean help?")
		fmt.Printf("\nTry: gogit undo help\n\n")
		os.Exit(0)
	} else if len(os.Args) == 2 && os.Args[1] == "undo"{
		fmt.Println("Are you sure you want to undo your commit to last commit?")
		prompt := promptui.Select{
			Label: "Select[Yes/No]",
			Items: []string{"Yes", "No"},
		}
		_, result, err := prompt.Run()
		if err != nil {
			log.Fatalf("%v\n", err)
			os.Exit(1)
		}
		if result == "Yes"{
			undoCMD := exec.Command("git","reset","HEAD~1")
			undoCMD.Stderr = &undoOut
			error := undoCMD.Run()
			if error != nil {
				fmt.Println(undoOut.String())
			}else{
				fmt.Println("Undo To Last Commit")
			}
		} else {
			fmt.Println("Undo Aborted!!")
			os.Exit(0)
		}
	}else if len(os.Args) == 4 && os.Args[2] == "-h"{
		uniqueCommit = os.Args[3]
		fmt.Println("Are you sure you want to undo to ",uniqueCommit,"commit?")
		prompt := promptui.Select{
			Label: "Select[Yes/No]",
			Items: []string{"Yes", "No"},
		}
		_, result, err := prompt.Run()
		if err != nil {
			log.Fatalf("%v\n", err)
			os.Exit(1)
		}
		if result == "Yes"{
			undoCMD := exec.Command("git","reset",uniqueCommit)
			undoCMD.Stderr = &undoOut
			error := undoCMD.Run()
			if error != nil {
				fmt.Println(undoOut.String())
			}else{
				fmt.Println("Undo Commit to ",uniqueCommit,"commit")
			}
		} else {
			fmt.Println("Undo Aborted!!")
			os.Exit(0)
		}
	}
}