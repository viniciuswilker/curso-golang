package enderecos

import "testing"

func TestTipoDeEndereco(t *testing.T) {
	// Test + NomeDaFuncao  recebe t *testing.T

	enderecoParaTeste := "Avenida Paulista"

	tipoDeEnderecoEsperado := "Avenida"

	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
		t.Error("O tipo recebido é diferente do esperado")
	}

}
