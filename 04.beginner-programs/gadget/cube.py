from collections import defaultdict

def main():
    # Create a dictionary to save the results of (a^3 + b^3) and the corresponding a, b pairs
    cube_sums = defaultdict(list)

    # Calculate all a^3 + b^3 and map them to the corresponding a, b pairs
    for a in range(1, 1000):
        for b in range(a, 1000):
            sum_ = a**3 + b**3
            for c, d in cube_sums[sum_]:
                print(f"a={a}, b={b}, c={c}, d={d}")
            cube_sums[sum_].append((a, b))

    print(len(cube_sums))

if __name__ == "__main__":
    main()