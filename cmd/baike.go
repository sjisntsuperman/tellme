package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

var (
  platform string
)

var openCmds = map[string]string{
  "windows": "cmd /c start",
  "darwin": "open",
  "linux": "xdg-open",
}

var baikeCmd = &cobra.Command{
  Use:  "baike",
  Aliases: []string{"bk", "wk", "wiki"},
  Short: "find things in baike site",
  Args: cobra.ExactArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("ready")
    err := findInBaike(args[0], platform)
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }
  },
}

func init() {
  rootCmd.AddCommand(baikeCmd)
  baikeCmd.Flags().StringVarP(&platform, "platform", "p", "baidu", "platform to fin")
}

func findInBaike(keyword, platform string) error {
  var link string
  if platform == "baidu" || platform == "bd" {
    link = fmt.Sprintf("https://baike.baidu.com/item/%s", keyword)
  }
  if platform == "hudong" || platform == "baike" || platform == "hd" {
    link = fmt.Sprintf("https://www.baike.com/wiki/%s", keyword)
  }
  if platform == "wikipedia" || platform == "wiki" || platform == "wp" {
    link = fmt.Sprintf("https://zh.wikipedia.org/wiki/%s", keyword)
  }
  if link == "" {
    return fmt.Errorf("invalid platform")
  }

  goos := runtime.GOOS
  openCmd := "open"
  openCmd, ok := openCmds[goos]
  if !ok {
    return fmt.Errorf("can not open link in %s", goos)
  }
  if err := exec.Command(openCmd, link).Start(); err != nil {
    return err
  }
  return nil
}
