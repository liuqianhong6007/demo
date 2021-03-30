package main

import (
	"context"
	"github.com/liuqianhong6007/demo/k8s/com"
	"github.com/liuqianhong6007/demo/k8s/watcher/protocol"
	"google.golang.org/grpc"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"log"
	"net"
)

func newServer(addr string) *Server {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	clientset := com.NewKubernetesClientset(true, "")
	grpcServer := grpc.NewServer()
	server := &Server{
		addr:       addr,
		lis:        lis,
		grpcServer: grpcServer,
		clientset:  clientset,
	}
	protocol.RegisterWatchServiceServer(grpcServer, server)

	return server
}

type Server struct {
	protocol.UnimplementedWatchServiceServer
	addr       string
	lis        net.Listener
	grpcServer *grpc.Server
	clientset  *kubernetes.Clientset
}

func (s *Server) Watch(in *protocol.MatchPodCondition, stream protocol.WatchService_WatchServer) error {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic: ", err)
		}
	}()

	namespace := in.GetNamespace()
	label := in.GetLabelSelector()
	if namespace == "" || label == "" {
		stream.Send(&protocol.MatchPodResponse{
			Status: protocol.MatchPodResponse_ParamError,
		})
	}

	watcher, err := s.clientset.CoreV1().Pods(namespace).Watch(context.Background(), metav1.ListOptions{
		LabelSelector: label,
	})
	if err != nil {
		stream.Send(&protocol.MatchPodResponse{
			Status: protocol.MatchPodResponse_UnknownError,
		})
	}

	for {
		select {
		case event := <-watcher.ResultChan():
			switch event.Type {
			case watch.Added:
				stream.Send(&protocol.MatchPodResponse{
					Status: protocol.MatchPodResponse_Ok,
				})
			case watch.Deleted:
			}

		}
	}

	return nil
}

func (s *Server) Listen() {
	defer func() {
		log.Println("game service grpc server stop")
		s.grpcServer.Stop()
		s.lis.Close()
	}()
	log.Println("game service grpc server listen at: ", s.addr)
	if err := s.grpcServer.Serve(s.lis); err != nil {
		panic(err)
	}
}
