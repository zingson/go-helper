package v3

type MktService struct {
	*Client
}

func (c *Client) Mkt() *MktService {
	return &MktService{Client: c}
}
