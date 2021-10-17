# hardpassword
Simple console password keeper/generator that not save password in database

## Quickstart:
1. Clone this repo and go to hardpassword directory
2. Run command `go build -o hardpass -ldflags "-X main.secretKey=YOUSECRETKEY" main.go` with replace *YOUSECRETKEY*
3. Or change secretKey manually right at source code and then run `go build -o hardpass main.go`
4. Use it

## Example:
> ./hardpass gmail.username
>
> B@sc/=bg2&Va:<zb`E

or just

> ./hardpass gmail
>
> pt3?+$Pv@R*BbmJ5W8

or 

> ./hardpass google
>
> kztkhujZE](D67Je,(

This command generates the password for the Google account. HardPassword does not save password. It uses hash functions for generate a password.

The generated password depends entirely on the input data. If you change at least one character, then the password will change completely. But for the same input data a password will always be  the same.

## Options:

>__-l__ password length (default 18)

>__-p__ password strength 1-4 (default 4)

## Example:

Change password strength to minimum value:
> ./hardpass -p 1 gmail
>
> 553559853857046019

As you can see it only uses digits

> ./hardpass -p 2 gmail
>
> 7tobjcqv6ly1am8zw3

> ./hardpass -p 3 gmail
>
> F725K0HueVJzBT4fwQ

You can generate a password of any length

> ./hardpass -p 3 -l 100 gmail
>
> RtuIPF46QI9D2v5gk4e9ymJapa1WVbyOm43ldiEhyPse8lK2suM1YSsLj75c9Nb41kDtTbCbJFntqUWqDXC0cKmG24kkTZJ8trjU