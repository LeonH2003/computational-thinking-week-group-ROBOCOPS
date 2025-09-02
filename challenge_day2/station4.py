<<<<<<< HEAD
def solution_station_4(num):
    if num < 2:
        return False
    for i in range(2, int(num ** 0.5) + 1):
        if num % i == 0:
            return False
    return True
=======
def solution_station_4(num):
>>>>>>> ba788208413770e609a2871cbb67f44e8a51b608
