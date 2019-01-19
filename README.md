# zkt
## A ZooKeeper Tree Printer

`zkt` prints the tree of a ZooKeeper instance.
It does not include the values of the znodes.

```sh
$ ./zkt --server=localhost
+- /
|  +- node-I
|  |  +- subnode-A
|  |  |  +- subsubnode-1
|  |  |  +- subsubnode-2
|  |  |  +- subsubnode-3
|  |  +- subnode-B
|  |  +- subnode-C
|  +- node-II
|  +- node-III
|  +- zookeeper
|  |  +- quota
```

## Flags

### `server`
`--server <hostname>` will specify the ZooKeeper instance to attach to.

The default value is `127.0.0.1:2181`

### `p`
`-p` configures zkt to print each node's full path rather than it's name.

```sh
$ ./zkt -p
+- /
|  +- /node-I
|  |  +- /node-I/subnode-A
|  |  |  +- /node-I/subnode-A/subsubnode-1
|  |  |  +- /node-I/subnode-A/subsubnode-2
...
```

The default value is `false`

### `debug`
`-debug` configures zkt to print logs, this may provide more information about the underlying ZooKeeper client library's attempts to connect to the ZooKeeper server.

## License

mit, please see LICENSE file.
