package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// from https://zetcode.com/golang/exec-command/
	{
		cmd := exec.Command("tr", "a-z", "A-Z")
		cmd.Stdin = strings.NewReader("and old falcon")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("translated phrase: %q\n", out.String())
	}
	{
		prg := "echo"
		arg1 := "there"
		arg2 := "are three"
		arg3 := "falcons"
		cmd := exec.Command(prg, arg1, arg2, arg3)
		stdout, err := cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Print(string(stdout))
	}
	{
		out, err := exec.Command("ls", "-l").Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(out))
	}
	{
		cmd := exec.Command("cat")
		stdin, err := cmd.StdinPipe()
		if err != nil {
			log.Fatal(err)
		}
		go func() {
			defer stdin.Close()
			io.WriteString(stdin, "an old falcon")
		}()
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", out)
	}
	{
		cmd := exec.Command("echo", "an old falcon")
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}
		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}
		data, err := ioutil.ReadAll(stdout)
		if err != nil {
			log.Fatal(err)
		}
		if err := cmd.Wait(); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", strings.ToUpper((string(data))))
	}
	// from https://pkg.go.dev/os/exec#example-Cmd.CombinedOutput
	{
		cmd := exec.Command("sh", "-c", "echo stdout; echo 1>&2 stderr")
		stdoutStderr, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", stdoutStderr)
	}
	{
		cmd := exec.Command("pwd")
		// Set Dir before calling cmd.Environ so that it will include an
		// updated PWD variable (on platforms where that is used).
		cmd.Dir = ".."
		cmd.Env = append(os.Environ(), "POSIXLY_CORRECT=1")
		out, err := cmd.Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", out)
	}
	{
		out, err := exec.Command("date").Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("The date is %s\n", out)
	}
	{
		cmd := exec.Command("sleep", "1")
		log.Printf("Running command and waiting for it to finish...")
		err := cmd.Run()
		log.Printf("Command finished with error: %v", err)
	}
	{
		cmd := exec.Command("sleep", "5")
		err := cmd.Start()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Waiting for command to finish...")
		err = cmd.Wait()
		log.Printf("Command finished with error: %v", err)
	}
	{
		cmd := exec.Command("sh", "-c", "echo stdout; echo 1>&2 stderr")
		stderr, err := cmd.StderrPipe()
		if err != nil {
			log.Fatal(err)
		}
		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}
		slurp, _ := io.ReadAll(stderr)
		fmt.Printf("%s\n", slurp)
		if err := cmd.Wait(); err != nil {
			log.Fatal(err)
		}
	}
}
