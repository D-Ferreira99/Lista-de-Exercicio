class Carro:
    def __init__(self, marca,modelo,ano):
        self.marca = marca
        self.modelo = modelo
        self.ano = ano

    def exibir(self):
        print(
            self.marca,
            self.modelo,
            self.ano
        )


carro1 = Carro('Lamborghini', 'Temerario', 2024)
carro2 = Carro('Fiat', 'Strada', 2020)

carro1.exibir()
carro2.exibir()