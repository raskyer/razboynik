package kernel

import (
	"fmt"
	"io"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"strconv"

	"github.com/eatbytes/razboy"
)

func LaunchRPC(r *RPCServer) {
	server := rpc.NewServer()

	server.Register(r)
	server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)

	l, e := net.Listen("tcp", ":"+strconv.Itoa(r.Port))

	if e != nil {
		fmt.Fprintln(os.Stderr, e)
	}

	for {
		conn, err := l.Accept()

		if err != nil {
			fmt.Fprintln(os.Stderr, e)
		}

		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

type RPCArgs struct {
	Addr   string
	Port   int
	Line   string
	Config razboy.Config
}

type RPCReply struct {
	Status bool
}

type RPCClient struct {
	Addr string
}

type RPCServer struct {
	Port    int
	Clients []RPCClient
}

func (rs RPCServer) Handshake(args RPCArgs, reply *RPCReply) error {
	fmt.Println(args)
	rs.Clients = append(rs.Clients, RPCClient{args.Addr})
	reply.Status = true

	return nil
}

func (rs RPCServer) Send(args RPCArgs, reply *RPCReply) error {
	r := razboy.CreateRequest(args.Line, &args.Config)
	_, err := razboy.Send(r)

	reply.Status = true

	return err
}

func (rc RPCClient) Call(m string, args RPCArgs) error {
	var (
		conn   io.ReadWriteCloser
		client *rpc.Client
		reply  RPCReply
		err    error
	)

	conn, err = net.Dial("tcp", rc.Addr)

	if err != nil {
		fmt.Println(err)
		return err
	}

	defer conn.Close()

	client = jsonrpc.NewClient(conn)
	err = client.Call(m, args, &reply)

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(reply)

	return nil
}
