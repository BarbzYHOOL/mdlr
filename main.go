package main

import (
	"github.com/exlinc/mdlr/mdlrf"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	initCmd        = kingpin.Command("init", "Create a template mdlr.yml file in the directory")
	listCmd        = kingpin.Command("list", "List the current modules from the mdlr.yml file in the directory")
	addCmd         = kingpin.Command("add", "Add a module to the mdlr.yml file in the current directory. Create the mdlr.yml file if it doesn't yet exist")
	addName        = addCmd.Flag("name", "The internal module name used in the mdlr.yml file. Must be unique to this mdlr.yml file").Short('n').Required().String()
	addType        = addCmd.Flag("type", "The module type -- defaults to git").Short('t').Default("git").String()
	addPath        = addCmd.Flag("path", "Path for the module within the project").Required().Short('p').String()
	addUrl         = addCmd.Flag("url", "URL for the module").Required().Short('u').String()
	addBranch      = addCmd.Flag("branch", "Name of the branch of the repo to use").Default("master").Short('b').String()
	addCommit      = addCmd.Flag("commit", "Long or short hash of the commit of the module to use").Short('c').String()
	removeCmd      = kingpin.Command("remove", "Remove the module from the mdlr.yml file. Use the --files flag to remove the files from the filesystem path as well")
	removeName     = removeCmd.Flag("name", "The name of the module to remove from the mdlr.yml file").Short('n').String()
	removeFiles    = removeCmd.Flag("files", "Remove the files in the module's path").Short('f').Bool()
	importCmd      = kingpin.Command("import", "Import the mdlr.yml modules. Use the --specific flag to specify a single module by name")
	importForce    = importCmd.Flag("force", "Force the import -- this will completely reset the path of the module and pull from the internet").Short('f').Bool()
	importSpecific = importCmd.Flag("specific", "Specify the name of the module to import").Short('s').String()
	updateCmd      = kingpin.Command("update", "Update the mdlr.yml modules. Use the --specific flag to specify a single module by name")
	updateForce    = updateCmd.Flag("force", "Wipe database and reset tables first").Short('f').Bool()
	updateSpecific = updateCmd.Flag("specific", "Specify the name of the module to update").Short('s').String()
	statusCmd      = kingpin.Command("status", "Get the status of the mdlr.yml modules. Use the --specific flag to specify a single module by name")
	statusSpecific = statusCmd.Flag("specific", "Specify the name of the module to update").Short('s').String()
)

func main() {
	kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version("0.1.1").Author("EXL Inc.")
	kingpin.CommandLine.Help = "mdlr"
	c, err := mdlrf.NewMdlrCtxForCmd()
	if err != nil {
		log.WithError(err).Fatal("Unable to get context for the command")
	}

	switch kingpin.Parse() {
	case "init":
		err = c.Init()
		if err != nil {
			log.WithError(err).Fatal("Unable to initialize the mdlr.yml file")
		}
		log.Info("Successfully initialized the mdlr.yml file!")
		log.Exit(0)
	case "list":
		// todo
	case "add":
		// todo
	case "remove":
		// todo
	case "import":
		// todo
	case "update":
		// todo
	case "status":
		// todo
	default:
		log.Error("Unknown command")
	}
}
