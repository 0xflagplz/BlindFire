package main

import (
    "fmt"
    "log"
    "os"
    "os/exec"

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

    var selection int
    fmt.Scanln(&selection)
    return selection
}

func DownloadFile(number int) {
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

func editFile() {
    if _, err := os.Stat(filename); err == nil {
        fmt.Printf("File exists\n")
    } else {
        fmt.Printf("File does not exist\n")
    }
    fmt.Println("Enter Domain to Append: ")
    var domain string
    fmt.Scanln(&domain)
    cmd := exec.Command("sed", "-i","s/$/@"+domain+"/",filename)
    err := cmd.Run()

    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("")
    fmt.Println("File Is Ready For User Enumeration")
}

func main() {
    const header = ".__   __.      ___      .___  ___.  _______   _______  _______ .__   __.     \n" +
        "|  \\ |  |     /   \\     |   \\/   | |   ____| /  _____||   ____||  \\ |  |    \n" +
        "|   \\|  |    /  ^  \\    |  \\  /  | |  |__   |  |  __  |  |__   |   \\|  |    \n" +
        "|  . `  |   /  /_\\  \\   |  |\\/|  | |   __|  |  | |_ | |   __|  |  . `  |     \n" +
        "|  |\\   |  /  _____  \\  |  |  |  | |  |____ |  |__| | |  |____ |  |\\   |     \n" +
        "|__| \\__| /__/     \\__\\ |__|  |__| |_______| \\______| |_______||__| \\__|     \n" +
        "                                                                         \n" +
        "                                                   @AchocolatechipPancake\n"

    fmt.Print(header)
    num := selectList()
    DownloadFile(num)
    editFile()
    os.Exit(3)
}
