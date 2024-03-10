import requests
import os

def get_weather():
    api_key = os.getenv('WEATHER_API_KEY')
    if not api_key:
        raise ValueError("API ключ не найден. Убедитесь, что вы задали переменную окружения WEATHER_API_KEY.")

    url = f"http://api.openweathermap.org/data/2.5/weather?id=524901&appid={api_key}&units=metric"
    response = requests.get(url)
    if response.status_code == 200:
        data = response.json()
        temperature = data['main']['temp']
        weather_description = data['weather'][0]['description']
        report = f"Текущая температура в Москве: {temperature}°C, погода: {weather_description}"
    else:
        report = "Ошибка при получении данных о погоде"

    print(report)
    # Запись отчета в файл
    with open('weather_report.txt', 'w') as file:
        file.write(report)

if __name__ == "__main__":
    print("Script started")
    get_weather()
