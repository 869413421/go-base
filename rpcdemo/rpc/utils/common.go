package utils

type MathArgs struct {
	A, B int
}

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Response struct {
	OK bool `json:"ok"`
	ID int  `json:"id"`
	Message string `json:"message"`
}
