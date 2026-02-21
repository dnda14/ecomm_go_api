package productsi

import{
	"net/http"

}

type handler struct{
	service Service

}

func NewHandler(service Service) *handler{
	return &handler{
		service:service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request){
	
	if err:= h.service.ListProducts(r.Context());err!=nil{
		slog.Error("server failed on mounting. err : %s","error",err)
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	products:= struct{
		Products []string `json:"products"`
	}
	json.Write(w,http.StatusOk,products)
}