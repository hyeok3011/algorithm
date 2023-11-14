def solution(temperature, t1, t2, a, b, onboard):
    total_power = 0
    target = max(t1 - temperature, temperature - t2)
    temperature_outside 
    for time in range(len(onboard)):
        if onboard[time]:  # 승객이 탑승 중일 때
            if temperature < target:
                temperature += 1
                total_power += a
            elif temperature > target + 1:
                temperature -= 1
                total_power += a
            else:
                total_power += b
        else:  # 승객이 탑승하지 않을 때
            if temperature > temperature_outside:
                temperature -= 1
            elif temperature < temperature_outside:
                temperature += 1
        # 'temperature_outside'는 실외 온도를 나타냅니다. 이는 문제에서 제공되지 않았으므로 추가적인 정보가 필요합니다.

    return total_power
