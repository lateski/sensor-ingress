import requests
api_key = 'api-secret'

def main():
    temp_file = open('/sys/class/thermal/thermal_zone0/temp', 'r')
    temp = temp_file.readline()
    temp_file.close()
    temp = float(temp)/1000.0
    submit_sensor_temp("raspi_cpu", temp)


def submit_sensor_temp(sensor,temp):
    headers = {'Authorization': 'Bearer '+api_key}
    payload = {'name': sensor, 'value': temp}
    requests.post( 'http://localhost:9100/sensor/',headers=headers, json=payload )


if __name__ == '__main__':
    main()
