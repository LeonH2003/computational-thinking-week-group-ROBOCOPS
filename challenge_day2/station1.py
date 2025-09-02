import math

def solution_station_1(num):
    #fibonacci sequence
    phi = (1 + math.sqrt(5)) / 2
    for n in range(0, num + 1):
        result = round(pow(phi, n) / math.sqrt(5))   
    return result    

