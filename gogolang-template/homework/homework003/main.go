package main

// cmd/cli/cli.go
import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func nsComdFunc(c *cli.Context) {
	ns, err := net.LookupNS("www.baidu.com")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(ns)
}
func cnameComdFunc(c *cli.Context) {
	ns, err := net.LookupCNAME("www.baidu.com")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(ns)
}
func ipComdFunc(c *cli.Context) {
	ns, err := net.LookupIP("www.baidu.com")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(ns)
}

func main() {
	// nsflag := flag.NewFlagSet("ns", flag.ExitOnError)
	// cnameflag := flag.NewFlagSet("cname", flag.ExitOnError)
	// ipflag := flag.NewFlagSet("ip", flag.ExitOnError)
	// switch os.Args[1] {
	// case "ns":
	// 	err := nsflag.Parse(os.Args[2:])
	// 	if err != nil {
	// 		log.Panic(err)
	// 	}
	// 	ns, _ := net.LookupIP("www.baidu.com")
	// 	fmt.Println(ns)
	// case "cname":
	// 	err := cnameflag.Parse(os.Args[2:])
	// 	if err != nil {
	// 		log.Panic(err)
	// 	}
	// 	cname, _ := net.LookupCNAME("www.baidu.com")
	// 	fmt.Println(cname)
	// case "ip":
	// 	err := ipflag.Parse(os.Args[2:])
	// 	if err != nil {
	// 		log.Panic(err)
	// 	}
	// 	ip, _ := net.LookupNS("www.baidu.com")
	// 	fmt.Println(ip)
	// }

	cliFlag := cli.NewApp()
	cliFlag.Name = "Website Lookup CLI"
	cliFlag.Usage = "Let's you query IPs, CNAMEs and Name Servers!"
	cliFlag.Version = "0.0.0"

	nsComd := cli.Command{
		Name:   "ns",
		Usage:  "Looks Up the NameServers for a Particular Host",
		Action: nsComdFunc,
	}
	cnameComd := cli.Command{
		Name:   "cname",
		Usage:  "Looks up the CNAME for a particular host",
		Action: cnameComdFunc,
	}
	ipComd := cli.Command{
		Name:   "ip",
		Usage:  "Looks up the IP addresses for a particular host",
		Action: ipComdFunc,
	}

	cliFlag.Commands = cli.Commands{nsComd, cnameComd, ipComd}
	err := cliFlag.Run(os.Args[:])
	if err != nil {
		log.Fatal(err)
	}
}
