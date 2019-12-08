package _1_oop

import "testing"

type Progarmmer interface {
	WriteHelloWorld() string
}

type GoPorgrammer struct {

}

func (g *GoPorgrammer)WriteHelloWorld() string {
	return "fmt.println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	var p Progarmmer
	p = new(GoPorgrammer)
	t.Log(p.WriteHelloWorld())
}