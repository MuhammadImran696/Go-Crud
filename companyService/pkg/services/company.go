package services

import (
	"context"
	"fmt"
	"net/http"

	"example.com/authProject/companyService/pkg/db"
	"example.com/authProject/companyService/pkg/models"
	pb "example.com/authProject/companyService/pkg/pb"
	"github.com/google/uuid"
)

type Server struct {
	H db.Handler
	pb.UnimplementedCompanyServiceServer
}

func (s *Server) CreateCompany(ctx context.Context, req *pb.CreateCompanyRequest) (*pb.CreateCompanyResponse, error) {
	var product models.Company
	product.Id = uuid.New()
	product.Name = req.Name
	product.Description = req.Description
	product.Employees = req.Employees
	product.Registered = req.Registered
	product.Type = req.Type

	if result := s.H.DB.Create(&product); result.Error != nil {
		return &pb.CreateCompanyResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.CreateCompanyResponse{
		Status: http.StatusCreated,
		Id:     product.Id.String(),
	}, nil
}

func (s *Server) FindOne(ctx context.Context, req *pb.GetOneRequest) (*pb.GetOneResponse, error) {
	var company models.Company
	fmt.Print(req.Id)
	if result := s.H.DB.First(&company, "id = ?", req.Id); result.Error != nil {
		return &pb.GetOneResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	data := &pb.GetOneData{
		Id:          company.Id.String(),
		Name:        company.Name,
		Description: company.Description,
		Employees:   company.Employees,
		Registered:  company.Registered,
		Type:        company.Type}

	return &pb.GetOneResponse{
		Status: http.StatusOK,
		Data:   data,
	}, nil
}
func (s *Server) DeleteCompany(ctx context.Context, req *pb.DeleteCompanyRequest) (*pb.DeleteCompanyResponse, error) {
	var company models.Company
	var exists bool
	s.H.DB.Model(company).
		Select("count(*) > 0").
		Where("id = ?", req.Id).
		Find(&exists)
	if !exists {
		return &pb.DeleteCompanyResponse{
			Msg: "Record not found",
		}, nil
	} else {
		err := s.H.DB.Delete(&company, "id = ?", req.Id)
		if err != nil {
			fmt.Println(err)
		}
	}
	return &pb.DeleteCompanyResponse{
		Msg: "Record Deleted",
	}, nil
}

func (s *Server) UpdateCompany(ctx context.Context, req *pb.PatchRequest) (*pb.PatchResponse, error) {
	var company models.Company
	var exists bool
	ID := req.Id
	s.H.DB.Model(company).
		Select("count(*) > 0").
		Where("id = ?", req.Id).
		Find(&exists)
	if !exists {
		return &pb.PatchResponse{
			Msg: "Record not found",
		}, nil
	} else {
		var product models.Company

		s.H.DB.First(&product, "id = ?", ID)
		if req.Name == "" {
			product.Name = product.Name
		} else {
			product.Name = req.Name
		}
		if req.Description == "" {
			product.Description = product.Description
		} else {
			product.Description = req.Description
		}
		if req.Employees == 0 {
			product.Employees = product.Employees
		} else {
			product.Employees = req.Employees
		}
		// if !req.Registered {
		// 	product.Registered = product.Registered
		// } else {
		// 	product.Registered = req.Registered
		// }
		product.Registered = req.Registered
		if req.Type == "" {
			product.Type = product.Type
		} else {
			product.Type = req.Type
		}

		err := s.H.DB.Save(&product)
		if err != nil {
			fmt.Println(err)
		}
	}
	return &pb.PatchResponse{
		Msg: "Record Updated",
	}, nil
}
