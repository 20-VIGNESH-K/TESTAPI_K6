package models

type RequestData struct {
	Message string `json:"message" binding:"required"`
}

// type ResponseData struct {
//     Reply string `json:"reply"`
// }