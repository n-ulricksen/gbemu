package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var flagCompare bool

var allowedFiletypes = map[string]bool{
	".txt": true,
	".log": true,
}

func main() {
	setupFlags()
	args := flag.Args()

	if flagCompare {
		if len(args) != 2 {
			printUsage()
			os.Exit(0)
		}

		// Check file type
		fp1, fp2 := args[0], args[1]
		ext1, ext2 := filepath.Ext(fp1), filepath.Ext(fp2)
		if !allowedFiletypes[ext1] || !allowedFiletypes[ext2] {
			var allowed []string
			for ft := range allowedFiletypes {
				allowed = append(allowed, ft)
			}
			fmt.Printf("allowed filetypes: %s\n", allowed)
			os.Exit(0)
		}

		// Load the log files
		log1, err := os.Open(fp1)
		if err != nil {
			fmt.Printf("could not open file %s\n", fp1)
			fmt.Println(err)
			os.Exit(0)
		}
		defer log1.Close()

		log2, err := os.Open(fp2)
		if err != nil {
			fmt.Printf("could not open file %s\n", fp2)
			fmt.Println(err)
			os.Exit(0)
		}
		defer log1.Close()

		// Scan both log files
		scanner1 := bufio.NewScanner(log1)
		scanner2 := bufio.NewScanner(log2)
		isScanning := true

		// Channels used to capture one line at a time from each log file
		lineChan1 := make(chan string)
		lineChan2 := make(chan string)

		go func() {
			for scanner1.Scan() {
				lineChan1 <- scanner1.Text()
			}
			if err := scanner1.Err(); err != nil {
				fmt.Printf("failed to scan line of log 1\n")
				fmt.Println(err)
				os.Exit(0)
			}
			close(lineChan1)
			isScanning = false
		}()
		go func() {
			for scanner2.Scan() {
				lineChan2 <- scanner2.Text()
			}
			if err := scanner2.Err(); err != nil {
				fmt.Printf("failed to scan line of log 2\n")
				fmt.Println(err)
				os.Exit(0)
			}
			close(lineChan2)
			isScanning = false
		}()

		memAddrRe := regexp.MustCompile(`\[0x[\w]{4}\]`)
		cycleRe := regexp.MustCompile(`\d+$`)

		const linesToPrint int = 8
		lines1 := make([]string, linesToPrint)
		lines2 := make([]string, linesToPrint)
		linesIdx := 0

		lineCount := 0
		for isScanning {
			lineCount++

			line1 := <-lineChan1
			line2 := <-lineChan2
			lines1[linesIdx] = line1
			lines2[linesIdx] = line2
			linesIdx = (linesIdx + 1) % linesToPrint

			memAddr1 := memAddrRe.FindString(line1)
			memAddr2 := memAddrRe.FindString(line2)
			cycle1 := cycleRe.FindString(line1)
			cycle2 := cycleRe.FindString(line2)

			if memAddr1 != memAddr2 || cycle1 != cycle2 {
				fmt.Printf("\nDIFF IN LOGS AT LINE %d:\n", lineCount)
				fmt.Println("-------------------------------------")
				fmt.Printf("%-18s%s\t%s\n", "LOG", "ADDRESS", "CYCLE")
				fmt.Printf("%-18s%s\t%s\n", fp1, memAddr1, cycle1)
				fmt.Printf("%-18s%s\t%s\n", fp2, memAddr2, cycle2)
				fmt.Println("-------------------------------------")
				fmt.Println()

				// Print previous lines from each log file
				fmt.Printf("%s:\n", fp1)
				for i := 0; i < linesToPrint; i++ {
					idx := (i + linesIdx) % linesToPrint
					fmt.Println(lines1[idx])
				}
				fmt.Println()
				fmt.Printf("%s:\n", fp2)
				for i := 0; i < linesToPrint; i++ {
					idx := (i + linesIdx) % linesToPrint
					fmt.Println(lines2[idx])
				}
				fmt.Println("-------------------------------------")

				os.Exit(0)
			}
		}

		fmt.Printf("Complete! Scanned %d lines.\n", lineCount)
	} else {
		printUsage()
		os.Exit(0)
	}
}

func setupFlags() {
	flag.BoolVar(&flagCompare, "c", false,
		"Specify two log files to compare")
	flag.Parse()
}

func printUsage() {
	fmt.Println("usage: ./logparse -c <log1.txt> <log2.txt>")
}
