class SaldoInsuficienteException(Exception):
    def __init__(self, saldo, valor_saque):
        super().__init__(f"Saque de {valor_saque} não liberado. Saldo disponível: {saldo}")

class ContaBancaria:
    def __init__(self, saldo):
        self.saldo = saldo

    def sacar(self, valor):
        if valor > self.saldo:
            raise SaldoInsuficienteException(self.saldo, valor)
        self.saldo -= valor
        print(f"Saque realizado com sucesso. Saldo atual: {self.saldo}")


conta = ContaBancaria(2500)

try:
    conta.sacar(3500)
except SaldoInsuficienteException as e:
    print(e)  