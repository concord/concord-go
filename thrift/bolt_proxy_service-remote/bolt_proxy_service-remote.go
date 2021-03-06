// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"bolt"
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  void updateTopology(TopologyMetadata topology)")
	fmt.Fprintln(os.Stderr, "  void updateSchedulerAddress(Endpoint e)")
	fmt.Fprintln(os.Stderr, "  void registerWithScheduler(ComputationMetadata meta)")
	fmt.Fprintln(os.Stderr, "  void setState(string key, string value)")
	fmt.Fprintln(os.Stderr, "  string getState(string key)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := bolt.NewBoltProxyServiceClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "updateTopology":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "UpdateTopology requires 1 args")
			flag.Usage()
		}
		arg58 := flag.Arg(1)
		mbTrans59 := thrift.NewTMemoryBufferLen(len(arg58))
		defer mbTrans59.Close()
		_, err60 := mbTrans59.WriteString(arg58)
		if err60 != nil {
			Usage()
			return
		}
		factory61 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt62 := factory61.GetProtocol(mbTrans59)
		argvalue0 := bolt.NewTopologyMetadata()
		err63 := argvalue0.Read(jsProt62)
		if err63 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.UpdateTopology(value0))
		fmt.Print("\n")
		break
	case "updateSchedulerAddress":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "UpdateSchedulerAddress requires 1 args")
			flag.Usage()
		}
		arg64 := flag.Arg(1)
		mbTrans65 := thrift.NewTMemoryBufferLen(len(arg64))
		defer mbTrans65.Close()
		_, err66 := mbTrans65.WriteString(arg64)
		if err66 != nil {
			Usage()
			return
		}
		factory67 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt68 := factory67.GetProtocol(mbTrans65)
		argvalue0 := bolt.NewEndpoint()
		err69 := argvalue0.Read(jsProt68)
		if err69 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.UpdateSchedulerAddress(value0))
		fmt.Print("\n")
		break
	case "registerWithScheduler":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "RegisterWithScheduler requires 1 args")
			flag.Usage()
		}
		arg70 := flag.Arg(1)
		mbTrans71 := thrift.NewTMemoryBufferLen(len(arg70))
		defer mbTrans71.Close()
		_, err72 := mbTrans71.WriteString(arg70)
		if err72 != nil {
			Usage()
			return
		}
		factory73 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt74 := factory73.GetProtocol(mbTrans71)
		argvalue0 := bolt.NewComputationMetadata()
		err75 := argvalue0.Read(jsProt74)
		if err75 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.RegisterWithScheduler(value0))
		fmt.Print("\n")
		break
	case "setState":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "SetState requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := []byte(flag.Arg(2))
		value1 := argvalue1
		fmt.Print(client.SetState(value0, value1))
		fmt.Print("\n")
		break
	case "getState":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetState requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetState(value0))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
