package models

type UserSignUpBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role,omitempty"` // Optional role field
}

type UserSignInBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}	

type MetaDataUpdate struct{

	AvatarId string `json:"avatarId" validate:"required"`	

}

type CreateSpace struct {
	Name string `json:"name" validate:"required"`
	Width int `json:"width" validate:"required"`
	Height int `json:"height" validate:"required"`
    	MapId string `json:"mapId"`
}


type AddElement struct {
	
	ElementId string `json:"element" validate:"required"`
	X int 	`json:"x" validate:"required"`	
	Y  int 	`json:"y" validate:"required"`	
	SpaceId string `json:"spaceId" validate:"required"`
}


type CreateElement struct {

	ImageUrl string `json:"imageUrl"`
	Width int `json:"width" validate:"required"`
	Height int `json:"height" validate:"required"`
	// not status its static
	Status bool`json:"status" validate:"required"`

}

type  UpdateElement struct {
	ImageUrl string `json:"imageUrl"`
	
}

type CreateAvatar struct {
	ImageUrl string `json:"imageUrl" validate:"required"`
	Name string `json:"Name" validate:"required"` 
}
type DeleteElement struct {
	ElementId string `json:"elementId" validate:"required"`
	SpaceId string `json:"spaceId" validate:"required"`
}

type CreateMap struct {
	Thumbnail   	 	string          	`json:"thumbnail" validate:"required"`
	Height       		int			`json:"height" validate:"required"`
	Width       		int			`json:"width" validate:"required"`
	Name         		string          	`json:"name" validate:"required"`
	DefaultElements 		[]MapElement  	`json:"defaultElements" validate:"dive"`
}
type MapElement struct {
	ElementID    string  `json:"elementId" validate:"required"`
	X            float64 `json:"x" validate:"required"`
	Y            float64 `json:"y" validate:"required"`
}  