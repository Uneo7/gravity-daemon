package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func Download(server Server, name string, url string) {

	LogToConsole(server, "[Downloader] Starting download")

	filePath := path.Join(server.Path, name)
	output, err := os.Create(filePath)

	if err != nil {
		fmt.Println("[Error] File creation failed", filePath, "-", err)
		LogToConsole(server, "[Downloader][Error] File creation failed")
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading:", url, "-", err)
		LogToConsole(server, "[Downloader][Error] Download failed")

		return
	}
	defer response.Body.Close()

	_, err = io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading:", url, "-", err)
		LogToConsole(server, "[Downloader][Error] Download failed")
		return
	}

	LogToConsole(server, "[Downloader] "+name+" successfully downloaded")
}
