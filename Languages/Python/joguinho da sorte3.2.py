import random
secret = random.randint (1, 100)
positiva = ["sim", "vamos", "claro", "bora", "s",]
negativa = ["nao", "não", "não to a fim", "negativo", "sai fora", "n"]
while True:
    pergunta = input("Vamos jogar um joguinho?").lower()
    if pergunta in positiva:
        print("Ótimo, vamos lá")
        print("Estou pensando em um número entre 1 e 100.")    
        tentativa = 0
        while True:
            try:
                chute = int(input("tente adivinhar o número que estou pensando!"))
                tentativa += 1
                if chute > 100:
                    print("Eu disse que é um número entre 1 e 100...")
                else:
                    if chute != secret:
                        print("haha, errou.")
                        if chute > secret:
                            print("Mas ta quase lá, o número que estou pensando é menor")
                        else: 
                            print("Quase lá, mas o número que estou pensando é maior")
                    else:
                        if tentativa == 1:
                            print("Uau, inacreditável, acertou de primeira, sorte ou telepatia??")
                            break
                        elif tentativa <= 5: 
                            print(f"Excelente, acertou em {tentativa} tentativas!!")
                            break
                        elif tentativa <= 10:
                            print(f"Impressionante, mas poderia melhorar. Acertou em {tentativa} tentativas")
                            break
                        else:
                            print("Exatamente, como adivinhou??")
                            print(f"Acertou em {tentativa} tentativas, nada mal.")
                            break

            except ValueError:
                print("Engraçado, mas isso não é um número .-. ")
        break
    elif pergunta in negativa:
        print("Poxa, que pena! talvez na próxima vez :(")
        break
    else:
        print("Não entendi o que disse, quer jogar ou não?")
