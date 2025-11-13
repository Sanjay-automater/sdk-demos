#!/usr/bin/env python3
import os
import requests
import sys

API_KEY = os.environ.get("API_KEY")
if not API_KEY:
    print("Please set the API_KEY environment variable.")
    sys.exit(1)

# Replace with real SDK import / client usage if available
# Example: from sdk import Client; client = Client(api_key=API_KEY)
ENDPOINT = "https://api.example.com/v1/demo"  # placeholder

def main():
    headers = {"Authorization": f"Bearer {API_KEY}", "Content-Type": "application/json"}
    payload = {"message": "hello from python demo"}
    try:
        resp = requests.post(ENDPOINT, json=payload, headers=headers, timeout=10)
        resp.raise_for_status()
        print("Response status:", resp.status_code)
        print("Body:", resp.text)
    except requests.RequestException as e:
        print("Request failed:", e)
        sys.exit(1)

if __name__ == "__main__":
    main()
