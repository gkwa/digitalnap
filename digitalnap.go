package digitalnap

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"

	"github.com/mitchellh/go-homedir"
)

func Main() int {
	slog.Debug("digitalnap", "test", true)

	filePath := "~/Downloads/client_secret_130386091184-6u5if2vcsa4sqqt1j5il2upsb53l5t7h.apps.googleusercontent.com.json"
	main(filePath)

	return 0
}

type GoogleCredentials struct {
	Web struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		// Add other fields as needed
	} `json:"web"`
}

var credentials GoogleCredentials

func main(filePath string) {
	expandedPath, err := homedir.Expand(filePath)
	if err != nil {
		fmt.Println("Error expanding ~ in file path:", err)
		return
	}

	file, err := os.Open(expandedPath)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer file.Close()

	fileContent, err := os.ReadFile(expandedPath)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	err = json.Unmarshal([]byte(fileContent), &credentials)
	if err != nil {
		fmt.Println("Error unmarshalling JSON data:", err)
		return
	}

	fmt.Println("Client ID:", credentials.Web.ClientID)
	fmt.Println("Client Secret:", credentials.Web.ClientSecret)
}
