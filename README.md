# topt
A simple program that reads TOTP URIs from standard input and generates OTP codes, outputs the codes in various formats (plain text, JSON, or table). 

## Installation
```
go install github.com/towsifkafi/topt@latest
```
If you want to build it yourself, feel free to clone the repository and run `go build`

## Usage

```sh
# passing the URI as an argument
topt "otpauth://totp/Example%20OTP?secret=verysecret&algorithm=SHA1&digits=6&period=30"

# pipe into topt
echo "otpauth://totp/Example%20OTP?secret=verysecret&algorithm=SHA1&digits=6&period=30" | topt

# Output: Name: Example OTP, OTP Code: 524777

# it also supports multiple URIs at once (should be separated by newlines)

# if you have it in clipboard
pbpaste | topt
# if you're in linux (requires xsel)
xsel --clipboard --output | topt
# get secrets from keyring
lssecret -s | grep otpauth:// | sed "s/Secret:\t//g" | topt --table
```

## Flags
<details open>
  <summary><code>topt --json</code> : Output in JSON format</summary>
<br>

```json
[
  {
    "name": "Example OTP",
    "otp_code": "005673"
  }
]
```
</details>

<details open>
  <summary><code>topt --table</code> : Output in a table</summary>
<br>

```json
+-------------+----------+
|    NAME     | OTP CODE |
+-------------+----------+
| Example OTP |   660968 |
+-------------+----------+
```
</details>
