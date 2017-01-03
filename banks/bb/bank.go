package bb

const id = 001
const aceite = "N"
const currency = 9
const currencyName = "R$"

type Bank struct {
	Agency 			int
	Account			int
	Convenio		int
	Contrato		int
	Carteira		int
	VariacaoCarteira	int
	FormatacaoConvenio	int
	FormatacaoNossoNumero	int
}