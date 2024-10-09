class Produto:
    def __init__(self, nome, preco):
        self.nome = nome
        self.preco = preco

    def __add__(self, other):
        if isinstance(other, Produto):
            return self.preco + other.preco
        return NotImplemented

    def __repr__(self):
        return f"Produto({self.nome}, {self.preco})"


produto1 = Produto("Açucar", 5.39)
produto2 = Produto("Arroz", 37.99)
preco_total = produto1 + produto2
print(f"Preço total: {preco_total}")  
