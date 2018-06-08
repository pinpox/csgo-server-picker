package main

import "fmt"

import "github.com/coreos/go-iptables/iptables"

// CHAIN specifies the chain that will be used by the program
var CHAIN = "OUTPUT"

// TABLE specifies the table that will be used by the program
var TABLE = "filter"

func main() {

	ipt, err := iptables.New()
	if err != nil {
		panic(err)
	}

	// block ip
	// iptables -A INPUT -s 10.10.10.10 -j DROP

	// block range
	// iptables -A INPUT -s 10.10.10.0/24 -j DROP

	// delete rule (replace -A with -D)
	// iptables -D INPUT -s 10.10.10.10 -j DROP

	fmt.Println("List existing rules")
	ListRules(ipt)

	fmt.Println("Bock ip 10.10.10.10")
	BlockServer(ipt, "10.10.10.10")

	fmt.Println("List rules again")
	ListRules(ipt)

	fmt.Println("Unblock ip 10.10.10.10")
	UnblockServer(ipt, "10.10.10.10")

	fmt.Println("List rules one last time")
	ListRules(ipt)
}

// ListRules prints all IP rules
func ListRules(ipt *iptables.IPTables) {

	// func (ipt *IPTables) List(table, chain string) ([]string, error)
	rules, err := ipt.List(TABLE, CHAIN)
	if err != nil {
		fmt.Println("Could not list rules")
		panic(err)
	}
	fmt.Println(rules)
}

// BlockServer blocks a specific IP
func BlockServer(ipt *iptables.IPTables, ip string) {

	// func (ipt *IPTables) Append(table, chain string, rulespec ...string) error
	// -I INPUT -s {IP-HERE} -j DROP
	ipt.AppendUnique(TABLE, CHAIN, "-s", ip, "-j", "ACCEPT")
}

// UnblockServer removes the block rule for a specific ip
func UnblockServer(ipt *iptables.IPTables, ip string) {

	ipt.Delete(TABLE, CHAIN, "-s", ip, "-j", "ACCEPT")

}

// Cleanup removes all created rules
func Cleanup(ipt *iptables.IPTables) {

}
