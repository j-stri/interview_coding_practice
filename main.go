package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

func main() {
	//
	//Let's start you off with a simpler task than the main interview,
	//given the time available, language unfamiliarity, etc.

	// **Problem statement:**

	// As you might know, in big sporting contests such as the Olympics,
	//there is usually an unofficial medal table,
	//where countries are sorted according to the number of medals received by their athletes.

	// There are several ways to order such tables.
	//One obvious way is to compare the total numbers of medals.
	//Another way, more common in Europe,
	//is to compare the number of gold medals first,
	//then the number of silvers, then bronze.

	// Your task is to write a program in Golang to implement the European ranking system.

	// You are given *n* - the number of medal sets.
	//For each medal set, you are given the list of three countries
	//who received the gold, silver and bronze medals from this set.
	//You have to construct and output the medal table.

	// The formal rules that should be used to compare countries are the following.

	// Country A should be placed higher than country B if:
	// • Country A has more gold medals than country B, or
	// • Country A has the same number of gold medals as B,
	//but more silver medals, or the countries have the same number of gold and silver medals,
	//but A has more bronze ones.

	// **Input definition**

	// The first line of the input contains integer *n* - the number of medal sets
	// (1 ≤ n ≤ 500).

	//Each of the next *n* lines contains IDs of three countries.
	//Each country ID is a string of exactly three uppercase letters (see examples).
	//They are supplied in order of gold, silver, bronze.

	// **Output definition**

	// Print the medal table, one line per country.
	//The first in each line is the place of a country in the standings.
	//It should be followed by the country ID.
	//Then print the numbers of gold, silver and bronze medals for this country.
	//Countries that have exactly the same number of gold, silver and bronze medals
	//should have the same ranking, with printout order being lexicographic.

	// Example input/output

	// 5
	// RUS GER FIN
	// GER USA SWE
	// IRL SWE GER
	// GER RUS USA
	// RUS IRL RUS

	// yields
	// 1 GER 2 1 1
	// 1 RUS 2 1 1
	// 3 IRL 1 1 0
	// 4 SWE 0 1 1
	// 4 USA 0 1 1
	// 6 FIN 0 0 1
	// I hope that's clear?

	//for parsing the standard input it's AI written code

	// Create a new scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)

	// Read the first line to get 'n', the number of lines to follow
	scanner.Scan()
	firstLine := scanner.Text()
	n, err := strconv.Atoi(strings.TrimSpace(firstLine))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading the number of lines:", err)
		return
	}

	input := make([]string, n)

	// Read the next 'n' lines
	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()
		// Process the line here
		input[i] = line

		fmt.Println("Received line:", line)
	}

	// Check for any errors that occurred during scanning
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from input:", err)
	}

	printEuroRank(input)

}

type Country struct {
	countrycode string
	gold        int
	silver      int
	bronze      int
}

func printEuroRank(input []string) {
	//n := len(input)

	outmap := make(map[string]*Country)

	for _, line := range input {
		for medal_kind, country := range strings.Split(line, " ") {
			_, ok := outmap[country]
			if !ok && country != "" {
				outmap[country] = &Country{countrycode: country, gold: 0, silver: 0, bronze: 0}
			}
			switch medal_kind {
			case 0:
				outmap[country].gold += 1
			case 1:
				outmap[country].silver += 1
			case 2:
				outmap[country].bronze += 1

			}

		}

	}

	countries := maps.Values(outmap)

	//alternative method without maps import:
	//v := make([]string, 0, len(m))
	// for  _, value := range m {
	// 	v = append(v, value)
	// }

	//ref: https://stackoverflow.com/questions/13422578/in-go-how-to-get-a-slice-of-values-from-a-map

	sort.SliceStable(countries, func(i, j int) bool {
		return countries[i].countrycode < countries[j].countrycode
	})

	sort.SliceStable(countries, func(i, j int) bool {

		return countries[i].gold > countries[j].gold ||
			(countries[i].gold == countries[j].gold && countries[i].silver > countries[j].silver) ||
			(countries[i].gold == countries[j].gold && countries[i].silver == countries[j].silver &&
				countries[i].bronze > countries[j].bronze)
	})

	//fmt.Println(countries)

	rank := 1
	tie := 0

	for i, v := range countries {

		if i > 0 &&
			countries[i].gold == countries[i-1].gold &&
			countries[i].silver == countries[i-1].silver &&
			countries[i].bronze == countries[i-1].bronze {
			tie = 1
		}

		fmt.Printf("%d %s %d %d %d\n", rank-tie, v.countrycode, v.gold, v.silver, v.bronze)

		rank += 1
		tie = 0

	}

}
