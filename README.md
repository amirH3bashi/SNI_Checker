
# 🔍 SNI Checker

A fast and minimal tool to check the availability, ping, packet loss, and TLS handshake of a list of SNI hostnames — sorted by best response time.

![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go&logoColor=white)
![Build](https://img.shields.io/badge/build-passing-brightgreen)
![License](https://img.shields.io/github/license/amirH3bashi/SNI_Checker)
![Platform](https://img.shields.io/badge/platform-Linux%20%7C%20macOS%20%7C%20WSL-lightgrey)

---

## ✨ Features

- 🟢 Check if port `443` is open
- 📶 Measure average ping and packet loss
- 🔐 Verify TLS handshake with SNI
- 📄 Outputs results sorted by best ping
- 📥 Reads from `sni.txt`
- 📤 Writes to `sni_valid.txt`
- 🚀 Easy one-liner installer

---

## 🚀 One-line Install & Run

Just run:

```bash
bash <(curl -Ls https://raw.githubusercontent.com/amirH3bashi/SNI_Checker/main/install.sh)
````

This will:

* Install Go (if not installed)
* Clone the project
* Build the binary
* Run `sni-checker` automatically

---

## 📥 Input Format

Place your SNI list in a file named `sni.txt` (already included in the repo):

```
google.com
cloudflare.com
example.com
...
```

There is no limit to the number of entries. It works fine with large input.

---

## 📤 Output Format

Results will be written to `sni_valid.txt` and sorted by lowest ping:

```
cloudflare.com | Ping: 12.3ms | Loss: 0.0%
google.com     | Ping: 18.1ms | Loss: 0.0%
example.com    | Ping: 45.6ms | Loss: 0.0%
```

You can use the result in other network tools or scripts.

---

## 🧱 Manual Build

If you want to clone and run manually:

```bash
git clone https://github.com/amirH3bashi/SNI_Checker.git
cd SNI_Checker
go mod tidy
go build -o sni-checker main.go
./sni-checker
```

---

## ⚙ Dependencies

This tool uses the following Go modules:

* [`go-ping`](https://github.com/go-ping/ping)
* [`fatih/color`](https://github.com/fatih/color)
* [`schollz/progressbar`](https://github.com/schollz/progressbar)

Install them with:

```bash
go mod tidy
```

---

## 📝 License

This project is licensed under the MIT License.

---

## 🤝 Contributing

Pull requests, issues, and stars are welcome!

---

## 📫 Contact

Made with ❤️ by [amirH3bashi](https://github.com/amirH3bashi)
