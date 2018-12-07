package rota

type Rotator interface {
	Today() string
	Next() string
	Reset()
}

type rotator struct {

}

func (r *rotator) Today() string {

}

func (r *rotator) Next() string {

}

func (r *rotator) Reset() {

}

func NewRotator() Rotator {
	return &rotator{}
}