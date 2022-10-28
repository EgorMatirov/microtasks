package formatter

type GetHello struct {
	Hello string
}

func NewSayHelloResponse(s string) GetHello {
	return GetHello{
		Hello: "hello, " + s,
	}
}
