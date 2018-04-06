import requests

class Client:
    def __init__(self, api_email, api_token):
        self.api_email = api_email
        self.api_token = api_token


    def _submit_batch(self, url, batch):
        # TODO: validate column_name
        batch_parts = batch.parts()

        guid = ""
        headers = {
                'user-agent': 'kentik-python-api/0.1',
                'Content-Type': 'application/json',
                'X-CH-Auth-Email': self.api_email,
                'X-CH-Auth-API-Token': self.api_token
        }

        # submit each part
        for batch_part in batch_parts:
            # submit
            resp = requests.post(url, headers=headers, data=batch_part.build_json(guid))

            # print the HTTP response to help debug
            print resp.text

            # break out at first sign of trouble
            resp.raise_for_status()
            guid = resp.json()['guid']
            if guid == None or len(guid) == 0:
                raise RuntimeError('guid not found in batch response')


    # submit a populator batch
    def submit_populator_batch(self, column_name, batch):
        url = 'https://api.kentik.com/api/v5/tagging/column/%s/populators' % column_name
        self._submit_batch(url, batch)


    # submit a tag batch
    def submit_tag_batch(self, batch):
        url = 'https://api.kentik.com/api/v5/tagging/tags'
        self._submit_batch(url, batch)
