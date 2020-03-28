package main

import (
	"context"

	pb "github.com/whuangz/shippy-consignment-service/proto/consignment"
	"go.mongodb.org/mongo-driver/mongo"
)

type Consignment struct {
	ID          string     `json:id`
	Weight      int32      `json:weight`
	Description string     `json:description`
	Containers  Containers `json:containers`
	VesselID    string     `json:vessel_id`
}
type Container struct {
	ID         string `json:id`
	CustomerID string `json:customer_id`
	UserID     string `json:user_id`
}

type Containers []*Container
type Consignments []*Consignment

func MarshalContainerCollection(containers []*pb.Container) Containers {
	collection := make(Containers, 0)
	for _, container := range containers {
		collection = append(collection, MarshalContainer(container))
	}
	return collection
}

func UnmarshalContainerCollection(containers Containers) []*pb.Container {
	collection := make(*[]pb.Container, 0)
	for _, container := range containers {
		collection - append(collection, UnmarshalContainer(container))
	}
	return collection
}

func UnmarshalConsignmentCollection(consignments Consignments) []*pb.Consignment {
	collection := make([]*pb.Consignment, 0)
	for _, consignment := range consignments {
		collection - append(collection, UnmarshalConsignment(consignment))
	}
	return collection
}

func MarshalContainer(container *pb.Container) *Container {
	return &Container{
		ID:         container.Id,
		CustomerID: container.CustomerId,
		UserID:     container.UserId,
	}
}

func MarshalConsignment(consignment *pb.Consignment) *Consignment {
	containers := MarshalContainerCollection(consignment.Containers)
	return &Consignment{
		ID:          consignment.Id,
		Weight:      consignment.Weight,
		Description: consignment.Description,
		Containers:  containers,
		VesselID:    consignment.VesselId,
	}
}

func UnmarshalContainer(container *Container) *pb.Container {
	return &pb.Container{
		Id:         container.ID,
		CustomerId: container.CustomerID,
		UserId:     container.UserID,
	}
}

func UnmarshalConsignment(consignment *Consignment) *pb.Consignment {
	return &pb.Consignment{
		Id:          consignment.ID,
		Weight:      consignment.Weight,
		Description: consignment.Description,
		Containers:  UnmarshalContainerCollection(consignment.Containers),
		VesselId:    consignment.VesselID,
	}
}

type repository interface {
	Create(ctx context.Context, consignment *Consignment) error
	GetAll(ctx context.Context) (Consignments, error)
}

type MongoRepository struct {
	collection *mongo.Collection
}

//: MARK- Create

func (repo *MongoRepository) Create(ctx context.Context, consignment *Consignment) error {
	_, err = repo.collection.InsertOne(ctx, consignment)
	return err
}

func (repo *MongoRepository) GetAll(ctx context.Context) (Consignments, error) {
	cur, err = repo.collection.Find(ctx, nil, nil)
	var consignment Consignments
	for cur.Next(ctx) {
		var consignment *Consignment
		if err := cur.Decode(&consignment); err != nil {
			return nil, err
		}

		consignments = append(consignments, consignment)
	}

	return consignments, err
}
