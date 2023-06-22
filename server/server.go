package server

import (
	"context"
	"github.com/offluck/abgorrence/common/models/endpoint"
	"github.com/offluck/abgorrence/common/models/relations"
	"net/http"
)

type Server struct {
	http.Server
	relations relations.Relations
}

func New(addr string) *Server {
	s := &Server{
		relations: relations.New(),
	}
	s.Addr = addr
	return s
}

func (s *Server) AddHandler(endpoint endpoint.Endpoint, handler func(context.Context, Request) (Response, error)) {
	http.Handle(
		endpoint.URL,
		Init[Request, Response](handler, s.relations.GetRelationsFor(endpoint)),
	)
}

func (s *Server) AddRelation(from endpoint.Endpoint, to endpoint.Endpoint) {
	s.relations.Add(from, to)
}

func (s *Server) DeleteRelation(from endpoint.Endpoint, to endpoint.Endpoint) {
	s.relations.Delete(from, to)
}

func (s *Server) IsRelationPresent(from endpoint.Endpoint, to endpoint.Endpoint) bool {
	return s.relations.IsRelationPresent(from, to)
}
