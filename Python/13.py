class Matematica:
    @staticmethod
    def fatorial(n):
        if n == 0 or n == 1:
            return 1
        elif n < 0:
            raise ValueError("Apenas Números Positivos")
        else:
            resultado = 1
            for i in range(2, n + 1):
                resultado *= i
            return resultado

print(Matematica.fatorial(8))  
print(Matematica.fatorial(22))  
