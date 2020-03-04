from darksky.api import DarkSky
from darksky.types import languages, units, weather
from geopy.geocoders import Nominatim

geolocator = Nominatim(user_agent="specify_your_app_name_here")
location = geolocator.geocode("Ciudad de Mexico, Mexico")

API_KEY = '34c1183964f3da3d546e551b2f863648'

darksky = DarkSky(API_KEY)

forecast = darksky.get_forecast(
    location.latitude,
    location.longitude,
    extend=False, # default `False`
    lang=languages.ENGLISH, # default `ENGLISH`
    values_units=units.AUTO, # default `auto`
    exclude=[weather.MINUTELY, weather.ALERTS], # default `[]`,
    #timezone='UTC' # default None - will be set by DarkSky API automatically
)


location = geolocator.reverse(f'{forecast.latitude}, {forecast.longitude}')
print(f'{location.raw["address"]["state"]}, {location.raw["address"]["country"]} = lat: {location.latitude} - lon: {location.longitude} ')
print(f'Temperature: {forecast.currently.temperature}°')
print(f'Current summary: {forecast.currently.summary}')
print(f'Timezone: {forecast.timezone}')

print('''\n  ---------------------------------------------------------------------
      Date      Temperature(avg)              Date summary
  ---------------------------------------------------------------------''')
for day in forecast.daily.data:
    avg = (day.temperature_max + day.temperature_min) / 2
    print(f'   {day.time.date()}        {round(avg, 1)}°           {day.summary}')
print('  ---------------------------------------------------------------------')
print(f'  Forecasted week summary: {forecast.daily.summary}')