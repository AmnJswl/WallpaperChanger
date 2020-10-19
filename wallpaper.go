package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/reujab/wallpaper"
)

func main() {
	fmt.Println("Wallpaper changer by Amn!")

	var t string
	var interval = 30
	var err error = nil

	if len(os.Args) < 2 {
		fmt.Println("Time Interval not provided. Defaulting to", interval, "Minute")
	} else {
		t = os.Args[1]
		interval, err = strconv.Atoi(t)
	}

	if err != nil {
		fmt.Println("Improper Argument!")
		interval = 30
	}

	fmt.Println("Setting time interval :", interval, "Minute")

	for {
		fmt.Println("Getting a new Random Image from Lorem Picsum")
		file, err := DownloadFile("1.jpg", "https://picsum.photos/1920/1080")
		if err != nil {
			Log("Unable to Download file")
		}
		location, err := filepath.Abs(file)
		result := wallpaper.SetFromFile(location)
		currentTime := time.Now()
		if result != nil {
			// Log(currentTime.Format("2006.01.02 15:04:05") + " : Error reaching Lorem Picsum!")
			Log("Unable to Set Wallpaper")
		}
		fmt.Println("Wallpaper Set: " + location)
		Log(currentTime.Format("2006.01.02 15:04:05") + " : Wallpaper Updated!")
		time.Sleep(time.Duration(interval) * time.Minute)
	}

}

// Log : For Logging purposes
func Log(text string) {
	logfile, err := os.OpenFile("wallpaper.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatalln("Unable to open file")
	}

	if _, err := logfile.Write([]byte(text + "\n")); err != nil {
		log.Fatalln(err)
	}

	if err := logfile.Close(); err != nil {
		log.Fatal(err)
	}
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) (string, error) {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return "", err
	}

	filename := out.Name()

	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return filename, err
}
