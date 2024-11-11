package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/pquerna/otp/totp"
)

type OTPResult struct {
	Name string `json:"name"`
	Code string `json:"otp_code"`
}

func main() {

	jsonOutput := flag.Bool("json", false, "output in JSON format")
	tableOutput := flag.Bool("table", false, "output in table format")
	flag.Parse()

	var results []OTPResult

	uris := flag.Args()
	if len(uris) == 0 {
		// no uri provided as arguments so read from stdin
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			uris = append(uris, strings.TrimSpace(scanner.Text()))
		}
		if err := scanner.Err(); err != nil {
			log.Fatalf("Error reading input: %v", err)
		}
	}

	for _, otpURI := range uris {
		if otpURI == "" {
			continue // skip empty lines
		}

		u, err := url.Parse(otpURI)
		if err != nil {
			log.Printf("Failed to parse URI: %v", err)
			continue
		}

		name := strings.TrimPrefix(u.Path, "/totp/")
		name = strings.TrimPrefix(name, "/")
		if name == "" {
			log.Println("Name not found in URI")
			continue
		}

		secret := u.Query().Get("secret")
		if secret == "" {
			log.Println("Secret not found in URI")
			continue
		}

		passcode, err := totp.GenerateCode(secret, time.Now())
		if err != nil {
			log.Printf("Failed to generate OTP for %s: %v", name, err)
			continue
		}

		result := OTPResult{Name: name, Code: passcode}
		results = append(results, result)
	}

	if *jsonOutput {

		jsonData, err := json.MarshalIndent(results, "", "  ")
		if err != nil {
			log.Fatalf("Failed to marshal JSON: %v", err)
		}
		fmt.Println(string(jsonData))

	} else if *tableOutput {

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "OTP Code"})
		for _, result := range results {
			table.Append([]string{result.Name, result.Code})
		}
		table.Render()

	} else {

		for _, result := range results {
			fmt.Printf("Name: %s, OTP Code: %s\n", result.Name, result.Code)
		}

	}
}
