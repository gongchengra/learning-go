from collections import defaultdict

def main():
    # Create a dictionary to save the results of (a^3 + b^3) and the corresponding a, b pairs
    cube_sums = defaultdict(list)

    # Calculate all a^3 + b^3 and map them to the corresponding a, b pairs
    for a in range(1, 1000):
        for b in range(a, 1000):
            sum_ = a**3 + b**3
            cube_sums[sum_].append((a, b))

    print(len(cube_sums))
    # Collect all combinations
    quads = []

    # Loop through the dictionary to find all solutions where c^3 + d^3 have the same result
    for pairs in cube_sums.values():
        if len(pairs) > 1:
            for i in range(len(pairs)):
                for j in range(i+1, len(pairs)):
                    # Add non-repeating combinations
                    quads.append((pairs[i][0], pairs[i][1], pairs[j][0], pairs[j][1]))

    # Globally sort all combinations
    quads.sort()

    # Print all the sorted combinations
    for q in quads:
        print(f"a={q[0]}, b={q[1]}, c={q[2]}, d={q[3]}")

if __name__ == "__main__":
    main()
