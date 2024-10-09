from abc import ABC, abstractmethod

class Funcionario(ABC):
    @abstractmethod
    def calcularSalario(self):
        pass

class FuncionarioHorista(Funcionario):
    def __init__(self, horas_trabalhadas, valor_por_hora):
        self.horas_trabalhadas = horas_trabalhadas
        self.valor_por_hora = valor_por_hora

    def calcularSalario(self):
        return self.horas_trabalhadas * self.valor_por_hora


class FuncionarioAssalariado(Funcionario):
    def __init__(self, salario_mensal):
        self.salario_mensal = salario_mensal

    def calcularSalario(self):
        return self.salario_mensal


horista = FuncionarioHorista(horas_trabalhadas=200, valor_por_hora=30)
assalariado = FuncionarioAssalariado(salario_mensal=3500)

print(f"Funcionário Horista: R${horista.calcularSalario():.2f}")
print(f"Funcionário Assalariado: R${assalariado.calcularSalario():.2f}")
