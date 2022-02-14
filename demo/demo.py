import requests

base_url = 'http://localhost:8080'
headers = {'Content-type': 'application/json', 'Accept': 'text/plain'}


def main():

    # step 1 - create job
    resp = requests.post(base_url + '/login', json={'username': 'admin', 'password': 'admin'}, headers=headers)
    token = resp.json()['token']
    print(token)
    assert resp.status_code == 200

    headers["Authorization"] = f"Bearer {token}"
    resp = requests.post(base_url + '/task', json={'task': 'newtask'}, headers=headers)
    print(resp)

    return



if __name__ == "__main__":
    main()
