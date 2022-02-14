import requests

base_url = 'http://localhost:8080'
headers = {'Content-type': 'application/json', 'Accept': 'text/plain'}


def main():

    # step 1 - create job
    resp = requests.post(base_url + '/login', json={'username': 'admin', 'password': 'admin'}, headers=headers)
    token = resp.json()['token']
    print(token)
    assert resp.status_code == 200

    
    return


if __name__ == "__main__":
    main()
