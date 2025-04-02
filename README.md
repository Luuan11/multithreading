# Multithreading

![go](https://github.com/user-attachments/assets/90f5ff8e-884b-4f08-974b-8a2fc2e47984)

### Desafio Golang Pós Go-Expert
  Desafio na qual consiste em utilizar conhecimentos com Multithreading e APIs para obter o resultado mais rápido entre duas APIs de CEP.

### 💬 Sobre: 
  APIs realizando requisições simultaneamente para as APIs:
 ```plaintext
A - https://brasilapi.com.br/api/cep/v1/01153000 + cep
B - http://viacep.com.br/ws/ + cep + /json
```
  a API que entregar a resposta mais rápida é exibida

### ✨ Executando:
  Comece rodando o main.go

```shell
❯ go run main.go 
```
  Após isso, iniciará a porta :8080 no localhost, só acessar e informar o valor após a "/", Exemplo:
  
```plaintext
http://localhost:8080/08080280
```

## Instalação

    - Clonar repositorio 
    $ git clone https://github.com/Luuan11/multithreading.git 

    - Instalar dependencias
    $ go mod tidy

    - Rodar projeto
    $ go run main.go
---

Made with 💜 by [Luan Fernando](https://www.linkedin.com/in/luan-fernando/).
