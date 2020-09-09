package server

import (
	"context"
	"encoding/json"
	"github.com/learn/go-programming-tour-book/tag-service/pkg/bapi"
	"github.com/learn/go-programming-tour-book/tag-service/proto"
)

type TagServer struct {
}

func NewTagServer() *TagServer {
	return &TagServer{}
}

func (t *TagServer) GetTagList(ctx context.Context, r *proto.GetTagListRequest) (*proto.GetTagListReply, error) {
	api := bapi.NewAPI("http://localhost:8000")
	body, err := api.GetTagList(ctx, r.GetName())
	if err != nil {
		return nil, nil
	}
	reply := proto.GetTagListReply{}
	err = json.Unmarshal(body, &reply)
	if err != nil {
		return nil, nil
	}
	return &reply, nil
}
