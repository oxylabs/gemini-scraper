import requests
from pprint import pprint

# Structure payload.
payload = {
    'source': 'gemini',
    'prompt': 'best supplements for better sleep',
    'parse': True,
    'geo_location': "United States",
    'callback_url': "https://your-server.com/oxylabs-callback"
}

# Get response.
response = requests.request(
    'POST',
    'https://data.oxylabs.io/v1/queries',
    auth=('USERNAME', 'PASSWORD'),
    json=payload,
)

# Push-Pull returns the job ID immediately - not the result.
# Once the job status is done, retrieve the parsed response.
pprint(response.json())
