class Configuracao:
    _instancia = None 

    @classmethod
    def get_instance(cls):
        if cls._instancia is None:
            cls._instancia = cls() 
        return cls._instancia

    def __init__(self):
        if not hasattr(self, 'inicializado'):
            self.opcoes = {}
            self.inicializado = True 

    def definir_opcao(self, chave, valor):
        self.opcoes[chave] = valor

    def obter_opcao(self, chave):
        return self.opcoes.get(chave)

# Exemplo de uso
config1 = Configuracao.get_instance()
config2 = Configuracao.get_instance()

config1.definir_opcao('Linguagem', 'Java')

print(config2.obter_opcao('Linguagem'))  
print(config1 is config2)               
