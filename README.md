# boleto
[![Coverage Status](https://img.shields.io/badge/coverage-55%25-brightgreen.svg)]()

Gerador de boletos para diversos bancos usando Go.

[![golang.sh-600x600.jpg](https://s27.postimg.org/coqxnki9f/golang_sh_600x600.jpg)](https://postimg.org/image/yb5y4lgtr/)

## Features
* Bancos suportados: Banco do Brasil, Bradesco, Caixa, Itau e Santander;
* Gera linha digitavel e imagem com código de barras;
* Apenas boletos registrados, conforme novas regras FEBRABAN;

## Usage

[Veja a documentação no GoDoc](https://godoc.org/github.com/kezzbr/boleto)

## TODO
* Criar layout dos boletos;
* Gerar remessas dos boletos registrados;

## Dependencies
* [boombuler/barcode](github.com/boombuler/barcode)
