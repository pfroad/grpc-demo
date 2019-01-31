package apis

import (
	"context"

	"airparking/order-bk/pkg/models"
	pb "airparking/order-bk/pkg/pb"
)

type OrderServer struct{}

// func (os *OrderServer) Search(ctx context.Context, sr *pb.SearchRequest) (*pb.SearchResponse, error) {

// }

func (os *OrderServer) Detail(ctx context.Context, ir *pb.IDRequest) (*pb.OrderResponse, error) {
	id := ir.Id
	orderModel := &models.Order{Id: uint(id)}

	order, err := orderModel.One()
	if err != nil {
		return nil, err
	}

	return &pb.OrderResponse{
		Id:     int64(order.Id),
		UserId: int64(order.UserId),
		Mobile: order.Mobile,
	}, nil
}
