class Animal:
    def __init__(self, nome):
        self.nome = nome
    def som(self):
        raise NotImplementedError("Metodo som não implementado")

class Cachorro(Animal):
    def som(self):
        return f"{self.nome}: Au au"

class Gato(Animal):
    def som(self):
        return f"{self.nome}: Miau"


cachorro = Cachorro("Rex")
gato = Gato("Malenia")
print(cachorro.som())  
print(gato.som())      