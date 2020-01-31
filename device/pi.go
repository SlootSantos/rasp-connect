package device

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

// Pi is a Raspberry Pi that
// can be connected to
type Pi struct {
	Nickname string
	Hostname string
	User     string
	Password string
	Port     string
}

// Connect method
// connects to the receiver raspberry via ssh
func (p *Pi) Connect() error {
	hostKey := createHostKey()

	sshConf := &ssh.ClientConfig{
		User: p.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(p.Password),
		},
		HostKeyCallback: hostKey,
	}

	connection, err := ssh.Dial("tcp", p.Hostname+":"+p.Port, sshConf)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	session, err := connection.NewSession()
	if err != nil {
		fmt.Println("Error", err)
	}

	defer session.Close()

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		session.Close()
		fmt.Errorf("request for pseudo terminal failed: %s", err)
	}

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	in, _ := session.StdinPipe()

	if err := session.Shell(); err != nil {
		log.Fatal("failed to start shell: ", err)
	}

	// Handle control + C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for {
			<-c
			fmt.Fprint(in, "\n")

			session.Close()
			os.Exit(0)
		}
	}()

	// Accepting commands
	for {
		reader := bufio.NewReader(os.Stdin)
		str, _ := reader.ReadString('\n')
		fmt.Fprintf(in, str)
	}
}

func (p *Pi) exec(cmd string) error {
	hostKey := createHostKey()

	sshConf := &ssh.ClientConfig{
		User: p.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(p.Password),
		},
		HostKeyCallback: hostKey,
	}

	connection, err := ssh.Dial("tcp", p.Hostname+":"+p.Port, sshConf)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	session, err := connection.NewSession()
	if err != nil {
		fmt.Println("Error", err)
	}

	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("hostname && ls"); err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())
	return nil
}

func (p *Pi) Exec(cmd string) error {
	return p.exec(cmd)
}

func (p *Pi) ExecAsync(cmd string, wg *sync.WaitGroup) error {
	err := p.exec(cmd)

	wg.Done()
	return err
}

func createHostKey() ssh.HostKeyCallback {
	rootDir := os.Getenv("HOME")
	knownhostsLoc := "/.ssh/known_hosts"

	hostKeyCallback, err := knownhosts.New(rootDir + knownhostsLoc)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return hostKeyCallback
}
