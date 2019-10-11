package ipfs

import (
	"fmt"
	"io/ioutil"
	"strings"
	"bytes"
	"log"
	"os"
	shell "github.com/ipfs/go-ipfs-api"
)

// This is the default URL for the local IPFS daemon
// Can set other path using the SetPath function
var ipfsPath = "localhost:50013333"

func getShell(url string) *shell.Shell {
	sh := shell.NewShell(url)
	return sh
}

// Set the IPFS path of the local IPFS daemon
func SetPath(path string) {

	ipfsPath = path
}

// AddFile adds a file to IPFS
// Returns the hash of the file, error
func AddFile(filePath string) (string, error) {

	content, err := ioutil.ReadFile(filePath)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Printf("File contents: %s", content)

	// Get Reader from file
        reader := bytes.NewReader(content)

	// Get shell
	sh := getShell(ipfsPath)
	// Add file to IPFS
        hash, err := sh.Add(reader)

        if err != nil {
                fmt.Fprintf(os.Stderr, "error: %s", err)
                os.Exit(1)
        }
        //fmt.Println("File added", hash)

	return hash, err
}

// AddString adds a string to IPFS
// Returns the hash of the string, error
func AddString(str string) (string, error) {

	// Get Reader from string
        reader := strings.NewReader(str)

	// Get shell
	sh := getShell(ipfsPath)

	// Add string to IPFS
        hash, err := sh.Add(reader)

        if err != nil {
                fmt.Fprintf(os.Stderr, "error: %s", err)
                os.Exit(1)
        }
        //fmt.Println("String added", hash)

	return hash, err
}

// GetFile gets a file from IPFS
// Returns the path of the file, error
func GetFile(hash string, outdir string, fileName string) (string, error) {

	// Get shell
	sh := getShell(ipfsPath)

	file := outdir + fileName

	return file, sh.Get(hash, file)
}


// GetString gets a string from IPFS
// Returns the string, error
func GetString (hash string) (string, error) {

	// Get shell
	sh := getShell(ipfsPath)

	// Get string from IPFS and store it in temp file to read it
	tmpFile := "/tmp/" + hash
	sh.Get(hash, tmpFile)

	// Read the string from temp file
	str, err:= ioutil.ReadFile(tmpFile)
        if err != nil {
                log.Fatal(err)
        }

	// Delete temp file
	os.Remove(tmpFile)

	return string(str), err
}
