package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"

	"github.com/cavaliergopher/grab/v3"
)

var filename string

func selectList() int {
	fmt.Println("Select Statistically Likely Format:")
	fmt.Println("  1. jjs")
	fmt.Println("  2. jjsmith")
	fmt.Println("  3. john.smith")
	fmt.Println("  4. john")
	fmt.Println("  5. johnjs")
	fmt.Println("  6. johns")
	fmt.Println("  7. johnsmith")
	fmt.Println("  8. jsmith")
	fmt.Println("  9. jsmith2")
	fmt.Println("  10. places")
	fmt.Println("  11. service-accounts")
	fmt.Println("  12. smith")
	fmt.Println("  13. smithj")
	fmt.Println("  14. smithj2")
	fmt.Println("  15. smithjj")
	fmt.Println("")
	fmt.Println("Number:")

	var selection int
	fmt.Scanln(&selection)
	return selection
}

func domainValidatiton(domain string) {
	matched, err := regexp.MatchString(`^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z]{2,3})$`, domain)
	if matched== false {
		
		fmt.Println("The answer is an Invalid Entry")
		fmt.Println("Would you like to add a domain? (Y or N)")
		var question string
		fmt.Scanln(&question)
		if question == "Y" || question == "y" {
			domainInput()
		}
		if question == "N" || question == "n" {
			quitter()
			fmt.Println(err)
		}

	} else {
		appendFile(domain)
	}

}

func quitter() {
	fmt.Println("Exiting Without Appending a Domain")
	os.Exit(3)
}

func appendFile(domain string) {
	cmd := exec.Command("sed", "-i", "s/$/@"+domain+"/", filename)
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("")
}

func downloadFile(number int) {
	entry := number - 1
	test := [15]string{"jjs.txt", "jjsmith.txt", "john.smith.txt", "john.txt", "johnjs.txt", "johns.txt", "johnsmith.txt", "jsmith.txt", "jsmith2.txt", "places.txt", "service-accounts.txt", "smith.txt", "smithj.txt", "smithj2.txt", "smithjj.txt"}
	var savedURL string
	savedURL = "https://raw.githubusercontent.com/insidetrust/statistically-likely-usernames/master/"
	resp, err := grab.Get(".", savedURL+test[entry])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Download saved to", resp.Filename)
	filename = test[entry]
}

func domainInput() {
	fmt.Println("Enter Domain to Append: ")
	var domain string
	fmt.Scanln(&domain)
	domainValidatiton(domain)

}

func main() {
	const header = ".______    __       __  .__   __.  _______   _______  __  .______       _______ 	\n" +
		"|   _  \\  |  |     |  | |  \\ |  | |       \\ |   ____||  | |   _  \\     |   ____|		\n" +
		"|  |_)  | |  |     |  | |   \\|  | |  .--.  ||  |__   |  | |  |_)  |    |  |__   		\n" +
		"|   _  <  |  |     |  | |  . `  | |  |  |  ||   __|  |  | |      /     |   __|  		\n" +
		"|  |_)  | |  `----.|  | |  |\\   | |  '--'  ||  |     |  | |  |\\  \\----.|  |____ 		\n" +
		"|______/  |_______||__| |__| \\__| |_______/ |__|     |__| | _| `._____||_______|		\n" +
		"                                                                                       \n" +
		"                                                   @AchocolatechipPancake\n"

	fmt.Print(header)
	num := selectList()
	downloadFile(num)
	domainInput()
	fmt.Println("File Is Ready For User Enumeration")
	os.Exit(3)
}
