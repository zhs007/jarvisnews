package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
	corebasedef "github.com/zhs007/jarviscore/basedef"
	"github.com/zhs007/jarvisnews/basedef"
)

func addStart(rootCmd *cobra.Command) {
	var daemon bool

	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start jarvisnews",
		Run: func(cmd *cobra.Command, args []string) {
			if daemon {
				command := exec.Command("./jarvisnews", "start")
				err := command.Start()
				if err != nil {
					fmt.Printf("start jarvisnews error. %v \n", err)

					os.Exit(-1)

					return
				}

				// fmt.Printf("jarvissh start, [PID] %d running...\n", command.Process.Pid)
				// ioutil.WriteFile("jarvissh.pid", []byte(fmt.Sprintf("%d", command.Process.Pid)), 0666)

				daemon = false

				os.Exit(0)

				return
			}

			fmt.Printf("jarvisnews start.\n")

			fmt.Printf("jarvisnews start, [PID] %d running...\n", os.Getpid())
			ioutil.WriteFile("jarvisnews.pid", []byte(fmt.Sprintf("%d", os.Getpid())), 0666)

			startServ()
		},
	}

	startCmd.Flags().BoolVarP(&daemon, "deamon", "d", false, "is daemon?")

	rootCmd.AddCommand(startCmd)
}

func addStop(rootCmd *cobra.Command) {
	var stopCmd = &cobra.Command{
		Use:   "stop",
		Short: "Stop jarvisnews",
		Run: func(cmd *cobra.Command, args []string) {
			strb, err := ioutil.ReadFile("jarvisnews.pid")
			if err != nil {
				fmt.Printf("read jarvisnews.pid error %v\n", err)

				os.Exit(-1)

				return
			}

			command := exec.Command("kill", string(strb))
			command.Start()

			time.Sleep(time.Duration(30) * time.Second)

			fmt.Printf("jarvisnews stop.\n")
		},
	}

	rootCmd.AddCommand(stopCmd)
}

func addRestart(rootCmd *cobra.Command) {
	var daemon bool

	restartCmd := &cobra.Command{
		Use:   "restart",
		Short: "Restart jarvisnews",
		Run: func(cmd *cobra.Command, args []string) {
			if daemon {
				strb, err := ioutil.ReadFile("jarvisnews.pid")
				if err == nil {
					fmt.Printf("stop jarvisnews %v ... \n", string(strb))

					command := exec.Command("kill", string(strb))
					command.Start()

					time.Sleep(time.Duration(30) * time.Second)
				}

				command := exec.Command("./jarvisnews", "start")
				err = command.Start()
				if err != nil {
					fmt.Printf("start jarvisnews error. %v \n", err)

					os.Exit(-1)

					return
				}

				// fmt.Printf("jarvissh start, [PID] %d running...\n", command.Process.Pid)
				// ioutil.WriteFile("jarvissh.pid", []byte(fmt.Sprintf("%d", command.Process.Pid)), 0666)

				daemon = false

				os.Exit(0)

				return
			}

			fmt.Printf("jarvisnews start.\n")

			fmt.Printf("jarvisnews start, [PID] %d running...\n", os.Getpid())
			ioutil.WriteFile("jarvisnews.pid", []byte(fmt.Sprintf("%d", os.Getpid())), 0666)

			startServ()
		},
	}

	restartCmd.Flags().BoolVarP(&daemon, "deamon", "d", false, "is daemon?")

	rootCmd.AddCommand(restartCmd)
}

func addVersion(rootCmd *cobra.Command) {
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "get jarvisnews version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("jarvisnews version is %v \n", jarvisnewsbasedef.VERSION)
			fmt.Printf("jarvis core version is %v \n", corebasedef.VERSION)
		},
	}

	rootCmd.AddCommand(versionCmd)
}

func startCmd() error {
	rootCmd := &cobra.Command{
		Use: "jarvisnews",
	}

	//--------------------------------------------------------------------------------------------------------------------
	// start

	addStart(rootCmd)

	//--------------------------------------------------------------------------------------------------------------------
	// stop

	addStop(rootCmd)

	//--------------------------------------------------------------------------------------------------------------------
	// restart

	addRestart(rootCmd)

	//--------------------------------------------------------------------------------------------------------------------
	// version

	addVersion(rootCmd)

	return rootCmd.Execute()
}
