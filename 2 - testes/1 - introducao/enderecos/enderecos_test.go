package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	enderecoEsperado string
}

func TestTipoDeEndereco(t *testing.T) {
	// Test + NomeDaFuncao  recebe t *testing.T

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia ABC", "Rodovia"},
		{"Praça das rosas", "Tipo Inválido"},
		{"Estrada Qualquer", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"Praça das rosas", "Tipo Inválido"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		{" ", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.enderecoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s", tipoDeEnderecoRecebido, cenario.enderecoEsperado)

		}
	}

}
