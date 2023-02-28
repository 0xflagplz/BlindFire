# BlindFire

**Automated Statistically Likely for your User Enumeration.**

**NOT** Windows Compatible

This was a repetitive Perimeter process of identifying the email schema, downloading the correct file, and editing. The correct file can now be selected and modifyed from a single tool!

https://github.com/insidetrust/statistically-likely-usernames

## Installation
```
$ go mod tidy
$ go build
$ ./NameGen
```

## Execute From Anywhere
```mv NameGen /usr/local/bin```


# Example Usage
```
$ ./NameGen
.__   __.      ___      .___  ___.  _______   _______  _______ .__   __.
|  \ |  |     /   \     |   \/   | |   ____| /  _____||   ____||  \ |  |
|   \|  |    /  ^  \    |  \  /  | |  |__   |  |  __  |  |__   |   \|  |
|  . `  |   /  /_\  \   |  |\/|  | |   __|  |  | |_ | |   __|  |  . `  |
|  |\   |  /  _____  \  |  |  |  | |  |____ |  |__| | |  |____ |  |\   |
|__| \__| /__/     \__\ |__|  |__| |_______| \______| |_______||__| \__|

                                                   @AchocolatechipPancake
Select Statistically Likely Format:
  1. jjs
  2. jjsmith
  3. john.smith
  4. john
  5. johnjs
  6. johns
  7. johnsmith
  8. jsmith
  9. jsmith2
  10. places
  11. service-accounts
  12. smith
  13. smithj
  14. smithj2
  15. smithjj
>> 3
Download saved to john.smith.txt
File exists
Enter Domain to Append:
>> domain.com

File Is Ready For User Enumeration
```

## Results
```$ head john.smith.txt```
```
john.smith@domain.com
david.smith@domain.com
michael.smith@domain.com
chris.smith@domain.com
mike.smith@domain.com
arun.kumar@domain.com
james.smith@domain.com
amit.kumar@domain.com
imran.khan@domain.com
jason.smith@domain.com
```
