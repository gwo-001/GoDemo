package model

type PostTestModel struct {
	Name string `binding:"required" json:"name"`
	Age int8 `json:"age"`
}
