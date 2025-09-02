def solution_station_4(num):
    if num < 2:
        return False
    for i in range(2, int(num ** 0.5) + 1): #checking if it is a prime number
        if num % i == 0:
            return False
    return True
