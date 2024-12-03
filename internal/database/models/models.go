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

	AvatarId string `json:"avatarId"`	

}

type CreateSpace struct {

	Name string `json:"name"`
	Dimension  string   `json:"dimension"`
}


type AddElement struct {
	
	Element string `json:"element"`
	X int 	`json:"x"`	
	Y  int 	`json:"y"`	
}


type CreateElement struct {

	ImageUrl string `json:"imageUrl"`
	Width int `json:"width"`
	Height int `json:"height"`
	Status bool`json:"status"`

}

type  UpdateElement struct {
	ImageUrl string `json:"imageUrl"`
	
}

type CreateAvatar struct {
	ImageUrl string `json:"imageUrl"`
	Name string `json:"Name"`
}

