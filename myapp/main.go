/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/spf13/myapp/cmd"
)

func main() {
	cmd.Execute()

	//var (
	//	topicFlag     string
	//	deadQueueFlag string
	//)
	//
	//var rootCmd = &cobra.Command{Use: "test"}
	//rootCmd.PersistentFlags().StringVarP(&topicFlag, "topic", "t", "default_topic", "Specify the topic")
	//rootCmd.PersistentFlags().StringVarP(&deadQueueFlag, "deadqueue", "d", "", "Specify the dead queue (optional)")
	//rootCmd.Run = func(cmd *cobra.Command, args []string) {
	//	// Your application logic here
	//	if deadQueueFlag == "" {
	//		fmt.Printf("Topic: %s\n", topicFlag)
	//	} else {
	//		fmt.Printf("Topic: %s, DeadQueue: %s\n", topicFlag, deadQueueFlag)
	//	}
	//}
	//
	//// Execute Cobra
	//if err := rootCmd.Execute(); err != nil {
	//	fmt.Println(err)
	//}
}
