def gcd(a: int, b: int) -> int:
    """Find the greatest common divisor of a and b
    by using Euclid's reciprocal division method

    Args:
        a (int): integer
        b (int): integer

    Returns:
        int: greatest common divisor of a and b
    """
    a, b = abs(a), abs(b)
    while b:
        a, b = b, a % b
    return a
