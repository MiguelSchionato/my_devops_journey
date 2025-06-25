print("Welcome to the rollercoaster!")
bill = 0

def questionary():
    int(height)
    if height >= 120:
        print("You can ride the rollercoaster")
        age = int(input("What's your age?\n"))
        if age < 12:
            bill = 5
            print("Child thickets are $5.")
        elif age <= 18:
            bill = 7
            print("Youth tickets are $7.")
        else:
            bill = 12
            print("Adult tickets are $12.")

        photo = input("Do you want a photo taken? Y or N\n").upper()
        if photo == "Y":
            bill += 3
        print(f"Your final bill is ${bill}")

    else:
        print("Sorry, you have to grow taller to ride the rollercoaster")


while True:
    try:
        height = int(input("What's your height in centimeters?\n"))
        questionary()
        break
    except:
        print("Sorry, not an available anwser")


