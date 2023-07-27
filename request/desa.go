package request

type (
	DesaRequest struct {
		TentangDesa string `json:"tentang_desa" binding:"required"`
		Visi        string `json:"visi" binding:"required"`
		Misi        string `json:"misi" binding:"required"`
	}
)