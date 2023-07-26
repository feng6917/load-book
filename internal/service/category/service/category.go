package service

import (
	"context"
	"errors"
	"load-book/internal/model/category"
	"load-book/internal/utils"

	v1 "load-book/api/load-book/v1"
	"load-book/internal/service/category/biz"
)

// Service is a health service.
type Service struct {
	v1.UnimplementedCategoryServer

	uc *biz.UseCase
}

// NewService new a health service.
func NewService(uc *biz.UseCase) v1.CategoryServer {
	return &Service{uc: uc}
}

func (s *Service) Get(ctx context.Context, req *v1.CategoryGetRequest) (*v1.CategoryGetReply, error) {
	res, err := s.uc.FindByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	item := v1.Category{}
	err = utils.Unmarshal2Message(&item, res)
	if err != nil {
		return nil, err
	}
	return &v1.CategoryGetReply{
		Item: &item,
	}, nil
}

func (s *Service) List(ctx context.Context, req *v1.CategoryListRequest) (*v1.CategoryListReply, error) {
	var pageLimit, pageOffset int
	if req.Page != nil {
		pageLimit = int(req.Page.Limit)
		pageOffset = int(req.Page.Offset)
	}
	res, err := s.uc.List(ctx, &category.Form{
		Limit:  pageLimit,
		Offset: pageOffset,
		Name:   req.Name,
	})
	if err != nil {
		return nil, err
	}
	items := make([]*v1.Category, len(res))
	if len(res) > 0 {
		for k, v := range res {
			temp := v1.Category{}
			err = utils.Unmarshal2Message(&temp, v)
			if err != nil {
				return nil, err
			}
			items[k] = &temp
		}
	}
	return &v1.CategoryListReply{Items: items}, nil
}

func (s *Service) Update(ctx context.Context, req *v1.CategoryUpdateRequest) (*v1.CategoryUpdateReply, error) {
	if req.Id < 1 {
		return nil, errors.New("request param fail")
	}
	err := s.uc.Update(ctx, &category.Form{
		ID:   req.Id,
		Name: req.Name,
		Desc: req.Desc,
	})
	return nil, err
}

func (s *Service) Delete(ctx context.Context, req *v1.CategoryDeleteRequest) (*v1.CategoryDeleteReply, error) {
	if req.Id < 1 {
		return nil, errors.New("request param fail")
	}
	err := s.uc.DeleteByID(ctx, req.Id)
	return nil, err
}

func (s *Service) Create(ctx context.Context, req *v1.CategoryCreateRequest) (*v1.CategoryCreateReply, error) {
	id, err := s.uc.Create(ctx, category.Form{
		Name: req.Name,
		Desc: req.Desc,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CategoryCreateReply{Id: id}, nil
}
