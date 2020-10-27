package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"github.com/apache/thrift/lib/go/thrift"
	//"thrift/tutorial/gen-go/tutorial"
	"thriftExample/gen-go/echo"
)

func main() {
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTCompactProtocolFactory()

	transport, err := thrift.NewTSocket(net.JoinHostPort("127.0.0.1", "9898"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}

	useTransport, err := transportFactory.GetTransport(transport)
	client := echo.NewEchoClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to 127.0.0.1:9898", " ", err)
		os.Exit(1)
	}
	defer transport.Close()

	req := &echo.EchoReq{Msg:"You are welcome."}
	var defaultCtx = context.Background()
	res, err := client.Echo(defaultCtx, req)
	if err != nil {
		log.Println("Echo failed:", err)
		return
	}
	log.Println("response:", res.Msg)
	fmt.Println("well done")
	retSum, err := client.Add(defaultCtx, 3,7)
	fmt.Println( "sum is --- --- ---", retSum, "---", err)

	work := echo.NewWork()
	//work.Op = echo.Operation_MULTIPLY
	work.Op = echo.Operation_DIVIDE
	work.Num1 = 98
	work.Num2 = 0
	quotient, err := client.Calculate(defaultCtx, 1, work)
	fmt.Println("add result: ", quotient, err)

	err = client.Zip(defaultCtx)
	fmt.Println(err)
}