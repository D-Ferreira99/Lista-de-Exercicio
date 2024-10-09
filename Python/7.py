class Professor:
    def __init__(self, nome):
        self.nome = nome
        self.escolas = []  
    def adicionar_escola(self, escola):
        if escola not in self.escolas:
            self.escolas.append(escola)
            escola.adicionar_professor(self) 
            print(f"O professor {self.nome} já leciona na escola {escola.nome}.")
       
    
    def listar_escolas(self):
        if self.escolas:
            print(f"O professor {self.nome} leciona nas seguintes escolas:")
            for escola in self.escolas:
                print(f"- {escola.nome}")
        else:
            print(f"O professor {self.nome} não leciona em nenhuma escola.")


class Escola:
    def __init__(self, nome):
        self.nome = nome
        self.professores = []  
    
    def adicionar_professor(self, professor):
        if professor not in self.professores:
            self.professores.append(professor)
        else:
            print(f"O professor {professor.nome} já está associado à escola {self.nome}.")
    
    def listar_professores(self):
        if self.professores:
            print(f"A escola {self.nome} tem os seguintes professores:")
            for professor in self.professores:
                print(f"- {professor.nome}")
        else:
            print(f"A escola {self.nome} ainda não tem professores associados.")

escola_a = Escola("Escola A")
escola_b = Escola("Escola B")
professorA = Professor("Thomas")
professorB = Professor("Lucas")

professorA.adicionar_escola(escola_a) 
professorB.adicionar_escola(escola_a)  
professorB.adicionar_escola(escola_b) 

escola_a.listar_professores()
escola_b.listar_professores()

professorA.listar_escolas()
professorB.listar_escolas()