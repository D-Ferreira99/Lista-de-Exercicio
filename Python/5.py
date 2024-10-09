class Animal:
    def __init__(self, nome):
        self.nome = nome
    
    def som(self):
         raise NotImplementedError("Metodo som não implementado")

class Cachorro(Animal):
    def som(self):
        return f"{self.nome} Au au"

class Gato(Animal):
    def som(self):
        return f"{self.nome} Miau"


def emitir_sons(animais):
    for animal in animais:
        print(animal.som())


cachorro = Cachorro("Rex")
gato = Gato("Malenia")
gatoB = Gato("Madri")

animais = [cachorro, gato, gatoB]
emitir_sons(animais)