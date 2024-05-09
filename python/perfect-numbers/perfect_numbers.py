def classify(number: int) -> str:
    """ A perfect number equals the sum of its positive divisors.

    :param number: int a positive integer
    :return: str the classification of the input integer
    """

    if number < 1:
        raise ValueError("Classification is only possible for positive integers.")

    aliquot = sum([n for n in range(1, number) if not number % n])

    if number < aliquot:
        return "abundant"
    elif number > aliquot:
        return "deficient"

    return "perfect"
