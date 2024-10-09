# Classe Motor
class Motor:
    def __init__(self, potencia):
        self.potencia = potencia  

class Carro:
    def __init__(self, modelo, ano, potencia_motor):
        self.modelo = modelo
        self.ano = ano
        self.motor = Motor(potencia_motor)  

    def exibir(self):
        print(f"Modelo: {self.modelo}, Ano: {self.ano}, Potência do Motor: {self.motor.potencia}")
    

# Exemplo de uso
meu_carro = Carro("Strada", 2015, 86)  
meu_carro.exibir()