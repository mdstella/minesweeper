package model

type SkeletonRequest struct {
	Parameter string `json:"param,omitempty"`
}

type SkeletonResponse struct {
	Message string `json:"message,omitempty"`
	Err error `json:"error,omitempty"`
}

// Implementing error method
func (r SkeletonResponse) Error() error { return r.Err }
