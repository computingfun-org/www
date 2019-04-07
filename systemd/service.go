package systemd

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// InstallFlag is the "systemd" flag pointer.
// Is set when flag.Parse is called.
var InstallFlag *bool

func init() {
	InstallFlag = flag.Bool("systemd", false, "Install systemd service 💾 ")
}

// InstallMain calls InstallDisplay then exits the application with the error code returned if the "systemd" was set.
// If the "systemd" flag wasn't set InstallMain does nothing.
// Must be called after flag.Parse.
func InstallMain() {
	if *InstallFlag {
		if InstallDisplay() == nil {
			os.Exit(0)
		}
		os.Exit(1)
	}
}

// InstallDisplay calls Install and logs the output in user friendly text.
func InstallDisplay() error {
	log.Println("💾  Installing systemd service:")
	if err := Install(); err != nil {
		log.Println("\t❌  Failed: " + err.Error())
		return err
	}
	log.Println("\t✔️  Success")
	return nil
}

// Install the systemd service file making this executable (at this location) a service that runs on startup.
func Install() error {
	path, err := os.Executable()
	if err != nil {
		return err
	}
	file := []byte("\n[Unit]\nDescription=Computing Fun web server\n[Service]\nExecStart=" + path + "\nWorkingDirectory=" + filepath.Dir(path) + "\n[Install]\nWantedBy=multi-user.target")
	return ioutil.WriteFile("/etc/systemd/system/cf-www.service", file, 0664)
}
