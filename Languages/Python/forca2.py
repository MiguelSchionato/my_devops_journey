import random

# Seus dicionários e listas aqui (mantenha como estão)
palavras = ["BIG","SMALL","LAZY","GOOD","BEAUTIFUL","YELLOW","HAPPY","RED","TALL","SAD","FRIENDLY","HEAVY","STRONG","SMART","LUCKY","CLASSROOM","MUSIC ROOM","CAFETERIA","LIBRARY","COURT","BATHROOM","SCIENCE LAB","COMPUTER LAB","PORTUGUESE","ART","HISTORY","ENGLISH","SCIENCE","MATH","INFORMATICS","GEOGRAPHY","SCHOOL BAG","PENCIL CASE","COLOUR PENCILS","SHARPENER","COMPASS","NOTEBOOK","BOOK","RUBBER","PEN","PENCIL","PAINT BRUSHERS","RULER","WATERCOLOUR PAINT","NOTEPAD","HIGHLIGHTER","GLUE","MARKERS","STAPLER","SCIOSSORS","CALCULATOR",]
dic = {"BIG":"Grande","SMALL":"Pequeno","LAZY":"Preguiçoso(a)","GOOD":"Bom (Boa)","BEAUTIFUL":"Bonito(a)","YELLOW":"Amarelo","HAPPY":"Feliz","RED":"Vermelho","TALL":"Grande (tamanho)","SAD":"Triste","FRIENDLY":"Amigavel (Amistoso)","HEAVY":"Pesado","STRONG":"Forte","SMART":"Inteligente","LUCKY":"Sortudo","CLASSROOM":"Sala de aula","MUSIC ROOM":"Sala de música","CAFETERIA":"CAFETERIA","LIBRARY":"Biblioteca","COURT":"Corte (Tribunal)","BATHROOM":"Banheiro","SCIENCE LAB":"Laboratório de ciências","COMPUTER LAB":"Laboratório de computação","PORTUGUESE":"Português","ART":"Arte","HISTORY":"História","ENGLISH":"Inglês","SCIENCE":"Ciências","MATH":"Matemática","INFORMATICS":"Informática","GEOGRAPHY":"Geografia","SCHOOL BAG":"Mochila escolar","PENCIL CASE":"Estojo","COLOUR PENCILS":"Lápis de cor","SHARPENER":"Apontador","COMPASS":"Compasso","NOTEBOOK":"Caderno","BOOK":"Livro","RUBBER":"Borracha","PEN":"Caneta","PENCIL":"Lápis","PAINT BRUSHERS":"Pincel","RULER":"Régua","WATERCOLOUR PAINT":"Tinta a base de água","NOTEPAD":"Caderno de notas","HIGHLIGHTER":"Marca texto","GLUE":"Cola","MARKERS":"Canetinha (Marcador)","STAPLER":"Grampeador","SCIOSSORS":"Tesoura","CALCULATOR":"Calculadora",}
positivo = ["sim","vamos","bora","yes","partiu","s"]
negativo = ["não","n","nao","negativo","no"]

forca = [r"""      

----------------
|              |
|
|
|
|
|
|
|
|
|""",
r"""

----------------
|              |
|              
|              O
|
|
|
|
|
|""",
r"""            

----------------
|              |
|              
|              O
|              |
|
|
|
|
|""",         
r"""            
----------------
|              |
|
|              O
|             \|
|              | 
|
|
|
|""",
r"""            
 
----------------
|              |
|              
|              O
|             \|/
|              |
|
|
|
|""",
r"""


----------------
|              |
|
|              O
|             \|/
|              |
|             /
|
|
|""",
r"""
 
 
----------------
|              |
|
|              O
|             \|/
|              |
|             / \
|
|
|"""]

dificuldade = 2 

while True:
    pergunta_ini = input("Vamos jogar um jogo?? ").lower()

    if pergunta_ini in positivo:
        # Reinicia as variáveis do jogo para uma nova partida
        traços = []
        erros = []
        aleatorio = random.choice(palavras)

        for letra_palavra in aleatorio:
            if letra_palavra == " ":
                traços.append(" ")
            else:
                traços.append("___")
        
        erros_cometidos = 0
        chances_totais = (len(forca) - 1) * dificuldade

        print("Oba, o jogo dessa vez é forca. Vou pensar em uma palavra e você tem 12 chances de acertar a minha palavra.")
        print("A dica é: É uma palavra em inglês que significa", dic[aleatorio])

        while "___" in traços and erros_cometidos < chances_totais:
            indice = erros_cometidos // dificuldade
            if indice >= len(forca):
                indice = len(forca) - 1
            
            print("\n" + "="*30)
            print(forca[indice])
            print("  ".join(traços))
            print(f"Erros: {' '.join(erros)}")
            print(f"Chances restantes: {chances_totais - erros_cometidos}")
            
            chute = input("Digite uma letra: ").upper()
            
            if len(chute) > 1 or not chute.isalpha():
                print("Por favor, digite apenas uma letra por vez, sem trapacear, hein!")
                continue
            
            if chute in erros or chute in traços:
                print("Você já tentou essa letra, tente outra.")
                continue
            
            acerto = False
            for ind, letra in enumerate(aleatorio):
                if letra == chute:
                    traços[ind] = chute
                    acerto = True
            
            if not acerto:
                print("Que pena, você errou.")
                erros_cometidos += 1
                erros.append(chute)
            else:
                print(f"Boa, a letra '{chute}' está na palavra!")

        indice_desenho_final = erros_cometidos // dificuldade
        if indice_desenho_final >= len(forca):
            indice_desenho_final = len(forca) - 1
        print(forca[indice_desenho_final])

        if "___" not in traços:
            print(f"Parabéns, você acertou! A palavra era '{aleatorio}'")
            print("Que significa:", dic[aleatorio])
        else:
            print(f"Mais sorte na próxima! A palavra era '{aleatorio}'")
            print("Que significa:", dic[aleatorio])

        print("="*30)
        
        rejogar = input("Você quer jogar mais uma partida?? ").lower()

        if rejogar in negativo:
            print("Que pena, talvez na próxima vez :(")
            break

    elif pergunta_ini in negativo:
        print("Que pena, talvez na próxima vez :(")
        break
    else:
        print("Não entendi o que você disse. Por favor, responda 'sim' ou 'não'.")