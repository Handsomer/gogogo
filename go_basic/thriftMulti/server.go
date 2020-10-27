package main

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"strconv"
	"thriftExample/gen-go/echo"
	"thriftExample/gen-go/shared"
)


type EchoServer struct {
	log map[int]*shared.SharedStruct
}

func (e *EchoServer) Zip(ctx context.Context)(err error){
	fmt.Println("run zip function")
	err = fmt.Errorf("%v", "zip no error return oooo")
	fmt.Println(err)
	return err
}

func (p *EchoServer) Calculate(ctx context.Context, logid int32, w *echo.Work) (val int32, err error) {
	fmt.Print("calculate(", logid, ", {", w.Op, ",", w.Num1, ",", w.Num2, "})\n")
	switch w.Op {
	case echo.Operation_ADD:
		val = w.Num1 + w.Num2
		break
	case echo.Operation_SUBTRACT:
		val = w.Num1 - w.Num2
		break
	case echo.Operation_MULTIPLY:
		val = w.Num1 * w.Num2
		break
	case echo.Operation_DIVIDE:
		if w.Num2 == 0 {
			ouch := echo.NewInvalidOperation()
			ouch.WhatOp = int32(w.Op)
			ouch.Why = "Cannot divide by 0"
			err = ouch
			return
		}
		val = w.Num1 / w.Num2
		break
	default:
		ouch := echo.NewInvalidOperation()
		ouch.WhatOp = int32(w.Op)
		ouch.Why = "Unknown operation"
		err = ouch
		return
	}
	entry := shared.NewSharedStruct()
	entry.Key = logid
	entry.Value = strconv.Itoa(int(val))
	k := int(logid)
	/*
	  oldvalue, exists := p.log[k]
	  if exists {
	    fmt.Print("Replacing ", oldvalue, " with ", entry, " for key ", k, "\n")
	  } else {
	    fmt.Print("Adding ", entry, " for key ", k, "\n")
	  }
	*/
	p.log[k] = entry
	return val, err
}

func (e *EchoServer) Add(ctx context.Context, num1 int32, num2 int32) (r int32, err error) {
	return num1 + num2, nil
}

func (e *EchoServer) Echo(ctx context.Context,req *echo.EchoReq) (*echo.EchoRes, error) {
	fmt.Printf("message from client: %v\n", req.GetMsg())
	res := &echo.EchoRes{
		Msg: "success",
	}
	return res, nil
}

func main() {
	transport, err := thrift.NewTServerSocket(":9898")
	if err != nil {
		panic(err)
	}

	handler := &EchoServer{}
	processor := echo.NewEchoProcessor(handler)

	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTCompactProtocolFactory()
	server := thrift.NewTSimpleServer4(
		processor,
		transport,
		transportFactory,
		protocolFactory,
	)
	if err := server.Serve(); err != nil {
		panic(err)
	}
}