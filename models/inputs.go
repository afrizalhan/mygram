package models

type RegisterInput struct {
	Age      int    `json:"age" binding:"required" example:"8"`
	Email    string `json:"email" binding:"required" example:"maryjane@gmail.com"`
	Password string `json:"password" binding:"required" example:"maryjanewats"`
	Username string `json:"username" binding:"required" example:"maryj"`
}

type UpdateInput struct {
	Email    string `json:"email" binding:"required" example:"johndoe@gmail.com"`
	Username string `json:"username" binding:"required" example:"johndoe"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required" example:"maryjane@gmail.com"`
	Password string `json:"password" binding:"required" example:"maryjanewats"`
}

type PhotoInput struct {
	Title    string `json:"title" binding:"required" example:"Foto Selfie"`
	Caption  string `json:"caption" binding:"required" example:"Foto di pantai"`
	PhotoURL string `json:"photo_url" binding:"required" example:"https://static.vecteezy.com/packs/media/vectors/term-bg-1-666de2d9.jpg"`
}

type CommentInput struct {
	Message string `json:"message" binding:"required" example:"Fotonya bagus!!"`
	PhotoID uint   `json:"photo_id" binding:"required" example:"3"`
}

type CommentUpdateInput struct {
	Message string `json:"message" binding:"required" example:"Fotonya Keren!!"`
}

type SocialMediaInput struct {
	Name           string `json:"name" example:"Github"`
	SocialMediaUrl string `json:"social_media_url" example:"github.com/username"`
}
