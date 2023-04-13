# BlindFire

**Automated Statistically Likely for your User Enumeration.**

**NOT** Windows Compatible

This was a repetitive Perimeter process of identifying the email schema, downloading the correct file, and editing. The correct file can now be selected and modified from a single tool!

https://github.com/insidetrust/statistically-likely-usernames

## Installation
```
$ go mod tidy
$ go build
$ ./BlindFire -h
```

## Execute From Anywhere
```mv BlindFire /usr/local/bin```


# Usage
```
└─$ ./BlindFire -h          
Usage of ./BlindFire:
  -domain string
        Domain To Append (default " ")
  -empty
        Do NOT append domain to the list
  -list string
        List To Download

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
