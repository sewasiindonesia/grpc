	package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "github.com/sewasiindonesia/grpc/slack/proto"
)

type GrpcClient interface {
	Close()
}

type Client struct {
	slackBotClient pb.SlackClient
	conn            *grpc.ClientConn
}

type Options struct {
	Address         string
	WithInterceptor bool // set true to generate log
}

func GetClient(o *Options) (*Client, error) {
	var (
		conn *grpc.ClientConn
		err  error
	)

	//remove round robin balancer for the moment
	//r, _ := naming.NewDNSResolverWithFreq(time.Second * 1)
	//balancer := grpc.RoundRobin(r)

	if o.WithInterceptor == true {
		conn, err = grpc.Dial(o.Address, grpc.WithInsecure(), WithClientInterceptor())
	} else {
		conn, err = grpc.Dial(o.Address, grpc.WithInsecure())
	}

	if err != nil {
		log.Printf("grpc error - connection to Eventapp GRPC failed %v", err)
		return nil, err
	}

	client := &Client{
		conn: conn,
	}

	client.crossSellClient = pb.NewCrossSellClient(conn)

	return client, nil
}

func WithClientInterceptor() grpc.DialOption {
	return grpc.WithUnaryInterceptor(clientInterceptor)
}

func clientInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...) // <==
	log.Printf("invoke remote method=%s duration=%s error=%v", method,
		time.Since(start), err)
	return err
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) SendSlackBot(ctx context.Context, req *pb.SendSlackBotRequest) (*pb.SendSlackBotResponse, error) {
		return c.slackBotClient.SendSlackMessage(ctx, req)
	}
