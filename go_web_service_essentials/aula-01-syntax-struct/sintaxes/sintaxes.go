package sintaxes

type Example struct {
	flag    bool
	counter int16
	pi      float32
}

func NovoZeroValue() Example {
	var e Example
	return e
}

func NovoComValores() Example {
	return Example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}
}

func NovoComCampoParcial() Example {
	return Example{
		flag: true,
	}
}

func NovoReunindoDadosAntes() Example {
	flagValue := true
	return Example{
		flag: flagValue,
	}
}
