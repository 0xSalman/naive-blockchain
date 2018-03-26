package grpc

import (
  "context"
  "fmt"
  "net"

  "github.com/salmana1/naive-blockchain/core"
  "google.golang.org/grpc"
)

type Server struct {
  blockchain *core.Chain
}

func Start(host string, port int) error {

  srv := grpc.NewServer()

  lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
  if err != nil {
    return err
  }

  newChain, err := core.NewChain()
  if err != nil {
    return err
  }
  RegisterBlockchainServer(srv, &Server{
    blockchain: newChain,
  })

  fmt.Printf("server listening on %s:%d\n", host, port)

  return srv.Serve(lis)
}

func (s *Server) AddBlock(ctx context.Context, req *AddBlockRequest) (*AddBlockResponse, error) {

  _, err := s.blockchain.AddBlock(req.Data)
  if err != nil {
    return nil, err
  }

  return new(AddBlockResponse), nil
}

func (s *Server) GetBlockchain(ctx context.Context, req *GetBlockchainRequest) (*GetBlockchainResponse, error) {

  resp := new(GetBlockchainResponse)

  for _, b := range *s.blockchain {
    resp.Blocks = append(resp.Blocks, &Block{
      Index:     b.Index,
      Timestamp: b.Timestamp.Unix(),
      Hash:      b.Hash,
      PrevHash:  b.PrevHash,
      Data:      b.Data,
    })
  }

  return resp, nil
}
