package create

import (
	"fmt"
	"github.com/renfy96/fly/internal/pkg/helper"
	"github.com/renfy96/fly/tpl"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

type Create struct {
	ProjectName        string
	CreateType         string
	FilePath           string
	FileName           string
	FileNameTitleLower string
	FileNameFirstChar  string
	IsFull             bool
}

func NewCreate() *Create {
	return &Create{}
}

var CreateCmd = &cobra.Command{
	Use:     "create [type] [name]",
	Short:   "Create a new application/bc",
	Example: "fly create bc user",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

	},
}
var CreateApplicationCmd = &cobra.Command{
	Use:     "application",
	Short:   "Create a new application",
	Example: "fly create application user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}
var CreateBCCmd = &cobra.Command{
	Use:     "bc",
	Short:   "Create a new bc",
	Example: "fly create bc user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}
var CreateAllCmd = &cobra.Command{
	Use:     "all",
	Short:   "Create a new application & bc",
	Example: "fly create all user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}

func runCreate(cmd *cobra.Command, args []string) {
	c := NewCreate()
	c.ProjectName = helper.GetProjectName(".")
	c.CreateType = cmd.Use
	c.FilePath, c.FileName = filepath.Split(args[0])
	c.FileName = strings.ReplaceAll(strings.ToUpper(string(c.FileName[0]))+c.FileName[1:], ".go", "")
	c.FileNameTitleLower = strings.ToLower(string(c.FileName[0])) + c.FileName[1:]
	c.FileNameFirstChar = string(c.FileNameTitleLower[0])

	switch c.CreateType {
	case "application":
		c.createApplication()
	case "bc":
		c.createBc()
	case "model":
		c.createModel()
	case "all":
		c.CreateType = "application"
		c.createApplication()
		c.CreateType = "bc"
		c.createBc()
		c.createModel()
		cmdExec := exec.Command("wire", "cmd/server/wire.go")
		if b, err := cmdExec.CombinedOutput(); err != nil {
			fmt.Println("wire err", err, string(b))
		}
		cmdExec = exec.Command("gofmt", "-w", "internal")
		if b, err := cmdExec.CombinedOutput(); err != nil {
			fmt.Println("fmt err", err, string(b))
		}
	default:
		log.Fatalf("Invalid handler type: %s", c.CreateType)
	}

}
func createFile(dirPath string, filename string) *os.File {
	filePath := dirPath + filename
	// 创建文件夹
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create dir %s: %v", dirPath, err)
	}
	stat, _ := os.Stat(filePath)
	if stat != nil {
		return nil
	}
	// 创建文件
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", filePath, err)
	}

	return file
}

func (c *Create) createApplication() {
	filePath := c.FilePath
	if filePath == "" {
		filePath = fmt.Sprintf("internal/%s/%s/", c.CreateType, strings.ToLower(c.FileName))
	}
	var (
		pPath = filePath + "params/"
		sPath = filePath + "service/"
	)

	genFile(c, pPath, "app-params", strings.ToLower(c.FileName))

	genFile(c, sPath, "app-service", strings.ToLower(c.FileName))
}
func (c *Create) createBc() {
	filePath := c.FilePath
	if filePath == "" {
		filePath = fmt.Sprintf("internal/%s/%s/", c.CreateType, strings.ToLower(c.FileName))
	}
	var (
		aPath = filePath + "application/"
		sPath = filePath + "domain/"
	)

	genFile(c, aPath, "bc-query", "query_service")
	genFile(c, aPath, "bc-command", "command_service")

	genFile(c, sPath+"agg/", "bc-agg", strings.ToLower(c.FileName))
	genFile(c, sPath+"agg/impl/", "bc-agg-impl", strings.ToLower(c.FileName))
	genFile(c, sPath+"repository/", "bc-repository", strings.ToLower(c.FileName))
	genFile(c, filePath+"infrastructure/repository/", "bc-infra-repo", strings.ToLower(c.FileName))
}
func (c *Create) createModel() {
	filePath := c.FilePath
	if filePath == "" {
		filePath = fmt.Sprintf("internal/repository/model/")
	}

	genFile(c, filePath, "model", strings.ToLower(c.FileName))
	// 在迁移中加入数据结构并生成对应orm
	_ = searchAndWriteInfConst(c.FilePath+"cmd/migration/gen.go", "model."+c.FileName+"{},", "ApplyBasic")
	cmdExec := exec.Command("go", "run", "cmd/migration/gen.go")
	if b, err := cmdExec.CombinedOutput(); err != nil {
		fmt.Println("gen err :", err, string(b))
	}
}
func genFile(c *Create, dirPath, tmp, fileName string) {
	f := createFile(dirPath, fileName+".go")
	if f == nil {
		log.Printf("warn: file %s%s %s", dirPath, strings.ToLower(fileName)+".go", "already exists.")
		return
	}
	defer f.Close()
	t, err := template.ParseFS(tpl.CreateTemplateFS, fmt.Sprintf("create/%s.tpl", tmp))
	if err != nil {
		log.Fatalf("create %s error: %s", fileName, err.Error())
	}
	err = t.Execute(f, c)
	if err != nil {
		log.Fatalf("create %s error: %s", c.CreateType, err.Error())
	}
	log.Printf("Created new %s: %s", c.CreateType, dirPath+strings.ToLower(c.FileName)+".go")
}
func searchAndWriteInfConst(fileName, content, infoName string) error {

	lineBytes, err := os.ReadFile(fileName)
	var lines []string
	if err != nil {
		fmt.Println(err)
	} else {
		contents := string(lineBytes)
		lines = strings.Split(contents, "\n")
	}
	var newLines []string

	for _, line := range lines {
		newLines = append(newLines, line)
		isIn, err := regexp.MatchString(infoName, line)
		if err != nil {
			log.Printf("匹配常量类里的interface报错 :%v", err)
			continue
		}
		if isIn {
			newLines = append(newLines, "\t\t"+content)
		}
	}

	file, err := os.OpenFile(fileName, os.O_WRONLY, 0666)
	defer file.Close()
	_, err = file.WriteString(strings.Join(newLines, "\n"))
	if err != nil {
		return err
	}
	return nil
}
