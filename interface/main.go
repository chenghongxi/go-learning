package main

import (
	v1 "go-learning/interface/namespace"
	"go-learning/interface/scaler"
)

// moremorehttps://github.com/kubernetes/client-go/blob/4841142cdc4ba44627994f8c25da76c2be2cdb5d/kubernetes/typed/apps/v1/apps_client.go#L27

type AppsInterface interface {
	v1.PixiuGetter
	scaler.ScalerGetter
}

type apps struct {
	svc string
}

func (p *apps) Pixiu(namespace string) v1.PixiuInterface {
	return v1.NewPixiu(p.svc, namespace)
}

func (p *apps) Scaler(namespace string) scaler.ScalerInterface {
	return scaler.NewScaler(p.svc, namespace)
}

func New(svc string) AppsInterface {
	return &apps{svc}
}

func main() {
	p := New("service")

	if err := p.Pixiu("pixiu-ns").Create("pixiu"); err != nil {
		panic(err)
	}

	if err := p.Scaler("scaler-ns").Create("scaler"); err != nil {
		panic(err)
	}
}
