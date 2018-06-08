package main

import "fmt"

import "github.com/coreos/go-iptables/iptables"

var blockingTable iptables.IPTables

var CHAIN = "CSGO"
var TABLE = "filter"

func main() {
	fmt.Println("vim-go")

	// block ip
	// iptables -A INPUT -s 10.10.10.10 -j DROP

	// block range
	// iptables -A INPUT -s 10.10.10.0/24 -j DROP

	// delete rule (replace -A with -D)
	// iptables -D INPUT -s 10.10.10.10 -j DROP

	fmt.Println("List existing rules")
	ListRules()

	fmt.Println("Bock ip 10.10.10.10")
	BlockServer("10.10.10.10")

	fmt.Println("List rules again")
	ListRules()

	fmt.Println("Unblock ip 10.10.10.10")
	UnblockServer("10.10.10.10")

	fmt.Println("List rules one last time")
	ListRules()
}

// ListRules prints all IP rules
func ListRules() {

	// func (ipt *IPTables) List(table, chain string) ([]string, error)
	rules, err := blockingTable.List(TABLE, CHAIN)
	if err != nil {
		fmt.Println("Could not list rules")
		panic(err)
	}
	fmt.Println(rules)
}

// BlockServer blocks a specific IP
func BlockServer(ip string) {

	// func (ipt *IPTables) Append(table, chain string, rulespec ...string) error
	// -I INPUT -s {IP-HERE} -j DROP

	rule := "iptables -A INPUT -s 65.55.44.100 -j DROP"
	blockingTable.Append("blockingTable", "csgo", rule)

}

// UnblockServer removes the block rule for a specific ip
func UnblockServer(ip string) {
}

// Cleanup removes all created rules
func Cleanup() {

}
