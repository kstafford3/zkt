package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

// Node holds relevant information for a znode in the zookeeper tree.
// Includes path, name, and depth of the node.
type node struct {
	path  string
	name  string
	depth int
}

// Config holds configuration data passed in by CLI arguments and flags.
type config struct {
	printPath bool
	server    string
	debug     bool
	depth     int
}

func main() {
	serverFlagPtr := flag.String("server", "127.0.0.1:2181", "The address of the zookeeper server to connect to")
	printPathParamPtr := flag.Bool("p", false, "Print each znode's path instead of its name")
	depthFlagPtr := flag.Int("depth", -1, "The maximum depth the tree will include.")
	debugFlagPtr := flag.Bool("debug", false, "Turn on logging")

	flag.Parse()

	config := config{
		printPath: *printPathParamPtr,
		server:    *serverFlagPtr,
		debug:     *debugFlagPtr,
		depth:     *depthFlagPtr,
	}

	if !config.debug {
		log.SetOutput(ioutil.Discard)
	}

	connection, _, err := zk.Connect([]string{config.server}, time.Second*10)
	if err != nil {
		panic(err)
	}

	ch := make(chan node)
	go printTree(ch, config)
	root := node{"/", "/", 0}
	walk(connection, ch, root, config)
	close(ch)
}

// Executes a depth-first traversal through the zookeeper tree starting with the provided parent node.
// Pushes each node onto the channel (excluding the parent node)
func walk(connection *zk.Conn, ch chan node, target node, config config) {
	ch <- target
	// only limit depth if the configured depth is >= 0
	if config.depth < 0 || target.depth < config.depth {
		children, _, err := connection.Children(target.path)
		if err != nil {
			panic(err)
		}

		for _, childName := range children {
			childPath := target.path
			// hacky way to tell if this is root
			if len(target.path) > 1 {
				childPath += "/"
			}
			childPath += childName

			child := node{
				path:  childPath,
				name:  childName,
				depth: (target.depth + 1),
			}
			walk(connection, ch, child, config)
		}
	}
}

// Prints the nodes in a channel in a tree structure.
// Assumes that the channel will discharge nodes in depth-first order
func printTree(ch chan node, config config) {
	for node := range ch {
		line := ""
		for level := 0; level < node.depth; level++ {
			line += "|  "
		}
		line += "+- "
		if config.printPath {
			line += node.path
		} else {
			line += node.name
		}
		fmt.Println(line)
	}
}
