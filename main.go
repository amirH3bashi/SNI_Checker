package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/go-ping/ping"
	"github.com/schollz/progressbar/v3"
)

type SNIResult struct {
	Host       string
	Ping       time.Duration
	PacketLoss float64
	Reachable  bool
}

func checkSNI(sni string) bool {
	dialer := &net.Dialer{Timeout: 5 * time.Second}
	conn, err := tls.DialWithDialer(dialer, "tcp", sni+":443", &tls.Config{
		ServerName:         sni,
		InsecureSkipVerify: true,
	})
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func measurePing(host string) (time.Duration, float64, error) {
	pinger, err := ping.NewPinger(host)
	if err != nil {
		return 0, 100.0, err
	}
	pinger.Count = 3
	pinger.Timeout = 5 * time.Second
	pinger.SetPrivileged(true)

	err = pinger.Run()
	if err != nil {
		return 0, 100.0, err
	}

	stats := pinger.Statistics()
	return stats.AvgRtt, stats.PacketLoss, nil
}

func readSNIList(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var sniList []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			sniList = append(sniList, line)
		}
	}
	return sniList, scanner.Err()
}

func isPortOpen(host string, port string) bool {
	timeout := 3 * time.Second
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}


func main() {
	inputFile := "sni.txt"
	outputFile := "sni_valid.txt"

	sniList, err := readSNIList(inputFile)
	if err != nil {
		color.Red("❌ Failed to read input file: %v", err)
		return
	}

	color.Cyan("🔍 Checking %d SNI hostnames...\n", len(sniList))
	bar := progressbar.NewOptions(len(sniList),
		progressbar.OptionSetDescription("Processing..."),
		progressbar.OptionShowCount(),
		progressbar.OptionClearOnFinish(),
	)

	var results []SNIResult

	for _, sni := range sniList {
		bar.Add(1)

		if !isPortOpen(sni, "443") {
			color.Red("❌ %s - port 443 closed", sni)
			continue
		}

		reachable := checkSNI(sni)
		pingTime, loss, err := measurePing(sni)

		if err != nil || loss >= 100 {
			color.Red("❌ %s - ping failed or 100%% loss (%v)", sni, err)
			continue
		}		

		if reachable {
			color.Green("✅ %s - ping: %v | loss: %.1f%%", sni, pingTime, loss)
			results = append(results, SNIResult{
				Host:       sni,
				Ping:       pingTime,
				PacketLoss: loss,
				Reachable:  true,
			})
		} else {
			color.Red("❌ %s - TLS handshake failed", sni)
		}
	}


	// مرتب‌سازی بر اساس پینگ (از کم به زیاد)
	sort.Slice(results, func(i, j int) bool {
		return results[i].Ping < results[j].Ping
	})

	// ذخیره در فایل خروجی
	file, err := os.Create(outputFile)
	if err != nil {
		color.Red("❌ Failed to create output file: %v", err)
		return
	}
	defer file.Close()

	for _, result := range results {
		line := fmt.Sprintf("%s | Ping: %v | Loss: %.1f%%\n", result.Host, result.Ping, result.PacketLoss)
		_, err := file.WriteString(line)
		if err != nil {
			color.Red("❌ Failed to write to file: %v", err)
		}
	}
	

	color.Yellow("\n🎉 Done! Valid SNI saved to '%s' (sorted by best ping)\n", outputFile)
}
