/*
Copyright © 2023 Yash Dixit
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
	"yashandstuff.com/imagician/utils"
)

// huntCmd represents the hunt command
var huntCmd = &cobra.Command{
	Use:   "hunt",
	Short: "hunts all the images in the folder specified",
	Long: `By hunt - it means that it will find all the images inside a folder specified - all levels deep. 
	For example:

	imagician hunt ./all_files/ ~/Desktop/cat-photos/

	-- 
	This command iterates through all files and folders in the all_files directory and creates
	a folder (if not exists already) named cat-photos and dumps all the images from all_files directory
	to cat-photos. 
	`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 2 {
			fmt.Println("just 2 command line params needed; the src folder & the dest folder - you mentioned " + strconv.Itoa(len(args)))
		}

		srcFolder := args[0]
		destFolder := args[1]

		log.Println("Starting the copying process")

		err := utils.CopyAllImagesFromDirectory(srcFolder, destFolder)

		if err != nil {
			fmt.Println("ahh, something went wrong - couldn't copy them files")
			fmt.Println("error: " + err.Error())
		}

		log.Println("Finished, all files copied from " + srcFolder + " - to " + destFolder)
	},
}

func init() {
	rootCmd.AddCommand(huntCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// huntCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// huntCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
