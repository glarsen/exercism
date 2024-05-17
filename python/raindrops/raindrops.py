def convert(number: int) -> str:
    sounds = {
        3: "Pling",
        5: "Plang",
        7: "Plong"
    }

    results = (
        sound for divisor, sound in sounds.items() if number % divisor == 0
    )

    return "".join(results) or str(number)
