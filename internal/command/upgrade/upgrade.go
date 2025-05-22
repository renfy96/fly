package upgrade

import (
	"fmt"
	"github.com/renfy96/fly/config"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

var UpgradeCmd = &cobra.Command{
	Use:     "upgrade",
	Short:   "Upgrade the fly command.",
	Long:    "Upgrade the fly command.",
	Example: "fly upgrade",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("go install %s\n", config.FLYCmd)
		cmd := exec.Command("go", "install", config.FLYCmd)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatalf("go install %s error\n", err)
		}
		fmt.Printf("\n🎉 fly upgrade successfully!\n\n")
	},
}
