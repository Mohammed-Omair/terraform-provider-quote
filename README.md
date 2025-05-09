# Terraform Provider: Random Quote 🌟  
_A lightweight Terraform provider to fetch random quotes from [Quotable.io](https://quotable.io)._  

[![Terraform Version](https://img.shields.io/badge/terraform-%3E%3D1.0-623CE4?logo=terraform)](https://terraform.io)
[![GitHub Release](https://img.shields.io/github/v/release/Mohammed-Omair/terraform-provider-quote)](https://github.com/Mohammed-Omair/terraform-provider-quote/releases)

---

## Features ✨  
- **Random Quotes**: Fetch a new quote on `terraform apply`.  
- **Simple Attributes**: Access `content` and `author` fields.  
- **Cross-Platform**: Pre-built binaries for **Windows, Linux, and macOS**.  
- **Educational**: Demonstrates Terraform provider development basics.  

---

## Installation  

### Using GitHub Releases (Recommended)  

#### Downloading The Right File (Linux)
```bash
ARCH="linux_amd64"  # For ARM64: "linux_arm64"

# Download provider binary
wget "https://github.com/Mohammed-Omair/terraform-provider-quote/releases/download/v0.1.0/terraform-provider-quote_0.1.0_${ARCH}.zip"

# Unzip
unzip "terraform-provider-quote_0.1.0_${ARCH}.zip"
```

#### Downloading The Right File (Windows)
```shell
$ARCH = "windows_amd64"  # For ARM64: "windows_arm64"

# Download provider binary
Invoke-WebRequest -Uri "https://github.com/Mohammed-Omair/terraform-provider-quote/releases/download/v0.1.0/terraform-provider-quote_${ARCH}.exe.zip" -OutFile "terraform-provider-quote_windows_amd64.exe.zip"

# Unzip
Expand-Archive -Path "terraform-provider-quote_windows_amd64.exe.zip" -DestinationPath .
```

#### Moving The Binary to The Right Directory
For windows, create and move the binary to the folder ```APPDATA\terraform.d\plugins\github.com\Mohammed-Omair\quote\0.1.0\windows_amd64```

For linux, create and move the binary to the folder ```~/.terraform.d/plugins/github.com/Mohammed-Omair/quote/0.1.0/linux_amd64/terraform-provider-quote_linux_amd64``` and also grant permissions using ```chmod +x ~/.terraform.d/plugins/github.com/Mohammed-Omair/quote/0.1.0/${ARCH}/terraform-provider-quote_linux_amd64```

### Local Development

- Step 1: Clone the repository: ```git clone https://github.com/Mohammed-Omair/terraform-provider-quote.git```
- Step 2: Build the provider: ```cd terraform-provider-quote && go build -o terraform-provider-quote```
- Step 3: Move the binary to the right directory. It is similar to the steps described in the previous method.

## Usage Example
### Terraform configuration (`main.tf`):  
```hcl
terraform {
  required_providers {
    quote = {
      source  = "github.com/Mohammed-Omair/quote"
      version = "0.1.0"                            
    }
  }
}

resource "quote_random" "example" {}

output "quote" {
  value = "✨ ${quote_random.example.content} — ${quote_random.example.author}"
}
```

### Sample Output:
```bash
Apply complete! Resources: 1 added, 0 changed, 0 destroyed.

Outputs:

quote = "✨ Flow with whatever is happening and let your mind be free. Stay centered by accepting whatever you are doing. This is the ultimate. — Zhuang Zhou"
```

## Troubleshooting ( Things I Learnt )
| Issue    | Solution |
| -------- | ------- |
| Antivirus Blocks Provider Binary  | Exclude the directory in which the plugin resides    |
| Certificate Errors | Added code to skip SSL verification    |
| Development override warning    | Safe to ignore in local testing    |

## Acknowledgments
- Quotes provided by [Quotable.io](https://quotable.io)
- Inspired by HashiCorp’s Provider Development Guide
