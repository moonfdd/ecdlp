from random import randint

def generate_lucas_pseudoprimes(t):
    assert t >= 64, "Smaller numbers you can debug yourself!"
    while True:
        k2 = randint(1, 200) * 2 + 1
        k3 = randint(1, 200) * 2 + k2
        if k2 % 5 == 0 or k3 % 5 == 0:
            continue
        if gcd(k2, k3) != 1:
            continue
        M = k2 * k3 * 20
        x = crt([7, pow(k3, -1, k2), pow(k2, -1, k3)], [20, k2, k3])
        k = int((2 ** (t - 1) // (k2 * k3)) ** (1 / 3)) // M - 1
        p1 = M * k + x
        p2 = k2 * (p1 + 1) - 1
        p3 = k3 * (p1 + 1) - 1
        if p2 % 5 not in [2, 3] or p3 % 5 not in [2, 3]:
            continue
        # Enumerating k
        while True:
            k += 1
            p1 = M * k + x
            if not is_prime(p1):
                continue
            p2 = k2 * (p1 + 1) - 1
            p3 = k3 * (p1 + 1) - 1
            n = p1 * p2 * p3
            if n.nbits() < t:
                continue
            if n.nbits() > t:
                break
            if not is_prime(p2) or not is_prime(p3):
                continue
            break
        if n.nbits() == t:
            break
    return n, p1, p2, p3
n, _, _, _ = generate_lucas_pseudoprimes(512)
print(n)
from Crypto.Math.Primality import lucas_test
print(lucas_test(int(n)))

