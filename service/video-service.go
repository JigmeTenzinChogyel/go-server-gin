package service

import "go-server/entity"

// videoService struct definition
type videoService struct {
	videos []entity.Video
}

// VideoService interface definition
type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

// Save method implementation
func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

// FindAll method implementation
func (service *videoService) FindAll() []entity.Video {
	return service.videos
}

// New function to create a new videoService instance
func New() VideoService {
	return &videoService{}
}

type Person struct {
	name string
}

// type Cook interface {
// 	bake() string
// }
