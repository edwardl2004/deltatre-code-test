import requests
import json


def test_search_word_default_list(base_url):
    url = base_url + "v1/words/search"

    _search_and_verify_result(url, "hello", 6, True)
    _search_and_verify_result(url, "simple", 5, True)
    _search_and_verify_result(url, "goodbye", 4, True)
    _search_and_verify_result(url, "filter", 1, True)
    _search_and_verify_result(url, "yes", 8, True)

    url = base_url + "v1/words"
    r = requests.get(url)
    j = r.json()
    assert r.status_code == 200
    assert j == {
        "list": [
            {
                "word": "yes",
                "count": "8"
            },
            {
                "word": "hello",
                "count": "6"
            },
            {
                "word": "simple",
                "count": "5"
            },
            {
                "word": "goodbye",
                "count": "4"
            },
            {
                "word": "filter",
                "count": "1"
            }
        ]
    }


def test_update_list(base_url):
    url = base_url + "v1/words"
    payload = {
        "words": ["golang", "world"]
    }
    r = requests.post(url, data=json.dumps(payload))
    assert r.status_code == 200
    j = r.json()
    assert j == {
        "status": "success",
        "message": ""
    }

    url = base_url + "v1/words/search"
    # search existing words
    _search_and_verify_result(url, "golang", 2, True)
    _search_and_verify_result(url, "world", 3, True)
    # search non-existing words
    _search_and_verify_result(url, "nonexistingword", 3, False)

    url = base_url + "v1/words"
    expected_response = {
        "list": [
            {
                "word": "world",
                "count": "3"
            },
            {
                "word": "golang",
                "count": "2"
            }
        ]
    }
    r = requests.get(url)
    j = r.json()
    assert r.status_code == 200
    for topsearch in expected_response["list"]:
        assert topsearch in j["list"]


def test_update_long_list(base_url):
    url = base_url + "v1/words"
    payload = {
        "words": ["grpc", "client", "api", "server", "protoc", "gateway", "runtime"]
    }
    r = requests.post(url, data=json.dumps(payload))
    assert r.status_code == 200
    j = r.json()
    assert j == {
        "status": "success",
        "message": ""
    }

    url = base_url + "v1/words/search"
    # search existing words
    _search_and_verify_result(url, "server", 2, True)
    _search_and_verify_result(url, "client", 3, True)
    _search_and_verify_result(url, "grpc", 3, True)
    _search_and_verify_result(url, "gateway", 4, True)
    # search non-existing words
    _search_and_verify_result(url, "nonexistingword", 3, False)

    url = base_url + "v1/words"
    expected_response = {
        "list": [
            {
                "word": "server",
                "count": "2"
            },
            {
                "word": "client",
                "count": "3"
            },
            {
                "word": "grpc",
                "count": "3"
            },
            {
                "word": "gateway",
                "count": "4"
            }
        ]
    }
    r = requests.get(url)
    j = r.json()
    assert r.status_code == 200
    for topsearch in expected_response["list"]:
        assert topsearch in j["list"]


def _search_and_verify_result(url, word, times, found):
    payload = {
        "word": word,
    }

    for i in range(times):
        r = requests.post(url, data=json.dumps(payload))
        j = r.json()
        assert r.status_code == 200
        assert j == {"found": found}
