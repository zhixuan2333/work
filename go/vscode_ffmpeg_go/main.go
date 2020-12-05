package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"regexp"
	"runtime"
)

//vscodeV get vscode version
func vscodeV() []byte {
	// Print Go Version
	out, err := exec.Command("code", "-v").Output()
	if err != nil {
		log.Printf("Get vscode version failed: %s\n", err.Error())
	}

	return out
}

// atLine get string by int
func atLine(f []byte, n int) (s string) {
	r := bytes.NewReader(f)
	bufReader := bufio.NewReader(r)
	for i := 1; ; i++ {
		line, _, err := bufReader.ReadLine()
		if err != nil {
			break
		}
		if i == n {
			s = string(line)
			break
		}
	}
	return s
}

// electron get version
func electron(version string) string {
	// get .yarnrc file
	r, err := http.Get("https://raw.githubusercontent.com/Microsoft/vscode/" + version + "/.yarnrc")
	if err != nil {
		log.Printf("get electron version failed: %s", err.Error())
	}
	defer r.Body.Close()

	// get .yarnrc version
	b, err := ioutil.ReadAll(r.Body)
	s := string(b)

	// re match version
	rule, err := regexp.Compile(`".*?"`)
	if err != nil {
		log.Printf("re rule is failed: %s\n", err.Error())
	}
	results := rule.FindAllString(s, -1)
	i := results[1]
	end := len(i) - 1
	re := i[1:end]

	return re
}

// get system OS info
func systemversion() string {
	OS := runtime.GOOS
	if OS == "windows" {
		return "win32"
	}
	if OS == "darwin" {
		return "darwin"
	}
	return "linux"

}

var commands = map[string]string{
	"windows": "start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

// Open open url
func Open(url string) error {
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}

	cmd := exec.Command(run, url)
	return cmd.Start()
}

func main() {
	vscode := vscodeV()
	Vversion := atLine(vscode, 1)
	arch := atLine(vscode, 3)
	yarnrc := electron(Vversion)
	OS := systemversion()
	fmt.Printf("vscode: %s\n", Vversion)
	fmt.Printf("arch: %s\n", arch)
	fmt.Printf("version: %s\n", yarnrc)
	fmt.Printf("OS: %s\n", OS)
	url := "https://github.com/electron/electron/releases/download/v" + yarnrc + "/electron-v" + yarnrc + "-" + OS + "-" + arch + ".zip"

	cmd := exec.Command("cmd /C start", url)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
	err = Open(url)
	if err != nil {
		log.Fatal("Open url failed: ", err)
	}

}
