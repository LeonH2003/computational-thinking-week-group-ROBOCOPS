from datetime import datetime

def solution_station_2(date):
    date = datetime.strptime(date, '%Y-%m-%d').date()
    date.weekday()
    if date.weekday() == 0:
        return "月曜日"
    elif date.weekday() == 1:
        return "火曜日"
    elif date.weekday() == 2:
        return "水曜日"
    elif date.weekday() == 3:
        return "木曜日"
    elif date.weekday() == 4:
        return "金曜日"
    elif date.weekday() == 5:
        return "土曜日"
    elif date.weekday() == 6:
        return "日曜日"
