package smsapi

type Collection interface {
	GetSize() uint
}

type CollectionMeta struct {
	Size uint `json:"Size"`
}

func (c *CollectionMeta) GetSize() uint {
	return c.Size
}
