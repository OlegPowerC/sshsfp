package main

import (
	"flag"
	"fmt"
	"golang.org/x/crypto/ssh"
	"strings"
	"os"
)

var (
	user = flag.String("u", "", "User name")
	host = flag.String("h", "", "Host")
	port = flag.Int("p", 22, "Port")
	passwd = flag.String("pw", "", "Password")
	interfaces = flag.String("i", "", "Interfaces separate by comma")
)

func main() {

	//var hostKey ssh.PublicKey
	flag.Parse()
	config := &ssh.ClientConfig{
		User: *user,
		Auth: []ssh.AuthMethod{
			ssh.Password(*passwd),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),

	}

	addr := fmt.Sprintf("%s:%d", *host, *port)
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		panic(err)
	}


	strone12 := "clear mac address-table dynamic"
	if len(*interfaces) > 3{
		interfacessplits := strings.Split(*interfaces,",")
		stprefix := "clear mac address-table dynamic interface"
		strone12 = ""
		for _,stint := range interfacessplits{
			strone12 = strone12+stprefix+" "+stint+","
		}
	}
	strone12 = strone12+",exit"
	strone1 := strings.Split(strone12,",")

	// Create sesssion
	sess, err := client.NewSession()
	if err != nil {
		fmt.Println("Error1")
	}
	defer sess.Close()

	// StdinPipe for commands
	stdin, err := sess.StdinPipe()
	if err != nil {
		fmt.Println("Error1")
	}

	// Enable system stdout
	// Comment these if you uncomment to store in variable
	sess.Stdout = os.Stdout
	sess.Stderr = os.Stderr

	// Start remote shell
	err = sess.Shell()
	if err != nil {
		fmt.Println("Error1")
	}

	for _, cmd := range strone1 {
		_, err = fmt.Fprintf(stdin, "%s\n", cmd)
		if err != nil {
			fmt.Println("Error1")
		}
	}

	// Wait for sess to finish
	err = sess.Wait()
	if err != nil {
		fmt.Println("Error1")
	}

}